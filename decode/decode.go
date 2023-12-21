package decode

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	kind          string
	DecodeCommand = &cobra.Command{
		Use:   "decode",
		Short: "Decode strings or file contents",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				err := errors.New("not found content")
				return err
			}

			if strings.Compare(kind, "base64") == 0 {
				res, err := base64.StdEncoding.DecodeString(args[0])
				if err != nil {
					return err
				}
				fmt.Printf(string(res))
				fmt.Println()
				return nil
			}

			if strings.Compare(kind, "") == 0 {

			}

			return errors.New("no valid decoding method is provided!")
		},
	}
)

func init() {
	DecodeCommand.Flags().StringVar(&kind, "type", "base64", "decode type")
}
