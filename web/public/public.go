// Code generated by go-bindata.
// sources:
// index.tpl
// list.tpl
// package.tpl
// subpackage.tpl
// DO NOT EDIT!

package public

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xeb\x93\xdb\xb6\x11\xff\x7e\x7f\xc5\x9a\x4e\x47\xb2\x6c\x91\x3a\xbb\x6e\x33\x3c\x92\x9d\xf3\xf9\x52\xb7\xe3\xa4\x19\xdf\x75\x3a\x9d\x34\x1f\x20\x62\x45\x22\x07\x02\x2c\x00\x4a\xa7\x6a\xf4\xbf\x77\xc0\x87\x44\xf1\xa1\xbb\x64\x1c\x7c\xd0\x03\xfb\xfc\xed\x2e\x96\x4b\x04\x2f\xa8\x8c\xcd\x36\x47\x48\x4d\xc6\xa3\x8b\xa0\xf9\x42\x42\xa3\x0b\x00\x80\xc0\x30\xc3\x31\xba\x16\x54\xc9\x0c\x29\x09\xbc\x6a\xa3\x22\x66\x68\x08\xc4\x29\x51\x1a\x4d\xe8\x14\x66\x35\xff\xd6\x69\x93\x04\xc9\x30\x74\x48\x61\x52\xa9\x1c\x88\xa5\x30\x28\x4c\xe8\xfc\x9d\x6c\xe1\x9e\x6c\xb9\x54\x03\xdc\x14\x75\xac\x58\x6e\x98\x14\x2d\x91\x83\x03\xc0\x34\x10\x48\x14\xc9\x53\x90\x2b\x30\x29\x02\x0a\xc3\x14\xc2\x83\x90\x1b\x01\x6b\xa6\xd9\x92\x23\x24\x12\x0a\xc1\xd6\xa8\x34\x3a\xd1\x45\x65\xa5\x52\x0c\x16\x70\xe8\x18\x7c\x34\xde\x2f\x64\x4d\xaa\xdd\xda\x13\xbb\x36\x4c\x50\xb9\x71\xa5\xe0\x92\x50\x08\x61\x55\x88\xd8\xba\x03\xd3\x57\xb0\x3b\x70\xd9\xb5\x26\xca\xba\x28\xae\x7a\xbb\xda\x10\x53\x68\x08\x81\xca\xb8\xc8\x50\x18\x37\x41\x73\xcb\xd1\xfe\xfc\xb0\xfd\x1b\x9d\x4e\x2a\x8e\xc9\xab\xbe\x6c\xa6\x93\x73\x82\x99\x4e\x86\xa4\xb8\x3c\x2b\xc5\x65\x4f\xea\x80\x8b\xe4\x39\x0a\xfa\x59\x26\x53\x66\x30\xeb\x82\x6c\x0c\x50\x79\x17\x2b\xc9\x39\x84\xd6\x96\xab\xcb\x3f\xf7\x32\x87\xa8\xf5\xff\x13\xb2\x24\x35\x30\x2f\xb7\x62\xce\x50\x98\xc3\xd6\xe5\x55\x4f\xaf\xe5\xaa\xac\xdf\xa4\x8c\xd3\xca\x7e\x9f\x8d\xad\x60\xda\x98\x1f\x72\xaf\x51\x75\xf4\x29\x7c\x8e\x4f\x7d\x43\xfb\x8b\xf1\x7f\xde\x6c\x34\xb8\x2b\xa9\xb2\xc9\x2b\x57\x0a\x5d\x2c\x33\x66\xce\xd6\x4c\x83\xe7\x85\x2d\x9c\x31\x30\x0a\x4d\xa1\x04\xac\x08\xd7\xf8\x94\x97\x07\x85\x99\x4e\xdc\x35\xe1\x05\x7e\x2d\xad\xd6\x43\x57\xa3\xa0\xd3\xa3\xea\xbe\xdc\x81\x06\x21\x4c\x26\x7d\xfa\xb8\xd5\xfd\xd5\xcc\xbb\xe8\xc2\xa8\x4e\xdf\x4f\x93\x7f\xe1\xf2\x4e\xc6\x0f\x68\x26\x3f\x8f\x55\x64\xae\xa4\x91\xd6\x6a\x6a\x4c\xae\xfd\x09\x84\xad\x03\xc0\x65\x4c\x6c\x06\xdc\x92\x2b\x96\x1c\xfe\x02\x93\x8d\xd6\x13\xf0\xed\xf7\x80\xa3\x56\xe5\x46\xdf\x48\x21\x30\x6e\xe7\x70\x30\x85\x76\x55\x47\xd8\x65\x42\xa0\xfa\x74\xff\xfd\x67\xeb\x4a\xb0\x8c\x6a\x0d\x4c\x24\xae\x1b\x78\xcb\x68\x72\x75\x31\x28\x6e\xc3\x0b\x21\x08\xdc\xc0\x01\xec\xb4\xc2\xf4\x1a\x26\xbe\xe7\x4d\xe0\xf5\x00\x9e\x54\x6a\x63\x19\xbc\x4d\xaf\x7b\x9c\x24\x4e\x0a\x99\xa3\x68\xe3\xc0\xb5\x19\x83\x52\x4b\x69\xc9\xd1\xe5\x32\x29\x59\x87\x75\x3f\x0d\x1c\x69\x8d\x7a\x48\x74\x3f\xbc\xed\xcd\x6a\x9f\x51\x29\xa9\x7e\x17\xa7\x6d\x7a\x6d\x7f\x69\x77\xc9\x58\x21\x31\x58\x9f\xe5\xe9\x84\xb2\xf5\x58\x48\xed\xb2\xd2\x3d\xd0\xb7\x5f\xbe\xfc\xe3\x8b\x0f\x36\x57\xb8\x2e\x13\x73\x06\xbd\x5d\x9d\x6e\x3b\x16\xa6\xce\xd1\x68\xe1\xb5\x61\x8a\xb9\xd4\xf8\x2b\xc2\xf4\x44\xa9\x4a\x01\xa5\x46\xea\x3e\xe1\xbc\x46\x73\xcf\x32\x94\x85\x99\x1e\x8e\xca\x1b\xb8\x5c\x2c\x16\x8b\x51\x24\xe7\x70\x64\xa8\x35\x49\x7e\x0d\x92\xf2\x09\x59\x49\xd9\xe7\x2b\xae\x8d\x4b\x89\x21\xae\xce\x39\x33\xd3\xc9\x7f\xc4\xb9\x0c\xae\xa4\x82\x69\x59\x08\x10\xc2\xe2\x0a\x18\x04\x07\x65\x2e\x47\x91\x98\xf4\x0a\xd8\xeb\xd7\xe7\x1c\x80\xaf\x52\x4b\x70\x52\x4f\xf7\xf8\x68\x3b\x4e\xe3\xca\x4f\xec\xe7\xf3\x92\xcf\xaa\x21\x18\x6c\xea\x30\x9c\x93\x81\xad\x43\x82\xa7\x1d\x03\x7b\x40\xae\x71\xa4\x29\xff\xb6\xb0\x0c\x1e\xad\x7f\xcb\x42\xc1\x52\xc9\x8d\x46\x3b\x7e\xa0\x06\x21\x0d\xe8\x22\xcf\xa5\x32\xc7\x9e\xa9\xc7\x8a\xf6\x6c\x94\x8e\x91\xa9\xa1\x07\x5e\x35\x09\xd6\x23\xa9\x36\x5b\x8e\xed\x59\x31\xd6\xba\x35\x24\x7a\x33\x3b\x2b\x77\x62\x20\xd7\xa8\x56\x5c\x6e\x7c\x48\x19\xa5\xd8\x9a\x0b\xf7\xad\xe3\xfc\x52\x1b\x85\x24\x9b\xdb\xe1\x96\x30\x81\xaa\xa3\x25\x27\x94\x32\x91\xf8\xb0\x38\xf5\x38\x23\x2a\x61\xa2\xb7\xbd\x61\xd4\xa4\x3e\x7c\xbb\x58\xe4\x8f\xa7\x94\xb4\x9c\x72\x7c\x78\xfb\xbe\x47\x5a\x92\xf8\x21\x51\xb2\x10\xd4\xb7\xc3\xf4\xb6\xe5\xe9\xd1\x4f\x3b\x52\xee\x46\xc5\x36\x29\x33\xf8\x2c\x0f\x5b\x80\xc0\x9b\x2d\xdc\xf7\x98\x41\xef\xb3\xdb\xf0\x7e\x13\xac\x63\x02\x48\x61\xe4\x20\x28\x3b\xab\x8d\x06\xbc\x71\xa8\x76\xe9\x39\xe0\xbc\x59\x2e\x35\xb3\x9d\xcb\x07\xb2\xd4\x92\x17\x06\xbb\x58\xbc\xd9\x52\x1a\x23\x33\x1f\x2e\xfb\x40\xbd\x19\xc7\x95\xf1\xc1\x82\x19\x8e\xc1\xbb\x7e\x0c\xbc\x59\xaf\xd6\x5a\xb2\xfb\xa6\xa0\x6d\x0d\x1f\xde\x7e\xce\x17\xf4\x40\x39\x7b\xb3\x95\x14\x66\xbe\x22\x19\xe3\x5b\x1f\x9c\x4f\xc8\xd7\x68\x58\x4c\xe0\x07\x2c\xd0\x79\x03\x87\x8d\x37\x70\xad\x18\xe1\x6f\x40\x13\xa1\xe7\x1a\x15\x5b\x75\xb1\x9c\xa8\x6a\xb1\x9d\x30\xcd\x37\xb8\x7c\x60\x66\x6e\xfd\x9b\x6b\xf6\x3f\x9c\x13\xfa\x4b\xa1\x8d\x6f\x1f\x31\x7f\xe8\xf0\x66\xfa\x59\x7c\xb1\xe4\x52\xf9\xf0\xf2\x5d\xb9\x86\x6a\x82\x74\x70\x1f\x24\xfe\xfc\xfe\xf6\xfa\xc3\xa9\xb6\xd2\x22\xc5\x58\x2a\x52\xe5\x5c\x48\x81\x83\x4a\x7d\x12\x1b\xb6\xee\xb6\xc9\x9e\x7c\x21\x28\x2a\xce\x86\x95\x14\xdd\x94\x70\xa6\xcd\xbc\xcc\xe4\xdc\x66\x72\xdc\xbc\x6b\x64\x3e\x17\x64\x3d\x7a\x80\xe7\x0d\xca\xdb\xc5\xed\x87\xef\xde\x77\xda\x83\x54\x14\x95\x0f\x97\xf9\x23\x68\xc9\x19\x85\x97\x1f\x2f\x6f\x2f\xbf\x5b\x8c\x9a\xca\x49\xfc\x60\x1f\xe4\x4f\x9b\x2b\x57\x57\xcf\xb1\x56\x03\xaf\xba\x7d\x08\x96\x92\x6e\xa3\x8b\x80\xb2\x35\xc4\x9c\x68\x1d\x3a\x35\xa4\xe6\xd2\xa0\xe0\xc7\xea\x0d\x38\x8b\x82\xf4\x32\x0a\x08\xa4\x0a\x57\xa1\xe3\x39\xed\x4b\x0b\x12\x05\x9e\xa5\x7a\x9c\x75\x64\x0e\xfc\x64\x29\x0b\xe3\x44\xd7\xf6\xab\x12\x68\x78\x03\xcf\x5a\x0a\x3c\xca\xd6\xd1\x45\x90\xbe\x8d\xee\x90\xa8\x38\x6d\xdf\x3d\x34\xb7\x0e\x7f\x95\x9c\x88\x04\xfe\x59\xdf\x3c\x04\x5e\xfa\xd6\x22\x51\x16\x97\x05\xc2\x68\xe8\x74\xbb\x7f\x03\x27\x7d\x17\xdd\x28\xb2\xe1\x70\x57\x32\x04\x5e\xfa\xae\xa6\x58\xc3\xd0\x1a\xd5\xaa\x61\xce\x87\x40\xe7\x44\xd4\x3a\xed\x8e\x13\x05\x9e\xdd\x8a\x6a\x5f\x1b\xe1\x92\x85\xcb\xc4\x39\x21\xbc\x98\xcf\x83\xb2\x19\x5a\xaa\xfd\xd1\x6a\x06\x01\x13\x79\xd1\xdc\x94\x54\x2f\xb5\x0e\x94\x2f\x79\xa1\x73\x87\x82\x3a\xe0\x8d\x30\xdb\xfa\x76\x4a\x95\x99\x4e\x1c\xb0\x67\x33\x74\xfe\xf4\x47\xc7\x6b\x62\x69\x2d\x45\xf3\xf9\x31\xa0\x65\x74\x76\x3b\xd8\x30\x93\xc2\x37\x9c\x18\xd4\x06\xfc\x10\xdc\x1b\x29\x56\x2c\x71\xbf\x27\xda\xa0\x72\x3f\x57\x84\xfd\xde\xf2\xb2\xd5\x81\x73\xbf\x2f\x43\x7b\x0c\x62\xcd\x58\xc6\x12\x29\xfc\x58\x15\xa6\x6e\xc5\xb3\xe0\x50\x56\x5b\xe8\x0c\x1e\x27\x27\xb2\x26\x14\x11\x09\xc2\x37\xf9\x43\x62\x9d\x69\x59\x6b\x97\xcf\x6e\x57\x72\xb8\x1f\xed\x04\x7a\x53\x4e\x3a\xf4\xda\x72\xc1\xb1\xb0\x1a\x9e\x1f\x89\x49\x61\xbf\x77\xa2\xce\xc6\xb1\xd8\x76\x3b\x40\x41\x1b\x1b\x27\x65\x77\x24\x1d\x7f\xfd\x50\x64\x4b\x54\x20\x57\x50\x1f\x3e\x0d\x4c\x00\x13\x14\x1f\x7d\xd8\xed\xc0\xfd\xf8\xc1\x6d\xd0\x7f\x46\x51\x46\xca\x06\xbb\xaa\xb2\xff\x16\x58\x20\x54\x33\xef\x81\xfd\x5e\x96\xc4\x13\xf6\xeb\xaa\x8d\x6d\xa4\x7a\x40\xa5\x4b\xd6\xd2\x46\x37\x41\x77\x86\x18\x0d\x8e\xc2\x4c\x1a\xd4\xce\xa9\x39\x0d\x9a\x89\x18\xed\x2b\x7b\xa2\x48\x66\x4b\x58\x19\xa4\x4f\x6a\x8b\x4b\xe1\x4a\x99\x57\x77\x04\xaf\xba\xa5\xfc\x7f\x00\x00\x00\xff\xff\x47\xda\x99\xaa\xbd\x14\x00\x00")

