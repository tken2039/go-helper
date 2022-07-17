package gohelper

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tken2039/go-helper/internal/gohelper"
)

type MakeTemplateOption struct {
	DirName string
}

func NewMakeTemplateCmd() *cobra.Command {
	opt := &MakeTemplateOption{}

	makeTemplateCmd := &cobra.Command{
		Use:   "make-template",
		Short: "make yml template for generating Go project",
		Run: func(cmd *cobra.Command, args []string) {
			err := runMakeTemplate(opt)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	makeTemplateCmd.Flags().StringVarP(&opt.DirName, "dir", "d", "", "directory path for output")

	return makeTemplateCmd
}

func runMakeTemplate(opts *MakeTemplateOption) error {
	tm := gohelper.NewTmplMaker(opts.DirName)

	return tm.Run()
}
