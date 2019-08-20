// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 8, 20, 8, 33, 1, 555538348, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(2019, 8, 8, 13, 51, 12, 390915853, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(2019, 8, 20, 8, 33, 1, 554377526, time.UTC),
			uncompressedSize: 6456,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5f\x6f\x1c\xb7\x11\x7f\xd7\xa7\x18\x9c\x1f\x6c\x03\x77\x7b\x52\x94\x04\xc5\xa6\x0a\x90\xd8\x89\xea\x3a\x96\x05\xab\x6e\xd1\xa7\x9a\xc7\x9d\xbb\x25\x8e\x4b\x6e\x39\xe4\x5d\x16\x42\xbe\x7b\x31\xe4\xfe\xe1\x4a\x27\x3b\xf0\x5b\xfd\x60\x49\xdc\xf9\x3f\xc3\xdf\xcc\x70\xb5\x5a\x9d\x89\x56\xfd\x13\x1d\x29\x6b\x4a\x10\x6d\x4b\xeb\xc3\xc5\xd9\x5e\x99\xaa\x84\xd7\xd8\x6a\xdb\x35\x68\xfc\x59\x83\x5e\x54\xc2\x8b\xf2\x0c\xc0\x88\x06\x4b\xd8\xea\xf0\xfb\xfd\x3d\xa8\x2d\x14\x37\xa2\x41\x6a\x85\x44\xf8\xe3\x8f\xfe\x7b\xfc\xb3\x84\xfb\xfb\xf9\xd7\xfb\x7b\x40\x53\x31\x19\xb5\x28\x59\x98\xc3\x56\x2b\x29\xa8\x84\x8b\x33\x00\x42\x8d\xd2\x5b\xc7\x5f\x00\x1a\xe1\x65\xfd\x9b\xd8\xa0\xa6\x74\x90\xeb\x66\x6a\xef\x84\xc7\x5d\x97\x3e\xfa\xae\xc5\x12\x3e\xa0\x74\x28\x3c\x9e\x01\x78\x6c\x5a\x2d\x3c\xf6\xc2\x32\x0f\xf8\x9f\x30\xc6\x7a\xe1\x95\x35\xa3\x70\x80\xd6\xd9\x06\x7d\x8d\x81\x0a\x65\xd7\xad\x75\xbe\x84\xc5\xe5\xf9\xe5\xc5\x02\x9e\x81\x47\xad\x33\x0a\xf0\x16\x48\x3a\xd1\x22\xac\x1b\xf4\x4e\x49\x62\xe7\x5a\xab\x8c\x7f\x4e\xc0\xcc\x45\x2f\x58\xcf\x7c\x78\xe0\x05\xc0\x10\x8b\xf8\x3b\xba\x83\x92\xf8\x93\x94\x36\x18\x7f\x33\x27\x04\x38\x58\x1d\x1a\x1c\x45\xad\x7a\x51\x3b\xe5\x57\x7b\xec\x46\x05\xc4\x51\xf0\x93\xc2\xe1\x64\x92\xb7\x62\x96\x2a\x26\x38\xa3\xaa\x70\x2b\x82\xf6\xef\x6c\x85\x25\x9c\x7f\x7b\x7e\x0e\xcf\xe0\x58\xa3\x81\x86\xad\xc1\x0a\x1c\x8a\x6a\x65\x8d\xee\x96\x70\x44\x38\x5a\xf3\xdc\xc3\x06\x41\x6c\x34\x72\x3c\x64\xdd\xd8\xea\xac\x17\xf8\x0c\xfe\x51\x2b\x02\x45\x20\xc0\x37\xed\x96\x20\x10\x56\xb0\xb5\x0e\x76\x68\xd0\x09\xaf\xcc\x0e\xee\xee\xfe\x06\x7b\xec\xa8\x80\x37\x06\xde\xfe\x85\xe0\xc7\x2b\xb8\x28\x2e\xce\x97\xa3\x94\x41\x77\x72\x81\x40\x38\xcc\xed\x20\xcb\xa6\x18\xc4\x0a\x04\x10\xb6\x82\x8b\xa2\x0f\x14\x1c\x71\x14\x23\x85\x81\xa3\x53\x9e\x0d\x2d\x4e\xc7\x6f\x87\x66\x0c\x06\x36\xad\xef\x5e\x2b\x97\x07\xb1\xc1\x4a\x85\xa6\x84\x77\xd8\x58\xd7\xe5\x7e\x22\x6c\xad\xd6\xf6\xc8\x1e\xf5\xaa\x15\x45\x57\x03\xf1\x99\x00\x19\xc8\xdb\x46\x71\x04\xf6\xc6\x1e\xcd\x7f\x6a\x4b\x9e\x46\x11\x5b\xa5\x71\x09\xc7\x5a\xc9\x1a\x3a\x1b\xe0\xa8\xb4\x4e\x4e\x79\x0b\x95\xe5\x7b\xc6\xc7\xcc\xc4\xbf\x38\xb0\x47\xc3\x66\x8f\x02\x1c\xb6\x16\x9c\xf0\x35\x3a\xf0\xb5\x30\xbd\xe2\x9d\xf2\x75\xd8\x80\xe5\x43\x04\xad\xf6\x58\xc0\xbf\x6d\x78\xae\x35\x08\x4d\x76\x50\x31\x0f\x36\x28\x0f\xca\x78\x1b\x79\xa4\x35\x5e\x28\x83\x6e\x09\x1b\xd4\xf6\x58\xc0\x1d\x4e\x51\xad\xbd\x6f\xa9\x5c\xaf\x2b\x2b\xa9\xe0\xc2\x92\x15\x5f\x1d\x34\x6b\xbe\x7a\xe4\xd7\xbb\xa0\x2a\xa4\x75\x20\x5c\xb5\x4e\x1d\x84\xc7\x58\x7a\xec\x48\x51\xfb\x46\x8f\x92\x86\x5c\x10\xd5\x2b\x69\xcd\x56\xed\xc6\x4f\x00\xe9\xe0\x9d\x68\xcb\xec\x30\xbf\x48\xab\x8c\xed\x6b\xf3\x52\xec\xc3\x06\xd7\x49\xc8\x54\x7e\x5f\xcc\xc9\x51\x51\xcd\x27\xb5\x38\x20\x08\xa8\xd4\x76\x8b\x8e\x41\x73\x90\xd0\xdf\xaa\x09\x18\x63\x0a\x92\xb8\x3c\x09\x0c\x2e\x07\x55\xe1\x10\xf6\xad\xda\x35\xa2\x9d\x0c\x51\xbe\x06\x61\x00\x8d\x77\x5d\xf4\xe1\x53\x22\xfa\xb4\x04\x61\x2a\x08\x46\xda\x86\xd1\x3a\xf2\x27\x6f\xdf\xc5\x74\x0a\x53\x8d\x52\xd0\x1c\xa2\x04\x85\xd4\xe7\xf3\x51\x06\x38\x0c\x5f\x91\x81\x8c\xed\x8b\x19\x88\x48\xe0\x2d\xa8\x86\x71\x12\xae\x6f\xaf\x23\x08\xc0\x0b\x76\x8b\xd4\xce\x28\x33\x29\x67\xe7\x0e\xe8\xd4\x56\xc9\x08\xd8\xd0\x06\xd7\x5a\x42\x7a\xf9\x27\x02\x39\x4a\x49\xf0\x91\xa2\xc8\x01\x62\x7d\x7f\x22\x70\x20\xdc\x6e\xba\xa6\x4f\x44\x6c\xd7\xee\x18\x3f\x28\x0b\xcd\x1c\x82\x9f\x3d\x01\xc2\x8f\xf9\x4e\x80\xf0\x10\xce\xf1\x26\x3e\xc2\xff\xac\x43\xf4\x51\x77\x18\x71\xd2\x58\x58\x94\xe9\x26\x2e\x40\x35\x62\x87\xa9\xfa\x99\xa1\x80\x5f\x95\xa9\xa2\xcf\x0d\xc3\x8a\x43\x39\x55\x6d\x82\x14\x8d\x82\x90\xc1\x23\xb2\x72\x12\x78\x4e\x00\xe1\xc7\x7b\x5f\x87\x4d\x51\x59\xb9\x47\x57\x48\xdb\xac\xdd\xfa\x88\xe2\x80\x47\xeb\xf6\xb4\x66\x25\x6b\x2f\xc6\xf0\x0d\xb9\xe4\x9e\xcf\xf3\x00\x6b\xf6\x62\x07\x6c\x6d\x31\xd2\x44\x55\x25\xf4\x42\x95\x5d\x27\x54\x89\x3f\xca\x8b\xe2\xe2\xb2\xb8\x9c\xd3\xde\x06\xad\x6f\xad\x56\xb2\x2b\xe1\xcd\xf6\xc6\xfa\x5b\x87\x94\x7b\xe2\x90\x6c\x70\x12\x29\xc7\x72\x87\xff\x0d\x48\x7e\x76\x06\x20\xdb\x50\xc2\x77\xe7\xcd\xec\xb0\x89\x70\x5f\xc2\xf7\xdf\xbe\x53\xd3\xa8\x60\x5d\xce\xbc\x9a\xb2\x73\x1b\xc7\x86\xcb\xf3\x4b\xee\x9e\xca\x6c\xad\x6b\x62\xd9\x0a\x3d\x52\x6b\x75\x40\x83\x44\xb7\xce\x6e\x30\xb7\x80\xc3\x7a\x3d\xef\xdc\x49\x55\x12\x38\x3f\x16\xbe\x2e\x61\x2d\x5a\x95\x22\x7d\xf8\x7e\xad\x2a\x34\x5e\xf9\xae\x68\xc3\x26\xa3\x55\x46\x79\x25\xf4\x6b\xd4\xa2\xbb\xe3\x3b\x5a\x51\x09\xdf\x65\x04\x5e\x35\x68\x83\x3f\xf1\x8d\x1b\xad\xfa\xff\x30\x35\xbb\xb8\xb3\xc4\x9c\x1e\x91\x20\xb5\xba\xdb\x64\x19\x7a\x19\x2d\xab\xd6\x44\x35\xcf\x7a\x36\x4d\x9f\xa0\x6d\x8f\x39\x3b\x4e\x19\x28\x93\x6a\xee\x39\x25\x1e\xa2\x7a\x3d\x83\xca\x21\x66\xef\x8d\xee\x4a\xf0\x2e\x20\x4b\xe3\x39\x28\xa2\xd4\xa6\x07\x77\xbe\x56\x2d\xba\xad\x75\x12\x59\x68\x1a\x7c\x78\xee\x79\xca\xf0\x7c\x36\x99\xdb\x7e\x10\xae\xb7\x3d\x91\x7d\x9d\xf9\xd9\x1d\x7d\x63\xa4\x0e\x11\x3d\x79\x7c\x4b\x4d\x6e\x40\xd6\x34\x1f\x7c\x61\x9c\x19\x06\x9a\x1f\x98\xf5\xc1\xa8\x31\x22\x2c\x54\x28\xb5\x70\x3c\xb6\x6d\xec\x21\x03\x80\xcf\x8c\x02\x09\x22\x73\xe7\x9d\xb5\x7e\x5d\x10\xd5\x4f\x3a\x20\xcc\x4c\xeb\x62\x6a\x53\x8b\xa4\x79\x39\x90\x64\x12\xd0\x1c\x94\xb3\x26\x36\x85\xd4\x6f\x17\x6f\x3f\xfe\xfc\xcb\xab\xf7\x37\xbf\xbe\xb9\x5e\xa4\x36\xb0\xe4\x78\xd8\x03\x3a\x37\xef\xd9\x99\x98\xd8\xe6\x36\x5d\xea\xa8\x5e\x9f\xf2\xf1\x51\xb3\x7d\xec\xe3\x54\x9c\x4c\xfc\xa4\xa3\xdc\xf7\x78\xf9\x18\xb4\x31\x4c\x67\xe3\x48\x6f\x5d\xcc\x49\x26\xe2\xe1\x50\x93\x27\x3d\x4e\x34\xc3\xf8\x2d\x0c\x08\xed\xd1\x19\x1e\xaf\x1f\x59\xbc\x75\xb6\xe1\xb2\x18\xa6\x96\x25\x08\xe2\x72\xeb\x3b\x2b\x87\x41\x5b\xb9\xa7\xc7\xc9\x46\x73\x28\x4f\xc4\x65\x0a\xf7\x2c\x2e\x07\xa1\x03\x3e\x8a\xc9\x97\x8a\xf8\x61\x0d\x0c\x7d\xf7\x33\x15\xc0\x6d\x7f\xde\xee\x3f\xd3\xf0\x9f\xa8\x4b\xa6\x4a\x13\xce\x8c\x6e\x8e\x0f\x93\xd1\xac\xb2\x9c\xf9\x90\xd2\x90\xd6\x34\xac\xb8\x11\x49\x21\x6b\xac\x38\xb2\x79\x6a\xc7\xc9\x92\x93\xc8\x61\x59\x66\x52\xac\xeb\x47\xc7\x8c\xa1\x5f\x33\x23\xe3\x32\x2a\xe1\xf5\x88\x42\xdb\xea\x8e\x03\x41\x79\x28\xa6\x01\xce\x1f\x2d\x5b\x19\x38\xa5\xb1\xe0\xe2\x4e\x1c\xf3\x00\xb5\x3d\xc6\x15\xd0\x1a\x83\xd2\xc7\xe1\xce\xcf\x43\xb7\x5a\x8d\x0e\xc4\xf9\x9f\x95\x5f\x8d\x47\x45\x3f\xf7\x14\x74\x90\x85\xd4\x81\x3c\xba\x82\xf1\x4b\xe7\x21\xf9\x48\xe9\xaa\x4d\xa1\x78\x95\x48\xdf\xdc\xce\x9c\xe2\x5b\x47\xe8\xe3\x8a\x39\x4f\xec\x64\xc3\x40\xcf\x8b\xbc\x77\x4c\x19\x97\xbe\x0c\x81\x73\x8b\x7b\xea\xab\xb3\xd9\xa0\xa5\x08\x9a\x40\x71\x09\x8e\xd1\x53\x58\xa5\x6a\xda\x44\x5c\x8f\x23\x4e\xdc\x7d\x5f\x0c\x0b\xe5\xcb\xdc\x96\xe1\x6e\xa5\x2a\xe4\xc9\x2c\x5b\x81\x67\x86\x30\x16\x26\x7c\x5f\x55\xca\x5d\x3d\x42\xfd\xdc\xac\x0f\xd9\x80\x35\x25\xef\xe3\x87\xdf\xd2\x8e\x2e\xcc\x2e\x7d\xbb\x56\x3e\xee\x8d\xa4\xbc\x75\xdd\x88\x56\xbf\xf2\x70\x38\x53\xce\x3d\x28\x38\x7d\x75\x7f\x0f\xc5\xb5\xf2\x2c\x29\x3e\xf5\xcc\x29\x36\x4e\x18\x59\x0f\x44\x3f\xc7\xbf\xd2\xa3\x8f\xda\xc6\x23\xbe\x1b\x74\x8a\x93\xe7\x03\xe6\xbb\x8b\x69\xa0\xbf\x5b\x65\x32\x86\xc5\x72\xd1\xbf\x1d\x69\xc2\x9c\x9d\xc7\xab\xc7\xad\xea\x28\x4c\x2c\x3f\x87\x9c\x55\x99\xa6\xfa\x46\x18\xb5\xe5\x79\x8f\x0b\x94\x54\x85\x2e\xf9\xfa\x60\x72\x8e\x3b\xaf\x25\x84\x60\x2a\x74\x0f\x02\xe8\x50\x0b\xaf\x0e\x18\xc7\x19\x1a\xd2\xbb\x9b\x05\xf1\x41\xc1\x8f\xce\x51\xd8\x54\xca\x5d\x2c\xd3\xcf\x6f\xc6\x87\xb0\x29\x38\xf1\xa1\xeb\x54\x70\xe2\xeb\xd1\x10\xd5\x81\xea\x84\x80\x8f\x84\xee\x14\x7f\x20\x74\x63\xe6\x98\x06\x4e\xf3\xff\xd2\x08\x75\xd2\x00\xe4\x0f\x83\x84\x81\x6a\x7a\xca\x3b\x09\xba\xc8\xf7\xf4\x68\x39\xa0\x68\xe2\xf3\x10\xc7\x89\xbb\x81\xf2\x0f\x16\xbc\x3c\x56\x3d\xae\xf6\xa8\x79\xf5\x19\x18\x1d\x38\x7a\x59\xcc\x75\xf5\xd7\x3d\x76\xa0\xaa\x1f\x47\xb2\xcf\xb4\xca\xcc\x2a\x16\x21\x7c\x70\x38\xdb\x32\x4f\xe8\x8a\x9f\xbb\xd5\x48\x4f\x33\x2c\x18\xa0\x10\x94\x87\x5a\x50\x84\x79\x6b\x74\x07\x42\x4a\xa4\x04\x97\x35\xa6\x87\x9a\x17\xc3\x9b\xc0\xa7\xad\xd0\x84\x9f\x5e\x9e\xd0\x36\xf0\xcf\x03\x4c\xde\x05\xe9\x93\xa2\x63\xdc\xf3\xb8\xef\x07\x0f\xd4\x19\x09\x1b\x6b\xf7\x7b\xc4\x96\xcb\x75\xd4\xb1\xd8\x29\xbf\x58\x42\x83\x82\x23\xc5\xd7\x1c\x44\x5c\xbc\xfa\x0a\x0e\x2d\x79\x87\xa2\x19\x4b\xf9\xa1\x35\x2c\x7a\x45\x5e\x78\xbc\xda\x29\xff\x74\xc2\x0d\xfe\xee\x87\xac\x67\x7d\x40\x18\x58\x0c\x3a\x16\x03\x4a\x67\x42\x5e\x60\xb1\x2b\x96\xf0\x2f\xde\x1c\xe1\x95\xb6\xa1\x7a\x59\xc4\x97\x03\x6f\xf7\x3c\xb4\x12\xb4\xc2\x79\x25\x83\x16\x6e\x88\x62\x2f\xe5\x61\x83\xe9\xb5\x5e\x1d\x89\x97\x53\xc9\xb2\x8a\xb8\x91\x16\x69\x25\x1d\x36\x90\x07\x6c\x51\xd1\x95\xd8\xc8\x8b\x6f\x2e\x1f\xff\x9f\x3b\x7c\x87\xee\x70\xe2\xc1\x97\x67\xad\xa9\xbb\x72\xa9\xfe\x90\xc3\xbc\xd8\x73\x7b\x48\xb9\x22\xf4\xd9\x2b\xf2\xf3\xec\x21\x3a\x7b\x51\x66\x17\xe3\xcb\x48\x9c\x77\xe6\x60\xac\x15\x79\x34\xab\xde\x84\xab\xf2\xf2\xfc\xf2\xe2\xac\xbf\xc6\x3f\x55\x95\x4a\xbb\x26\x83\xf8\x4f\x3c\xc3\xcc\xf0\x72\xfa\x3e\xf5\xf1\xfb\x7b\x70\xb1\x25\x7c\x81\x7b\x15\x9f\xf3\x67\x57\x7f\xfa\x6d\x50\xf0\xbe\xed\xc5\xbf\xbe\xb9\x1b\x1a\x30\x2d\xfb\xb9\x30\xb8\xbe\x1d\x83\xa9\xac\x27\xb0\x91\x18\x1a\xd1\xc5\x1d\x5d\x1f\xa6\xd7\x1a\x43\xda\xda\x7d\x68\x41\x11\x05\x24\xb0\x06\xc8\x36\x08\x6f\xc3\x06\x9d\x41\x8f\xc4\xd2\x43\x4b\xd3\x63\x4c\x65\x68\x78\x06\x58\xdc\x58\x83\x8b\xfc\xcb\xab\x68\x40\xfe\x1c\x93\x94\xd3\xfc\x85\x66\x98\xef\xa2\x7d\xb3\x2f\xe3\xe8\xb9\xb8\x58\x9c\xfd\x2f\x00\x00\xff\xff\xbd\x69\x00\x69\x38\x19\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(2019, 8, 8, 13, 51, 12, 391220692, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(2019, 8, 8, 13, 51, 12, 391706708, time.UTC),
			uncompressedSize: 874,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcd\x6e\x9c\x40\x10\x84\xef\x3c\x45\x49\x7b\x0d\x1b\x61\x69\x2f\xdc\xa2\x38\x89\x2c\x25\xd6\x5e\x9c\x7b\x7b\x68\xf0\x28\xf3\x97\xe9\x66\xb3\x04\xf9\xdd\xa3\xd9\x5f\x36\xf6\x9c\x80\xaa\xaf\xa7\xa6\x80\xba\xae\xab\x15\x3c\x7b\x43\xe6\x85\x3b\x74\x9c\x5c\x9c\x3c\x07\xc5\x28\xdc\xe1\x79\xc2\x57\x37\xee\xa1\x11\x07\x47\xb5\x82\x89\x41\xc9\x06\xce\xb0\x9e\x06\x86\x67\xa5\x8e\x94\xd6\x15\x25\xfb\x93\xb3\xd8\x18\x5a\x50\x4a\xf2\x71\xd7\x54\xbf\x6c\xe8\x5a\xdc\x5f\xc6\x56\x67\x7b\x5b\x01\x81\x3c\xb7\xd7\xdd\xe7\x19\xb6\xc7\xfa\x91\x3c\x4b\x22\xc3\x78\x7d\x3d\x99\x0e\xb7\x2d\xe6\xf9\x56\x9d\x67\x70\xe8\x8a\x4d\x12\x9b\x32\x31\x73\x72\xd6\x90\xb4\x68\x2a\x40\xd8\xb1\xd1\x98\x8b\x02\x78\x52\xf3\xf2\x9d\x9e\xd9\xc9\xf1\xc1\x9b\x00\x15\xa0\xec\x93\x23\xe5\x13\xb2\x08\x5b\x96\xbb\xa1\xdf\xe3\x81\x73\x94\xb2\x2e\x5d\x5d\x98\xfa\x5d\xa6\xac\x43\x9b\x0b\xa1\x6d\xd6\x9b\x75\xb3\xb9\xd5\xb7\xa3\x73\xdb\xe8\xac\x99\x5a\x3c\xf4\x8f\x51\xb7\x99\xa5\xd4\x7a\x76\x51\x1e\x16\xf9\x6a\xd4\x1e\x9b\xe6\x0e\xc0\x0a\x3f\x68\x6f\xfd\xe8\xcb\x0e\x31\x4f\xe5\x95\x8e\xc2\x1f\x60\x03\x3c\x0f\xf4\x3c\x29\xcb\x12\x7c\xc0\xc6\xe3\x06\x14\xfb\x97\xd1\xc7\x8c\x18\x18\x56\xd9\x2f\xed\x09\x4d\x73\xd7\x34\x58\xe1\x9e\x7b\x1a\x9d\x22\xc5\x7c\xcd\xb5\x2a\x9e\xdd\xee\x78\xf9\x14\x4c\xf4\x87\x8f\x4c\x23\x06\x56\xb8\x38\x08\x62\x0f\x26\xf3\x82\xcc\xbf\x47\x16\x05\x85\x0e\x99\x25\xc5\x20\xbc\xbe\x0c\x2a\x53\x6f\x4e\x78\xec\xd3\x38\xcb\x41\xaf\x07\x58\x74\xbf\x8d\x59\xdb\x63\xba\x8b\x2c\x6c\xc6\x6c\x75\xfa\x1c\x83\xf2\x5e\xdb\x05\x97\xc7\xf0\x49\x9e\x84\xf3\xff\xcc\x49\xfa\x96\xe3\x98\xde\x6a\xe4\x5c\xfc\xb3\xcd\x76\x67\x1d\x0f\xfc\x45\x0c\x39\xd2\xc3\xaf\xd0\x93\x13\xae\xfe\x05\x00\x00\xff\xff\x5d\x9a\x63\xab\x6a\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(2019, 8, 8, 13, 51, 12, 391977285, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
