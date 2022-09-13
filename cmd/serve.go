package cmd

import (
	"api/internal/http"
	"fmt"
	"github.com/spf13/cobra"
)

// art generator: https://patorjk.com/software/taag/#p=display&f=Small&t=SHORT-ME
var banner = `
  ___ _  _  ___  ___ _____    __  __ ___ 
 / __| || |/ _ \| _ \_   _|__|  \/  | __|
 \__ \ __ | (_) |   / | ||___| |\/| | _| 
 |___/_||_|\___/|_|_\ |_|    |_|  |_|___|
`

var serveCommand = &cobra.Command{
	Use:              "serve",
	TraverseChildren: true,
	Short:            "Start the api server",
	Long:             `Starts the api server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\u001B[1;36m" + banner + "\u001B[0m")

		http.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCommand)
}
