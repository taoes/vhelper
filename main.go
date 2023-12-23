package main

import (
	"errors"
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
	"vhelper/net/web"
	"vhelper/net/ws"
	"vhelper/safe/decode"
	"vhelper/safe/encode"
	"vhelper/tool/jsonFormat"
	"vhelper/tool/random"
	"vhelper/tool/time"
)
import _ "github.com/spf13/cobra"

var (
	lang string
	zone string
)

var rootCmd = &cobra.Command{
	Use:   "config",
	Short: "Terminal command-line assistant for development engineers",
	Long:  "Terminal command-line assistant for development engineers",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(lang) > 0 {
			viper.Set("lang", lang)
			fmt.Println("The current language has been updated to " + lang)
		}

		err := viper.WriteConfig()
		return err
		return errors.New("No valid command found, please refer to the following documentation for usage")
	},
}

func init() {
	initConfig()
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
	rootCmd.AddCommand(random.Command)
	rootCmd.AddCommand(jsonFormat.Command)
	rootCmd.AddCommand(proxy.Command)
}

func initConfig() {
	dir, err := os.UserConfigDir()
	if err != nil {
		os.Exit(-1)
	}
	viper.AddConfigPath(dir)
	viper.SetConfigName("vhelper")
	viper.SetConfigType("toml")
	viper.ReadInConfig()
	rootCmd.PersistentFlags().StringVar(&lang, "lang", "", "language")

}

func main() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(-1)
	}
}
