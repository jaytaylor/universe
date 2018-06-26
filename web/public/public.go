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

var _indexTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\xd1\x6a\xe3\x3a\x10\x86\xef\xf3\x14\x73\x4c\x6f\x6b\xd3\xf6\xe6\x50\x64\x43\x36\x85\x85\xa5\xbb\x14\xda\x7d\x80\x89\x35\xb1\x44\x64\xc9\x3b\x1a\x27\x1b\x8c\xdf\x7d\x91\xed\xa4\x69\x59\x58\xdf\x58\xcc\x7c\xfa\x47\xff\x8c\xa4\xfe\xd3\xa1\x96\x53\x47\x60\xa4\x75\xd5\x4a\x9d\x7f\x84\xba\x5a\x01\x00\x28\xb1\xe2\xa8\x5a\x7b\xcd\xa1\x25\x8d\xaa\x98\x03\x73\xb2\x25\x41\xa8\x0d\x72\x24\x29\xb3\x5e\x76\xb7\xff\x67\xd7\x29\x8f\x2d\x95\x19\xf6\x62\x02\x67\x50\x07\x2f\xe4\xa5\xcc\xbe\xe1\x09\xde\xf0\xe4\x02\xff\x85\xd6\x14\x6b\xb6\x9d\xd8\xe0\xaf\xb6\x5c\x0e\x00\x36\x02\x42\xc3\xd8\x19\x08\x3b\x10\x43\x40\x5e\x2c\x13\xec\x7d\x38\x7a\x38\xd8\x68\xb7\x8e\xa0\x09\xd0\x7b\x7b\x20\x8e\x94\x55\x2b\x55\xcc\x96\xd4\x36\xe8\x53\x32\x78\x77\x6d\xc9\xdc\xa5\xd0\x7d\xf5\x4a\xc8\xb5\xb9\xd6\x3c\xab\x7d\x0d\x0e\x7d\x03\x3f\x17\x45\x55\x98\xfb\x24\xc6\xd5\x6a\x18\xe0\x68\xc5\xc0\x8d\x43\xa1\x28\xf0\x58\x42\xbe\x09\x7e\x67\x9b\xfc\x3b\x46\x21\xce\x9f\xe7\xc4\x38\x26\xd6\xee\x2e\xe4\x38\xae\x94\xb6\x87\xa5\x03\xe6\xa1\x5a\xc0\x0d\xe3\xd1\x91\x86\x17\xac\xf7\xd8\x50\x54\x85\x79\x58\xa0\xde\x41\x94\x93\xa3\x32\x73\x36\xca\xed\xb4\xbe\x4d\xf3\x7b\x04\x1f\x7c\x32\x3a\x0c\xc0\xe8\x1b\x82\x9b\x6e\xdf\xa4\xc3\x5c\x55\x83\xe5\x53\xce\x56\xc3\x30\x11\xf9\x13\x0a\xe6\x1b\x26\x14\xd2\xeb\x44\x81\x42\x30\x4c\xbb\x32\x2b\xce\xcc\x0b\x8a\x81\x71\xcc\xaa\x4f\x01\x55\x60\xa5\x0a\x67\xa7\xaa\xe4\xf5\xb9\x86\x2a\xfa\x74\x89\x8a\xc9\xdc\x7b\xea\x7d\xf5\xa3\x6f\xb7\xc4\x69\x7c\xdd\xe2\x11\xac\x07\xeb\x35\xfd\x7e\x84\x61\x80\xfc\xe9\x4b\x7e\x76\xff\x4c\x7e\xea\x54\xea\xf5\xd4\x19\xf8\xd5\x53\x4f\xe0\xc8\x37\x62\x2e\xf8\x5b\x98\x92\x1f\xf0\x75\x2d\xf6\x40\x70\x0c\xbc\x27\x8e\x13\x3a\xd5\xf8\x3c\xa0\x57\x41\x89\x90\x31\xb5\x41\x28\x66\x1f\xcb\x45\x88\xd6\xd7\x04\x1d\x87\x86\xb1\x85\x28\xc8\x42\xfa\x9f\x6a\xf5\xb4\x79\x16\x2b\x96\x5b\x57\xcc\xcf\xeb\x4f\x00\x00\x00\xff\xff\x0f\x9f\x77\x2f\x76\x03\x00\x00")

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

	info := bindataFileInfo{name: "index.tpl", size: 886, mode: os.FileMode(436), modTime: time.Unix(1529878619, 0)}
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

