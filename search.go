package tmdb

import (
	"net/url"
)

func (c *Client) searchMovie(params url.Values) error {
	if params == nil {
		return ErrNoParams
	}
	return nil
}

func (c *Client) searchCompanies(params url.Values) error {
	if params == nil {
		return ErrNoParams
	}
	return nil
}

func (c *Client) searchPeople(params url.Values) error {
	if params == nil {
		return ErrNoParams
	}
	return nil
}
