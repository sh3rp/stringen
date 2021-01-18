package cmd

import (
	"fmt"

	"github.com/sh3rp/stringen"
	"github.com/spf13/cobra"
)

var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "Generate SHA256 checksum hash",
	Run: func(cmd *cobra.Command, args []string) {
		material := stringen.GenRandomCharacters(20, stringen.CharTypeAlphaNumericSpecial)
		hash := stringen.GenSha256Hash(material)
		fmt.Println(hash)
	},
}
