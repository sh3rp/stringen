package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a random UUID",
	Run: func(cmd *cobra.Command, args []string) {
		uu, err := uuid.NewRandom()
		if err != nil {
			fmt.Printf("Error: %+v", err)
			return
		}
		fmt.Println(uu.String())
	},
}
