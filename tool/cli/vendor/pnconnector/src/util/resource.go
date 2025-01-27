// Code generated for package util by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../resources/events.yaml
package util

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

var _ResourcesEventsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5d\xdd\x52\x1c\x39\xb2\xbe\x9f\xa7\x50\x70\xb1\x3b\x13\x01\x3e\xa6\x7f\x6c\xe0\x0e\xf3\xe3\xc3\x86\x31\x5e\xda\xf6\x9e\x73\x35\xa1\xae\x52\x77\x6b\xa8\x92\x6a\x24\x15\xd0\xb3\xb1\xef\xc2\xb3\xf0\x64\x27\xf4\x53\x3f\x5d\x52\x55\x67\xab\x3d\x73\x26\x62\x22\x0c\xa8\xf4\x65\xa6\xa4\x54\x2a\x95\x99\x3a\x3a\x3a\xfa\xe9\x23\xc7\xd9\x19\xba\x7a\x24\x4c\x49\xf4\x8d\xd1\x05\x4d\xb0\xa2\x9c\xfd\xf4\x9d\x08\x49\x39\x3b\x43\x07\x8f\x6f\x4e\x47\x07\x3f\xe5\x3c\x2d\x33\x22\xcf\x7e\x42\xe8\x08\x31\x9c\x93\x33\x74\x70\x71\x77\x7b\x7b\xf7\xf9\xe0\x27\x84\x10\x4a\x78\xc9\xd4\x19\x3a\x3d\x3d\x35\x3f\xd2\x74\xa6\xb0\x50\x67\xe8\xad\xfb\xf1\x8a\xa5\x67\x08\x35\x7f\x67\x0b\x7e\x66\xfe\xa5\xfb\x4b\x78\x4a\xaa\xa6\xfa\xbf\x8c\x3c\x92\xec\x0c\x1d\xdc\x7c\xbe\xbe\x3b\xa8\x7f\x9b\x13\x29\xf1\x52\x03\xcf\xca\x24\x21\x52\x36\x7f\x2a\x04\x9f\x67\x24\x3f\x43\x07\xcd\xef\x24\xcf\x4a\x65\x59\x38\xd8\xa0\xfa\xd3\xcd\x06\xc9\xc7\x6f\xdf\x6e\x92\x7c\xfc\xf6\x6d\x87\xea\xe3\xb7\x03\x64\xd7\xcd\x5b\x94\x6b\xc2\x7d\xba\xe9\x9c\x2f\xb8\x44\x54\x22\xa9\x91\x48\xea\xd1\xef\x13\xef\x83\x1d\xef\x0a\xa6\x88\xc8\x29\xc3\xb1\x78\x23\x18\x5e\x92\x51\x24\x89\x78\x24\x42\x63\x52\x46\x15\xc5\x19\xfd\x23\x12\x74\x0c\x03\x65\xe4\x49\x03\x13\xa6\x34\x68\xc2\x19\x23\x49\x2c\x9f\x13\x30\x9f\x0e\x2e\xa5\x72\x3f\xc4\x29\x0c\x51\x90\xdf\x4b\x22\x15\xca\xe5\x52\xc3\x0a\x92\x10\xfa\x18\x09\xf9\x0e\x0a\x29\x0b\xce\x24\xa9\x30\x25\x61\x2a\x06\xef\x18\xb8\x32\x5a\x93\xa7\x10\xa4\xc0\x82\xb2\x25\x22\xcf\x34\x0e\x14\xb8\x42\x5a\xa0\x6a\x25\x08\x4e\xd1\x6f\x9c\xb2\x38\xc1\x1e\x03\x57\x09\x61\x78\x9e\x11\x24\x48\x29\xc9\x11\x4e\x53\x11\x05\xe6\xad\x8e\x7f\x9d\xdf\x7f\x06\x80\xa1\x05\xa6\x59\x24\x83\xde\xf2\xb8\xba\xbf\xbf\xbb\xf7\x41\x25\x4f\x1e\x88\x42\x89\x20\x66\x23\xd9\x07\xd2\x5b\x1f\xc3\x90\x73\xca\x52\x3d\x6f\xf6\x40\xf4\x96\xc7\x30\x62\x46\xa5\x22\x7b\xb1\xf8\x1e\x08\x48\x0a\x9e\x65\x3f\x44\xa8\x27\xbb\xb1\x88\x93\x84\x14\x6a\x1f\xc0\x53\x20\x60\x8e\x9f\x2b\xe5\x4a\x84\xe0\x51\x2b\x63\xe4\xe9\x9b\x3e\x30\xfb\x0f\xad\xd8\xf6\x9c\x33\x23\x4f\xdb\x6c\x81\xb4\xfa\x7b\x4f\x50\x4f\xdd\x84\x35\xc0\xd7\x9b\xdb\xab\x4b\x74\xf7\xed\x6b\x14\x88\xa7\x66\x7a\x38\xbb\xf9\xfc\xfd\xfc\xd3\xcd\x25\xfa\x72\x7e\x7f\x7e\x1b\x83\x34\x06\x6e\x13\x5f\xee\x66\xe8\x66\x86\x3e\x7c\x9b\xfd\x2f\x0c\xa6\x36\xfa\x6e\x3e\xf0\xeb\xbb\x19\x9a\x29\xac\x08\xba\xc5\x0c\x2f\x89\xd8\xb0\x02\x47\x9e\x15\x78\xec\x59\x81\xa3\x21\x2b\xf0\xd8\x37\xcc\x2e\xaf\x3e\x7c\xfb\x18\x58\x59\x86\x88\x84\x33\x45\x9e\x15\xc2\x69\x1a\x35\x07\x8e\xa1\x86\x99\x83\x5b\x61\xb6\x8c\x02\x1a\x8d\x3d\xd5\x7f\x7f\xf5\xe5\xee\xfe\xeb\xaf\x5f\xef\xcf\x2f\xae\x02\x96\xa7\x93\xf5\x5a\x2a\x92\xa3\x7b\x92\xf0\x47\x22\xd6\xe8\x86\x15\x82\x2f\x05\x91\x72\xc7\xb1\xfb\xce\xb3\x32\x27\xa1\x41\x1b\x77\x07\x6d\xe4\x99\xee\xa3\xa1\x41\x1b\x81\x4d\x77\x4b\x83\x55\xc0\x31\x52\x1c\x81\xed\x76\x87\x94\x92\x8c\xc4\x22\x01\x27\x86\x43\xca\xb5\x2c\x23\x91\x80\x66\xba\x43\x2a\xd9\x3e\x58\x40\xfb\xdc\x61\x99\x65\x85\x14\x47\x6a\x45\xcc\x3e\x1d\x85\x09\xb4\xd0\x2d\xe6\xeb\x8b\x20\x39\x7f\x24\x29\x5a\x08\x9e\x6b\xe0\xd7\x97\x58\x64\xdf\x70\xee\x51\xf0\x2b\x52\x9d\x0f\x48\x8a\x1e\xdd\xdc\xe1\x44\x22\xc6\x95\xb6\xa0\x03\xf0\xe6\x23\xd7\xf4\x89\xaa\x95\x11\x91\xd7\x89\x5e\x7c\x88\x8b\xea\xc7\x9b\xcb\xbe\x6e\x9b\x73\xf6\x15\x53\xc6\x96\xd6\xca\x4d\x08\x92\xa8\xfe\xbe\xf0\x42\x37\x4d\x56\x24\x79\xd0\x3b\xa1\x6a\x28\x6a\x49\xac\x25\x0e\x6f\xf1\xf4\x88\x83\x73\x94\x63\xb6\x76\x9d\xf9\x9a\xa6\x5e\xc8\x98\x19\x56\xe6\xf5\x9a\x46\x73\x92\xe0\x52\x12\x43\x4b\x8e\x9f\x69\x5e\xe6\x88\x95\xf9\x9c\x08\xc4\x17\x55\x87\x48\xad\xb0\x32\x5f\xb7\xbe\xa4\x12\x91\xe7\x84\x90\xb6\x22\x6f\xa4\x72\x4f\x94\x58\x3b\x86\xcd\x04\xd1\x0c\x97\xfa\xdc\xa8\xa9\x16\x15\xad\x08\xe7\xdc\x1e\x7a\xa4\xd2\x2d\xaa\xce\x37\x39\x69\x89\x04\x68\x02\x5c\x19\xca\x6c\x77\xe6\x38\xe0\xe0\x24\xfd\x83\x84\xa7\x86\x37\x15\x74\x53\xc7\xa1\x44\xe7\x42\xe0\x35\x4a\x70\x81\x13\xaa\xd6\x01\x7e\x2f\xf4\xa0\x1a\x29\x4a\xbb\x03\x54\x6d\x11\x66\x29\x32\xb2\x58\x62\xca\x3c\x86\x7c\xdb\x2d\xcc\xd0\xf7\xd6\x9c\xa2\x12\x29\xce\x91\x5c\x71\x31\x3c\xcf\x4d\x6b\xa2\xe7\xa7\x1d\x2f\xd5\xfd\xa8\x3b\x8b\xf1\xc6\x97\x73\xa2\x9e\x08\x61\x68\x64\x78\x18\x4d\xa7\x7a\x3f\x15\x38\x51\x44\xf8\x23\xe3\x5b\x84\x5b\x19\x79\x7d\xa9\x58\xc9\x38\x5b\xf6\xce\x5a\x9f\x8b\xce\x07\xc3\x5c\x98\xb9\xdb\x5a\xc5\x66\x62\x6c\x63\x06\x38\xcd\x3a\xa3\x92\x96\x45\x46\x93\xe0\x5e\x89\xce\x37\x94\x0f\x6e\xda\xda\xaf\x71\xa6\xcf\xe2\x6b\xbb\x12\xe4\x00\x6b\x29\x5d\x2c\x88\xd0\xe7\x85\x16\x93\x3e\x03\xc0\xc3\xf2\x37\x66\x8f\x3a\xed\x35\xd2\xea\x70\xcb\x60\x68\x6b\x0e\x53\x26\x11\x46\x52\x09\xab\xd2\xb0\xf1\x12\x69\x51\xe3\x2c\xe3\x4f\x41\xe5\xd0\xa8\x4c\x6f\xa0\x72\x42\x94\xf4\xfe\x24\xca\x2c\xa0\x0c\x7c\x03\x3a\xcc\xe4\x7d\xbd\xb4\xcd\x9a\x36\xf3\x07\x8b\xe5\x4e\x8a\xa0\x9a\x77\x1b\xdf\x85\xd5\xdd\x86\x7e\x2f\xb8\x94\x34\xac\x80\x5a\x8c\x00\xd7\x4e\x80\x11\x99\xe3\x2c\xdb\x9d\x91\xd7\x97\xcd\x0f\x43\x8a\x2c\xa7\xcc\x6c\x07\x7a\x1c\x13\x5f\x8d\x1a\xa5\x20\x34\xdf\x3e\x3f\xc0\xe5\xd3\xe1\xc7\x4e\x1a\xba\x0c\x79\xa3\x82\x1c\xbd\xbe\x98\xef\xea\xc5\x6d\x3e\xb6\x0b\x6c\x9e\xf1\xe4\x61\x53\xdd\x37\x3c\xde\xb0\xa2\x54\x1b\xbc\x28\xae\xb7\xb6\xbc\xcc\x14\x2d\x32\xa2\xb7\x3e\xaf\x83\x86\xbd\xc9\x6e\x3a\xbb\x5a\xda\x7d\x66\x20\x3a\x57\x8a\xe4\x85\xd2\x44\x98\x36\x8d\x02\xab\x96\x53\x5f\x17\x0d\x4b\x9f\xb9\x5a\x99\x39\xc7\x51\xca\x7d\x8a\x77\x53\xce\x15\x5c\xbf\xe9\xda\xa6\xd9\xb5\xea\xa7\xda\xef\x06\x4a\xf7\xd4\x3f\xa8\xb0\x05\xf7\xe9\xbe\x70\x86\x8d\x3d\x3f\x54\x62\x72\xf4\xf4\x12\x6f\x4d\x65\xf7\x0d\xee\xfb\xaa\xa1\xb5\xc5\xb3\xfb\xa8\x32\x62\xcc\xba\x77\x6c\x6e\x5a\x76\x1e\x4b\xef\x7c\x67\xcd\x86\xdf\xa7\xe6\xe9\x1a\xd3\x4c\x43\x59\x93\xa8\x82\xca\x89\xc2\x31\xa6\xf5\x3b\xdf\x61\x33\x0c\xcb\x0b\xc2\xf6\x06\xf5\xf4\xc0\x16\x50\xe3\x8f\xde\x17\xd4\x77\xe0\x0c\x83\x3e\x09\xfa\x03\xe4\xeb\x7b\x8a\xc3\xa8\xdf\x1b\x9c\xd7\x17\xe3\x0e\x61\x0a\xcd\x05\x7f\x20\x2c\x06\xf7\x3d\x74\x3a\xdd\x92\x9c\xeb\x2d\xca\x2a\x73\xca\xd9\xeb\xcb\x02\xd3\xac\x14\xa1\xf5\x81\x12\x2c\xdd\x3a\x96\x2b\x5e\x66\x29\x62\xe4\x51\x1f\x09\x92\xa4\x14\xe8\x08\xad\x08\x2e\x5a\x5d\xa1\x6e\x4f\xcd\x9a\xf9\xda\x6b\xf9\xbe\x87\xce\xc8\x1b\xf6\x88\x33\x9a\x22\xca\x52\xf2\xdc\xe3\x25\xdd\x4e\xb2\xf9\xfa\x67\x37\xca\x34\xfd\x05\x51\x6d\x84\x30\x9c\x65\x6b\xb4\x14\x98\xb9\x23\x0d\xb5\x60\xc1\x4d\xc3\xb6\x47\x19\x5f\xd2\xe4\xf5\xa5\x4d\x48\x8b\xab\x5d\xa7\xbc\x91\xe2\x9b\xd7\x17\x46\x9e\x5e\x5f\xea\xa3\x62\x04\x83\xdf\xec\xbd\x87\xe2\x68\x49\x1f\x49\x73\xea\x3c\x44\x29\x91\x85\x9e\xe2\x2d\xab\xca\xb8\x92\x2a\x43\x2d\xc7\xcf\xf1\xfc\x42\x57\x9b\xde\xbf\x25\xb6\xe7\x60\x47\x44\xc7\xd6\x85\xb3\x7a\xee\x1d\xe4\xab\x9e\xb7\x9a\xd1\x30\xae\x4e\xfc\x1d\x1e\x0b\x46\x5b\x67\x8e\x9a\xaf\xbb\x52\x69\x73\xe1\x77\x2e\x91\xc0\x2c\x64\x54\x9e\xa3\x47\x9c\x95\x04\x65\x44\x9a\x93\x34\xdb\xb4\xae\x0a\x73\x0e\xd0\x43\xa7\xfb\xb0\x4d\x9f\xb0\xac\x8c\x6c\x90\x89\xd6\x7c\xd9\x3e\xa9\x57\x66\xfa\xc6\x09\xf4\x8d\xc7\xec\x29\x9c\x59\x7b\x0c\x36\x73\xa6\xcf\x30\x68\x5c\x0d\x1d\x4f\x03\x17\x28\xe3\x38\x25\xa9\x19\x35\x5e\xaa\xea\xb2\xbe\xdf\x38\xa8\x95\x87\xdb\x61\xad\x9d\x61\x3f\xf3\xd9\xf0\x6d\x9c\x3e\x36\x9c\x93\xf6\x1a\x97\x59\xe0\x00\x5d\x71\xc0\xf3\x5c\x4b\xae\xe1\xa4\x20\x62\xc1\x45\xae\x15\x85\x1d\xc3\xd9\xd7\xbb\x2f\xd6\xd3\x0c\xd0\xd4\xa7\xbe\x51\xdc\x47\xdf\x25\x67\x6e\x6e\xf7\x68\xbb\x19\xd7\xcb\x47\xff\x4d\xa2\x1c\xaf\xdd\xc2\x48\x4b\x51\x1f\x3b\x04\x4f\x88\x94\xfa\x47\xbe\x68\xbb\xba\x0e\xed\x6c\xd0\x4b\xa6\x9c\x4b\xfd\x3b\xa6\xf4\x5e\x2f\xac\x22\xcf\xdd\xf0\x3e\x71\xf1\x80\x9e\x48\x96\xbd\x09\x1d\xdf\x34\x30\x5a\x70\x61\x49\x40\x2b\xcc\xd2\x4c\x43\xe1\x4c\x0f\xec\x72\x85\xa8\xaa\xc4\x66\x29\x33\xbc\x94\x92\x08\x64\x21\x13\xcf\x2e\x3a\xf6\x7d\xd2\xbd\xe2\xd1\x94\xdb\x11\xd4\x08\x7d\xde\x69\xdf\xac\x6e\x2b\x0d\xc6\x5b\xdd\x74\xbb\x68\x58\x9d\x55\x6d\x24\xca\x4b\xd9\xf1\x9c\x2d\xb8\x70\xb6\xa6\x66\xbe\xc7\x65\x15\xb8\x21\xe9\x65\xac\x63\x7e\xaf\xb0\x44\x92\xa8\x9a\x50\x86\xd8\xef\xec\x4f\x22\x12\x3e\x39\xeb\x8d\x4b\x29\x9c\xac\xcc\xf9\x5c\x16\x38\x31\x3b\x4f\x2d\xd2\x5e\xad\x45\x17\xe6\xca\xb3\xfe\x4a\x1a\x5d\x27\x0b\x92\xd0\x05\x25\x69\x35\x87\x3b\x63\xa3\xa7\xe6\xcf\xe4\xf9\x0d\x3a\xca\xd1\x68\xfa\xee\x97\xf6\x45\xc9\xed\x87\xfb\xe0\xdd\x96\x17\xe1\x34\x1a\x7b\xd7\x24\xe3\xa1\x6b\x92\x31\xf4\x9a\xe4\xdc\x1a\x42\xda\xa8\x33\x36\x96\xb4\x81\x5a\x11\xe6\xdc\x18\x7a\x5f\x92\xcf\x05\x4a\xb1\xc2\xd5\xe9\x40\x2f\x6a\x63\xc7\x46\x81\x02\xaf\x4e\x6a\x50\x9c\xa6\x7b\x22\x02\xaf\x50\x0a\x2c\xa8\x5a\x5b\x8f\xca\x5e\x62\x05\x5e\xa3\xb8\x39\x57\x96\x34\xdd\x1f\x14\x7a\x9b\x91\x92\x47\x9a\x58\xef\xc7\x82\x97\x2c\xe6\x9e\x68\x0c\xbd\x2a\xd8\x10\xa8\xb6\xda\xa3\xc0\x80\xee\x1d\x5f\x9a\xd1\x88\x40\x77\x66\x3d\x43\xf7\x13\xa6\x37\x5b\xfa\xd1\x1e\xc8\x7a\x4f\x30\xef\xb6\x2d\x0c\x86\x8d\xe9\x35\xe7\x5c\x21\x41\x12\x2e\xd2\x01\x58\xbd\xbb\xe1\xb9\xdb\x76\xab\x40\x3a\xdb\xc1\x86\x73\xb7\x51\xce\x76\x78\xbc\x26\x2d\x32\xbd\x20\x21\x28\x99\xce\x55\xd2\x13\xf6\xf1\x27\x90\xea\xbb\x85\x7b\x22\x38\xba\x7a\x3b\x76\x76\xfa\xee\xdb\xbe\x60\x98\x6a\x7a\xda\xbb\xa5\xb8\x78\x9f\xd1\xd8\x77\xaf\x0e\xe0\x69\x8b\x15\x47\x07\x17\x8d\xc6\x63\x68\x3c\xcc\x52\x5b\x2a\xad\x15\x0f\x97\x66\xbd\x95\x5f\x1a\x4d\x08\xdb\xcc\x27\xde\x66\x3e\x19\xda\xcc\x27\xd0\xcd\xfc\xc6\xc6\xf0\x22\x99\xe0\x18\x87\xcc\x04\xba\x83\xdf\x93\x68\x04\xe0\x76\xfd\x89\x4a\xe5\x36\x97\x28\x18\xe0\x1e\x6d\x60\xcc\x72\x7c\x7d\xd9\x03\x2d\xb8\x43\xff\x57\x58\xc9\x5c\x53\xa6\xf5\xca\xe3\xcf\x25\x4d\x7f\x89\x42\x03\x07\x59\xea\xa3\x74\xec\xda\x99\xf8\x81\xc7\x3d\x30\xce\x04\x30\x3e\xd7\x78\x34\x68\x58\xa5\x43\x4b\x32\x2e\xe3\xb5\xd0\xe4\xad\x17\x53\x39\x68\xde\xbc\xbe\x0c\xec\x57\x76\xe1\xb7\xee\x77\xec\x7d\x4c\x15\x59\xa2\x8f\x08\xfe\xd7\x03\xa4\x79\xd1\x97\x61\xd2\x3e\x73\x66\x20\x8c\xf7\xe5\x32\x3c\x77\xd1\x1d\xcb\xd6\x28\x25\xfa\xb0\x43\xd2\x5a\x76\xd6\xc3\xe2\x08\x04\x90\x04\x8d\x09\x4f\x31\xc9\x39\xb3\xa9\x12\x31\xa3\x02\x0d\x03\xaf\x71\x78\x11\x05\x03\xd4\x41\x4e\x5c\xf6\xac\x98\x13\xa6\x15\x92\x8a\xcc\x1d\x18\x4d\x7c\x33\x70\x08\x55\xab\xa3\x1f\x03\x0b\x3c\x3d\xb8\x13\x71\xbc\x12\xf4\x6d\xc1\x3e\xfe\xf6\x05\x82\x5a\x73\x6d\x8e\xe2\x55\x85\x1f\xf0\xdd\xa7\x2a\x7e\x0c\x1c\x50\x33\x95\xb5\xc7\xdc\x2c\x37\x94\x73\x46\x15\x17\x34\x14\x8b\x82\xb3\xac\xf5\x77\xb7\x7c\x24\xc2\xa2\x76\xd7\xbc\xbe\x88\x92\x31\xca\x96\x87\x88\x0b\x6d\xd6\xba\xe6\xb2\x2f\x72\xcc\xa7\x1b\xa8\xb6\xda\x74\xf3\x62\x1b\xd9\xaf\x2f\x3e\xdd\xaf\x2f\x2d\xc2\x4d\x2f\x05\x49\xa3\xe9\x06\xc7\x9f\xb7\x94\x4e\x41\x52\x54\x32\xf2\x5c\x98\x55\x99\xad\x7d\xd2\x21\x8d\x07\x68\x82\xc6\x25\xdb\xcc\x1b\xb2\xc7\x62\xf2\x83\x85\x7a\xa0\x04\xc9\x08\x96\xbb\x41\xd5\xb6\xb1\x71\xbd\x6f\x89\x03\x9e\x7a\x36\xf1\xfb\x21\x9b\x78\xea\xdb\xc4\x7d\x27\x25\x73\xda\x72\x57\xfa\x29\xae\x03\x11\xde\x44\xc8\x6b\xea\x5b\xc8\x10\x54\x13\x48\xb0\x17\x2c\xf4\xd4\x64\x61\x7b\x6e\x92\x76\x01\x84\x9e\x9c\x2c\xe0\x66\xa0\x69\x1c\x22\x34\x5d\xe9\x72\x7f\x97\xd3\x14\x6c\x43\xab\x15\x69\x85\x94\x4a\x99\x36\x6e\xea\x2a\xc2\x70\x6c\xdc\x89\xad\x73\xed\x2e\x74\x40\x8d\x6c\xf6\x98\x57\x5b\xcb\x5e\xe7\xef\x29\xd8\xd0\x9e\x49\x17\xd9\x54\xf1\xbb\x34\xde\x5a\x61\x6f\x01\xff\xfd\x1f\x34\x5f\xab\x40\xb8\x2e\x84\x04\x68\x42\xd3\xac\x2d\xee\xfa\xa6\x54\x13\x15\x87\x0b\xcd\x6b\xba\xfd\x70\xff\xfa\x62\xc2\x39\xa2\xc5\xec\x9b\xcc\xfd\x58\x2e\x88\x23\x1e\x0b\xaa\x93\xcc\x9d\xe1\x91\xe2\x19\x11\x98\x25\xc4\x28\x56\xb4\x27\x9f\x50\xc5\xf4\x2f\xc1\xd9\xd2\xa3\x20\x27\x6a\xc5\x53\xa4\xd6\x45\xcc\xee\x35\xf5\x6d\xea\x1e\xf4\x83\x7f\xff\x07\x7d\xc1\x42\x51\x73\x1b\x53\x5f\xcb\x18\xb6\xfd\x9c\x74\x08\x32\x54\x5b\x35\xc8\x9c\x99\x3b\xe4\x7d\x40\xa1\x5a\xcb\x08\xfb\xf5\xc5\xea\x66\xf2\x68\xb2\x9f\xa3\xf4\x24\x38\xbd\xd2\x85\x9e\xd8\xc8\x04\x9c\x21\x9c\xa6\x82\x48\xb9\xc7\xc4\x82\x2a\x09\x3b\xb1\xf4\x0c\x32\xb7\xba\xd6\xb7\xda\x63\xa5\x98\x40\x0e\xd3\xf4\x67\xdd\x76\x5e\x2e\x16\x5a\xb1\xdb\xfb\xe0\x14\x2b\x7c\x24\x15\x17\x78\x49\x7e\x69\xdd\xea\xcd\xd7\x46\xf7\xb4\x3b\xae\xef\xa1\x71\xa2\x4a\x9c\x55\xbf\x35\x3d\x1b\x7b\xac\x8a\x73\x0e\xdd\x40\x37\x11\x10\xb6\x7d\x5f\x24\xea\xd4\x37\x4d\xc3\x27\x2a\x73\x85\x5d\x59\x66\xa8\x39\x37\xc6\x88\x1d\x1a\x09\xcf\xb4\xe2\xc8\x31\x65\xe6\xc2\x73\x3f\x23\x74\x0a\x8d\x58\x67\xdc\x8d\x42\x38\x4b\x06\x82\x04\xbc\x8b\x61\x84\xa4\xee\xe2\x7d\x41\x85\x34\x39\x7f\x7a\x9e\x58\x9f\xc9\x1e\xac\x42\xf3\x26\xbd\x19\xb7\xc2\x12\xcd\xb5\xb9\x11\x99\x7d\x38\x9a\xfa\x9e\xfe\x5d\xa1\xbb\x1e\xa3\x5d\xc0\x81\x43\x6c\x6f\x5a\xcc\x15\x7d\x9a\x7a\xa4\xf8\x0b\xda\x23\x36\x2e\x43\xc1\xa7\x17\x38\x51\x1a\x7a\xdd\xd4\xd8\x4e\xb2\xd6\x41\x6d\x62\x37\x73\xb9\xaa\x74\x0c\x13\x12\x1e\x5e\x59\x03\x54\x43\x9d\x3d\x2d\x02\xad\x1f\x20\xc9\x08\x16\x91\x43\x0b\xdd\x95\x36\xc7\x76\xeb\x6a\xaa\xae\xc9\x02\x61\x32\x03\xc4\x40\x37\xac\x5d\x89\x29\xd9\x03\xe3\x4f\x6c\x63\xe4\x9e\xcc\xd6\x13\xbe\xea\xf3\x29\x83\x1a\xde\x9b\x94\x0d\x4d\xfc\x66\xd6\x38\xaa\x5a\xc7\x4d\x17\xe5\x65\x64\x08\xa0\x2e\x10\xac\x3e\x30\x73\xa4\xc2\xaa\x94\xa8\x2c\xd2\xc8\xdc\xda\xe9\x14\xa8\x8c\xac\xeb\xc0\x65\x81\x9f\xa1\x2b\xbd\x7e\xd1\x67\x2e\x72\x9c\x1d\xf8\x9d\x02\x1d\xbb\xc1\x4e\x2f\xc9\x52\xe0\x94\xa4\x81\x6e\x81\x9e\xdb\x60\xb7\xb7\xd4\x44\xb2\x05\x7a\x05\xae\xd5\x60\xaf\x1f\x4c\xcc\x75\xa0\x53\xa0\x13\xb6\xd3\x69\xaf\x40\x81\xc5\x67\x3a\xdd\x0d\x88\xd2\x5b\x04\xa0\x0e\xef\xc9\xbc\xa4\x59\x1a\x96\xa3\x67\x28\x82\xba\x9c\x29\x5e\xf8\x9d\xf9\xb1\xb3\x5b\x17\xa9\xcb\x1f\xd9\x58\x67\x1d\x15\xd6\x8e\x89\x33\x1b\x69\x7f\xd4\x68\x9d\x8f\x62\x3e\x34\xb6\x87\x4f\x24\xf4\xbc\xd7\x10\xb9\xd1\x6d\x1f\x99\x55\x44\x75\x2f\x71\x00\xd2\xa0\xc7\x41\xa0\xfc\x56\x74\xb9\x32\x15\x86\x28\x37\xb1\x3f\xbf\xf1\xb9\x26\xb3\x2a\x4b\x40\xd9\x32\x64\x61\xab\x4e\xe8\x2d\x46\x4f\x2b\x9a\x79\xc1\x16\xd3\x13\xe8\xf1\x11\x2c\xc8\x10\xbd\xb2\x4c\x56\x08\x4b\x24\xea\x49\xbc\x17\x0b\xb5\x53\xd5\x2d\x8a\x2d\xa1\x06\x27\x9e\x5b\xf5\x64\xc8\xad\x7a\x02\x0d\x35\x68\x96\xe4\x25\x99\x97\xcb\x8c\x2f\xfd\xae\x80\xea\xb8\xe9\xaa\x72\x98\xfb\x5d\x01\x55\x70\xab\xab\x8d\x98\x9d\x56\x4f\x83\xa5\x31\x90\xdf\xa5\x1b\xb6\x7a\xc0\x90\x20\x45\x93\x86\xdc\xea\x18\x5c\x1b\x6c\xb3\xbf\x88\xdd\xf3\x04\x5a\x14\xcc\x1d\x62\x13\x93\xe1\x5a\x18\xfb\x20\xe3\xc9\x43\xd4\x96\x7d\xe2\x7b\x0e\x41\xa0\xaf\x2f\x54\xa2\x92\xc5\xe3\x42\x6f\x9f\x2b\xc1\x26\x3c\x2f\x4c\xd4\x96\x0b\x7f\x5c\x94\x59\xe0\x8a\x06\x02\x0c\xd5\xb3\x07\x15\xb4\x20\xb2\xcc\x54\x43\x8a\x4b\x2d\x8a\x71\xf6\x9c\x40\x03\x15\x7b\xc1\xdd\x5a\x8a\x03\x87\x9e\x93\x79\x0d\xa7\xb0\x58\x92\x40\x46\x82\x5a\x75\x5c\x13\x36\x57\xc9\xa6\x25\x64\xdc\xa6\x4f\x62\xb6\x46\x45\xe5\x9f\x83\xd0\x07\x34\x9e\x18\xb7\x47\x79\x55\xd3\xe9\x13\x68\xab\xed\x38\x92\x52\x67\xba\x40\x68\x80\x9e\x7b\xae\xeb\xdd\xc3\x5e\xd5\x3a\x4a\x0e\xab\x88\x33\x53\x5d\x92\x17\x9d\x64\x0c\xd8\xcd\xda\x05\x67\x0b\xba\x84\x45\x9d\x9d\x7a\x5b\xc1\x50\x6d\xcf\xba\xf9\x56\x21\x27\x86\x06\xb4\xa0\xa6\x72\x1d\x4e\x51\xca\x59\x8c\x73\xe4\x14\x1a\x7e\xb6\x24\xca\x65\x0c\x45\x23\x01\xb7\x26\xeb\xd4\xb4\x58\xf1\xd1\xd4\xa7\xd0\xed\xab\x0d\x17\x19\xd0\x79\x1a\x88\xdc\xee\xcb\x86\xa8\x42\x56\xdd\x00\x1a\x77\x64\xac\xc3\xf6\x34\x10\xc5\xbd\x0d\xf7\x81\xac\xe3\xf1\x02\x17\xf7\xe1\xe4\xbd\xf6\xf4\x34\x41\x6a\xb1\x92\x0d\xd4\x8d\xdb\x8a\xe8\xb2\xfc\xe3\xb9\x84\xa6\x64\x6e\x60\xee\x75\xa7\x73\x1a\xa8\x22\xb7\x5d\xb2\x7b\x86\xea\x9e\x8e\xa0\x39\xc8\xbf\x49\xce\x50\xca\x93\x4a\x63\xf3\xf9\x6f\x24\x09\xec\x3b\xd5\x24\x6b\x94\xc5\xdf\xec\xfa\x32\x09\xe9\xe6\x17\xd6\xb4\xde\x9c\x18\x87\xad\x0b\x31\xf4\xb7\x30\x5b\x41\x4d\x7c\x8b\x8b\x62\x5b\x85\xba\xb1\x57\xec\x6c\x7c\xdc\x51\xc1\x4d\xcc\x84\x0d\xf0\xe6\xdb\xba\x1c\x79\x5d\x76\x13\x83\xea\x2e\x3f\x5e\x0c\xef\x10\x63\x2f\x2e\x79\x3c\x14\x97\x3c\x06\xc7\x25\x7f\xbc\x40\x4a\xd0\xe5\x92\x44\xb9\x18\xc7\xe0\xb8\xe4\x8f\x17\xf1\x75\x9a\xc7\xe0\xd8\xe4\x8f\x17\xaf\x2f\x71\x1b\xcf\x18\x1c\x98\xfc\xf1\xa2\x2e\x78\x15\x19\x4e\x39\x06\x87\x6d\x7e\xa7\x89\xa2\x79\x6d\xae\x27\x9c\x49\x25\xca\x44\xc5\x2c\xe4\x31\x38\x88\xf3\x13\xc7\xda\x66\x7d\x24\x42\x12\x94\xe3\x88\x48\xce\x31\x38\x92\xd3\x60\xd9\xcd\xd5\x94\x46\xd9\xb5\x9e\xe1\x2d\x51\xf8\x7a\xb6\xb1\x78\xa6\xdd\xc5\x33\xf1\xd6\xf6\x64\x68\xf1\x4c\x02\x85\x0c\x7b\x42\xaf\x6e\xed\x3f\x9a\x3b\xce\x54\x9f\xb8\x97\xa1\x98\xb9\xad\x32\x9b\x04\x72\x55\xc3\x32\xf3\x50\x35\x17\x5a\x19\xb6\x2d\x74\x38\xee\x28\x90\x23\x1b\x3c\x52\xdc\xd4\x75\xc8\xd1\xed\xf5\x6c\xa3\xfe\xc2\x6e\x78\xc0\x9b\xcc\x5b\xe3\xcf\xd9\x0f\x0a\x78\x54\xfb\xe6\x9c\x47\xfb\x81\x01\x8f\x66\xcd\xa9\x83\xb0\xdf\x4b\x52\x12\x53\x7a\x3d\x97\x81\x40\xcb\xde\xa6\xfa\x57\xb7\xd7\xb3\x43\x9b\x5a\xed\x0e\x52\xe4\xb9\xc0\x2c\x45\x0b\x41\x6c\xb1\xf1\x7f\x42\x88\x06\x66\xa5\x5d\xf0\xbc\xc0\x89\x1a\x28\x98\xdc\x6e\x92\xf0\x32\x4b\xd9\xdf\x4d\x78\x90\x56\xc8\x28\x2d\x4d\x64\xa9\x09\x44\x63\x26\x99\xdb\x50\x69\x92\x75\xff\x1c\x2a\x63\x48\xa8\x01\x03\x49\xba\x7d\x51\x42\xd7\x33\x6d\xef\x44\x17\xe3\x9c\x04\x92\x73\x87\xa1\xf6\xa8\xfc\x39\x09\x24\xe5\xf6\x81\x75\x95\x4c\xa4\xfd\x38\x09\x64\xe5\xf6\x95\x3e\x76\x91\x2a\x05\x16\x38\x27\xad\x0a\x7a\xbb\xc1\x41\xe3\x7f\x98\x49\x26\xc4\x69\x84\x07\x6a\x32\x06\x87\x29\x56\x37\x0c\xf1\x48\xbb\xdf\x9b\x36\x0f\x46\x44\x01\xee\x7e\x1d\x3a\x37\x51\xe0\x65\x84\xa5\x30\x19\x83\x63\x0e\x19\x47\x39\x49\x29\x36\xb3\xf1\xf6\x7a\x16\x05\x06\x0d\x34\x74\x35\xe4\x36\xaa\xf4\xcd\x69\xcc\xb6\x1e\x48\xd8\xde\x2a\x4f\x7d\xe6\x89\x82\x82\xaa\x12\x5a\x17\x3f\xda\xc7\x72\x08\xa4\x6c\x6f\x65\xcd\x55\x16\x58\x34\xb7\x3e\xbb\x21\xee\x7e\x31\xb4\xa0\x2e\x24\x2c\x1a\x13\xaa\x51\x2a\x90\xfa\x00\x5c\x90\xa8\x27\x30\x26\x81\xf4\xed\xbf\x80\x4d\xa8\xa2\x51\x82\x36\xd3\xb4\x61\x3a\xd6\x6e\x1a\x83\x1f\x6e\x68\xcd\x22\x93\x64\xb8\x17\xb7\x50\xb5\xd3\x92\xb0\xf5\x38\xec\x85\x0a\xd5\x3f\x0d\xaa\x4b\x74\xdf\x07\x15\x9c\x49\xd3\x9e\xc0\xd5\xc6\x15\x7b\x3d\x34\x19\x83\x5f\x73\x68\x87\x6a\xe1\xfd\x26\x71\xc0\x2b\xb7\x0d\xd3\xc6\x78\xbb\xea\xe6\xd1\xb8\x50\xbd\xf4\x24\xb8\x6a\x6a\xa6\xda\x3a\x69\xca\x5c\xbd\x6c\xd4\x2a\x08\x57\x94\xd6\x3b\x9f\xc0\x39\x4a\xa9\x7c\x80\x10\x05\x55\x5c\x8d\x30\x24\x21\x0f\xfd\x82\xd8\x99\x00\xa8\x1a\x33\xe3\xfe\x17\x09\x05\xaa\xe6\xfe\xd2\x91\x82\xaa\x40\x17\xf9\xda\x73\x02\xdb\x15\x15\xaa\x03\xab\xd8\x8a\x1f\x85\x0b\xd5\x82\x95\xc9\xb7\x37\x22\xb8\x78\x87\x67\xa8\xe4\x44\x61\xe4\x62\xcf\x63\x14\x03\xb8\x8a\x47\x77\x73\xdb\x1b\x18\x5c\xce\xa3\x05\x54\xcd\xf4\xe8\xd3\x0a\xb8\xae\x47\x58\xd0\x3d\xa5\x75\x21\xc0\xbb\xeb\x3a\x63\xbe\xec\x0b\xbb\xbb\xa1\xd6\x1a\xde\x3d\x70\xc1\xe9\x69\x8d\xbb\xa3\xae\xa7\x66\x25\x1e\x05\x0b\xd5\x53\x38\xcb\xb9\x54\x68\x51\x06\xca\x88\x43\x70\xa0\x9a\xa9\xb2\x52\x8c\x50\xa3\xac\x94\x31\x54\x19\x99\xb7\x38\xb0\xc2\x19\x5f\x86\x4a\xba\xee\x00\xe9\x57\xfa\xee\x7b\xa9\xcc\x39\xfd\x62\x1f\xb7\x9a\x8c\xfd\x12\xdd\xc3\xf3\xd3\x6c\x6d\x8f\x58\x50\x5e\x4a\xad\x05\x64\xdc\xc1\x70\xb2\xbb\x19\x26\xf1\x23\x41\x94\xf1\x34\x6a\x62\x4e\x76\xd7\x3b\x05\x2f\xec\xe3\x32\x18\x19\x29\x47\xc1\xee\xae\x75\x8a\x52\xae\x4c\x60\x4d\x34\x6a\xe0\x72\xef\xe2\xfe\xe6\xeb\xcd\xc5\xf9\x27\x1f\xf8\xa2\x99\xb1\xbb\x94\x83\xae\x2f\x58\x66\x6b\x79\x64\xf2\xd4\x36\xee\x58\x4e\xba\x57\x2c\x5e\x8e\xf8\x64\xfa\x7e\xe0\x8a\x25\x90\x23\xde\x73\xc5\x62\x55\xe4\xaf\x86\x84\x08\x59\x05\xd2\x3d\x7b\x80\x6c\x24\xd4\xaf\x38\x8d\xc7\xf2\xcf\x59\x43\xaf\x96\xfd\xea\x9e\x11\xdb\x05\xaf\x1e\x96\x4f\x7c\xe9\xbf\xba\xd6\x1d\x93\x13\x6f\x4c\x86\xae\xbd\x02\x21\xd6\x3d\x99\x35\x34\x53\x44\x20\xb9\x66\x0a\xf7\x55\xce\x06\x48\xcb\x0f\x96\xee\x8d\x64\x73\x88\xf6\x68\x6a\x5d\x0e\x05\x56\xab\x28\x54\xe0\x8d\x50\x93\xc1\x91\xf1\xe5\x91\x69\x6a\xcf\x47\x6a\x67\xe5\x5f\x0f\xda\x3f\xf9\x2c\xfc\x5c\x5e\x67\xe0\xde\x79\x8b\xe9\xdd\xd0\xc0\xbd\x83\x5e\x1c\x6a\xe2\x6d\xce\x69\xc1\x33\x9a\xac\x37\x1e\xf5\xbb\xfb\x82\xd5\xea\x88\x3d\xe6\x8b\xe1\x50\x84\xa9\x77\x9b\x3a\xed\xbe\xe8\xdc\xed\x74\x21\x4c\x41\xfa\x2d\x01\xd1\x53\xef\x91\xc0\x69\x6f\x08\x86\xeb\x78\x8e\x93\x87\xed\xfd\x7a\x71\x18\xd3\xee\x3b\x76\x31\xa5\xe2\xa6\x9e\xca\x9b\x0e\x2d\xaf\xe9\x14\x7c\x18\xde\x5a\xb9\xd3\x15\x5a\x18\xac\x5e\x85\x66\x09\x66\x1b\x49\xb1\xad\x92\xf9\x35\xbf\x36\xce\xb1\xcb\x68\x87\xd3\x53\x6f\xc4\x4f\x87\xc2\x13\x4f\xc1\x0f\x01\xce\xbe\x9c\x5f\x5c\x6d\xfe\x0d\xb6\x90\x2e\x8f\xce\x97\xdd\x2d\x49\xd3\x1d\x78\x7c\xbc\xfb\xee\xe4\x20\xe9\xfa\x03\xe8\x8e\xfe\x8d\xa5\x64\x41\x19\x49\x37\xc2\xa1\xda\x3d\x41\xc7\xfb\x1f\x92\x33\xf4\x75\x5d\x90\x4e\x60\x55\x37\x31\xb9\xba\x93\xf8\xc0\xd3\xb5\x69\xef\x63\xfa\xc9\x62\xbd\xf7\x1b\xd5\xc3\xa9\x36\x9e\xca\x18\x42\x5f\xee\x66\x5e\x97\xc7\xe0\xeb\x8b\x6f\x0c\x97\x6a\xc5\x05\xfd\x83\xa4\xe8\x9b\x24\xfd\x8c\x9c\xbb\x76\x36\xd5\xff\xbf\x09\x4e\x89\x2f\xbf\x63\xb0\xd7\xd2\xc8\xc3\x08\x71\x58\x7e\xa6\xdd\x17\xbc\x36\xb9\xfe\xd6\x99\xeb\xa3\x82\x3d\x05\xff\x73\xe4\xb6\x88\xa3\x9b\xd4\xf1\xb0\x05\x7f\xe3\x8b\x9f\xcb\x92\xa6\xbf\xa0\xef\x38\x2b\xfd\x71\x0c\xc4\x2d\xf4\x65\x55\xd8\x29\x71\x6e\x9e\xdb\x72\x45\xf9\x0b\x2c\xa5\xb5\x35\x03\x03\x1a\x7a\x61\xbf\xaf\xc8\x87\x7b\xcf\xb7\xa7\x5a\xbc\x7b\xaf\x5d\x0f\xa1\xfb\x5d\x5d\x78\xe5\xcb\xdd\xcc\x50\x62\x15\xa9\xa9\x94\x36\x53\x38\x09\x78\x6a\x90\x20\x36\xf4\x3a\x48\x2a\x78\xee\x05\x46\xbf\xa6\xf3\x16\x67\xae\xbe\xfe\x3f\x64\x28\x84\x5d\xcf\x55\x17\x82\x64\xba\x09\xcf\x8c\x1d\xde\x43\xa6\x1f\xf8\x02\x5d\x34\xd2\xe9\xa1\xeb\x2f\x90\x1f\xbc\x66\x2d\xe3\x6a\x65\xde\x54\x74\xa5\x72\x9b\x5c\xd5\xde\x6c\x18\x84\x3b\x1f\x19\xbb\xa8\x6e\xff\x06\xfd\x0b\x53\x65\xe3\x93\xd4\xdf\x65\x95\xfc\xd1\xdc\x34\xb6\xe8\x04\x1f\x88\xcf\xa5\xe4\x09\x35\x35\xe8\xb5\x8c\x12\x9c\x65\xbd\x4e\xc9\xaa\x81\x3e\xc7\xaa\x52\x68\x4d\x6d\x75\x9d\xa9\x2f\x21\xab\x5c\xd9\x80\x4c\xbb\x0f\x44\x48\x9e\x13\xa4\xa8\x57\x14\xf8\xf8\x18\xae\x70\xbf\xb7\x5f\x48\xa9\xde\x7f\xcc\x68\x1e\x78\x96\xdf\x0e\x7f\x96\xf1\x27\x89\x38\xcb\xd6\x68\x34\x7d\xd7\xbc\x99\x59\xb9\x6f\xde\xa0\xc2\xd6\xfd\x72\x95\x8b\x03\xef\x41\x76\x79\x72\x19\xf8\xe6\x8d\x36\x22\xaa\x68\x76\x2e\x2c\x55\x5d\xe6\x42\x0f\xa7\xf6\x30\xc7\x99\x79\xdc\x8c\x7e\xe0\xd7\x48\x12\x51\x65\x3c\xa7\x04\x77\xd3\xbc\x8e\x43\x6f\x53\xf6\xf4\x9a\xd8\x37\x5a\xb4\x8d\x43\x59\x8a\x6c\x32\x0b\x6d\x5e\xea\x6a\xf5\xe9\x3f\x36\xd2\xd3\xa7\x11\xed\x50\x4e\x65\x3d\x08\xdf\x9d\xc0\x3b\xef\x93\x9a\x01\xb1\xd3\x42\xf7\x15\x08\xfd\xe9\x66\x65\x06\x56\xe6\x18\xae\x84\xef\x89\x2c\x38\x93\x76\xfe\xf1\x52\x6d\x04\x3f\xc2\xcc\xa0\x91\x6f\x06\x8d\x06\xcd\xa0\xd1\x5b\x70\x69\xf4\xcb\xe1\xb2\xe8\xed\x44\x78\xdd\xd2\x09\xbd\x10\x44\xb6\xcf\xbc\x8d\xc8\x92\xea\x39\x07\x93\x9c\xe4\x3e\x69\x3a\x31\x75\x64\x58\xf3\xb0\x2b\x61\xaa\x79\x38\xb5\x26\x7f\x87\x3d\xee\x9f\x25\x11\xeb\xb8\x4d\xee\x86\x2d\xb2\xf2\xf9\xf2\x83\x59\x7a\x66\x20\x06\xd4\x73\xd5\xd8\x27\x15\xbc\xc8\xbe\xd2\x9c\x34\xb1\x51\x7d\x04\x57\x41\x54\xca\xb4\x26\x82\xf2\xb4\xa9\x98\x13\x20\xb0\x74\x53\xcb\x5a\x80\xc7\xf9\xe1\x34\x3f\x3c\xd6\xff\xaf\x0e\xdf\xad\x0e\x8f\x47\xab\xc3\xd1\x64\x75\xf8\x3e\x3d\x1c\xbf\x4d\xdb\x73\xef\xf3\xf7\xdb\x2b\x84\xd3\x9c\xb2\x2d\x91\xeb\xfe\xe4\x1b\xbf\xed\x1e\x18\xdb\x22\x69\x3e\xd8\x2e\x12\x43\x44\xa2\xb2\x5e\x69\x5c\x60\xf6\x77\x9b\x68\xc0\x1e\xb5\xf6\x55\x22\xeb\x5a\x9a\x1a\x0f\xea\x24\xfd\xc4\x97\xa8\xc0\xcb\xbe\xac\x91\x16\x5c\xe6\x5a\x06\xc0\xa0\xae\xd1\x8b\x7a\xc3\x1c\x86\x73\x8f\xf2\xb4\x36\x58\xe9\x0d\xd4\xf9\xff\xf3\x40\xfd\xe0\x71\xf8\xb1\x72\xee\x95\xe2\xff\x05\x00\x00\xff\xff\xbb\xfc\xb7\xcd\xd5\x8b\x00\x00")

func ResourcesEventsYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesEventsYaml,
		"../resources/events.yaml",
	)
}

func ResourcesEventsYaml() (*asset, error) {
	bytes, err := ResourcesEventsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../resources/events.yaml", size: 35797, mode: os.FileMode(420), modTime: time.Unix(1618193862, 0)}
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
	"../resources/events.yaml": ResourcesEventsYaml,
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
	"..": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"events.yaml": &bintree{ResourcesEventsYaml, map[string]*bintree{}},
		}},
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
