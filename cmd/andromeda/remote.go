package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"jaytaylor.com/andromeda/crawler"
	"jaytaylor.com/andromeda/db"
	"jaytaylor.com/andromeda/domain"
	"jaytaylor.com/andromeda/pkg/openssl"
)

func newRemoteCmd() *cobra.Command {
	remoteCmd := &cobra.Command{
		Use:   "remote",
		Short: "RPCs",
		Long:  "RPC operations",
	}

	remoteCmd.AddCommand(
		newRemoteEnqueueCmd(),
		newRemoteCrawlerCmd("crawler"),
	)

	return remoteCmd
}

func newRemoteEnqueueCmd() *cobra.Command {
	var onlyIfNotExists bool

	remoteEnqueueCmd := &cobra.Command{
		Use:   "enqueue <pkg-path-1> [<pkg-path-2> ...]",
		Short: "Add items to the to-crawl queue (or pass '-' to read from STDIN)",
		Long:  "Add items to the to-crawl queue (or pass '-' to read from STDIN)",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				cfg    = crawler.NewConfig()
				r      = crawler.NewRemote(CrawlServerAddr, cfg)
				now    = time.Now()
				tokens = args
			)

			if err := configureRemoteCrawler(r); err != nil {
				log.Fatalf("main: configuring remote crawler: %s", err)
			}

			if len(args) > 0 && args[0] == "-" {
				// Tokenize STDIN.
				s := bufio.NewScanner(os.Stdin)
				expr := regexp.MustCompile("[ \t\n\r]+")
				tokens = []string{}
				for s.Scan() {
					txt := s.Text()
					// To prevent submission of empty-string entries, remove leading or
					// trailing separator characters.
					for _, sep := range []string{" ", "\t", "\n", "\r"} {
						txt = strings.Trim(txt, sep)
					}
					tokens = append(tokens, expr.Split(txt, -1)...)
				}
				if err := s.Err(); err != nil {
					log.Fatal(err)
				}
			}

			toCrawls := make([]*domain.ToCrawlEntry, len(tokens))
			for i, arg := range tokens {
				toCrawls[i] = &domain.ToCrawlEntry{
					PackagePath: arg,
					Reason:      EnqueueReason,
					SubmittedAt: &now,
					Force:       EnqueueForce,
				}
			}
			opts := db.NewQueueOptions()
			opts.Priority = db.DefaultQueuePriority
			opts.OnlyIfNotExists = onlyIfNotExists

			n, err := r.Enqueue(toCrawls, opts)
			if err != nil {
				log.Fatal(err)
			}
			log.WithField("num-submitted", len(tokens)).WithField("num-added", n).Info("Enqueue operation successfully completed")
		},
	}

	remoteEnqueueCmd.Flags().IntVarP(&db.DefaultQueuePriority, "priority", "p", db.DefaultQueuePriority, "Priority level to use when adding items to the queue")
	remoteEnqueueCmd.Flags().StringVarP(&EnqueueReason, "reason", "r", EnqueueReason, "Reason to use for to-crawl entries")
	remoteEnqueueCmd.Flags().BoolVarP(&EnqueueForce, "force", "f", EnqueueForce, "Force queue insertion and crawl")
	remoteEnqueueCmd.Flags().BoolVarP(&onlyIfNotExists, "only-if-not-exists", "", onlyIfNotExists, "Only enqueue if the package does not already exist")
	addRemoteFlags(remoteEnqueueCmd)

	return remoteEnqueueCmd
}

