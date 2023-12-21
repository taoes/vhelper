package time

import (
	"github.com/spf13/cobra"
)

var (
	format      string
	TimeCommand = &cobra.Command{
		Use:   "random",
		Short: "Randomly generate numbers or unique IDs",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
)

func init() {
	TimeCommand.Flags().StringVar(&format, "type", "number", "Random  Type")
}
