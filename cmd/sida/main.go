package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sida [OPTIONS]",
		Short: "A static site generator",
	}

	return cmd
}

func main() {
	cmd := newCommand()
	cmd.SetOutput(os.Stdout)
	newCmd := newNewCommand()
	initCmd := newInitCommand()

	cmd.AddCommand(newCmd, initCmd)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
