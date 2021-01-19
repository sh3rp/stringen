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
		for i := 0; i < count; i++ {
			uu, err := uuid.NewRandom()
			if err != nil {
				fmt.Printf("Error: %+v", err)
				return
			}
			fmt.Println(uu.String())
		}
	},
}

func init() {
	uuidCmd.Flags().BoolVarP(&toLower, "lower", "l", false, "Output ULID in all uppercase")
	uuidCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of ULIDs to generate")
}
