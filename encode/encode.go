package encode

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	kind          string
	EncodeCommand = &cobra.Command{
		Use:   "encode",
		Short: "Encode strings or file contents",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				err := errors.New("not found content")
				return err
			}

			if strings.Compare(kind, "base64") == 0 {
				res := base64.StdEncoding.EncodeToString([]byte(args[0]))
				fmt.Printf(res)
				fmt.Println()
				return nil
			}

			if strings.Compare(kind, "") == 0 {

			}

			return errors.New("no valid encoding method is provided")
		},
	}
)

func init() {
	EncodeCommand.Flags().StringVar(&kind, "type", "base64", "encode type ")
}
