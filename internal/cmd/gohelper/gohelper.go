package gohelper

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tken2039/go-helper/internal/cmd/version"
)

var cmdName = "gohelper"

func NewRootCmd(ver string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   cmdName,
		Short: "This tool is supposed to solve the hassles associated with the Go language. (Maybe :D)",
	}

	rootCmd.AddCommand(version.NewVersionCmd(ver))
	rootCmd.AddCommand((NewGenProjectCmd()))

	return rootCmd
}

func HandleCmdErr(err error) {
	os.Exit(1)
}
