// Package deploy Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// deploy/staticresources/aro.openshift.io_clusters.yaml
// deploy/staticresources/master/role.yaml
// deploy/staticresources/master/rolebinding.yaml
// deploy/staticresources/master/serviceaccount.yaml
// deploy/staticresources/namespace.yaml
// deploy/staticresources/worker/role.yaml
// deploy/staticresources/worker/rolebinding.yaml
// deploy/staticresources/worker/serviceaccount.yaml
package deploy

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _aroOpenshiftIo_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xcd\x72\xdb\x38\x12\xbe\xeb\x29\xba\x3c\x07\xcf\xd4\x5a\xd2\x26\x7b\xd9\xd5\x2d\xe5\xcc\x4c\x69\x77\x36\x71\xc5\xde\x5c\x92\x1c\x5a\x60\x8b\xec\x35\x08\x70\x81\xa6\x14\x65\x6a\xde\x7d\xab\x01\x52\x7f\x26\x63\x3b\x55\x33\xc3\x8b\x2d\x10\xe8\x3f\x74\x7f\xf8\xd0\x9c\x4c\xa7\xd3\x09\x36\xfc\x9e\x42\x64\xef\x16\x80\x0d\xd3\x67\x21\xa7\xbf\xe2\xec\xfe\xef\x71\xc6\x7e\xbe\x79\xb1\x22\xc1\x17\x93\x7b\x76\xc5\x02\xae\xdb\x28\xbe\x7e\x47\xd1\xb7\xc1\xd0\x6b\x5a\xb3\x63\x61\xef\x26\x35\x09\x16\x28\xb8\x98\x00\xa0\x73\x5e\x50\x87\xa3\xfe\x04\x30\xde\x49\xf0\xd6\x52\x98\x96\xe4\x66\xf7\xed\x8a\x56\x2d\xdb\x82\x42\xd2\xd0\xeb\xff\xbe\x75\xf7\xce\x6f\xdd\x0f\x13\x00\x13\x28\x49\xb8\xe3\x9a\xa2\x60\xdd\x2c\xc0\xb5\xd6\x4e\x00\x1c\xd6\xb4\x00\x63\xdb\x28\x14\xe2\x0c\x83\x9f\xf9\x86\x5c\xac\x78\x2d\x33\xf6\x93\xd8\x90\x51\xb5\x65\xf0\x6d\xb3\x80\x07\xef\xb3\x84\xce\xb2\xce\xab\x2c\x2c\x8d\x58\x8e\xf2\xaf\xe3\xd1\x5f\x38\x4a\x7a\xd3\xd8\x36\xa0\x3d\xa8\x4e\x83\x91\x5d\xd9\x5a\x0c\xfb\xe1\x09\x40\x34\xbe\xa1\x63\xa9\xb1\x5d\x85\x2e\x64\x9d\xde\x28\x28\x6d\x5c\xc0\xaf\xbf\x4d\x00\x36\x68\xb9\x48\xde\xe6\x97\x6a\xee\xab\x9b\xe5\xfb\xbf\xdd\x9a\x8a\x6a\xcc\x83\x00\x05\x45\x13\xb8\x49\xf3\x7a\xe1\xc0\x11\xa4\x22\xc8\x33\x61\xed\x43\xfa\xd9\x9b\x08\xaf\x6e\x96\xdd\xea\x26\xf8\x86\x82\x70\x6f\x81\x3e\x47\x9b\xbf\x1f\x3b\xd3\x73\xa9\x86\xe4\x39\x50\xe8\x76\x53\x56\xd8\x6d\x1a\x15\x10\xb3\x6a\xbf\x06\xa9\x38\x42\xa0\x26\x50\x24\x97\x13\xe0\x48\x2c\xe8\x14\x74\xe0\x57\xff\x25\x23\x33\xb8\xa5\xa0\x42\x20\x56\xbe\xb5\x85\xe6\xc8\x86\x82\x40\x20\xe3\x4b\xc7\x5f\xf6\x92\x23\x88\x4f\x2a\x2d\x0a\x75\x5b\xd1\x3f\xec\x84\x82\x43\xab\x21\x6c\xe9\x0a\xd0\x15\x50\xe3\x0e\x02\xa9\x0e\x68\xdd\x91\xb4\x34\x25\xce\xe0\xdf\x3e\x10\xb0\x5b\xfb\x05\x54\x22\x4d\x5c\xcc\xe7\x25\x4b\x9f\xee\xc6\xd7\x75\xeb\x58\x76\xf3\x94\xb4\xbc\x6a\xc5\x87\x38\x2f\x68\x43\x76\x1e\xb9\x9c\x62\x30\x15\x0b\x19\x69\x03\xcd\xb1\xe1\x69\x32\xdc\xa5\x6c\x9f\xd5\xc5\x77\xfb\x8d\xbe\x3c\xb2\x54\x76\x9a\x10\x51\x02\xbb\x72\x3f\x9c\x72\x6f\x34\xee\x9a\x83\xba\xbb\xd8\x2d\xcb\xf6\x1f\xc2\xab\x43\x1a\x95\x77\x3f\xde\xde\x41\xaf\x34\x6d\xc1\x69\xcc\x53\xb4\x0f\xcb\xe2\x21\xf0\x1a\x28\x76\x6b\x0a\x79\xe3\xd6\xc1\xd7\x49\x22\xb9\xa2\xf1\xec\xa4\xcb\x24\x26\x77\x1a\xf4\xd8\xae\x6a\x16\xdd\xe9\xff\xb5\x14\x45\xf7\x67\x06\xd7\xa9\xe8\x61\x45\xd0\x36\x05\x0a\x15\x33\x58\x3a\xb8\xc6\x9a\xec\x35\x46\xfa\xdd\xc3\xae\x11\x8e\x53\x0d\xe9\xe3\x81\x3f\xc6\xaa\xd3\x89\x39\x5a\xfb\xe1\x1e\x4a\x06\x77\xa8\xab\xc0\xdb\x86\xcc\x49\x65\x14\x14\x39\x68\xf6\x0a\x0a\x69\xce\x1f\xa3\xcb\x78\x2d\xea\x53\x92\xa3\x0d\xfe\xe2\xcb\x92\x5d\x79\xfa\x6a\x7c\x91\x3e\x68\xc2\x1b\x85\xc5\x07\x2f\x46\x22\xd0\x3f\xc6\xbb\x35\x97\x03\x18\xb0\x57\x8a\xa2\x15\xb6\x80\xcb\x0f\x7f\x9d\xfe\xe3\xd3\x5f\x66\xf9\xcf\xe5\x73\x15\xd5\xde\xb1\x78\x7d\xf5\xf3\xf5\xed\x8f\x6e\xc3\xc1\xbb\x9a\x9c\x0c\xe9\x24\xd7\xd6\x43\xe3\x53\x78\xcd\x58\x3a\x1f\x85\x4d\xbc\x09\xbe\x18\x9c\x73\x77\x8e\x11\xcf\xb5\xee\x1d\x95\x23\xc1\x78\xa2\x8c\x3b\x72\x38\xec\xd9\x57\x05\xa4\x53\xa9\x41\xf3\xdc\x5d\x1c\xcc\x5c\xd8\x83\x23\xc9\x75\x45\xe6\x9e\xc2\x73\xf2\x29\xb2\x0c\x0d\x03\xb0\x50\x3d\xf8\xe2\x11\xe7\xfa\xd7\x18\x02\xee\x9e\xea\x40\x8f\x6a\xcb\xe2\x5c\xe3\x49\x21\xf6\x6c\x64\xf9\xba\x3f\x0d\x5f\x7d\x69\x03\x1d\x2d\xcf\x87\x13\x1d\x1d\xd1\x4f\x32\x3c\x52\xd8\xb0\xa1\x9b\xc0\xce\x70\x83\xf6\x39\x11\xac\x31\x81\x43\xbb\x72\x24\x0f\xed\x7f\x34\x5e\x5b\x1f\xee\x0f\xcb\xff\xec\xad\x18\x06\xc7\x4c\x61\x1e\x83\xc7\x34\xeb\x04\x20\xfd\x4a\x03\xfb\x6d\x08\x69\xbc\x2b\xf8\x88\x59\x8e\x29\xdf\x4f\xeb\x8e\x50\x92\xa4\xa7\x1f\x06\x76\x51\xd0\x19\x8a\xb3\xc9\x93\xa2\x7a\x22\xfd\xe2\x20\xe7\x70\xae\x66\x6a\xa3\x9e\x25\xe2\x73\x42\x76\x2e\x63\xf6\xf5\x5c\x99\x3e\x47\xa6\x62\x20\x5d\xb3\xe7\xe0\x50\x93\xa9\xd0\x71\xac\x53\x2d\xbb\x82\x0a\xe5\x42\x7a\xc6\x46\x2a\x60\x5b\x91\xd3\x80\x0e\x08\x2d\x48\x90\x6d\xdc\x1b\x71\x30\x4b\x75\xe8\x41\x8d\xd0\x04\xf6\x81\x21\x31\x6e\xf0\x01\xb6\x89\x86\xa5\x77\x4d\x63\xcf\x33\x23\xa5\x81\x07\xb4\xf6\x10\xbb\x24\x1e\x4a\xde\x90\x03\xa5\x2b\x33\xf8\xe8\x8e\xfd\xe9\x98\xdd\x8a\x00\x8b\x82\x86\xe0\x5a\x3c\xd0\xe7\xc6\xb2\x61\xb1\xbb\x4c\x01\x77\x47\x7b\x0f\x52\xa1\xa8\xb3\x21\x26\x6a\x67\x7c\xdd\x78\x97\xa2\x6d\x52\xb0\x56\xbe\x1d\x02\xfb\x80\x52\x25\x5a\x83\x2e\xb1\x14\x0e\x99\x2d\xf9\x48\x27\xd2\x53\x2c\x13\x05\xd2\x03\x3b\x11\x20\xaf\x2b\x07\x44\x1e\xc5\x30\xce\xe0\xad\x33\xd4\xe5\x74\x71\x95\x92\xba\x26\x74\xaa\x24\x85\xe4\x90\x1f\x06\x1d\x64\x5e\x34\x20\x53\x37\xb7\xa4\x02\x30\xac\x58\x02\x06\xb6\x3b\x98\x02\xeb\x6c\xe3\x6b\x8a\xd0\x60\x90\x1e\xbb\x5e\xdd\x2c\x33\xbf\xad\x30\x97\x51\xc4\x7a\x48\xe8\x0a\xcd\xfd\x16\x43\x11\xa7\x69\xf6\xda\x87\xfc\x4b\x63\x87\xc2\x2b\xb6\x2c\x29\xd4\x86\x82\xeb\x32\x64\x97\xdd\x4e\xfa\x86\x7c\xdf\x5b\x30\xbb\x78\xf0\xfa\x6b\x20\x08\x60\x31\xca\x5d\x40\x17\xb9\xbf\xcc\x0d\x23\xd6\xda\x87\x1a\x65\x01\x4a\x1d\xa7\xc2\x83\x9e\x3d\x8a\x6b\x35\xc5\x88\xe5\x88\x86\x47\xd6\x06\xc2\x38\x7c\xea\x8f\x41\xcb\xbb\xb4\x42\xf1\xe5\xac\x38\x11\xbc\xa3\xe9\xd6\x87\xe2\xea\x40\x80\x07\x05\xc3\xd9\x6d\x69\x7f\x4e\xa1\x50\xe9\xc3\x4e\x7f\x1b\x6c\x23\xed\x5f\xb4\x21\x90\x93\x0e\x7b\x87\xe0\x44\x9f\xa5\x0c\x58\x95\x20\x83\x5d\xda\x79\x56\x89\xad\x34\xad\x5c\x41\x6c\x4d\x05\x18\x93\xcd\x96\xdd\x98\xa1\x7a\x67\x37\x62\xa1\x54\x24\xed\x96\x6a\x7e\xb1\x83\xd8\xd6\x35\x06\xfe\x92\xd2\xdf\x64\x13\x3b\x74\x48\xc6\x8f\xd8\xf9\xc8\x86\x3c\x3c\x5e\x9e\xbc\x34\xbd\x7e\x7c\x27\x0f\x30\x7e\xb7\x6b\xa8\xe7\x0e\xba\x78\x1f\xee\x7d\x1d\x27\x57\xe3\x48\x6c\x64\xd7\xb0\x41\x6b\x77\x5a\xfa\xfd\x86\x17\x7a\x86\x17\x0a\xac\xb1\xf2\x41\xa0\xa9\x42\xba\x07\x1d\x43\x64\x52\x36\x26\xb5\x43\x4f\x76\x05\x6b\x3e\x74\xa7\x25\x27\xc8\x87\x8f\x17\xb8\x72\x5a\x33\x76\x2a\xa1\xa5\x8f\x17\xd0\x78\x8b\x81\x65\x37\x83\x9f\xfc\x10\x80\xe9\x43\x9f\xb1\x6e\x2c\x5d\x01\x9f\xfb\xd7\x6b\x89\xf9\x54\x41\x15\xc7\x66\x97\xf3\x28\xf5\x27\xae\xc6\x9c\x4f\xd6\x70\xcc\x5d\x8c\x8f\x17\x60\x30\xa6\x60\x36\xc1\xaf\x70\x65\x77\x69\x86\xda\x7a\x05\xd1\x9f\xaa\xfd\xba\xe7\x2b\x2d\x04\x6b\xa9\x80\x8f\x17\x4b\xd7\x89\x1f\x40\x20\x78\x2c\x23\xf2\x11\x40\x03\x2c\x6c\xda\xa5\xd9\xc0\x0b\x95\xf8\x60\x78\x94\xaa\x8e\x93\x2a\x45\x47\x14\x1f\x46\x2e\x59\xa3\x86\x07\xb2\x7a\x8d\x7e\x9b\x74\x3d\xa8\x83\xa7\xd0\x94\xbc\xf4\x1d\xad\x29\x15\x64\xea\xc2\x21\xbb\x08\xe4\x7c\x5b\x56\xe9\x32\xae\xa8\x9b\x12\xd1\x83\x25\x81\x9d\x6f\x87\x98\xa6\xd3\x8b\xb0\x68\x2e\xd7\xbe\xe0\x75\xde\xd2\x40\xdd\xb9\xd9\x35\x74\x9e\x79\x32\x0c\xf7\x9e\x46\x5c\x79\x75\xb3\xec\x3b\x4e\x7d\x6d\x86\xec\xd7\x80\xde\xaf\x86\x35\x3f\x6b\x26\x5b\xdc\xa0\x54\x4f\xd0\x7d\xb9\x5c\x77\xbe\x26\x12\xe1\xb5\x38\x98\x0c\x9d\x30\xbc\x44\x8a\x08\xf5\xa6\x31\x92\xd4\x4a\xec\x9c\x70\xa0\x6e\xc5\x55\xee\xba\x74\xcd\x9d\x43\x13\x4c\xb7\x08\x30\x57\x13\xfc\xf3\xf6\xed\x9b\xf9\xcf\x7e\x44\x64\xf2\x02\xd0\x18\x8a\x1d\xc1\xd4\x2b\xf5\x01\xd2\xbb\x4e\xc4\x6d\xa2\x9e\x35\x3a\x5e\x53\x94\x59\xa7\x83\x42\xfc\xf0\xf2\xd3\xd8\x11\xf2\x93\x0f\x0f\xd0\x62\xdf\x4a\xea\x13\x8a\x63\x0e\xc7\x5e\x22\x6c\x59\x2a\x1e\x2b\x6b\x05\x95\xa2\x73\x3b\x93\x4d\xc1\x7b\x02\xdf\xb9\xdb\x12\x58\xbe\xa7\x05\x5c\x68\xb6\x1d\x99\xf9\xab\xde\x8c\x7f\x1b\xae\x7b\x80\xef\xb7\x15\x05\x82\x0b\x9d\x74\x91\x8d\xdb\x77\x0c\x75\xec\x08\xcb\x3b\x23\x13\xa9\x94\xc0\x65\x49\x61\x90\x95\x42\x07\x6b\xb4\x21\x27\x3f\x68\xda\xf3\x1a\x9c\x3f\x12\x91\x04\xeb\xee\x35\x64\x78\xcd\x54\x3c\x30\xfa\xc3\xcb\x4f\xa3\x16\x9f\xc6\x4b\xb1\x97\x3e\xc3\xcb\x0c\xed\x0a\x9c\xbe\xf8\x61\x06\x77\x29\x3b\x76\x4e\xf0\xb3\x6a\x32\x4a\x5e\xc7\x22\xeb\x9d\xc2\xac\x87\x0a\x37\x04\xd1\xd7\x04\x5b\xb2\x76\xda\x11\x54\xd8\x62\x62\x12\xfd\xc6\x69\xbe\x61\xcf\x2d\xc7\xb3\xb5\xef\xd3\xde\xbd\x7d\xfd\x76\x91\x2d\xd3\x84\x2a\x13\xe5\x51\x4e\xbb\x66\x87\x36\x9d\x8c\xb9\x7f\x98\xb2\x71\xf4\x90\x8c\x6d\x4e\x1f\xf1\x1d\xef\xed\x8f\xb2\x75\x2b\x6d\xa0\xd9\x50\x4b\xe9\xd1\x3a\x3e\x6f\xa1\x1e\x9e\x81\x66\xea\x39\x70\xfc\x49\x2d\xc9\x27\x3b\xe7\x46\x3a\x7a\xe7\xce\xbd\x39\xca\xf2\xaf\x3a\xa7\x1c\x2e\x38\x12\x4a\xfe\x15\xde\xc4\x79\xba\x07\x34\x12\xe7\x7e\x43\x61\xc3\xb4\x9d\x6f\x7d\xb8\x67\x57\x4e\x35\x35\xa7\x39\x07\xe2\x3c\xf5\xa6\xe6\xdf\xa5\x3f\xdf\xec\xcb\x68\x73\x6b\xc8\xa1\x34\xf9\x8f\xf0\x4a\xf5\xc4\xf9\x37\x39\xd5\x37\x98\x9e\x7e\x8e\x5d\xde\x66\xc0\x30\xe7\x6b\xb5\x2c\xb6\x15\x9b\xaa\xff\x88\xd2\x61\xec\x48\x31\x71\x84\x1a\x8b\x0c\xcd\xe8\x76\xbf\x7b\x2a\x6b\x40\x33\xaf\xdf\x4d\xbb\x8f\x79\x53\x74\x85\xfe\x1f\x39\x8a\x8e\x7f\x53\x04\x5b\x7e\x52\xf9\xfe\x67\xf9\xfa\x8f\x49\xf0\x96\xbf\xa9\x56\x9f\x4d\x0b\x07\x16\x9c\x0d\xed\x3f\x8a\x6e\x5e\xa0\x6d\x2a\x7c\x71\x18\x4b\x74\x6a\xda\x7d\x07\x3d\x7a\x9d\x1b\x96\x54\x2c\x40\xaf\x04\x79\x40\x7c\xd0\x1b\x71\x1e\x39\x5c\xa9\x94\x33\x34\x42\xc5\x9b\xf3\x2f\xa1\x17\xf9\xd0\xea\x3f\x75\xa6\x9f\x47\x3d\x37\xf8\xf0\x69\x92\xa5\x52\xf1\xbe\xb7\x46\x07\xff\x1f\x00\x00\xff\xff\x5d\xa2\x37\x65\x4c\x1e\x00\x00")

