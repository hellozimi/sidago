package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type initCmd struct {
	cmd     *cobra.Command
	rootCmd *cobra.Command
	command
}

func (i *initCmd) Command() *cobra.Command {
	return i.cmd
}

func newInitCommand(rootCmd *cobra.Command) command {
	c := &initCmd{rootCmd: rootCmd}
	cmd := &cobra.Command{
		Use:   "init [PATH]",
		Short: "Creates a new sida",
		Long: `Creates a new sida at the target path.
The new site will be generated with the correct content to begin writing.`,
		RunE: c.runInit,
	}

	c.cmd = cmd

	cmd.Flags().Bool("force", false, "forces initialization in a non-empty directory")

	return c
}

func (i *initCmd) runInit(c *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("needs to provide a path")
	}

	p, err := filepath.Abs(filepath.Clean(args[0]))

	if err != nil {
		return err
	}

	force, _ := c.Flags().GetBool("force")

	return i.doRunInit(p, force)
}

func (i *initCmd) doRunInit(basepath string, force bool) error {

	dirs := []string{
		filepath.Join(basepath, "posts"),
		filepath.Join(basepath, "pages"),
		filepath.Join(basepath, "layout"),
	}

	s, err := os.Stat(basepath)
	if err != nil && os.IsExist(err) {
		return err
	}

	if s != nil {
		if !s.Mode().IsDir() {
			return fmt.Errorf("target path exists but isn't a directory")
		}
	}

	for _, dir := range dirs {
		// Mkdir dir
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	if err := i.makeConfig(basepath); err != nil {
		return err
	}

	fmt.Printf("🚀 Your new sida is now created in %s\n\n", basepath)

	return nil
}

func (i *initCmd) makeConfig(basepath string) error {
	cfg := []byte(`
title = "My new Sida"
`)

	return ioutil.WriteFile(filepath.Join(basepath, "config.toml"), cfg, 0777)
}
