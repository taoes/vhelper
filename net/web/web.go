package web

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	host    string
	path    string
	Command = &cobra.Command{
		Use:   "web",
		Short: "Start the static resource service on the local port",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("The static resource service will be started on port  %s For path %s .....", host, path)
			fmt.Println()
			http.Handle("/", http.FileServer(http.Dir(path)))
			err := http.ListenAndServe(host, nil)
			return err
		},
	}
)

func init() {
	Command.Flags().StringVar(&host, "host", ":9999", "Web service host")
	Command.Flags().StringVar(&path, "path", ".", "Static Resource path")
}
