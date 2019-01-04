package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

type newCmd struct {
	cmd     *cobra.Command
	rootCmd *cobra.Command
	command
	newType string
}

func (n *newCmd) Command() *cobra.Command {
	return n.cmd
}

func newNewCommand(rootCmd *cobra.Command) command {
	n := &newCmd{rootCmd: rootCmd}
	cmd := &cobra.Command{
		Use:   "new [OPTIONS]",
		Short: "Creates a new post or page",
	}

	n.cmd = cmd
	cmd.RunE = n.runNew

	return n
}

func (n *newCmd) runNew(c *cobra.Command, args []string) error {
	fmt.Printf("%v: %v\n", args, c.Parent().PersistentFlags().Lookup("path").Value)
	return nil
}
