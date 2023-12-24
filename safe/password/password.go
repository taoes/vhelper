package password

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
	"strconv"
	"strings"
	"time"
)

var (
	count   int
	len     int
	numshow bool
	Command = &cobra.Command{
		Use:   "password",
		Short: "生成随机密码",
		RunE: func(cmd *cobra.Command, args []string) error {
			rand.Seed(uint64(time.Now().Unix()))
			letter := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
			str := []byte(letter)
			for i := 0; i < count; i++ {
				res := ""
				if numshow {
					res += strconv.Itoa(i) + " "
				}
				for i := 0; i < len; i++ {
					res += fmt.Sprintf("%c", str[rand.Intn(strings.Count(letter, "")-1)])
				}
				fmt.Printf("%s\n", res)
			}
			return nil
		},
	}
)

func init() {
	Command.Flags().IntVar(&count, "count", 1, "生成数量")
	Command.Flags().BoolVar(&numshow, "numshow", false, "显示序号")
	Command.Flags().IntVar(&len, "len", 20, "密码长度")

}
