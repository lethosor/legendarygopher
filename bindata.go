// Code generated by go-bindata.
// sources:
// assets/css/main.css
// assets/templates/artifacts.html
// assets/templates/entities.html
// assets/templates/events.html
// assets/templates/figure.html
// assets/templates/figures.html
// assets/templates/index.html
// DO NOT EDIT!

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

var _assetsCssMainCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x2b\x28\xca\x2f\x48\x2d\x52\xa8\xe6\xe2\x2c\x49\xad\x28\xd1\x2d\x29\x4a\xcc\x2b\x4e\xcb\x2f\xca\xb5\x52\x48\x4e\x2c\xc8\x2c\x49\xcc\xc9\xac\x4a\xb5\xe6\xaa\xe5\x02\x04\x00\x00\xff\xff\x6d\x88\x75\x57\x29\x00\x00\x00")

func assetsCssMainCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsCssMainCss,
		"assets/css/main.css",
	)
}

func assetsCssMainCss() (*asset, error) {
	bytes, err := assetsCssMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/css/main.css", size: 41, mode: os.FileMode(436), modTime: time.Unix(1451762230, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesArtifactsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x51\xcd\x4e\xb4\x30\x14\xdd\x7f\x4f\x71\xbf\xce\x5a\x9a\x99\xd9\x96\x26\xc6\x31\x6a\x62\xd4\x85\x89\x71\x79\x85\x3b\x53\x62\x81\xa6\xbd\x0b\x09\xe1\xdd\xed\x14\x32\x54\x23\x1b\x0e\x39\x7f\x9c\x56\xfd\x3f\x3c\xdf\xbc\xbe\xbf\xdc\x82\xe1\xd6\xea\x7f\x6a\x7e\x41\x7c\x94\x21\xac\x67\x98\x3e\xb9\x61\x4b\xfa\xbe\x6f\x49\xc9\x19\xaf\x9c\x6d\xba\x4f\x30\x9e\x8e\xa5\x90\x18\x02\x71\x90\x55\x08\xb2\xc5\xa6\x2b\x22\x10\xe0\xc9\x96\x22\xf0\x60\x29\x18\x22\x16\xc0\x83\xa3\x52\x30\x7d\xf1\x59\x29\x96\x4a\xb9\x76\xaa\x8f\xbe\x1e\xb2\x0a\xb3\xd5\x8f\x74\xa2\xae\x46\x3f\xc0\x5d\xef\x0c\xf9\x28\xdf\xe6\x8a\x9d\xbe\xf6\xdc\x1c\xb1\xe2\x10\xa9\xdd\x4a\x8d\xa3\xc7\xee\x44\x50\xbc\xf5\xde\xd6\xc5\x45\x35\x4d\x99\x7b\x0f\x4d\x5d\x0a\x5c\xb8\xab\x71\x84\xe2\xe1\x00\xd3\x24\xa0\xb2\x71\x53\x29\x9c\xef\x1d\x79\xb1\xc6\x26\x1f\x2e\xbb\x37\x7f\x38\xf5\xe6\x82\x95\xc4\x9f\xc6\x33\xf3\x84\x2d\x41\xfe\x13\xd2\xec\xb3\x41\xee\x77\x73\x4a\x63\x6a\x53\x9e\xd3\x90\x0d\x8c\x07\xb3\x04\x29\x39\x9f\x5c\x0c\x4b\x77\xf9\x1d\x00\x00\xff\xff\x83\xaa\x6c\x8f\xe3\x01\x00\x00")

func assetsTemplatesArtifactsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesArtifactsHtml,
		"assets/templates/artifacts.html",
	)
}

