package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"os"
	"vhelper/img/qrcode"
	"vhelper/math/hex"
	"vhelper/net/agent"
	"vhelper/net/ip"
	"vhelper/net/proxy"
	"vhelper/net/search"
	"vhelper/net/web"
	"vhelper/net/webpage"
	"vhelper/net/ws"
	"vhelper/safe/decode"
	"vhelper/safe/encode"
	"vhelper/safe/password"
	"vhelper/tool/jsonFormat"
	"vhelper/tool/random"
	"vhelper/tool/time"
)
import _ "github.com/spf13/cobra"

var (
	lang      string
	searchEng string
)

var rootCmd = &cobra.Command{
	Use:               "",
	DisableAutoGenTag: true,
	Short:             "为开发者提供的CLI助手,So you Know, love Command Line MORE!",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(lang) > 0 {
			viper.Set("lang", lang)
			fmt.Println("应用程序语言更新为" + lang)
		}

		if len(searchEng) > 0 {
			viper.Set("searchEngineName", searchEng)
			fmt.Println("默认搜索引擎更新为：" + searchEng)
		}

		err := viper.WriteConfig()
		fmt.Println()
		return err
	},
}

func init() {
	initConfig()
	initCmdFlag()
	rootCmd.AddCommand(web.Command)
	rootCmd.AddCommand(time.TimeCommand)
	rootCmd.AddCommand(encode.EncodeCommand)
	rootCmd.AddCommand(decode.DecodeCommand)
	rootCmd.AddCommand(qrcode.QrCodeCommand)
	rootCmd.AddCommand(ip.Command)
	rootCmd.AddCommand(ws.Command)
	rootCmd.AddCommand(agent.Command)
	rootCmd.AddCommand(hex.Command)
	rootCmd.AddCommand(random.Command)
	rootCmd.AddCommand(password.Command)
	rootCmd.AddCommand(jsonFormat.Command)
	rootCmd.AddCommand(search.Command)
	rootCmd.AddCommand(proxy.Command)
	rootCmd.AddCommand(webpage.Command)
}

func initCmdFlag() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.SetVersionTemplate("V0.1")
	rootCmd.PersistentFlags().StringVar(&lang, "lang", "", "国际化语言")
	rootCmd.PersistentFlags().StringVar(&searchEng, "searchEngine", "", "默认搜索引擎")
	rootCmd.PersistentFlags().String("author", "", "作者: 燕归来兮 https://www.zhoutao123.com")
}

func initConfig() {
	dir, err := os.UserConfigDir()
	if err != nil {
		os.Exit(-1)
	}
	viper.AddConfigPath(dir)
	viper.SetConfigName("vhelper")
	viper.SetConfigType("toml")
	_ = viper.ReadInConfig()
}

func main() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(-1)
	}
}