func indexTplBytes() ([]byte, error) {
	return bindataRead(
		_indexTpl,
		"index.tpl",
	)
}

func indexTpl() (*asset, error) {
	bytes, err := indexTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.tpl", size: 5309, mode: os.FileMode(436), modTime: time.Unix(1532989243, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _listTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x6b\xc3\x30\x0c\x85\xef\xfe\x15\x6f\x66\xc7\x75\xb9\x8e\xe1\x18\x76\xdd\xa9\x87\xfd\x01\x35\x56\x63\xd3\x44\x0e\x8e\xda\x11\x4c\xfe\xfb\x68\xc3\x4a\x07\x3b\xd9\x42\xef\x7b\xd2\x93\x7b\x0a\xb9\xd3\x65\x62\x44\x1d\x07\x6f\xdc\xef\xc3\x14\xbc\x01\x00\xa7\x49\x07\xf6\xb5\x62\x60\xc1\x2b\xd6\x15\x7b\xea\x4e\xd4\xf3\x8c\x1d\x3e\x24\x94\x3c\x72\x20\xd7\x6c\xba\x8d\x19\x59\x09\x5d\xa4\x32\xb3\xb6\xf6\xac\xc7\xdd\x9b\x7d\x6c\x09\x8d\xdc\x5a\x3a\x6b\xcc\xc5\xa2\xcb\xa2\x2c\xda\xda\x4f\x5a\xf0\x45\xcb\x90\xcb\x3f\xea\xc0\x73\x57\xd2\xa4\x29\xcb\x03\x72\x5f\x00\x69\x06\xa1\x2f\x34\x45\xe4\x23\x34\x32\x58\x34\x15\xc6\x49\xf2\xb7\xe0\x92\xe6\x74\x18\x18\x7d\xc6\x59\xd2\x85\xcb\xcc\xd6\x1b\xd7\x6c\x49\xdd\x21\x87\xc5\x1b\x53\x2b\x0a\x49\xcf\x78\x9e\x4e\xfd\x9e\x34\xbe\xdc\x7e\x78\x6f\x6f\xd1\x8d\x23\xc4\xc2\xc7\xd6\x36\xb5\xde\x35\x58\x57\xeb\xff\xd6\xae\xa1\xab\x69\xf1\x57\x47\x96\x70\x65\x8d\x6b\xb6\x29\xae\xd9\xae\xfc\x13\x00\x00\xff\xff\x59\xb1\xc1\x65\x7d\x01\x00\x00")

