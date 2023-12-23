package random

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
	"strconv"
	"strings"
	"time"
)

var (
	format   string
	minValue int64
	maxValue int64
	count    int
	len      int
	numshow  bool
	letter   string
	Command  = &cobra.Command{
		Use:   "random",
		Short: "Randomly generate numbers or unique IDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			rand.Seed(uint64(time.Now().Unix()))

			if format == "string" {
				str := []byte(letter)
				for i := 0; i < count; i++ {
					res := ""
					for i := 0; i < 10; i++ {
						res += fmt.Sprintf("%c", str[rand.Intn(strings.Count(letter, "")-1)])
					}
					fmt.Printf("%s\n", res)
				}
				return nil
			}

			if format == "number" {
				for i := 0; i < count; i++ {
					randomValue := rand.Int63n(maxValue) + minValue
					if !numshow {
						fmt.Printf(strconv.FormatInt(randomValue, 10))
					} else {
						fmt.Printf("%d %s", i, strconv.FormatInt(randomValue, 10))
					}
					fmt.Println()
				}
				return nil
			}

			if format == "float" {
				for i := 0; i < count; i++ {
					lenStr := "%." + strconv.Itoa(len) + "f\n"
					result := rand.Float32()*float32(maxValue) + float32(minValue)
					if !numshow {
						fmt.Printf(lenStr, result)
					} else {
						fmt.Printf(lenStr, i, result)
					}
				}
				return nil
			}

			return nil
		},
	}
)

func init() {
	Command.Flags().StringVar(&format, "type", "number", "随机类型: string,number,float")
	Command.Flags().Int64Var(&minValue, "min", 0, "随机数字最小范围")
	Command.Flags().Int64Var(&maxValue, "max", 100, "随机数字最大范围")
	Command.Flags().IntVar(&count, "count", 1, "数量")
	Command.Flags().BoolVar(&numshow, "numshow", false, "是否输出序号")
	Command.Flags().StringVar(&letter, "letter", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "随机字符串取样范围")
	Command.Flags().IntVar(&len, "len", 20, "随机数据的输出长度")

}
