package jsonFormat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Command = &cobra.Command{
		Use:   `json`,
		Short: `json`,
		Long:  `JSON`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("未给定JSON内容")
			}
			var out bytes.Buffer
			err := json.Indent(&out, []byte(args[0]), "", "\t")

			if err != nil {
				return err
			}

			_, _ = out.WriteTo(os.Stdout)
			fmt.Println()
			return nil
		},
	}
)