func assetsTemplatesArtifactsHtml() (*asset, error) {
	bytes, err := assetsTemplatesArtifactsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/artifacts.html", size: 483, mode: os.FileMode(436), modTime: time.Unix(1451763920, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesEntitiesHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x51\xc1\x4e\xeb\x30\x10\xbc\xbf\xaf\xd8\xb7\x3d\x63\xab\xf4\x6a\xe7\x42\x2b\x84\x84\x80\x03\x12\xe2\x68\xea\x6d\x6d\xe1\x38\x91\xed\x03\x56\xd4\x7f\xc7\x49\x4a\xe3\xaa\xb9\xec\x4e\x76\x46\xb3\xb3\x16\xff\xb7\xaf\x0f\xef\x9f\x6f\x3b\x30\xa9\x75\xcd\x3f\x31\x17\x28\x9f\x30\xa4\xf4\xdc\x4e\x30\xd9\xe4\xa8\xd9\xf9\x52\x2d\x45\xc1\x67\xbc\xcc\x9d\xf5\xdf\x60\x02\x1d\x24\x72\x15\x23\xa5\xc8\xf7\x31\xf2\x56\x59\xcf\x4a\x83\x10\xc8\x49\x8c\x29\x3b\x8a\x86\x28\x21\xa4\xdc\x93\xc4\x44\x3f\x69\x64\xe2\xd9\x96\x2f\xbe\xe2\xab\xd3\xb9\xb2\x30\xeb\xe6\x99\x8e\xe4\xb5\x0a\x19\x1e\xbb\xde\x50\x28\xf4\x75\xcd\xb8\xaf\x36\x2c\xe0\x32\x19\x86\xa0\xfc\x91\x80\x7d\x74\xc1\x69\xf6\x47\x3a\x9d\x2a\x86\x3d\x00\x7b\x51\x2d\x41\xf5\x57\x98\x0d\x58\x2d\x91\x46\x41\xbe\x1b\x06\x60\x4f\xdb\x42\x40\xd8\xbb\x92\x52\x62\x1f\xba\x9e\x02\x2e\x4e\x93\x4a\x9d\x2f\xb1\xba\xd1\x35\xab\x4b\x2f\xb8\xba\x96\x8d\x93\x9b\x05\xb8\xd9\xd4\x31\x4a\xfa\xab\xa5\x17\x2c\xf8\x7c\xae\xa2\x98\x1e\xf1\x37\x00\x00\xff\xff\x19\x0d\x84\xdf\xdc\x01\x00\x00")

func assetsTemplatesEntitiesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesEntitiesHtml,
		"assets/templates/entities.html",
	)
}

func assetsTemplatesEntitiesHtml() (*asset, error) {
	bytes, err := assetsTemplatesEntitiesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/entities.html", size: 476, mode: os.FileMode(436), modTime: time.Unix(1451765413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesEventsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x51\xcb\x4e\xc3\x30\x10\xbc\xf3\x15\x8b\xcb\x95\x58\xa5\x37\x64\xe7\x42\xcb\x43\x42\x02\x09\x24\xd4\xa3\x69\x96\xd8\xc2\x75\x22\x7b\xa1\x44\x51\xfe\x1d\xe7\x41\xe3\x92\xcb\xee\x64\x66\x34\xbb\x6b\x71\xbe\x7e\xba\x79\xdd\x3e\x6f\x40\xd3\xde\xe6\x67\x62\x2c\x10\x3f\xa1\x51\x15\x63\x3b\x40\x32\x64\x31\xdf\x7c\xa3\xa3\x20\xf8\x88\x66\xd6\x1a\xf7\x09\xda\xe3\x87\x64\x5c\x85\x80\x14\xf8\x2e\x04\xbe\x57\xc6\x65\xb1\x61\xe0\xd1\x4a\x16\xa8\xb1\x18\x34\x22\x31\xa0\xa6\x46\xc9\x08\x7f\xa8\x57\xb2\x29\x94\xcf\xa9\xe2\xbd\x2a\x9a\x24\x42\x2f\xf3\x47\x2c\xd1\x15\xca\x37\x70\x57\xd5\x1a\x7d\x94\x2f\x53\xc5\x55\x7e\x6f\x02\x55\xde\xec\x94\x85\xbf\x51\xe3\xdf\xa3\xa4\x6d\x0f\x86\x34\x5c\x1c\xe0\x5a\x42\xf6\x56\x79\x5b\x74\x5d\xc2\x7a\xe5\x4a\x8c\x74\x36\x9a\x13\x4e\xe8\x15\x98\x42\x32\xec\x89\xcb\xb6\x85\xec\x61\x0d\x5d\xc7\x60\x67\xe3\xbe\x92\xd5\xbe\xaa\xd1\xb3\x39\x6a\x30\xa9\xe9\x26\x8b\xff\xb6\x7c\x71\xec\x05\x57\xa7\xae\xc8\xc4\x09\x6e\x4d\xf9\xe5\x11\xa6\x3a\x28\x7b\x26\x7b\x21\x45\xd8\x03\xe3\x06\xbc\x45\xe5\x21\x1d\x94\xeb\x55\xba\x70\x3c\xd8\xc9\x8a\x33\x16\x7c\xbc\x70\x74\x0c\xaf\xfe\x1b\x00\x00\xff\xff\x43\xf4\xef\x7c\x0d\x02\x00\x00")

func assetsTemplatesEventsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesEventsHtml,
		"assets/templates/events.html",
	)
}

func assetsTemplatesEventsHtml() (*asset, error) {
	bytes, err := assetsTemplatesEventsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/events.html", size: 525, mode: os.FileMode(436), modTime: time.Unix(1451765103, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesFigureHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x51\x4d\x4f\xc3\x30\x0c\xbd\xf3\x2b\x4c\xd4\x73\xa3\xb1\xdb\x94\xe6\xc2\xc6\x87\x84\x04\x42\x48\x88\x63\x58\xbd\x25\x22\x6b\xa7\xc4\x63\x54\x55\xff\x3b\x6e\xb3\x69\x81\x5e\x9a\xf8\xbd\xe7\xf7\xec\xa8\xeb\xe5\xf3\xed\xdb\xc7\xcb\x0a\x2c\xed\xbc\xbe\x52\xe9\x07\xfc\x29\x8b\xa6\x4e\xc7\xe9\x4a\x8e\x3c\xea\xbe\x87\xf2\xce\x6d\x0f\x01\x61\x18\x94\x4c\xc5\x0b\xc9\xbb\xe6\x0b\x6c\xc0\x4d\x25\xa4\x89\x11\x29\xca\x75\x8c\x72\x67\x5c\x53\xf2\x41\x40\x40\x5f\x89\x48\x9d\xc7\x68\x11\x49\x00\x75\x7b\xac\x04\xe1\x0f\x8d\x4c\x71\xf2\x96\x17\x73\xf5\xd9\xd6\x5d\x66\x61\x67\xfa\x09\xb7\xd8\xd4\x26\x74\x70\xdf\xee\x2d\x06\xa6\xcf\x72\xc6\x8d\x7e\x70\x91\xda\xe0\xd6\xc6\x43\x4a\xbb\x80\x7f\xc9\x99\x94\x29\xe6\x7a\xf5\x8d\x0d\x45\xae\xcf\xb3\xfa\xc1\x5f\x2e\x7d\x5f\x6c\x60\x51\x9d\x9b\x0c\x43\x8e\x1c\x27\xe4\xbd\x0d\xbe\xfe\x03\x04\xd3\x6c\x11\x0a\x1c\xe1\xe2\x78\x92\x26\x2b\x28\x36\xe5\xe3\x32\x63\xf3\xf2\xf4\xd8\xaa\x7c\xe5\xe1\x30\x4c\x2c\x96\x8e\x61\x19\xc9\x9a\x32\x9c\xcb\xe4\x39\xa5\x92\x69\x55\x3c\xc4\xf4\x8a\xbf\x01\x00\x00\xff\xff\x74\x93\x4f\x47\xdd\x01\x00\x00")

func assetsTemplatesFigureHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesFigureHtml,
		"assets/templates/figure.html",
	)
}

func assetsTemplatesFigureHtml() (*asset, error) {
	bytes, err := assetsTemplatesFigureHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/figure.html", size: 477, mode: os.FileMode(436), modTime: time.Unix(1451801113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesFiguresHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x51\x4d\x4f\xc3\x30\x0c\xbd\xf3\x2b\x4c\x76\xa6\xd1\xd8\x35\xed\x85\xf1\x25\x21\xc1\x01\x09\x71\x0c\xab\xd7\x44\x64\x6d\x65\x1b\x89\xaa\xea\x7f\x27\x6b\xab\x36\x02\x91\x4b\x9e\xe3\xf7\x6c\x3f\xc7\x5c\xee\x9f\x6f\x5e\xdf\x5f\x6e\xc1\xc9\x29\x14\x17\x66\xba\x20\x1e\xe3\xd0\x96\x13\x1c\x43\xf1\x12\xb0\xb8\xf3\xd5\x17\x21\x1b\x3d\x85\x6b\x3a\xf8\xfa\x13\x1c\xe1\x31\x57\xda\x32\xa3\xb0\x3e\x30\xeb\x93\xf5\x75\x16\x81\x02\xc2\x90\x2b\x96\x2e\x20\x3b\x44\x51\x20\x5d\x8b\xb9\x12\xfc\x96\x33\x53\xcd\x5d\xf5\xda\xd6\x7c\x34\x65\x97\xb4\x70\xdb\xe2\x09\x2b\xac\x4b\x4b\x1d\xdc\x37\xad\x43\x8a\xf4\x6d\xca\xb8\x2e\x1e\x3c\x4b\x43\xfe\x60\x03\x2c\xb3\xc6\xe7\x85\xd3\xf7\x64\xeb\x0a\x21\x7b\x6b\x28\x94\xd9\xcc\x19\x86\xa4\xc8\x0e\x7c\x99\xab\xe3\x98\xb9\xea\x7b\xc8\x1e\xf7\x30\x0c\x0a\x0e\x21\x1a\xcb\x55\x4b\x4d\x8b\xa4\xd6\x92\xa3\xca\xce\xe6\x37\x7f\x74\xc5\x66\xc1\x46\xdb\x7f\x64\x7a\x92\xb1\x4e\x74\x67\xf8\x5b\x14\xcd\xec\x52\x33\x71\x1b\xf3\xec\x46\x4f\xeb\x8a\x8c\xf1\x0f\x7f\x02\x00\x00\xff\xff\x7e\xb5\x9d\x8b\xdb\x01\x00\x00")

func assetsTemplatesFiguresHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesFiguresHtml,
		"assets/templates/figures.html",
	)
}

func assetsTemplatesFiguresHtml() (*asset, error) {
	bytes, err := assetsTemplatesFiguresHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/figures.html", size: 475, mode: os.FileMode(436), modTime: time.Unix(1451801812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x31\x4b\xc7\x30\x10\xc5\x77\x3f\xc5\xd9\x49\x97\x86\xee\x67\x40\xb4\xea\x20\xe8\x20\x88\x63\x34\xd7\x26\x90\x36\x12\x53\xa1\x94\x7e\x77\xd3\xa4\xb5\x85\x7f\xf9\xd3\x2e\xf7\xae\xef\xe5\xc7\x71\x87\x97\xf7\x2f\x77\x6f\x1f\xaf\x25\x28\xdf\x18\x7e\x81\xa9\x40\xf8\x50\x91\x90\x49\xc6\xd6\x6b\x6f\x88\x3f\xd9\x86\x90\x25\x9d\x62\x6c\xcd\xe1\xa7\x95\xfd\xe6\x89\x2a\xf8\x33\xd5\xd4\x4a\xe1\x7a\x78\xb4\xdf\x8a\x5c\x88\x17\x9b\x44\x67\xd6\x26\xfe\x30\x9a\xa3\x00\xe5\xa8\xba\xc9\x98\x70\x5e\x57\xe2\xcb\xff\x64\xfc\x76\x91\xc8\x04\x87\xab\x61\x00\x43\x2d\xe4\xef\xd6\x19\x99\xff\x9b\x30\x8e\xd7\xc8\x02\xe3\x0c\x94\xda\x30\xbd\xa6\xc0\x2c\x67\xb5\x87\x5c\xbc\x43\xc4\xdf\xc0\x9c\x78\xb1\xee\xd2\xa2\x73\x84\x55\xe9\xba\x73\xd3\x70\x0f\x49\xec\xd1\x66\xeb\x14\x87\x6c\xd9\x27\xb2\x74\x8a\xb0\xee\x78\xd0\xbf\x00\x00\x00\xff\xff\x36\xcd\xb6\x39\xe8\x01\x00\x00")

func assetsTemplatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesIndexHtml,
		"assets/templates/index.html",
	)
}

func assetsTemplatesIndexHtml() (*asset, error) {
	bytes, err := assetsTemplatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/index.html", size: 488, mode: os.FileMode(436), modTime: time.Unix(1451797714, 0)}
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
	"assets/css/main.css": assetsCssMainCss,
	"assets/templates/artifacts.html": assetsTemplatesArtifactsHtml,
	"assets/templates/entities.html": assetsTemplatesEntitiesHtml,
	"assets/templates/events.html": assetsTemplatesEventsHtml,
	"assets/templates/figure.html": assetsTemplatesFigureHtml,
	"assets/templates/figures.html": assetsTemplatesFiguresHtml,
	"assets/templates/index.html": assetsTemplatesIndexHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"main.css": &bintree{assetsCssMainCss, map[string]*bintree{}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			"artifacts.html": &bintree{assetsTemplatesArtifactsHtml, map[string]*bintree{}},
			"entities.html": &bintree{assetsTemplatesEntitiesHtml, map[string]*bintree{}},
			"events.html": &bintree{assetsTemplatesEventsHtml, map[string]*bintree{}},
			"figure.html": &bintree{assetsTemplatesFigureHtml, map[string]*bintree{}},
			"figures.html": &bintree{assetsTemplatesFiguresHtml, map[string]*bintree{}},
			"index.html": &bintree{assetsTemplatesIndexHtml, map[string]*bintree{}},
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

