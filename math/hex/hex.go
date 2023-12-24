package hex

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	base    int
	Command = &cobra.Command{
		Use:   "hex",
		Short: "提供数据之间的进制转换",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("未给定需要转换的数字")
			}
			sv := args[0]
			tenNumber, err := strconv.ParseInt(sv, base, 32)
			if err != nil {
				return err
			}
			fmt.Printf("十六进制: %s\n", strconv.FormatInt(tenNumber, 16))
			fmt.Printf("十进制: %s\n", strconv.FormatInt(tenNumber, 10))
			fmt.Printf("八进制: %s\n", strconv.FormatInt(tenNumber, 8))
			fmt.Printf("二进制: %s\n", strconv.FormatInt(tenNumber, 2))

			return nil
		},
	}
)

func init() {
	Command.Flags().IntVar(&base, "base", 10, "参数的进制: 16,10,8,2")

}
