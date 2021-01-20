package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sh3rp/stringen"
	"github.com/spf13/cobra"
)

var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "Generate SHA256 checksum hash",
	Run: func(cmd *cobra.Command, args []string) {
		if useStdin {
			reader := bufio.NewReader(os.Stdin)
			buf := make([]byte, 0, 4*1024)
			var data string
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
				data += base64.StdEncoding.EncodeToString(buf)
				if err != nil && err != io.EOF {
					log.Fatal(err)
				}
			}
			hash := stringen.GenSha256Hash(data)
			fmt.Println(hash)
		} else {
			for i := 0; i < count; i++ {
				material := stringen.GenRandomCharacters(20, stringen.CharTypeAlphaNumericSpecial)
				hash := stringen.GenSha256Hash(material)
				fmt.Println(hash)
			}
		}
	},
}

func init() {
	sha256Cmd.Flags().IntVarP(&count, "count", "c", 1, "Number of hashes to generate (if not using stdin)")
	sha256Cmd.Flags().BoolVarP(&useStdin, "in", "i", false, "Use standard input as the hashing material")
}
