package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type rootOptions struct {
	path string
}

var opts = rootOptions{}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sida [OPTIONS]",
		Short: "A static site generator",
	}

	cmd.PersistentFlags().StringVarP(&opts.path, "path", "p", "./", "base path for your sida location")

	return cmd
}

func main() {
	cmd := newRootCommand()
	cmd.SetOutput(os.Stdout)
	newCmd := newNewCommand(cmd)
	initCmd := newInitCommand(cmd)
	generateCmd := newGenerateCommand(cmd)
	versionCmd := newVersionCommand(cmd)

	cmd.AddCommand(
		newCmd.Command(),
		initCmd.Command(),
		generateCmd.Command(),
		versionCmd.Command(),
	)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
