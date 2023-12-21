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
		Short: "en",
		Long:  "字符串解码",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				err := errors.New("未给定需要解码的字符串")
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

			return errors.New("无效的解码方式")
		},
	}
)

func init() {
	DecodeCommand.Flags().StringVar(&kind, "type", "", "解码方式")
}