func aroOpenshiftIo_clustersYamlBytes() ([]byte, error) {
	return bindataRead(
		_aroOpenshiftIo_clustersYaml,
		"aro.openshift.io_clusters.yaml",
	)
}

func aroOpenshiftIo_clustersYaml() (*asset, error) {
	bytes, err := aroOpenshiftIo_clustersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aro.openshift.io_clusters.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\x31\x6f\x32\x31\x0c\x86\xf7\xfc\x8a\x88\x3d\x87\xbe\xed\xd3\xad\x1d\xba\x57\x55\x77\x93\x33\x60\x91\xc4\x91\xed\xd0\x96\x5f\x5f\xdd\xc1\xa9\x12\x54\x14\xda\xa5\xd3\xd9\x96\xfd\x3e\xaf\x2f\x89\x0b\x21\x38\xa8\xf4\x82\xa2\xc4\xa5\xf7\xb2\x82\xd8\x41\xb3\x2d\x0b\x1d\xc0\x88\x4b\xb7\xfb\xaf\x1d\xf1\x72\xff\xcf\xed\xa8\x0c\xbd\x7f\x48\x4d\x0d\xe5\x89\x13\xba\x8c\x06\x03\x18\xf4\xce\xfb\x28\x38\x0d\x3c\x53\x46\x35\xc8\xb5\xf7\xa5\xa5\xe4\xbc\x2f\x90\xb1\xf7\x20\x1c\xb8\xa2\x80\xb1\x84\x0c\xa3\x86\x93\x96\x50\x7b\x17\x3c\x54\x7a\x14\x6e\x55\x47\xa5\xe0\x17\x0b\xe7\xbd\xa0\x72\x93\x88\xa7\x5a\xe4\xb2\xa6\x4d\x86\xaa\x53\x3a\x8a\x6a\x85\x88\xc7\x54\x51\xf6\x14\x11\x62\xe4\x56\x6c\xac\xed\x51\x56\xf3\xe8\x68\x0d\xa7\x70\xc0\x84\xa7\x70\x83\x36\x7d\x13\xe9\x31\x68\x75\x98\xdb\x5e\xc1\xe2\xf6\x36\x5f\x8a\x51\xf0\x0a\xf1\x02\x53\x27\xed\xdb\x80\x50\xa7\x75\xcf\x90\x03\x60\xe6\xa2\xd7\xa8\xbf\xdb\x13\x84\x3b\xae\x58\x74\x4b\x6b\xeb\x88\xbf\x38\x8d\xe3\x2d\xf8\x99\x81\xbb\xfe\xc0\x3d\x56\x3e\x93\xe5\x9a\x0a\x24\x3a\xfc\x2d\x8b\x4b\x35\xb0\x76\xe6\x68\x66\x5f\x20\x2f\x40\x8a\xb1\x09\xd9\xfb\x37\xb4\xb9\x2d\x72\x31\x7c\xb3\xc8\x45\x4d\x80\xee\x7c\x16\x27\x13\x1f\x01\x00\x00\xff\xff\x53\x74\x48\x50\x23\x04\x00\x00")

func masterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterRoleYaml,
		"master/role.yaml",
	)
}

