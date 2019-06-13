package cmd

import (
	"os"

	"github.com/jbpratt78/tmdb"
	"github.com/spf13/cobra"
)

var mediaType, timeWindow string

var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "trending tmdb",
	Run: func(cmd *cobra.Command, args []string) {
		apikey, exists := os.LookupEnv("API_KEY")
		if exists {
			c := tmdb.New(apikey)
			err := c.GetTrending(mediaType, timeWindow)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	trendingCmd.Flags().StringVarP(&mediaType, "media-type", "m", "", "Media type of trending")
	trendingCmd.Flags().StringVarP(&timeWindow, "time-window", "t", "", "Time window of trending")
	rootCmd.AddCommand(trendingCmd)
}
