package tmdb

import "fmt"

var trendingURI = BaseURI + "/trending"

func (c *Client) GetTrending(mediaType, timeWindow string) error {
	if mediaType == "" {
		mediaType = "all"
	}
	if timeWindow == "" {
		timeWindow = "week"
	}

	v, err := c.constructURI(trendingURI + "/" + mediaType + "/" + timeWindow)
	if err != nil {
		return err
	}
	s, err := c.request(*v, "GET")
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
