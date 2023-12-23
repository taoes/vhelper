package encode

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
	"strconv"
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

			if strings.Compare(kind, "url") == 0 {
				res := url.QueryEscape(args[0])
				fmt.Printf("%s", res)
				fmt.Println()
				return nil
			}

			if strings.Compare(kind, "unicode") == 0 {
				textQuoted := strconv.QuoteToASCII(args[0])
				fmt.Printf("%v", textQuoted[1:len(textQuoted)-1])
				fmt.Println()
				return nil
			}

			if strings.Compare(kind, "md5") == 0 {
				md5Value := md5.Sum([]byte(args[0]))
				fmt.Printf("%x", md5Value)
				fmt.Println()
				return nil
			}

			if strings.Compare(kind, "sha1") == 0 {
				o := sha1.New()
				o.Write([]byte(args[0]))
				sha1Value := hex.EncodeToString(o.Sum(nil))
				fmt.Printf("%v", sha1Value)
				fmt.Println()
				return nil
			}

			return errors.New("no valid encoding method is provided")
		},
	}
)

func init() {
	EncodeCommand.Flags().StringVar(&kind, "type", "base64", "Encode type: base64, url, unicode, md5, sha1 ")
}
