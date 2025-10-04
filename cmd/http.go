package cmd

import (
	"fmt"

	server2 "github.com/raffreitas/fc-hex-go/adapters/web/server"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start the HTTP web server",
	Long: `Starts the HTTP web server for the application.

This command initializes and runs the web server that exposes the HTTP APIs
following hexagonal architecture patterns. The server will run continuously
and listen for HTTP requests until interrupted.

The server will start on the default configured port and be ready to
handle HTTP requests to the application endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebserver()
		server.Service = &productService
		fmt.Println("Webserver has been started")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
