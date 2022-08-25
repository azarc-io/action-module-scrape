package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/azarc-io/action-module-scrape/temp/module_v1"
	svg "github.com/h2non/go-is-svg"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"syscall"
)

const (
	maxImageSize = 500000
	schema       = "schema"
)

//********************************************************************************************
// FILE
//********************************************************************************************

func LoadDirs(log Logger, path string, cb func(path, name string)) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("listing %s: %s", path, err.Error())
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		cb(fmt.Sprintf("%s/%s", path, dir.Name()), dir.Name())
	}
}

func LoadFile(log Logger, file string, mustLoad bool) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		if !mustLoad && errors.Is(err, syscall.ENOENT) {
			return nil
		}
		log.Fatalf("reading file: %s", err.Error())
	}
	return data
}

func LoadFileString(log Logger, file string, mustLoad bool) string {
	return string(LoadFile(log, file, mustLoad))
}

//********************************************************************************************
// YAML
//********************************************************************************************

func ParseYaml(log Logger, file string, v interface{}) {
	data := LoadFile(log, file, true)
	if err := yaml.Unmarshal(data, v); err != nil {
		log.Fatalf("unmarshal file: %s", err.Error())
	}
}

//********************************************************************************************
// SCHEMA
//********************************************************************************************

func LoadSchema(log Logger, file string) string {
	v := LoadFile(log, file, true)
	if _, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(v)); err != nil {
		log.Fatalf("loading schema: %s", err.Error())
	}
	return string(v)
}

//********************************************************************************************
// IMAGE
//********************************************************************************************

func LoadImage(log Logger, file string, mustLoad bool) *module_v1.Image {
	fp, err := os.Open(file)
	if err != nil {
		if !mustLoad && errors.Is(err, syscall.ENOENT) {
			return &module_v1.Image{}
		}
		log.Fatalf("opening file: %s", err)
	}
	defer func() {
		if err = fp.Close(); err != nil {
			log.Fatalf("could not close %s: %s", file, err.Error())
		}
	}()

	cfg, tp, err := image.DecodeConfig(fp)
	if err != nil {
		if !errors.Is(err, image.ErrFormat) {
			log.Fatalf("error loading %s: %s", file, err)
		}

		buf := LoadFile(log, file, true)
		if !svg.IsSVG(buf) {
			log.Fatalf("%s uses an unsupported image format", file)
		}
		return &module_v1.Image{Encoding: module_v1.ImageEncoding_B64SVG, Data: base64.StdEncoding.EncodeToString(buf)}
	}
	if cfg.Width*cfg.Height > maxImageSize {
		log.Fatalf("error loading %s's HxW cannot exceed %d", file, maxImageSize)
	}
	img := &module_v1.Image{Data: base64.StdEncoding.EncodeToString(LoadFile(log, file, true))}
	switch tp {
	case "png":
		img.Encoding = module_v1.ImageEncoding_B64PNG
	case "jpeg":
		img.Encoding = module_v1.ImageEncoding_B64JPG
	default:
		log.Fatalf("unknown image encoding %s for %s", tp, file)
	}
	return img
}
