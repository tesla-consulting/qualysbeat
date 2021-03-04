package main

import (
	"os"

	"github.com/tesla-consulting/qualysbeat/cmd"

	_ "github.com/tesla-consulting/qualysbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
