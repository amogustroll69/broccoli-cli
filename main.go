package main

import (
	"strconv"
	"time"

	"github.com/krillissue/broccoli"
	"github.com/spf13/cobra"
)

var (
	application_port string
	storage_folder   string
	cors_origin      string
)

func main() {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve Broccoli",
		Long:  "Serve a Broccoli file server",
		Run: func(cmd *cobra.Command, args []string) {
			port, err := strconv.ParseUint(application_port, 0, 16)

			if err != nil {
				panic(err)
			}

			defer broccoli.NewBroccoli(uint16(port), storage_folder, cors_origin)
			for {
				time.Sleep(time.Second * 10)
			}
		},
	}

	cmd.Flags().StringVarP(&application_port, "port", "p", "8080", "Port to expose webserver on.")
	cmd.Flags().StringVarP(&storage_folder, "folder", "f", "uploads", "Folder to store uploaded data on.")

	rootCmd := &cobra.Command{Use: "broccoli-cli"}
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
