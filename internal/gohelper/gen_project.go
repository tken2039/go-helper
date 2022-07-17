package gohelper

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

type PrjGenerator struct {
	PrjName    string
	ConfigFile string
}

func NewPrjGenerator(prjName, config string) *PrjGenerator {
	pg := &PrjGenerator{}

	pg.PrjName = prjName
	pg.ConfigFile = config

	return pg
}

type prjConfig struct {
	GoMod      goModConfig `yaml:"gomod"`
	WithCmd    bool        `yaml:"withCmd"`
	HttpServer bool        `yaml:"httpServer"`
}

type goModConfig struct {
	Repository string `yaml:"repository"`
	GoVersion  string `yaml:"goVersion"`
}

func NewPrjConfig(repository, goVersion string, withCmd, httpServer bool) *prjConfig {
	pc := &prjConfig{GoMod: goModConfig{}}

	pc.GoMod.Repository = repository
	pc.GoMod.GoVersion = goVersion

	pc.WithCmd = withCmd
	pc.HttpServer = httpServer

	return pc
}

func NewPrjConfigPlain() *prjConfig {
	return &prjConfig{}
}

func (pc *prjConfig) fromYmlBytes(b []byte) error {
	err := yaml.Unmarshal(b, pc)
	if err != nil {
		return err
	}

	return nil
}

var (
	//go:embed static/go_src_plain.go.tmpl
	goSrcPlain []byte

	//go:embed static/go_src_with_http_server.go.tmpl
	goSrcWithHTTPServer []byte
)

const (
	cmdDir      = "/cmd"
	goExtension = ".go"
	goMod       = "go.mod"
)

func (pg *PrjGenerator) Run() error {
	fmt.Println("Start creating a project.")
	fmt.Println("Project Name:", pg.PrjName)

	conf := NewPrjConfigPlain()
	if pg.ConfigFile != "" {
		fmt.Println("Read configration:", pg.ConfigFile)
		confBytes, err := readFile(pg.ConfigFile)
		if err != nil {
			return err
		}
		err = conf.fromYmlBytes(confBytes)
		if err != nil {
			return err
		}
		fmt.Printf("[config] repository: %v\n", conf.GoMod.Repository)
		fmt.Printf("[config] go version: %v\n", conf.GoMod.GoVersion)
		fmt.Printf("[config] cmd dir: %v\n", conf.WithCmd)
		fmt.Printf("[config] http server: %v\n", conf.HttpServer)
	}

	goSrcDir := pg.PrjName
	if conf.WithCmd {
		goSrcDir = fmt.Sprintf("%v/%v/%v", goSrcDir, cmdDir, pg.PrjName)
	}

	err := makeDir(goSrcDir)
	if err != nil {
		return err
	}

	var goSrc []byte
	if conf.HttpServer {
		goSrc = goSrcWithHTTPServer
	} else {
		goSrc = goSrcPlain
	}
	filePath := fmt.Sprintf("%v/%v%v", goSrcDir, pg.PrjName, goExtension)
	err = writeFile(filePath, goSrc)
	if err != nil {
		return err
	}

	if conf.GoMod.Repository != "" && conf.GoMod.GoVersion != "" {
		goModContent := makeGoMod(conf.GoMod.Repository, pg.PrjName, conf.GoMod.GoVersion)
		goModPath := makePath(pg.PrjName, goMod)
		err = writeFile(goModPath, []byte(goModContent))
		if err != nil {
			return err
		}
	}

	fmt.Println("Finished.")

	return nil
}

func makeGoMod(repository, prjName, version string) string {
	content := fmt.Sprintf("module %v/%v\n\ngo %v", repository, prjName, version)
	return content
}