var _packageTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xcd\x6e\xdb\x38\x10\xbe\xeb\x29\xa6\x42\x8e\xb5\x85\x20\x3d\x2c\x0c\x5a\x40\x13\x23\xd8\x04\xdd\x34\xb0\xdb\x5e\x16\x7b\xa0\xc5\xb1\x49\x44\x22\x05\x92\x8a\x23\x18\x7e\xf7\x05\x49\x49\xa6\x6c\x1d\x7a\xe8\xc5\x26\x87\xdf\xcc\x7c\xf3\x6b\x93\x4f\x4c\x15\xb6\xad\x11\xb8\xad\xca\x3c\x21\xfd\x17\x52\x96\x27\x00\x00\xc4\x0a\x5b\x62\x7e\x3c\xc2\xfc\x95\x5a\x0e\xa7\x13\xcc\xe0\xab\x64\x5a\x55\xc8\x28\xc9\xc2\x73\x80\x56\x68\x29\x14\x9c\x6a\x83\x76\x99\x36\x76\x37\xfb\x2b\x8d\x9f\x24\xad\x70\x99\xd2\xc6\x72\xa5\x53\x28\x94\xb4\x28\xed\x32\x7d\xa6\x2d\xfc\xa0\x6d\xa9\xf4\x04\x9a\xa1\x29\xb4\xa8\xad\x50\x32\x52\x19\x08\x80\x30\x40\x61\xaf\x69\xcd\x41\xed\xc0\x72\x04\x94\x56\x68\x84\x37\xa9\x0e\x12\xde\x85\x11\xdb\x12\x61\xaf\xa0\x91\xe2\x1d\xb5\xc1\x34\x4f\x48\x16\x02\x24\x5b\xc5\xda\x3c\x49\x08\xbf\xf5\x11\xae\xb1\x56\x2f\xb4\x42\x38\x9d\x48\xc6\x6f\xdd\x0b\x13\xef\x79\x42\x28\x70\x8d\xbb\x65\x9a\xa5\x79\x46\x32\x9a\x27\xc7\x23\x68\x2a\xf7\x08\x37\x35\xd5\x28\xad\xcf\xcd\x62\xe9\x92\xd4\x5f\x0d\x9c\x4e\x91\xe6\xf1\x18\x63\xfb\x64\xa6\xf9\x85\x7c\x70\x1f\x9c\xa0\x64\xde\x4c\x16\x78\x70\xdd\x73\x7a\x5a\x2d\xc0\x51\x7e\x5a\x79\x74\x78\xfe\x34\x9b\x11\x7e\x97\xbf\xd2\xe2\x8d\xee\x11\xd6\x4a\xd9\x80\xea\xbc\x91\x8c\xdf\xe5\xb3\x59\x1e\x4c\xfc\x5c\x7f\x5b\xc0\x40\xd0\xc1\x7e\xae\xbf\x39\x4e\xe0\x05\x52\xed\x54\x59\xaa\x83\xa7\xd8\x3d\x39\x5a\x91\x33\xf0\x76\x1c\xe5\xe0\x66\x20\xef\xc4\xe0\x1d\x0d\xa8\xef\x07\x89\x3a\xc0\xfc\xf1\x02\xe7\x4e\xbf\x1e\x36\x01\xf0\xeb\x61\x13\x45\xe5\x3e\x1f\x85\x36\x16\x0c\xa2\x0c\x08\x7f\xdf\x20\xca\xaf\xf6\x02\x79\xce\xf8\xfb\x6d\x16\xc5\x9e\xe6\xcf\x9b\xef\x2f\x71\x00\x23\x34\xb7\xb6\x36\x8b\x2c\xdb\x2b\xa6\x8a\xb9\xd2\xfb\xb1\xee\x4a\x15\x4d\x85\xd2\x52\xd7\x88\xb1\x91\x50\x91\x28\xe9\x4f\x72\xa7\x74\xd5\xe1\xf8\x9d\xaf\xa2\xd8\xc1\x7c\x45\x2d\xf5\xa5\x74\x7a\xee\xe2\xbb\x2d\x44\x33\x5c\x2f\x62\xf1\xf2\x07\x55\x55\xc2\x9a\x08\xd9\x49\xa6\xc0\xf7\x9a\xca\x82\x63\x8c\xee\x45\x53\xf0\x1f\x74\x1f\x43\xdd\x75\xd2\x6a\x6b\xc7\x26\xdd\x7d\x0a\xf8\xa8\xf4\x5b\x0c\xf4\xf7\x33\xd0\xe5\xe2\x20\x2c\x87\x1b\xd3\x6c\x5f\xdf\xf6\x6e\x62\x84\x64\xf8\xd1\xc1\x37\xcd\xb6\x4b\xa3\x81\x34\x75\xd9\x0a\xc9\xeb\xe0\xf3\xa7\xaa\x56\xda\x9a\x21\x8d\x09\xe1\x5f\xf2\x4e\x48\x32\xfe\x25\x4f\x48\x53\x82\xb1\x6d\x89\xcb\xb4\x14\xc6\xce\xfc\x79\xe6\x36\xdc\x02\xa4\x92\x6e\xf8\xcf\xa3\x2b\xaa\xda\x31\x98\xb0\xee\xb7\x50\x29\xf2\xf1\xf8\x3a\x7c\x3f\xb2\xe1\x1c\x3a\xa1\x14\xe3\x59\x6d\xdc\x0a\x0d\x11\x9f\xc5\x5d\x28\x54\xb2\x21\xfa\xde\xf1\x1a\x29\xf3\x83\x13\xda\x29\xee\x10\xf7\xb0\x00\x52\x6b\xbf\x82\xaf\x14\x48\xe6\x5e\xae\x7c\x4d\x9d\x42\xa7\x46\xcd\x38\xca\xb6\x07\xdc\x39\x1f\x25\xca\xc9\x77\x78\x41\x63\x91\x41\x2f\x5a\x0c\xdd\xdd\x65\x33\x70\x73\x13\xf3\x39\xae\x6f\x6c\xe6\x55\xa3\xb5\xed\xf5\x5e\x3c\xab\x0e\xf9\x1d\x89\xfc\x36\x24\x5b\x3d\x1d\x63\x1f\x54\xa8\x1f\xb2\xfb\xe0\x61\x68\x0d\x64\x70\xdf\x92\xae\x3b\x26\xea\x7f\xa1\xf8\x07\x6a\xcf\x47\x4c\x3d\x95\x07\x4d\x0f\x25\xfc\x2d\x8c\x55\xba\x0d\xbd\x1a\x11\x61\x1f\x9f\xe1\xa6\xf0\x10\x47\xa8\x83\x0d\x7d\xde\xdd\xff\xf5\xce\xd9\x07\x9c\x4e\xff\xc1\xaa\xd1\x7e\xc9\xf8\x69\x0b\xba\xf3\x5e\x76\x31\x9a\x53\xea\x1b\x4b\x5d\xc8\xb1\xf6\xb3\xda\x76\xd2\xab\xa5\x3a\x65\xe1\x51\x48\x61\xf8\x95\x89\x5e\xfc\x5b\x36\x36\x4d\x51\x20\xb2\x6b\x1e\xbd\xfc\x37\x6c\xfc\x83\xc6\xf8\x86\x1c\x9b\xe8\xc5\x91\x85\xcb\xb2\x64\xe1\xe7\x9f\x64\xe1\x5f\xcf\xff\x01\x00\x00\xff\xff\x64\xf6\xc7\xf3\x0d\x09\x00\x00")

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

	info := bindataFileInfo{name: "package.tpl", size: 2317, mode: os.FileMode(436), modTime: time.Unix(1529888211, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _subpackageTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x41\x6b\xe3\x3c\x10\xbd\xfb\x57\xcc\x67\x7a\xfc\x62\x53\xda\xc3\x12\x14\x43\xdb\x50\x36\x85\x5d\x42\xd3\xdb\xb2\x07\xc5\x9e\x58\x22\xb6\x64\x24\xb9\xad\x31\xf9\xef\x8b\x24\x3b\x51\x9c\xb0\xf4\xb0\x97\x36\x7e\x7a\xef\xe9\xcd\x68\x64\x93\xff\x0a\x99\x9b\xae\x41\x60\xa6\xae\xb2\x88\x8c\xff\x90\x16\x59\x04\x00\x40\x0c\x37\x15\x66\x7d\x0f\xc9\x9a\x1a\x06\x87\x03\xcc\xe0\x41\x14\x4a\xd6\x58\x50\x92\xfa\x65\x4f\xad\xd1\x50\xc8\x19\x55\x1a\xcd\x22\x6e\xcd\x6e\xf6\x2d\x0e\x97\x04\xad\x71\x11\xd3\xd6\x30\xa9\x62\xc8\xa5\x30\x28\xcc\x22\x7e\xa1\x1d\xbc\xd1\xae\x92\xea\x0a\xbb\x40\x9d\x2b\xde\x18\x2e\x45\x20\x39\x06\x00\xae\x81\x42\xa9\x68\xc3\x40\xee\xc0\x30\x04\x14\x86\x2b\x84\xbd\x90\x1f\x02\xde\xb9\xe6\xdb\x0a\xa1\x94\xd0\x0a\xfe\x8e\x4a\x63\x9c\x45\x24\xf5\x05\x92\xad\x2c\xba\x2c\x8a\x08\xbb\x0d\x2b\x24\x29\xbb\xb5\x68\xc1\xdf\xb3\x88\x50\x60\x0a\x77\x8b\x38\x8d\xb3\x94\xa4\x34\x8b\xfa\x1e\x14\x15\x25\xc2\x4d\x43\x15\x0a\xe3\x54\xf3\x05\x24\xeb\x7d\x99\xac\x8f\x90\x3e\xfa\x05\x1e\x7d\x1f\xaa\x46\x42\x9c\x4d\xf0\x9f\xb4\x46\x17\xc4\x6f\x87\xa2\x70\x36\xa9\x4f\xc4\x94\xcb\x7c\x97\xf9\xcd\x60\x4d\xf3\x3d\x2d\x91\xa4\xec\x2e\xf3\xa9\x57\xcb\x39\xb8\x82\xf6\x65\xb2\x5a\x3a\x27\x2f\xb5\x7f\x37\x88\x73\x38\x8b\x34\x24\x3f\x65\x09\x01\x1b\x22\x94\x1f\x95\xcc\x98\x46\xcf\xd3\xb4\x94\x85\xcc\x13\xa9\xca\x34\xe8\x61\x9c\x2d\x65\xde\xd6\x28\x0c\xb5\x67\x17\x9a\x04\xf1\x5d\x6e\x58\x89\x9d\x54\xf5\xc0\xb3\x35\xf4\x3d\xf0\x9d\x0f\xb1\xa4\x86\xba\xda\xad\xd6\x3e\x24\xaf\xd8\xc8\x53\x75\x47\x68\x52\xa4\xc3\x9f\x64\x5d\x73\xa3\x27\xec\x01\xbd\x26\x78\x54\x54\xe4\x0c\xa7\x8a\x11\xbe\x26\x79\xa3\xe5\x94\x6e\xa1\xab\xee\x9d\xb9\xb4\xb6\xd8\x35\xf2\xb3\x54\xfb\x29\xd9\x61\x27\xf2\x38\xa1\xec\x3e\x5b\xd5\x8d\x54\x46\x93\x94\xdd\x67\x11\x69\x2b\xd0\xa6\xab\x70\x11\x57\x5c\x9b\x99\xfb\x3d\xb3\xd7\x7c\x0e\x42\x0a\x7b\x03\x4e\x33\xcc\xeb\xc6\x0d\xef\xa6\xdd\x26\x83\x8b\xed\xb7\xbb\x87\x15\xcf\xce\x47\xd7\x92\xc7\x71\xf5\xbf\xfd\xc1\x56\xfc\x7c\x4e\x5b\xfb\x12\xf1\x21\x4f\xf0\x78\xac\x76\xa7\x57\xa4\x85\x1b\x71\x3f\x0e\xe1\xe9\xda\x85\x39\x90\x46\xf9\xb7\xce\x19\x9b\xa4\x16\xbe\x70\xf6\x23\x15\xfa\xfb\x4a\xb0\x78\xec\xfc\xfa\xd8\x21\x2c\xe0\xb1\x23\x43\x93\xfe\xde\x86\x51\xfc\x0f\x3a\x61\xe3\x05\x7d\xb0\xfb\x3f\x29\xfa\x51\xc1\x77\xae\x8d\x54\x9d\x3f\xb6\x20\x4c\xf1\xf9\x3f\xdc\xe4\x8e\x62\x43\x0d\xb4\xe3\x3d\x18\x9e\x7f\xb9\xcd\x8b\x4f\x38\x1c\x7e\xc3\xb2\x55\xee\x06\xb9\x91\xf1\xda\x64\xc4\x26\xf3\x75\x4d\xbe\x31\xd4\x96\x1c\xaa\x5f\xe4\x76\x40\x1f\xcc\x17\x1c\x9e\xb9\xe0\x9a\x5d\x58\x8c\xf0\x97\x3c\x36\x6d\x9e\x23\x16\x97\x39\x46\xfc\x0b\x1e\x3f\x50\x6b\x5a\x0e\x17\xed\x64\x31\xc2\x81\xc3\xf4\x58\x52\xff\x39\x20\xa9\xff\x0a\xfe\x09\x00\x00\xff\xff\x74\x74\xf6\xf2\x1d\x07\x00\x00")

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

	info := bindataFileInfo{name: "subpackage.tpl", size: 1821, mode: os.FileMode(436), modTime: time.Unix(1529888162, 0)}
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

