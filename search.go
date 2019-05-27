package tmdb

import (
	"fmt"
)

var searchURI = BaseURI + "/search"

func (c *Client) SearchMovie(query string) error {
	if query == "" {
		return ErrNoQueryParam
	}

	searchURI += "/movie"

	var options = map[string]string{
		"query": query,
	}

	v, err := c.constructURI(searchURI, options)
	if err != nil {
		return err
	}

	res, err := c.request(*v, "GET")
	if err != nil {
		return err
	}

	fmt.Println(HandleResponse(res))

	return nil
}

func (c *Client) searchCompany(query string) error {
	if query == "" {
		return ErrNoQueryParam
	}

	searchURI += "/company"

	return nil
}

func (c *Client) searchPeople(query string) error {
	if query == "" {
		return ErrNoQueryParam
	}

	searchURI += "/people"

	return nil
}
