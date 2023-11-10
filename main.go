package main

import (
	"log"
	"os"
	"time"

	"github.com/krillissue/broccoli"
	"github.com/spf13/cobra"
)

var (
	host_addr      string
	storage_folder string
	cors_origin    string
)

func main() {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve Broccoli",
		Long:  "Serve a Broccoli file server",
		Run: func(cmd *cobra.Command, args []string) {
			info, err := os.Stat(storage_folder)

			if err != nil {
				log.Fatalln(err)
			}

			if !info.IsDir() {
				log.Fatalln("Folder must be a directory")
			}

			broccoli.NewBroccoli(host_addr, storage_folder, cors_origin)

			for {
				time.Sleep(time.Second * 10)
			}
		},
	}

	cmd.Flags().StringVar(&host_addr, "host", "0.0.0.0:8080", "Host address to expose webserver on.")
	cmd.Flags().StringVarP(&storage_folder, "folder", "f", "server_files", "Folder containing stuff for the webserver to host.")
	cmd.Flags().StringVarP(&cors_origin, "origin", "o", "*", "CORS origin for the webserver to work on.")

	rootCmd := &cobra.Command{Use: "broccoli-cli"}
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