func masterRoleYaml() (*asset, error) {
	bytes, err := masterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\xb1\x4e\xc5\x30\x0c\x45\x77\x7f\x85\x7f\x20\x45\x6c\x28\x1b\x30\xb0\x3f\x24\x76\xbf\xd4\xa5\xa6\xad\x1d\x39\x4e\x87\x7e\x3d\xaa\x40\x2c\x48\x9d\xef\x39\xf7\x40\x4a\x09\xa8\xca\x07\x7b\x13\xd3\x8c\x7e\xa7\x32\x50\x8f\xd9\x5c\x0e\x0a\x31\x1d\x96\xa7\x36\x88\x3d\xec\x8f\xb0\x88\x8e\x19\x5f\xd7\xde\x82\xfd\x66\x2b\xbf\x88\x8e\xa2\x9f\xb0\x71\xd0\x48\x41\x19\x10\x95\x36\xce\x48\x6e\xc9\x2a\x3b\x85\x79\xda\xe8\x14\xc0\x6d\xe5\x1b\x4f\x27\x44\x55\xde\xdc\x7a\xbd\x08\x02\xe2\xbf\xde\xe5\x7d\xeb\xf7\x2f\x2e\xd1\x32\xa4\x5f\xf3\x9d\x7d\x97\xc2\xcf\xa5\x58\xd7\xb8\x94\x7f\xb6\x56\xa9\x70\x46\xab\xac\x6d\x96\x29\x12\x1d\xdd\xf9\x0f\x86\xef\x00\x00\x00\xff\xff\xea\x5c\x27\x5f\x2f\x01\x00\x00")

func masterRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterRolebindingYaml,
		"master/rolebinding.yaml",
	)
}

