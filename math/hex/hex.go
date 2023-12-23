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
		Short: "Provide the ability to convert resources such as numbers",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("未给定需要转换的数字")
			}
			sv := args[0]
			tenNumber, err := strconv.ParseInt(sv, base, 32)
			if err != nil {
				return err
			}
			fmt.Printf("HEX Value: %s\n", strconv.FormatInt(tenNumber, 16))
			fmt.Printf("DEC Value: %s\n", strconv.FormatInt(tenNumber, 10))
			fmt.Printf("OCT Value: %s\n", strconv.FormatInt(tenNumber, 8))
			fmt.Printf("BIN Value: %s\n", strconv.FormatInt(tenNumber, 2))

			return nil
		},
	}
)

func init() {
	Command.Flags().IntVar(&base, "base", 10, "Current Base: 16,10,8,2")

}
