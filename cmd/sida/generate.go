package main

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/hellozimi/sidago/internal/builder"
	"github.com/hellozimi/sidago/internal/builder/config"

	"github.com/hellozimi/sidago/sida"
	"github.com/spf13/cobra"
)

type generateCmd struct {
	buildDir     string
	skipsBaseUrl bool
	cmd          *cobra.Command
	rootCmd      *cobra.Command
	command
}

func (g *generateCmd) Command() *cobra.Command {
	return g.cmd
}

func newGenerateCommand(rootCmd *cobra.Command) command {
	c := &generateCmd{rootCmd: rootCmd}
	cmd := &cobra.Command{
		Use:   "generate [OPTIONS]",
		Short: "Generates html files from markdown",
		RunE:  c.runGenerate,
	}

	c.cmd = cmd

	cmd.Flags().StringVarP(&c.buildDir, "build", "b", "./build", "pass to generate the content in any other path than ./build - relative to project cwd")
	cmd.Flags().BoolVarP(&c.skipsBaseUrl, "skip-baseurl", "s", false, "skips url transformation to add base url (defaults to false)")

	return c
}

func (g *generateCmd) runGenerate(c *cobra.Command, args []string) error {
	path := g.rootCmd.PersistentFlags().Lookup("path").Value.String()

	p, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return err
	}

	ok := sida.IsSida(p)
	if !ok {
		return fmt.Errorf("target directory isn't a Sida")
	}

	cfg, err := config.FromFile(p, "config")
	if err != nil {
		return err
	}

	baseURL, err := fixBaseURL(cfg.GetString("base_url"), g.skipsBaseUrl)
	if err != nil {
		return err
	}

	cfg.Set("base_url", baseURL)

	s := builder.NewSida(p, cfg)
	return s.Build()
}

func fixBaseURL(s string, skipBaseURL bool) (string, error) {
	if skipBaseURL {
		return "/", nil
	}

	if !strings.HasSuffix(s, "/") {
		s = s + "/"
	}
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
