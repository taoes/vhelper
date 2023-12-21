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
		Short: "en",
		Long:  "字符串编码",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				err := errors.New("未给定需要编码的字符串")
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

			return errors.New("无效的编码方式")
		},
	}
)

func init() {
	EncodeCommand.Flags().StringVar(&kind, "type", "", "编码方式")
}
