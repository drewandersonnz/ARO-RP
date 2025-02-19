// Code generated for package rbac by go-bindata DO NOT EDIT. (@generated)
// sources:
// staticresources/clusterrole.yaml
// staticresources/clusterrolebinding.yaml
package rbac

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clusterroleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x1a\x4b\xcf\xe3\xb8\xed\x9e\x5f\x61\x6c\x0f\x0b\x14\xf8\x32\x28\x7a\x29\xbe\x1e\x77\x8b\xa2\x40\xd1\x05\x06\xd3\xde\x19\x99\xb1\xb9\x91\x45\x8d\x44\xe5\x9b\xf4\xd7\x17\x92\x25\xc7\xce\xcb\x8e\xd3\x99\x53\x6c\x8a\xe2\x4b\x14\x5f\xce\x1f\xaa\x5f\xb8\xc6\xaa\x41\x83\x0e\x04\xeb\x6a\x77\xaa\x5a\x50\x87\x4f\x0d\x9a\x9a\xbc\xe2\x23\xba\x93\x02\xd5\xe2\x5f\xab\x5f\x7f\xab\xfe\xf5\xdb\x97\xea\x6f\xbf\xfe\xe3\xcb\x76\x03\x96\xfe\x83\xce\x13\x9b\xf7\xca\xed\x40\x6d\x21\x48\xcb\x8e\xfe\x0b\x42\x6c\xb6\x87\xbf\xf8\x2d\xf1\xa7\xe3\x9f\x36\x07\x32\xf5\x7b\xf5\x8b\x0e\x5e\xd0\x7d\x66\x8d\x9b\x0e\x05\x6a\x10\x78\xdf\x54\x95\x72\x98\x36\x7c\xa1\x0e\xbd\x40\x67\xdf\x2b\x13\xb4\xde\x54\x95\x81\x0e\xdf\x2b\x7f\xf2\x82\xdd\x3b\x38\x7e\xf3\x0e\x37\x2e\x68\xf4\xef\x9b\xb7\x0a\x2c\xfd\xdd\x71\xb0\x3e\x12\x79\xab\x7e\xfa\x69\x53\x55\x0e\x3d\x07\xa7\x30\xc3\x14\x77\x96\x0d\x1a\xf1\x02\x12\x3c\xfa\x4d\x55\x1d\xd1\xed\xf2\x72\x83\x92\x7e\x35\x79\x59\x4a\xd0\xec\xa9\xe9\xc0\xfa\xf4\x8a\xa6\xb6\x4c\x46\xf2\xdb\x11\xcb\xa3\xa6\x8e\xc4\x81\x69\xb0\x7f\x8f\x9a\x78\x0b\xaa\xbc\x72\x9d\x9f\x6c\x34\xa0\x17\x34\x72\x64\x1d\x3a\x54\x1a\xa8\xbb\xbd\x94\xa1\x5c\x0f\x0f\x82\x9d\xd5\x20\x79\xc5\xa1\xd5\xa4\x92\x29\x15\x1b\x71\xac\x35\xba\xb2\xd4\x6b\xf1\x35\xb0\x40\x0f\xf2\xe8\x8e\xa4\x10\x94\xe2\x50\xa4\xce\xb0\x47\x56\x8a\x0f\x1f\x20\xaa\x5d\x66\xaf\x28\xed\x27\xcd\xcd\x35\xc5\xab\xed\x50\x77\xe4\xa3\x33\x39\x6c\xc8\x8b\x1b\x3b\xd1\x35\xe1\x2e\x08\x08\x99\xe6\x03\x77\x2d\xf3\xa1\x3f\x97\xd0\x6f\xea\x95\x39\x82\xa6\xfa\x21\xce\x0a\x1d\xc1\x12\x7e\x13\x34\x51\x4e\x7f\x57\x38\x15\xbc\x70\x57\x80\x35\xee\xc9\xd0\x6b\x4c\x17\xd9\x04\x2c\xbd\x76\x82\x99\x00\xba\x2d\x5b\x34\xbe\xa5\xbd\xdc\x63\xe4\xf0\x6b\x40\x2f\x83\xf3\xac\xe2\x96\x6e\xd1\xf5\x0d\xcb\xae\xeb\xf0\x48\x7e\x38\xce\x1a\xb0\x63\xe3\x31\xbb\x6a\x8d\x56\xf3\xa9\x1b\x2e\x5c\x76\xfe\x61\x3d\x5e\x78\xdc\x07\x9d\x01\x2b\xc5\x9b\xb1\xc3\x59\x88\xde\xb7\x7e\x20\xa7\xa5\x97\xca\xf1\x0c\x65\xd5\x47\xe5\xb5\xa2\x07\x69\xd1\x48\x0e\x3b\x77\x3d\x53\xf8\x80\x26\x9e\x27\x7e\x5c\x30\x4a\xc1\x1f\x6f\x13\xbe\x4c\x25\xd7\x74\x3d\xea\xbd\x0f\xbb\xdf\x51\x09\x28\x85\xde\x9f\x79\x4c\x16\x53\xce\x98\xac\xdd\xde\xf4\xb4\x60\x8b\x6c\xeb\x58\xe3\x8e\x4c\x4d\xa6\xf1\x97\xf0\xec\xbd\x97\x18\x65\x69\x71\xb2\x7a\x46\xac\xf2\x7a\xc3\x64\x3f\xc4\x2c\x23\x6d\x1d\x7a\x71\xa4\x5e\x09\x8e\x41\xd8\x2b\xd0\x64\x9a\x6b\x4e\x49\x24\x36\x02\xda\x72\x5d\x30\x5f\x71\xf6\xc2\x6a\xd9\xc1\x4f\x39\xbe\x55\x1d\xa8\x96\x0c\xbe\x2c\xc8\x2e\x81\xaf\xb9\x3a\x36\xbf\xf3\xae\xe7\x95\x1f\xd6\x50\x0f\xa4\xeb\x19\x05\x13\xce\x39\xe8\x65\xc0\xf7\x66\xb8\x34\xea\x29\x74\x42\xfb\x18\x94\xf0\x41\x92\x1e\x21\x51\x63\x92\x33\xa6\x9c\xb6\x52\x0b\xa5\x39\xd4\xca\x61\x1d\xe3\x21\xe8\x39\x0f\x19\x10\xfd\x8b\x6c\xd3\x21\xcc\x27\xec\x3e\xb1\xfb\x1b\x51\xfb\xe2\x00\x07\xef\x65\x1b\xfb\x01\x76\x13\xe0\xb1\x2f\xf8\x7d\xe1\xed\x87\x18\x56\x1b\x9f\x9f\xf6\x08\x12\x1c\x36\x43\x65\x4a\x1d\x94\x42\x98\xcc\xde\x81\x17\x17\x54\x44\x29\xb0\x18\x07\xca\x6e\x83\xf2\xc1\xee\xd0\xbf\x70\x14\x35\x3f\x66\x71\xda\x90\x3d\xdc\x3a\x8e\xc1\x6a\x78\xf9\x46\x99\x82\x57\x2d\xd6\x61\xfd\xf5\xca\x6a\xcd\x9d\x60\x8f\xa5\x34\xd5\xfc\x61\x34\x43\x3d\x31\x4a\x2c\x13\x9d\x01\xad\xb9\xd1\x64\x0e\x93\xb5\x2b\x80\xe1\xec\x8a\x97\xa6\xb5\x3a\x34\x34\x05\x7d\x0d\xa4\x0e\x5e\xc0\xc9\x04\x7c\x82\x4e\x7b\xe8\xec\xe3\xbc\xf1\x58\xeb\x58\x78\x59\x0d\x26\xa9\x9e\x8c\x3d\x63\x03\xcb\x75\x3e\x2e\xc5\xc6\xa0\x12\x3a\x92\x9c\x54\x8b\xea\xb0\x5a\x0a\x76\x35\x99\xc7\x69\x5f\x23\x3c\xee\xe5\x1e\x30\x18\x3a\xda\xbb\xd4\x87\xa6\x4e\xaf\xaf\xa4\xfb\x56\xf0\x3e\x8b\xd2\x29\xae\x21\x3d\xf4\x1f\xd7\x74\xc7\x37\x69\x05\xe9\xbd\xe6\x8f\xec\x06\xdb\x73\x2f\x70\x4f\x89\x88\x1d\xaf\x5a\x07\xe5\x0a\x12\x3b\x92\x93\xc6\x23\xea\xff\x47\xb3\xd5\xa2\xee\x66\x1c\x30\xa2\xa8\x16\x9c\x38\xb4\xec\x49\xd8\xd1\x5a\xe5\x53\x90\x9a\x61\x37\x0e\x64\xe9\x51\x1c\x42\xf7\xdd\x19\x26\x2e\x03\xef\xb9\xec\xf7\x24\x5d\x81\x66\xa4\x51\x7e\x5b\x5c\x77\xa6\x4d\xb9\x3b\x3d\x2d\x0d\x1b\xe3\xc2\x21\x11\xb0\x2e\x98\xd5\xe1\x3a\x3b\xfd\x52\xe6\xb5\xf1\x0e\x15\xbb\xb5\xd5\x4a\xbc\x0e\xca\xd0\x56\x19\xb5\xbf\xc9\x20\x47\xc4\x37\x10\x01\xd5\xc6\x9e\xed\xed\xe5\x11\x40\xae\x1c\x67\x54\xcb\x58\x2d\x82\x96\x76\x88\xc2\x03\x7c\xfa\xb2\xba\x31\xce\x04\x26\x37\x7c\xfe\xc0\x05\xc8\xa0\x73\xc1\x08\x75\x38\x76\x80\x73\xc3\x3f\x86\x1e\xc2\x0e\x35\xca\x18\x34\xe1\x6b\x99\xf5\x0d\xf0\x5a\x95\x50\x40\xff\xf9\x76\xe1\x09\x0e\xd3\x72\xcb\xfe\x5c\x6b\xf4\x93\x89\xdc\xae\xad\x63\xe8\x48\xdd\x4f\x10\xa3\x01\x21\x3f\x74\xd4\x6b\xca\xd4\xcc\x8c\x89\xbc\xb0\x4b\xb7\x7f\xa8\xcf\x32\x24\xd7\x76\x03\x85\xb5\xba\xb1\x49\x81\xd8\x34\x5b\xc5\x0e\xd9\x6f\x15\x77\x37\x6a\x52\x8d\x4e\x3a\x30\x31\x7e\x8c\x8f\x79\x0c\x1f\x4c\x90\x69\x0e\xf6\xdf\xe1\xf0\xd8\xa1\xb4\x18\xfc\x15\x20\xf5\xfc\xbd\x7a\xfd\x50\x6c\x42\x43\x5a\x30\x9c\x70\xd6\x86\x9d\x7c\xcd\x97\x75\x82\x93\xa2\x16\x53\xbc\xca\x20\xcb\x9a\x54\xa9\x5c\x93\x8f\x85\x9d\x29\x53\x2c\x83\x32\x99\x1c\xbf\x26\xe6\x92\xe8\xd8\xcb\xe6\x38\x48\xb1\x7e\xd9\x68\x0f\xf4\x9a\x08\xd1\x21\xee\xf9\x64\x8e\xe1\x4a\xc3\xd0\x07\xdc\xec\x0a\x46\xd6\x5a\x23\x08\xd7\x78\x57\x84\x12\x9a\x06\x11\x56\x30\x58\x68\xe4\x5b\xfd\xd7\x45\xfb\x38\x69\xba\x94\xa7\xda\xd1\xd0\xbe\x5d\x44\xcf\x73\x0b\xa6\x3c\x79\x03\xd6\xb7\x2c\x97\x5f\x00\xce\xdd\x19\x8a\xaa\x47\xa9\x37\x22\xc6\x62\x34\xc9\x37\x71\xc6\x72\x26\x17\x94\x62\x64\xbe\xe8\x27\x23\xe8\x8c\x36\xb9\xbb\x71\x69\xd2\x8f\x65\xd0\x8d\x88\x53\xae\xe6\xb4\x01\x2c\x86\xbc\x60\x39\xc0\xef\xf0\xcd\x77\x5e\xc1\x24\xc2\xbd\x78\xac\xfe\x61\x40\x53\x20\xa0\xb9\xc9\xb0\xf1\xf9\x65\x61\x26\xbd\x33\x19\x2f\xa0\x53\xbb\x35\xbd\x67\x8a\x4d\x4d\x67\xbf\x28\xe0\x26\xc9\x33\x15\xa7\x57\x2d\xec\xbc\x72\x64\x5f\x88\xd8\x16\xd4\x21\x5a\x67\xbb\x4c\xd1\x8c\xde\x81\xa1\xfd\xcc\xe0\xe2\x9a\x55\xf4\xb0\xd3\xcd\x4e\xb2\x26\xef\x42\x52\x63\x17\xea\xa6\x84\xc0\x98\xff\x50\x85\xd8\x58\xbc\x76\xf9\x6d\x3f\xcf\xdc\xce\x8f\xc4\x33\x66\x1e\x63\xac\x66\x97\xa6\x13\xb3\xac\x12\xd6\xed\x11\xd0\x1d\x0b\x3e\x43\x78\xa5\xf0\xe9\x3b\xe1\xec\x4c\xc9\x6a\xc2\xba\xcc\xb2\x2f\xbf\x30\x2e\xf6\x88\x25\xbc\x9e\x65\xf2\x40\xb5\xbb\x1f\xaa\x7f\xe4\xf8\xfe\x91\x7c\x31\xf5\xce\x0e\xd0\x83\xac\xa5\x9f\x67\x44\xcf\xb7\x6a\x6b\x78\xf5\x81\xff\x51\xe2\x2f\x43\x83\x97\xd2\x6e\x09\x10\x5b\x32\xfd\xd0\x6d\xce\x7e\x60\x1a\x04\xad\x59\xbd\x52\xe7\x0e\x5c\x9f\x66\x76\xde\x9b\xb2\xef\xb7\x98\xc4\xbc\x38\xa0\xd5\x43\xa1\x92\xf6\xb7\x39\xd1\xdd\xb5\x77\xfe\x2b\x41\xa9\x12\x46\xf5\xd6\xc5\x4a\x2e\x0b\x6e\x2c\xad\x15\x71\x46\xb2\xcb\x22\xc7\xd3\xb9\x09\x8a\x85\x4d\xbf\x5d\x81\x05\x45\x42\xd3\xce\xe5\x5a\x8f\x73\xf3\xbd\x52\xdc\xf2\xdf\x8a\xb9\xef\x12\x8e\x0f\xe8\x0a\x72\x4a\xea\xa6\x64\xff\xc7\xd0\xb5\x72\x05\x83\x73\x1f\x4b\xac\xe3\x3d\x95\xe0\x94\x36\xac\x64\x16\xfc\xec\x7f\x01\x46\x75\x09\xa5\xda\x75\x38\x9b\xb8\xf9\xfb\xf1\x8d\x28\x99\xe1\xa9\x03\x6b\x6f\xf7\xe2\x57\x94\x3f\x5a\x74\x08\x3b\x0e\x32\x33\xc5\x21\x7b\x1e\x31\xf0\x11\x9d\xee\x79\xa4\xdb\x4c\xd6\x61\xac\xe9\x9e\x8a\x1e\x86\xcd\xe7\xcc\xe2\xdf\x9f\xff\x99\xb1\x7f\xfe\xe3\xcf\xd7\xdb\xff\x17\x00\x00\xff\xff\x20\xd5\x12\xb9\x91\x25\x00\x00")

func clusterroleYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterroleYaml,
		"clusterrole.yaml",
	)
}

func clusterroleYaml() (*asset, error) {
	bytes, err := clusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clusterrolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcd\xbd\x0a\xc2\x40\x0c\x07\xf0\xfd\x9e\x22\x2f\xd0\x8a\x9b\xdc\xa8\x83\x7b\x41\xf7\xb4\x8d\x1a\xdb\x26\x47\x92\x13\xf4\xe9\x45\x70\x93\x3a\xff\x3f\x7e\x58\xf8\x4c\xe6\xac\x92\xc1\x7a\x1c\x5a\xac\x71\x53\xe3\x17\x06\xab\xb4\xd3\xce\x5b\xd6\xcd\x63\x9b\x26\x96\x31\xc3\x61\xae\x1e\x64\x9d\xce\xb4\x67\x19\x59\xae\x69\xa1\xc0\x11\x03\x73\x02\x10\x5c\x28\x83\x3f\x3d\x68\xc9\x68\xda\xb8\x51\x32\x9d\xa9\xa3\xcb\x27\xc7\xc2\x47\xd3\x5a\xfe\x58\x09\xe0\x87\x5a\x7b\xf6\xda\xdf\x69\x08\xcf\xa9\xf9\x8e\x4e\x4e\xb6\xd6\x7e\x07\x00\x00\xff\xff\xc4\xb6\x1b\x05\xeb\x00\x00\x00")

func clusterrolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_clusterrolebindingYaml,
		"clusterrolebinding.yaml",
	)
}

func clusterrolebindingYaml() (*asset, error) {
	bytes, err := clusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"clusterrole.yaml":        clusterroleYaml,
	"clusterrolebinding.yaml": clusterrolebindingYaml,
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
	"clusterrole.yaml":        {clusterroleYaml, map[string]*bintree{}},
	"clusterrolebinding.yaml": {clusterrolebindingYaml, map[string]*bintree{}},
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
