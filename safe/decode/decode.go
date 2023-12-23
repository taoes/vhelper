package decode

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
	"strconv"
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

			if strings.Compare(kind, "url") == 0 {
				res, err := url.QueryUnescape(args[0])
				if err != nil {
					return err
				}
				fmt.Printf("%s", res)
				fmt.Println()
				return nil
			}

			if strings.Compare(kind, "unicode") == 0 {
				str, err := strconv.Unquote(strings.Replace(strconv.Quote(args[0]), `\\u`, `\u`, -1))
				if err != nil {
					return err
				}
				fmt.Printf("%v", str)
				fmt.Println()
				return nil
			}

			return errors.New("no valid decoding method is provided!")
		},
	}
)

func init() {
	DecodeCommand.Flags().StringVar(&kind, "type", "base64", "Decode type: base64, url, unicode")
}
