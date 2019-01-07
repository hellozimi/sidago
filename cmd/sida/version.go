package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Build, Version string
)

type versionCmd struct {
	cmd     *cobra.Command
	rootCmd *cobra.Command
	command
}

func (v *versionCmd) Command() *cobra.Command {
	return v.cmd
}

func newVersionCommand(rootCmd *cobra.Command) command {
	c := &versionCmd{rootCmd: rootCmd}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Displays the current version flags",
		Run:   c.run,
	}
	c.cmd = cmd

	return c
}

func (v *versionCmd) run(cmd *cobra.Command, args []string) {
	fmt.Printf("sida %s - build %s\n\n", Version, Build)
}