func masterRolebindingYaml() (*asset, error) {
	bytes, err := masterRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _masterServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\x8e\x02\x31\x0c\x05\xd0\xde\xa7\xf0\x05\x5c\x6c\x9b\x6e\xcf\x80\x44\xff\x95\xf9\x88\x08\xc5\x8e\x1c\xcf\x14\x9c\x9e\x06\x51\xbf\x27\x66\x26\x58\xe3\xce\xdc\x23\xbc\xe9\xf5\x27\xaf\xe1\x47\xd3\x1b\xf3\x1a\x9d\xff\xbd\xc7\xe9\x25\x93\x85\x03\x85\x26\xaa\x8e\xc9\xa6\xc8\xb0\x58\x4c\x54\xa4\x4d\xec\x62\x7e\x6d\x2f\x74\x36\x8d\x45\xdf\xcf\xf1\x28\xc3\xfb\x4c\xfe\xb2\x7c\x02\x00\x00\xff\xff\x5b\x98\x41\x31\x75\x00\x00\x00")

func masterServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_masterServiceaccountYaml,
		"master/serviceaccount.yaml",
	)
}

func masterServiceaccountYaml() (*asset, error) {
	bytes, err := masterServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xca\xb1\x0d\xc2\x40\x0c\x05\xd0\xde\x53\x78\x01\x17\xb4\x37\x04\x25\xfd\x27\xf9\x11\x56\x88\xef\x64\x5f\x28\x32\x3d\x4a\xf9\xa4\x27\x66\x26\x18\xfe\x62\x96\xf7\x68\xfa\x7b\xc8\xee\xb1\x36\x7d\xe2\x60\x0d\x2c\x94\x83\x13\x2b\x26\x9a\xa8\x06\x0e\x36\xed\x83\x51\x1f\xdf\xa6\xe1\x3a\x93\xd6\x07\x13\xb3\xa7\xd4\xe0\x72\xb7\xcd\x03\x5f\xbf\x98\x75\xcb\x74\x3f\xdf\xcc\xe0\x64\xc9\x3f\x00\x00\xff\xff\x44\x6f\xf6\xda\x72\x00\x00\x00")

func namespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_namespaceYaml,
		"namespace.yaml",
	)
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\x31\x4e\x04\x31\x0c\x45\x7b\x9f\xc2\x17\xc8\xac\xe8\x50\x5a\x0a\x7a\x84\xe8\xbd\x33\x86\xb5\x26\x13\x47\xb6\xb3\x2b\x71\x7a\x34\xd9\x9d\x06\x4a\xaa\x3c\x59\xff\xe7\x3f\x48\x29\x01\x35\xf9\x60\x73\xd1\x9a\xd1\xce\x34\x4f\xd4\xe3\xa2\x26\xdf\x14\xa2\x75\x5a\x9f\x7d\x12\x3d\x5d\x9f\x60\x95\xba\x64\x7c\x29\xdd\x83\xed\x4d\x0b\xc3\xc6\x41\x0b\x05\x65\x40\x9c\x8d\x47\xe1\x5d\x36\xf6\xa0\xad\x65\xac\xbd\x14\x40\xac\xb4\x71\x46\x32\x4d\xda\xd8\x28\xd4\xd2\x4d\x6d\x65\x03\xeb\x85\x3d\x43\x42\x6a\xf2\x6a\xda\x9b\xef\x3f\xa5\x3d\x3b\x69\xe3\xea\x17\xf9\x8c\x49\x14\x10\x8d\x5d\xbb\xcd\xfc\x48\xcc\x77\x0b\x07\xc4\x2b\xdb\xf9\xb8\xee\x0e\x3c\x70\xe1\xc2\x0f\xfc\xe2\x18\x6f\x11\xbf\x43\xa3\x98\x2f\x83\x7a\x5b\x8e\xc2\x6d\x1c\xff\xa1\x72\xf2\xa0\xe8\xbf\x8c\x8e\xed\x3f\x93\x3f\x01\x00\x00\xff\xff\x32\xe1\x82\x0f\x7b\x01\x00\x00")

func workerRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRoleYaml,
		"worker/role.yaml",
	)
}

func workerRoleYaml() (*asset, error) {
	bytes, err := workerRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerRolebindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8d\x31\x4e\xc6\x30\x0c\x46\x77\x9f\xc2\x17\x48\x11\x1b\xca\x06\x0c\xec\x3f\x12\xbb\x9b\xba\xd4\xb4\xb5\x23\xc7\x29\x52\x4f\x8f\x10\x88\x05\xa9\xf3\xf7\xde\xf7\x20\xa5\x04\x54\xe5\x8d\xbd\x89\x69\x46\x1f\xa9\x0c\xd4\x63\x31\x97\x93\x42\x4c\x87\xf5\xa1\x0d\x62\x77\xc7\x3d\xac\xa2\x53\xc6\xe7\xad\xb7\x60\xbf\xd9\xc6\x4f\xa2\x93\xe8\x3b\xec\x1c\x34\x51\x50\x06\x44\xa5\x9d\x33\x92\x5b\xb2\xca\x4e\x61\x9e\x3e\xcd\x57\x76\x70\xdb\xf8\xc6\xf3\x37\x44\x55\x5e\xdc\x7a\xbd\x08\x02\xe2\xbf\xde\xe5\x7d\xeb\xe3\x07\x97\x68\x19\xd2\xaf\xf9\xca\x7e\x48\xe1\xc7\x52\xac\x6b\x5c\xca\x3f\x5b\xab\x54\x38\xa3\x55\xd6\xb6\xc8\x1c\x89\xce\xee\xfc\x07\xc3\x57\x00\x00\x00\xff\xff\x21\x49\xf8\xf0\x2f\x01\x00\x00")

func workerRolebindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerRolebindingYaml,
		"worker/rolebinding.yaml",
	)
}

