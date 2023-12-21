package time

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
	"vhelper/utils"
)

var (
	format      string
	TimeCommand = &cobra.Command{
		Use:   "time",
		Short: "Outputs past or future time data",
		RunE: func(cmd *cobra.Command, args []string) error {
			currentTime := time.Now() //当前时间

			if len(args) >= 1 {
				m, err := time.ParseDuration(args[0])
				if err != nil {
					_ = fmt.Errorf("after time format error\n")
					return err
				}
				currentTime = currentTime.Add(m)
			}
			currentYear := currentTime.Year()     //当前年
			currentMonth := currentTime.Month()   //当前月
			currentDay := currentTime.Day()       //当前日
			currentHour := currentTime.Hour()     //当前小时小时
			currentMinute := currentTime.Minute() //当前分钟
			currentSecond := currentTime.Second() //当前秒

			//打印结果
			fmt.Printf("Second timestamp: %d \nMillsecond Timestamp: %d\n", currentTime.Unix(), currentTime.UnixMilli())
			fmt.Printf("Time ：%d-%s-%s %s:%s:%s\n", currentYear,
				utils.AddPrefix(int(currentMonth), "0"),
				utils.AddPrefix(currentDay, "0"),
				utils.AddPrefix(currentHour, "0"),
				utils.AddPrefix(currentMinute, "0"),
				utils.AddPrefix(currentSecond, "0"))
			fmt.Println()
			return nil
		},
	}
)

func init() {
	TimeCommand.Flags().StringVar(&format, "format", "", "Time Format")
}
