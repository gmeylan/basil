package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var artichautCmd = &cobra.Command{
	Use:   "artichaut",
	Short: "arti short",
	Long:  `arti long`,
	Run: func(cmd *cobra.Command, args []string) {
		artichaut()
	},
}

func artichaut() {
	fmt.Println("Arti")

}

func init() {
	rootCmd.AddCommand(artichautCmd)
}
