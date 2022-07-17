package main

import cmd "github.com/tken2039/go-helper/internal/cmd/gohelper"

var version string

func main() {
	if err := cmd.NewRootCmd(version).Execute(); err != nil {
		cmd.HandleCmdErr(err)
	}
}
