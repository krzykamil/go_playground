/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "See the list as is",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		auth_key := viper.GetString("NOTION_KEY")
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.notion.com/v1/pages/4adbdaac-f4bb-4fe8-8c8e-d40ed2a862ef", nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("Authorization", "Bearer "+auth_key)
		req.Header.Add("Notion-Version", "2022-06-28")
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}
		log.Println(string([]byte(body)))

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
