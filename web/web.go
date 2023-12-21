package web

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	host       string
	path       string
	WebCommand = &cobra.Command{
		Use:   "web",
		Short: "Start the static resource service on the local port",
		Long:  "Start the static resource service on the local port",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("The static resource service will be started on port  %s For path %s .....", host, path)
			fmt.Println()
			http.Handle("/", http.FileServer(http.Dir(path)))
			_ = http.ListenAndServe(host, nil)
			return nil
		},
	}
)

func init() {
	WebCommand.Flags().StringVar(&host, "host", ":1234", "Web service host")
	WebCommand.Flags().StringVar(&path, "path", ".", "Static Resource path")
}
