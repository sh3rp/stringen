package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/sh3rp/stringen"
	"github.com/spf13/cobra"
)

var stringToEncode string

var encodeBase64Cmd = &cobra.Command{
	Use:   "e64",
	Short: "Encode to a base64 string based on a passed parameter or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		var str string
		if useStdin {
			str = stringen.EncodeBase64(bufio.NewReader(os.Stdin))
		} else {
			str = stringen.EncodeBase64(bufio.NewReader(bytes.NewBufferString(stringToEncode)))
		}
		fmt.Println(str)
	},
}

var decodeBase64Cmd = &cobra.Command{
	Use:   "d64",
	Short: "Decode to a base64 string based on a passed parameter or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		var str string
		if useStdin {
			str = stringen.DecodeBase64(bufio.NewReader(os.Stdin))
		} else {
			str = stringen.DecodeBase64(bufio.NewReader(bytes.NewBufferString(stringToEncode)))
		}
		fmt.Println(str)
	},
}

func init() {
	encodeBase64Cmd.Flags().StringVarP(&stringToEncode, "string", "s", "", "Number of characters to generate")
	encodeBase64Cmd.Flags().BoolVarP(&useStdin, "in", "i", false, "Encode from stdin")
	decodeBase64Cmd.Flags().StringVarP(&stringToEncode, "string", "s", "", "Number of characters to generate")
	decodeBase64Cmd.Flags().BoolVarP(&useStdin, "in", "i", false, "Encode from stdin")
}
