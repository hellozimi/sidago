package main

import (
	"path/filepath"

	"github.com/hellozimi/sidago/internal/builder"
	"github.com/hellozimi/sidago/internal/builder/config"

	"github.com/hellozimi/sidago/sida"
	"github.com/spf13/cobra"
)

type generateCmd struct {
	path string
	cmd  *cobra.Command
}

func newGenerateCommand() *cobra.Command {
	c := &generateCmd{}
	cmd := &cobra.Command{
		Use:   "generate [OPTIONS]",
		Short: "Generates html files from markdown",
		RunE:  c.runGenerate,
	}

	c.cmd = cmd

	cmd.Flags().StringVarP(&c.path, "path", "p", "./build", "pass to generate the content in any other path than ./build")

	return cmd
}

func (g *generateCmd) runGenerate(c *cobra.Command, args []string) error {
	path := args[0]
	p, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return err
	}

	_, err = sida.IsSida(p)
	if err != nil {
		return err
	}

	cfg, err := config.FromFile(p, "config")
	if err != nil {
		return err
	}

	s := builder.NewSida(p, cfg)
	s.Build()

	return nil
}
