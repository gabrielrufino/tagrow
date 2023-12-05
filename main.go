package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{}
	root.AddCommand(&cobra.Command{
		Use:       "up [segment]",
		Short:     "Updates the semantic version",
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		ValidArgs: []string{"major", "minor", "patch"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Update %s\n", args[0])
		},
	})
	root.Execute()
}