func listTplBytes() ([]byte, error) {
	return bindataRead(
		_listTpl,
		"list.tpl",
	)
}

func listTpl() (*asset, error) {
	bytes, err := listTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "list.tpl", size: 381, mode: os.FileMode(436), modTime: time.Unix(1529871495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _packageTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\xcc\x0a\x39\xec\x02\xb1\xd5\x20\x39\x14\x86\x2c\x20\x1f\x08\x9a\xa0\xcd\x1a\x71\xba\x3d\x14\x3d\xd0\xd2\xd8\x22\x2c\x91\x02\x49\xc5\x31\x5c\xfd\xf7\x82\x1f\xa2\x29\xcb\x59\x04\xe8\xee\x25\x11\x67\xde\xcc\x3c\x3e\x92\x43\x7a\xbf\x87\xb3\x7a\xb3\x86\xe9\x0c\x26\xd0\xb6\xa3\xe4\x53\xce\x33\xb5\xab\x11\x0a\x55\x95\xe9\x28\xe9\xfe\x21\xc9\xd3\x11\x00\x40\xa2\xa8\x2a\x31\xdd\xef\x61\x32\x27\xaa\x80\xb6\x85\x31\x5c\xb3\x5c\xf0\x0a\x73\x92\xc4\xd6\x6d\xa1\x15\x2a\x02\x59\x41\x84\x44\x35\x8b\x1a\xb5\x1a\xff\x1a\x85\x2e\x46\x2a\x9c\x45\xa4\x51\x05\x17\x11\x64\x9c\x29\x64\x6a\x16\x3d\x92\x1d\xbc\x90\x5d\xc9\xc5\x09\x74\x8e\x32\x13\xb4\x56\x94\xb3\x20\xc4\x13\x00\x2a\x81\xc0\x5a\x90\xba\x00\xbe\x02\x55\x20\x20\x53\x54\x20\x6c\x18\xdf\x32\x78\xa5\x92\x2e\x4b\x84\x35\x87\x86\xd1\x57\x14\x12\xa3\x74\x94\xc4\x76\x82\xc9\x92\xe7\xbb\x74\x34\x4a\x8a\x0b\x33\xc3\x67\xac\xf9\x13\xa9\x10\xda\x36\x89\x8b\x0b\xed\xc9\xe9\x6b\x3a\x4a\x08\x14\x02\x57\xb3\x28\x8e\xd2\x38\x89\x49\x3a\xda\xef\x41\x10\xb6\x46\x38\xab\x89\x40\xa6\x8c\x36\x5a\xd5\xb9\x1f\x4a\x23\xb0\x8f\x34\xd2\x7b\x67\x27\x66\x94\x1e\xd9\x7d\x79\x5b\x04\x59\x6e\xd2\xc4\x96\x47\x21\x3a\x4e\x0f\x77\x53\xd0\x94\x1f\xee\x0c\xda\xba\x3f\x8d\xc7\x49\x71\x99\xce\x49\xb6\x21\x6b\x84\x67\xce\x95\x45\xb9\x6a\x49\x5c\x5c\xa6\xe3\x71\x6a\x53\xe8\xd9\x4a\xaa\xb8\xd8\x39\xa4\x27\xab\x43\xfe\x7c\xfe\x5d\xf3\x03\x63\x60\x7c\xc5\xcb\x92\x6f\x0d\x5d\xe7\xd2\x14\x5d\xe1\xfd\x1e\xb6\x54\x15\x70\xb6\xc5\xa5\x76\x4e\x67\x66\x9b\x4d\xfe\xb2\xc3\xb6\xd5\x08\xba\x02\x86\x1e\x62\xfc\xce\x69\xd8\x2c\x78\x23\x32\x84\x8c\xe7\x08\x5b\x5c\x42\x49\xd9\xa6\x4f\xa9\x0b\xed\x54\xf3\xc3\x3e\x13\x27\xd9\xe1\xcb\xe8\x02\xa6\x88\x56\xd7\x2a\xe2\x75\xd6\x66\x30\x9a\x78\xd4\xd7\x2d\x43\x61\x61\xe6\xf3\x08\xa7\xbf\xbe\xdd\x2e\x2c\xe0\xdb\xed\x22\x58\x00\xfd\xf7\x9e\x0a\xa9\x40\x22\x32\x8b\x30\xe3\x05\x22\xbb\x56\x47\xc8\xc3\xe6\x78\xbd\x88\x83\x65\x8a\xd2\xc7\xc5\xd7\xa7\x60\x56\x7d\x74\xa1\x54\x2d\xa7\x71\xbc\xe6\x39\xcf\x26\x5c\xac\xfb\xb1\x77\x3c\x6b\x2a\x64\x8a\xe8\x33\x13\x26\xb1\x9b\x27\xd8\x1f\x0f\x6c\xc5\x45\xe5\x70\xc5\x65\xea\x56\x69\x72\x47\x14\xf1\xeb\xa2\x07\xe6\x60\xd8\xd9\xf8\xe1\xd1\x5c\x8c\xfd\x96\x57\x15\x55\x32\x40\x3a\xcb\x29\xf0\x8d\x20\x2c\x2b\x30\x44\x77\xa6\x53\xf0\x17\xb2\x0e\xa1\x7a\x78\x32\xeb\x4e\xa1\x7c\xe1\x8a\x94\x01\x78\x2e\x50\xa9\xdd\xc1\x75\x2a\xf0\x9e\x8b\x4d\x58\xc0\x8c\x0f\xc0\x4e\x9b\x87\xaa\xe6\x42\x61\x7e\xb3\x33\x0a\x69\x4d\x93\xe2\x2a\xed\xcc\x70\xb3\xd3\x29\x4a\x64\x47\x50\xa8\xad\xe8\x27\x9c\xff\x42\x5d\x36\x82\x94\x10\x45\x10\xc9\x48\xd7\x2c\xae\xd2\x51\xd2\x94\x61\x9f\xa1\x55\x7d\x0e\x67\x02\x57\xd2\xf4\x99\x3e\x0f\xd3\x35\x4b\x9a\xf6\xdb\x0d\xad\x6a\x7f\x58\xec\x37\x7c\x76\xf5\x4d\xa2\xc9\xb3\xce\xd6\xb6\xfa\x8c\xa3\x40\x96\xe1\xd0\x3b\x24\xf7\xc5\xee\xa9\x92\xf6\x1b\x94\x63\xdb\x9d\x39\xdf\x12\x48\x59\x5a\xae\xd2\xb7\x05\xa3\xef\xf5\xc1\xee\xfb\xc3\x5a\xc1\x67\x53\xff\x10\xf4\x05\x7e\x39\x08\xed\xf6\xf1\x55\x7a\x5d\x96\xe0\x00\x49\xec\xc4\x02\xa9\x76\x25\xce\xa2\x92\x4a\x35\x36\xdf\x63\x7d\xb3\x4d\x81\x71\xa6\x9b\x7e\x4f\x4a\xc3\x85\xf4\x28\x7c\x54\xc3\xf7\xa7\xff\xbd\x1e\x14\xea\xf1\x82\x52\xbd\xa7\x49\xe8\x3b\xa9\x4b\x00\x78\x5f\x1b\x0d\xfa\x31\x02\x1d\xf1\xf9\xb9\x22\x05\x32\xc9\x66\x39\xb7\x0f\x15\xca\x72\x7c\x73\x87\x72\xd1\x2c\x5d\xf3\x92\x7a\x43\x7a\x81\x1c\x7c\x12\x30\x3d\x52\xe5\xff\x8b\x31\x2c\xf1\x13\xc4\x70\xf3\x21\x2c\xf7\x12\x74\x85\x9f\x91\xe4\xe6\xce\x0a\xa6\xe6\xba\xb1\x76\x4c\x21\xa9\x85\x79\xa8\x0d\x02\x92\x58\x7b\xbe\x23\xfc\xa0\xfe\x50\xed\x43\xb3\xbb\x4c\xbb\x1e\x76\x02\x04\x4f\x28\x75\x1b\x9c\xf7\xbb\xdd\x00\x39\x6c\x2b\x53\x7f\x05\x39\xdd\xed\x2c\xf4\xb5\x76\x1e\x6e\x87\x30\x8d\xed\xeb\xc3\x77\xd6\x21\xd4\xaf\x44\xcf\x64\x5e\x57\xc9\x52\xbc\xf7\x5e\xe8\xba\xfa\xad\x20\xdb\x12\x7e\xa3\x52\x3f\x92\xec\xd6\x09\xf6\x45\xfe\x76\x0e\x67\x99\x81\x68\x5e\x0e\xe6\x2f\x4e\x37\xfe\xdb\xec\x83\xfc\x0d\xda\xf6\x1f\xb8\x6b\x84\xb9\x6e\xcd\x3d\x63\x63\x27\x9d\xed\xe8\x52\x3a\x15\xbe\x50\x44\xf7\xfc\x30\xfa\x91\x2f\x9d\x75\xf0\xbc\x38\x95\xe1\x9e\x32\x2a\x8b\x41\x8a\xce\xfc\xa1\x1c\x8b\x26\xcb\x10\xf3\x21\x8f\xce\xfe\x81\x1c\x7f\xa0\x94\x7a\x05\x8f\x52\x74\xe6\x43\x86\xe0\xc4\xd8\xf7\x4b\x6c\xdf\xec\x49\x6c\x7f\xaa\xfc\x17\x00\x00\xff\xff\x2a\x9d\x7b\x70\xd2\x0c\x00\x00")

func packageTplBytes() ([]byte, error) {
	return bindataRead(
		_packageTpl,
		"package.tpl",
	)
}

func packageTpl() (*asset, error) {
	bytes, err := packageTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "package.tpl", size: 3282, mode: os.FileMode(436), modTime: time.Unix(1533014847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _subpackageTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x6b\xeb\x38\x14\xdd\xfb\x57\xdc\x31\x6f\x39\xb1\x79\xb4\x8b\x21\x28\x86\xb6\xa1\x4c\x0a\x33\x84\xa6\xbb\x61\x16\x8a\x75\x63\x89\xd8\x92\x91\xe4\xb6\xc6\xe4\xbf\x0f\x92\xec\x44\x71\xc3\xd0\xc5\xdb\xb4\xf1\xd1\x3d\x47\xe7\x7e\xd9\xe4\x37\xa6\x4a\xdb\xb7\x08\xdc\x36\x75\x91\x90\xe9\x1f\x52\x56\x24\x00\x00\xc4\x0a\x5b\x63\x31\x0c\x90\x6d\xa9\xe5\x70\x3a\xc1\x02\x1e\x24\xd3\xaa\x41\x46\x49\x1e\x8e\x43\x68\x83\x96\x42\xc9\xa9\x36\x68\x57\x69\x67\x0f\x8b\x3f\xd2\xf8\x48\xd2\x06\x57\x29\xed\x2c\x57\x3a\x85\x52\x49\x8b\xd2\xae\xd2\x17\xda\xc3\x1b\xed\x6b\xa5\x6f\x44\x33\x34\xa5\x16\xad\x15\x4a\x46\x94\xb3\x01\x10\x06\x28\x54\x9a\xb6\x1c\xd4\x01\x2c\x47\x40\x69\x85\x46\x38\x4a\xf5\x21\xe1\x5d\x18\xb1\xaf\x11\x2a\x05\x9d\x14\xef\xa8\x0d\xa6\x45\x42\xf2\x90\x20\xd9\x2b\xd6\x17\x49\x42\xf8\xcf\x38\x43\x92\xf3\x9f\x0e\x65\xe2\xbd\x48\x08\x05\xae\xf1\xb0\x4a\xf3\xb4\xc8\x49\x4e\x8b\x64\x18\x40\x53\x59\x21\xfc\x68\xa9\x46\x69\x3d\x6b\xb9\x82\x6c\x7b\xac\xb2\xed\x19\x32\x67\xbd\x48\x63\x18\x62\xd6\x14\x90\x16\x33\xfc\x6f\xda\xa0\x37\x12\xae\x43\xc9\xbc\x4c\x1e\x1c\x71\xed\x3d\xdf\x15\xe1\x32\xd8\xd2\xf2\x48\x2b\x24\x39\xbf\x2b\x82\xeb\xcd\x7a\x09\x3e\xa1\x63\x95\x6d\xd6\x5e\x29\x50\xdd\xdf\x1d\xe2\x12\xae\x2c\x8d\xce\x2f\x5e\x62\xc0\x99\x88\xe9\x67\x26\xb7\xb6\x35\xcb\x3c\xaf\x14\x53\x65\xa6\x74\x95\x47\x35\x4c\x8b\xb5\x2a\xbb\x06\xa5\xa5\xae\x77\xb1\x48\x64\xdf\xfb\x86\x8d\x3c\x28\xdd\x8c\x71\x2e\x87\x61\x00\x71\x08\x26\xd6\xd4\x52\x9f\xbb\xe3\xba\x87\xec\x15\x5b\x75\xc9\xee\x0c\xcd\x92\xf4\xf8\x93\x6a\x1a\x61\xcd\x2c\x7a\x44\x6f\x11\x1e\x35\x95\x25\xc7\x39\x63\x82\x6f\x51\xde\x68\x35\x0f\x77\xd0\x4d\xf5\xde\xa2\x79\x53\x96\xd6\x33\xc2\x56\xa3\xb5\xfd\xe5\xf8\x16\xf9\x59\xe9\xe3\xfc\x22\x8f\x5d\x82\xa7\x89\xe5\xf7\xc5\xa6\x69\x95\xb6\x86\xe4\xfc\xbe\x48\x48\x57\x83\xb1\x7d\x8d\xab\xb4\x16\xc6\x2e\xfc\xef\x85\x5b\xfb\x25\x48\x25\xdd\x46\x5c\x66\x5a\x34\xad\x1f\xe6\x5d\xb7\xcf\x46\x15\x57\x7f\xbf\x97\xb5\x28\xae\x47\xd9\x05\x4f\xe3\x1b\x7e\x87\x46\xd7\xe2\x7a\x6e\x3b\xf7\x52\x09\x26\x2f\xf0\xd4\x66\x77\xd3\x2b\x52\xe6\x47\x3e\x8c\x47\xdc\x6d\x77\xb0\x04\xd2\xea\xf0\x16\xba\x8a\x26\xb9\x83\xbf\x28\x87\x11\x8b\xf5\x43\x26\xc8\x1e\xfb\x70\x3e\x55\x08\x19\x3c\xf6\x64\x2c\xd2\xff\x97\x61\x22\xff\x82\x4a\x38\x7b\x51\x1d\xdc\xfd\x4f\x9a\x7e\xd4\xf0\xa7\x30\x56\xe9\x3e\xb4\x2d\x32\xc3\x3e\x7f\x87\x1f\xa5\x0f\x71\xa6\xc6\xb0\xf3\x5e\x8c\xcf\xff\xf8\xcb\xd9\x27\x9c\x4e\xff\xc2\xba\xd3\x7e\xa3\xfc\xc8\x04\x6e\x36\x61\xb3\xf9\xba\x45\xdf\x59\xea\x52\x8e\xd9\x2f\x6a\x3f\xa2\x0f\xf6\x1b\x0a\xcf\x42\x0a\xc3\xbf\x48\x4c\xf0\xb7\x34\x76\x5d\x59\x22\xb2\xaf\x3e\x26\xfc\x1b\x1a\x7f\xa1\x31\xb4\x1a\x77\xfa\x22\x31\xc1\x91\xc2\xbc\x2d\x79\xf8\x3c\x90\x3c\x7c\x15\xff\x0b\x00\x00\xff\xff\x5c\x2e\x02\xd5\x2d\x07\x00\x00")

func subpackageTplBytes() ([]byte, error) {
	return bindataRead(
		_subpackageTpl,
		"subpackage.tpl",
	)
}

func subpackageTpl() (*asset, error) {
	bytes, err := subpackageTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "subpackage.tpl", size: 1837, mode: os.FileMode(436), modTime: time.Unix(1531357019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"index.tpl": indexTpl,
	"list.tpl": listTpl,
	"package.tpl": packageTpl,
	"subpackage.tpl": subpackageTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.tpl": &bintree{indexTpl, map[string]*bintree{}},
	"list.tpl": &bintree{listTpl, map[string]*bintree{}},
	"package.tpl": &bintree{packageTpl, map[string]*bintree{}},
	"subpackage.tpl": &bintree{subpackageTpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

