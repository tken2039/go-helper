package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func NewVersionCmd(version string) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "show version informations.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("App version: ", version)
			fmt.Println("Go version :", runtime.Version())
		},
	}

	return versionCmd
}
