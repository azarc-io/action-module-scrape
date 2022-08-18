package util

import (
	"encoding/base64"
	"errors"
	"fmt"
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

const maxImageSize = 500000

type Logger interface {
	Fatalf(msg string, args ...any)
}

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

func LoadSchema(log Logger, file string) []byte {
	v := LoadFile(log, file, true)
	if _, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(v)); err != nil {
		log.Fatalf("loading schema: %s", err.Error())
	}
	return v
}

//********************************************************************************************
// IMAGE
//********************************************************************************************

func LoadImage(log Logger, file string, mustLoad bool) string {
	fp, err := os.Open(file)
	if err != nil {
		if !mustLoad && errors.Is(err, syscall.ENOENT) {
			return ""
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
		return fmt.Sprintf("data:image/svg+xml;base64, %s", base64.StdEncoding.EncodeToString(buf))
	}
	if cfg.Width*cfg.Height > maxImageSize {
		log.Fatalf("error loading %s's HxW cannot exceed %d", file, maxImageSize)
	}
	return fmt.Sprintf("data:image/%s;base64, %s", tp, base64.StdEncoding.EncodeToString(LoadFile(log, file, true)))
}
