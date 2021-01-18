package cmd

import (
	"fmt"

	"github.com/sh3rp/stringen"
	"github.com/spf13/cobra"
)

var randCmd = &cobra.Command{
	Use:   "rand",
	Short: "Generate a random string",
	Run: func(cmd *cobra.Command, args []string) {
		str := stringen.GenRandomCharacters(20, stringen.CharTypeAlphaNumericSpecial)
		fmt.Println(str)
	},
}
