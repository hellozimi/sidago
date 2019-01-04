package main

import (
	"github.com/spf13/cobra"
)

type command interface {
	Command() *cobra.Command
}
