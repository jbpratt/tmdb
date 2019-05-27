package cmd

import (
	"log"
	"os"

	"github.com/jbpratt78/tmdb"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search tmdb",
	Run: func(cmd *cobra.Command, args []string) {

		apikey, exists := os.LookupEnv("API_KEY")

		if exists {
			c := tmdb.New(apikey)
			err := c.SearchMovie(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
