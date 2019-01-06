package main

import (
	"fmt"

	"github.com/hellozimi/sidago/internal/content"
	"github.com/spf13/cobra"
)

type newCmd struct {
	cmd     *cobra.Command
	rootCmd *cobra.Command
	command
	force   bool
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
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("requires post or page argument followed by the wanted name")
			}
			if args[0] != "post" && args[0] != "page" {
				return fmt.Errorf("requries post or page as first argument")
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&n.force, "force", "f", false, "forces creation and overwrites existing file")

	n.cmd = cmd
	cmd.RunE = n.runNew

	return n
}

func (n *newCmd) runNew(c *cobra.Command, args []string) error {
	path := n.rootCmd.PersistentFlags().Lookup("path").Value.String()
	f := n.force
	pageType := args[0]
	pageName := args[1]
	cc := content.New(path, pageType, pageName)

	return cc.Execute(f)
}
