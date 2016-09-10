package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of pick",
		Long:  `The version command prints the version of pick`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("pick %s\n", Config.Version)
		},
	})
}
