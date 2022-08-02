package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/sh3rp/stringen/pkg/codec"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var claimsStr string
var secret string
var tokenStr string

var ejwtCmd = &cobra.Command{
	Use:   "ejwt",
	Short: "Generate a JWT token",
	Run: func(cmd *cobra.Command, args []string) {
		claims := make(map[string]interface{})

		if claimsStr != "" {
			claimKVs := strings.Split(claimsStr, ",")

			for _, kv := range claimKVs {
				tokens := strings.Split(kv, "=")
				claims[tokens[0]] = tokens[1]
			}
		}

		token := codec.GenerateJWTToken(secret, claims)

		fmt.Printf("%s", token)
	},
}

var djwtCmd = &cobra.Command{
	Use:   "djwt",
	Short: "Decode a JWT token",
	Run: func(cmd *cobra.Command, args []string) {
		if useStdin {
			reader := bufio.NewReader(os.Stdin)
			buf := make([]byte, 0, 4*1024)
			tokenStr = ""
			for {
				n, err := reader.Read(buf[:cap(buf)])
				buf = buf[:n]
				if n == 0 {
					if err == nil {
						continue
					}
					if err == io.EOF {
						break
					}
					log.Fatal(err)
				}
				tokenStr += base64.StdEncoding.EncodeToString(buf)
				if err != nil && err != io.EOF {
					log.Fatal(err)
				}
			}
			data, err := codec.DecodeJWTToken(secret, tokenStr)
			if err != nil {
				fmt.Printf("Error decoding token: %+v\n", err)
			} else {
				fmt.Println(data)
			}
		} else {
			data, err := codec.DecodeJWTToken(secret, tokenStr)
			if err != nil {
				fmt.Printf("Error decoding token: %+v\n", err)
			} else {
				fmt.Println(data)
			}
		}
	},
}

func init() {
	ejwtCmd.Flags().StringVarP(&claimsStr, "claims", "c", "", "Claims in the form of comma-delimited key-value pairs using '=' to separate the pairs (example: \"key1=value1,key2=value2,...\")")
	ejwtCmd.Flags().StringVarP(&secret, "secret", "s", "secret", "Secret used to encode/decode the JWT token")

	djwtCmd.Flags().StringVarP(&secret, "secret", "s", "secret", "Secret used to encode/decode the JWT token")
	djwtCmd.Flags().StringVarP(&tokenStr, "token", "t", "", "Token to decode")
	djwtCmd.Flags().BoolVarP(&useStdin, "in", "i", false, "Use stdin to decode token")
}
