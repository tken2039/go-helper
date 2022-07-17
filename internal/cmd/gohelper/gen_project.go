package gohelper

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tken2039/go-helper/internal/gohelper"
)

type GenProjectOption struct {
	PrjName string
	Config  string
}

func NewGenProjectCmd() *cobra.Command {
	opt := &GenProjectOption{}

	genPrdCmd := &cobra.Command{
		Use:   "init",
		Short: "generate Go project",
		Run: func(cmd *cobra.Command, args []string) {
			err := runGenProject(opt)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	genPrdCmd.AddCommand(NewMakeTemplateCmd())

	genPrdCmd.Flags().StringVarP(&opt.PrjName, "name", "n", "", "project name (creating directory with the name specified in this parameter)")
	genPrdCmd.Flags().StringVarP(&opt.Config, "config", "c", "", "fine name of cofig yml")

	return genPrdCmd
}

func runGenProject(opts *GenProjectOption) error {
	if opts.EmptyPrjName() {
		return errors.New("error: prj name is empty")
	}

	pg := gohelper.NewPrjGenerator(opts.PrjName, opts.Config)

	return pg.Run()
}
