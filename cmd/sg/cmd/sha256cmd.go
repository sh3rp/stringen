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
		for i := 0; i < count; i++ {
			material := stringen.GenRandomCharacters(20, stringen.CharTypeAlphaNumericSpecial)
			hash := stringen.GenSha256Hash(material)
			fmt.Println(hash)
		}
	},
}

func init() {
	sha256Cmd.Flags().IntVarP(&count, "count", "c", 1, "Number of ULIDs to generate")
}
