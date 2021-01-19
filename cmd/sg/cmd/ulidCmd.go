package cmd

import (
	"fmt"
	"strings"

	"github.com/sh3rp/stringen"
	"github.com/spf13/cobra"
)

var toLower bool

var ulidCmd = &cobra.Command{
	Use:   "ulid",
	Short: "Generate a random ULID string",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			str := stringen.GenUlid()
			if toLower {
				str = strings.ToLower(str)
			}
			fmt.Println(str)
		}
	},
}

func init() {
	ulidCmd.Flags().BoolVarP(&toLower, "lower", "l", false, "Output ULID in all uppercase")
	ulidCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of ULIDs to generate")
}
