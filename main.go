package main

import (
	"os"

	"github.com/fs015/qualysbeat/cmd"

	_ "github.com/fs015/qualysbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
