package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jbpratt78/tmdb"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search tmdb",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			log.Fatal("must supply at least 2 arguments")
		}
		apikey, exists := os.LookupEnv("API_KEY")

		if exists {
			c := tmdb.New(apikey)
			switch {
			case args[0] == "movie":
				s, err := c.SearchMovie(args[1:])
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(s)
			case args[0] == "company":
				s, err := c.SearchCompany(args[1])
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(s)
			case args[0] == "tv":
				s, err := c.SearchTv(args[1:])
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(s)
			case args[0] == "person":
				s, err := c.SearchPerson(args[1:])
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(s)
			case args[0] == "collection":
				s, err := c.SearchCollection(args[1:])
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(s)
			case args[0] == "keyword":
				s, err := c.SearchKeyword(args[1:])
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(s)
			case args[0] == "multi":
				fmt.Println("TODO")
			default:
			}
		} else {
			log.Fatal("no api key found")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
