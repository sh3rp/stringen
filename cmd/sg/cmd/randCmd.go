package cmd

import (
	"fmt"
	"github.com/sh3rp/stringen/pkg/codec"

	"github.com/spf13/cobra"
)

var length int
var strType string

var randCmd = &cobra.Command{
	Use:   "rand",
	Short: "Generate a random string",
	Run: func(cmd *cobra.Command, args []string) {
		var t int
		switch strType {
		case "a":
			t = codec.CharTypeAlpha
		case "an":
			t = codec.CharTypeAlphaNumeric
		case "ans":
			t = codec.CharTypeAlphaNumericSpecial
		}
		str := codec.GenRandomCharacters(length, codec.CharType(t))
		fmt.Println(str)
	},
}

func init() {
	randCmd.Flags().IntVarP(&length, "length", "l", 20, "Number of characters to generate")
	randCmd.Flags().StringVarP(&strType, "type", "t", "ans", "Type of characters to use (a: alpha, an: alphanumeric, ans: alphanumeric and symbols)")
}
