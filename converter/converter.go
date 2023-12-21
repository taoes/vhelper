package converter

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

var (
	kind    string
	Command = &cobra.Command{
		Use:   "converter",
		Short: "Provide the ability to convert resources such as numbers and images",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 || strings.Compare(args[0], "") == 0 {
				return errors.New("Please enter a valid content")
			}
			// TODO
			return nil
		},
	}
)

func init() {
	Command.Flags().StringVar(&kind, "type", "number", "Converter Type, [number,picture,translate]")

}
