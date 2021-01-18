package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sg",
	Short: "String generator",
	Long:  "Generate all manner of strangs",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(sha256Cmd)
	rootCmd.AddCommand(uuidCmd)
	rootCmd.AddCommand(randCmd)
}
