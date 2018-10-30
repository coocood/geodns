// Code generated by go-bindata.
// sources:
// status.html
// DO NOT EDIT!

// +build !devel

package main

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

var _statusHTML = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xc1\x72\xdb\x36\x10\x3d\x4b\x5f\x81\xe1\x34\xb7\x1a\x0c\xd3\x68\xa6\x4e\x21\x5e\xe2\x4c\x7a\xa8\x52\x27\x6a\x7b\xe8\x0d\x24\x56\x26\xa6\x24\xa0\x01\xd6\x76\x5c\x8e\xfe\xbd\x03\x90\xa0\x28\x92\x56\x65\x66\x72\x11\x89\x07\xbc\x05\xde\xbe\x15\x00\xb2\x02\xab\x32\x65\x05\x70\x91\x32\x94\x58\x42\xfa\x11\xf4\xcd\xa7\x2d\xa9\x6b\x42\xff\x02\x63\xa5\x56\xe4\x70\x60\x71\xd3\xb9\x64\xa5\x54\xff\x90\xc2\xc0\x6e\x1d\xc5\xb1\x02\x14\x8a\xd3\x4c\x6b\xb4\x68\xf8\x3e\x17\x8a\xe6\xba\x8a\xf1\x51\x22\x82\xb9\xea\x3a\xe2\x37\xf4\x27\x9a\xc4\xb9\xb5\x71\x87\x5d\xe5\xba\xca\xa4\x02\x41\x2b\xa9\x68\x6e\x6d\x44\x0c\x94\xeb\xc8\xe2\x53\x09\xb6\x00\xc0\xe8\xd2\xf9\x3c\xf0\xc8\x31\x2f\xc2\x44\x60\xee\x4b\xe0\xea\x38\xdb\xd9\x49\x7c\x2b\x5d\x2e\x50\xd0\x7f\xb5\x02\xc5\x2b\xf8\x11\x05\x75\x69\x01\x43\x6a\xb2\xd3\x0a\xaf\x1e\x41\xde\x15\xf8\x8e\x64\xba\x14\xbf\x90\xc3\x92\xc5\x2d\x8d\x65\x5a\x3c\xa5\xcb\x25\x13\xf2\x81\xe4\x25\xb7\x76\x1d\xe5\x5a\x21\x97\x0a\x4c\xe4\x3a\x8a\x24\xfd\x58\xea\x8c\x97\x2c\x2e\x92\xc1\x48\xa3\x1f\xdd\x98\x45\x1f\xb3\x7b\xae\x7e\xf6\xe8\x82\x21\xcf\x4a\x08\x1d\x4d\xc3\xff\x5e\x65\xda\x08\x30\x20\xda\x66\xae\x95\x00\x65\x41\x44\xe9\x72\xe1\x78\xc6\x3f\x17\x0c\x45\x60\x4b\xb5\xd3\xa4\x11\x15\xb5\x0b\x22\x16\x39\x5a\x16\xa3\xe8\x46\xa7\x09\xa9\xa4\x3a\x85\x56\x63\x28\x99\xc0\x36\xc0\x8f\x08\x8b\xdd\x12\x86\x6b\x49\x3f\xdf\x83\x91\x30\x98\xb2\xae\xf7\x46\x2a\xdc\x91\xe8\x15\x7d\xb3\x8b\x08\x6d\x56\x47\xdb\xc1\xf4\x0b\x47\x48\x7c\x19\xbe\x90\xb4\x9a\x43\x4a\x66\xb1\x9c\xf8\x1e\xaf\x91\xef\x9f\xce\x9f\xef\x6b\xe6\xaf\xd2\xa2\xbe\x33\xbc\x1a\xe4\x75\x73\xd6\x21\x8f\x5c\xbf\x7e\x35\x00\xae\x47\x00\x1d\x42\x1b\xfe\xf5\x14\xd8\xa2\xb8\x81\x87\xff\xf3\x7e\x2b\x55\x0e\xae\xe4\x0c\x3e\x9b\x5f\x11\x11\xd2\xe5\xb7\xd3\x45\x37\x52\x5d\x6a\x4a\x8f\x74\x6a\xc9\x85\xac\xdb\x1c\xaf\x5f\x93\x79\xbc\xeb\xb9\xbc\xb3\xc4\xe7\x72\xc2\xbf\xbe\x7c\xb6\xc6\x29\x32\xaa\xd4\x91\x59\x5f\x20\x07\xf5\x52\x9f\x1a\xd2\x2c\xb7\x02\x75\x8e\x67\x2d\x77\xa6\x73\x3d\xf6\x0c\xff\xfa\xec\x19\x2e\x06\xd9\x33\xbc\x6c\xa9\x67\x1c\xed\x6d\x3e\x2c\x16\xf2\x61\xea\xa0\x79\xfb\x4d\x07\x4d\x28\x99\xde\xce\x14\x36\xa5\x3f\xf7\x28\x2b\x38\x2e\xca\x4b\xa2\x0d\x4a\x6f\xf8\xd3\x16\x8d\x54\x77\xc7\x65\x77\x3b\xe6\xf3\x21\x6f\x4b\x8e\x3b\x6d\xaa\x61\xd0\x80\x4f\x04\x1b\x67\xa0\x7b\x16\x49\xfa\xb7\x56\xee\x2c\x72\xa7\xf2\xac\x04\xf8\xb5\x32\x2c\x52\x16\x63\xd1\xbc\xf9\xe3\xf3\x1e\xe1\x88\xac\x5a\xc4\xf6\x06\x4d\x60\xae\xf2\xc9\xe7\xdb\x6d\x8b\x34\x1e\xd6\x35\xf9\x01\xf5\xfe\xf7\x3d\xba\x9b\xd8\xbb\x35\xa1\x7f\x84\xd6\xe1\xe0\xba\x0d\x57\x77\x40\x08\xf5\x4a\x88\xc3\x5c\x02\xfb\xa7\x44\xeb\x93\x20\xb9\x2e\x9d\xe1\xeb\x55\xe8\x0d\x97\x9d\xc8\xe5\xf0\x13\xaf\xa0\xcb\x5f\xa8\x20\x16\x2c\x19\x1d\xdc\x93\x15\xba\x01\x34\x32\xb7\xcf\x9d\xdb\x17\x73\x56\x33\x38\xc9\x1c\xd2\xc9\x6e\x33\x21\xfa\x83\xbb\x0b\xbf\x44\xf9\x07\xa1\xec\x1c\xf5\x43\xde\xa5\x62\x46\xf3\xcd\x25\x4e\x66\xc2\x55\x5f\xae\xef\x15\x5a\x5f\x7a\x81\xfb\x1b\xcf\xa0\xdc\xba\x9b\xa3\xab\xc6\xf7\xcd\x80\x5e\x99\xfa\x2a\xac\x6b\x22\x77\x1d\xfd\x70\x68\xfe\xd9\xa3\xc2\x3c\xad\xcc\xd4\x87\xb6\xc3\x7f\x71\x28\xf3\x93\x68\x83\x6b\x51\x1b\xe2\xad\xab\x65\x1f\xa5\x55\xd3\x6e\x11\x7e\x95\x13\xf7\xb3\xba\x06\x25\x5c\xbc\xf0\xd2\xaa\x2e\x25\x0c\x65\xbf\xf7\xd8\xa5\xba\xdb\x00\x97\x0a\x6f\x82\x9f\x51\xde\x8f\xf7\xbd\xa4\x2f\xbb\xb7\xe3\xbe\x79\xdc\x36\x9b\xef\x1c\x16\xfb\x8f\xc6\xe5\x7f\x01\x00\x00\xff\xff\x91\x22\x37\x07\x3c\x0e\x00\x00")

func statusHTMLBytes() ([]byte, error) {
	return bindataRead(
		_statusHTML,
		"status.html",
	)
}

func statusHTML() (*asset, error) {
	bytes, err := statusHTMLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "status.html", size: 3644, mode: os.FileMode(420), modTime: time.Unix(1414238098, 0)}
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
	"status.html": statusHTML,
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
	"status.html": &bintree{statusHTML, map[string]*bintree{}},
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
