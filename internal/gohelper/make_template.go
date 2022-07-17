package gohelper

import (
	_ "embed"
	"strings"
)

type TmplMaker struct {
	DirName string
}

func NewTmplMaker(dirName string) *TmplMaker {
	tm := &TmplMaker{}

	tm.DirName = dirName

	return tm
}

//go:embed static/gohelper.yml.tmpl
var tmplYml []byte

const tmplYmlFileName = "gohelper_tmpl.yml"

func (tm *TmplMaker) Run() error {
	targetPath := makePath(tm.DirName, tmplYmlFileName)

	err := writeFile(targetPath, tmplYml)
	if err != nil {
		return err
	}

	return nil
}

func makePath(dirName, fileName string) string {
	if dirName == "" {
		return fileName
	}

	d := dirName
	if !strings.HasSuffix(d, "/") {
		d = d + "/"
	}

	path := d + fileName

	return path
}
