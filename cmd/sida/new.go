package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

type newCmd struct {
	newType string
}

func newNewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new [OPTIONS]",
		Short: "Creates a new post or page",
	}

	cmd.RunE = runNew

	return cmd
}

func runNew(c *cobra.Command, args []string) error {
	fmt.Printf("%v\n", args)
	return nil
}