func newRemoteCrawlerCmd(aliases ...string) *cobra.Command {
	remoteCrawlerCmd := &cobra.Command{
		Use:     "remote-crawler",
		Aliases: aliases,
		Short:   "Start a remote crawler worker",
		Long:    "Launch a remote crawler gRPC worker",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				cfg       = crawler.NewConfig()
				r         = crawler.NewRemote(CrawlServerAddr, cfg)
				stopCh    = make(chan struct{})
				runDoneCh = make(chan struct{}, 1)
			)

			if err := configureRemoteCrawler(r); err != nil {
				log.Fatalf("main: configuring remote crawler: %s", err)
			}

			if len(args) > 0 {
				// Crawl specified packages.
				if err := r.Do(args, stopCh); err != nil {
					log.Fatal(err)
				}
				return
			}

			go func() {
				r.Run(stopCh)
				runDoneCh <- struct{}{}
			}()

			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
			select {
			case s := <-sigCh:
				log.WithField("sig", s).Info("Received signal, shutting down remote crawler")
				stopCh <- struct{}{}
				log.Info("Exiting")
				return

			case <-runDoneCh:
				log.Info("Remote crawler run finished")
				log.Info("Exiting")
				return
			}
		},
	}

	addCrawlerFlags(remoteCrawlerCmd)

	addRemoteFlags(remoteCrawlerCmd)

	return remoteCrawlerCmd
}

func addRemoteFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringVarP(&CrawlServerAddr, "addr", "a", CrawlServerAddr, "Crawl server host:port address spec")

	cmd.Flags().StringVarP(&TLSCertFile, "cert", "c", TLSCertFile, "SSL/TLS public key certifcate file for mutual CA-based authentication, or in the case of a client connecting over HTTPS, this will be the SSL/TLS public-key certificate file belonging to the SSL-terminating server.")
	cmd.Flags().StringVarP(&TLSKeyFile, "key", "k", TLSKeyFile, "SSL/TLS private key certificate file for mutual TLS CA-based authentication")
	cmd.Flags().StringVarP(&TLSCAFile, "ca", "", TLSCAFile, "ca.crt file for mutual TLS CA-based authentication")
	cmd.Flags().BoolVarP(&AutoTLSCert, "auto-cert", "C", AutoTLSCert, "Use OpenSSL to automatically fill in the SSL/TLS public key of the gRPC server")

	return cmd
}

func configureRemoteCrawler(r *crawler.Remote) error {
	if AutoTLSCert {
		log.WithField("addr", r.Addr).Debug("Automatically downloading SSL/TLS certificate public-key")
		var err error
		if TLSCertFile, err = openssl.DownloadCertPubKey(r.Addr); err != nil {
			return err
		}
	}

	if len(TLSCertFile) == 0 && len(TLSKeyFile) == 0 && len(TLSCAFile) == 0 {
		return nil
	}

	var creds credentials.TransportCredentials
	if len(TLSCAFile) > 0 || len(TLSKeyFile) > 0 {
		// Implies intent to use mutual TLS CA authentication.
		log.WithField("cert", TLSCertFile).WithField("key", TLSKeyFile).WithField("ca", TLSCAFile).Debug("Mutual TLS CA authentication activated")

		if len(TLSCertFile) == 0 || len(TLSKeyFile) == 0 || len(TLSCAFile) == 0 {
			return errors.New("mutual TLS CA authentication required all 3 certificate flags: public.cert, private.key, and ca.cert")
		}

		// Load the client certificates from disk.
		certificate, err := tls.LoadX509KeyPair(TLSCertFile, TLSKeyFile)
		if err != nil {
			return fmt.Errorf("could not load TLS key pair cert=%v key=%v: %s", TLSCertFile, TLSKeyFile, err)
		}

		// Create a certificate pool from the certificate authority.
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(TLSCAFile)
		if err != nil {
			return fmt.Errorf("could not read ca certificate: %s", err)
		}

		// Append the certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			return errors.New("failed to append ca certs")
		}

		creds = credentials.NewTLS(&tls.Config{
			// N.B.: ServerName _must_ align with the Common Name (CN) in the certificate.
			ServerName:   strings.Split(CrawlServerAddr, ":")[0],
			Certificates: []tls.Certificate{certificate},
			RootCAs:      certPool,
		})
	} else {
		// Implies intent to use SSL/TLS transport.
		log.WithField("cert", TLSCertFile).Debug("Client SSL/TLS transport activated")

		var err error

		if creds, err = credentials.NewClientTLSFromFile(TLSCertFile, ""); err != nil {
			return err
		}
	}
	r.DialOptions = append(r.DialOptions, grpc.WithTransportCredentials(creds))
	return nil
}
