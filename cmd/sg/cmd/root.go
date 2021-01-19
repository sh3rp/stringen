package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var count int

var rootCmd = &cobra.Command{
	Use:   "sg",
	Short: "String Generator",
	Long:  "Generates a variety of strings",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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
	rootCmd.AddCommand(encodeBase64Cmd)
	rootCmd.AddCommand(decodeBase64Cmd)
	rootCmd.AddCommand(ulidCmd)
}
