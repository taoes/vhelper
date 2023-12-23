package agent

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

var (
	Command = &cobra.Command{
		Use:   "agent",
		Short: "Analyze browser agent content",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 || strings.Compare(args[0], "") == 0 {
				return errors.New("Please enter a valid agent")
			}
			// TODO
			return nil
		},
	}
)

func init() {

}
