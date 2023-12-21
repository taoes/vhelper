package ws

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

var (
	url     string
	Command = &cobra.Command{
		Use:   "ws",
		Short: "Establish a WebSocket link for easy debugging of network services",
		RunE: func(cmd *cobra.Command, args []string) error {
			if strings.Compare(url, "") == 0 {
				return errors.New("Please enter a valid WebSocket address")
			}
			// TODO
			return nil
		},
	}
)

func init() {
	Command.Flags().StringVar(&url, "url", "", "WebSocket Url Address, eg: 'ws://xxxx' OR 'wss://xxx'")

}