func workerRolebindingYaml() (*asset, error) {
	bytes, err := workerRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _workerServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\x31\xaa\x03\x41\x08\x06\xe0\xde\x53\x78\x01\x8b\xd7\xda\xbd\x33\x04\xd2\xcb\xec\x1f\x22\xcb\xea\xe0\xb8\x1b\xc8\xe9\xd3\x84\xd4\xdf\x47\x22\x42\x36\xfd\x8e\x5a\x9e\xa1\x7c\xfd\xd1\xee\xb1\x29\xdf\x50\x97\x0f\xfc\x8f\x91\x67\x34\x1d\x68\xdb\xac\x4d\x89\x39\xec\x80\xb2\x55\x4a\x4e\x94\x75\x96\xbc\xb2\x76\xd4\xd7\xd6\xb4\x01\xe5\x9c\x88\xf5\xf4\x47\x8b\xbd\xcf\xc2\x2f\xd3\x27\x00\x00\xff\xff\x5c\x51\x06\x72\x75\x00\x00\x00")

func workerServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_workerServiceaccountYaml,
		"worker/serviceaccount.yaml",
	)
}

func workerServiceaccountYaml() (*asset, error) {
	bytes, err := workerServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "worker/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"aro.openshift.io_clusters.yaml": aroOpenshiftIo_clustersYaml,
	"master/role.yaml":               masterRoleYaml,
	"master/rolebinding.yaml":        masterRolebindingYaml,
	"master/serviceaccount.yaml":     masterServiceaccountYaml,
	"namespace.yaml":                 namespaceYaml,
	"worker/role.yaml":               workerRoleYaml,
	"worker/rolebinding.yaml":        workerRolebindingYaml,
	"worker/serviceaccount.yaml":     workerServiceaccountYaml,
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
	"aro.openshift.io_clusters.yaml": {aroOpenshiftIo_clustersYaml, map[string]*bintree{}},
	"master": {nil, map[string]*bintree{
		"role.yaml":           {masterRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {masterRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {masterServiceaccountYaml, map[string]*bintree{}},
	}},
	"namespace.yaml": {namespaceYaml, map[string]*bintree{}},
	"worker": {nil, map[string]*bintree{
		"role.yaml":           {workerRoleYaml, map[string]*bintree{}},
		"rolebinding.yaml":    {workerRolebindingYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {workerServiceaccountYaml, map[string]*bintree{}},
	}},
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
