package tmdb

import (
	"encoding/json"
	"net/url"
)

var searchURI = BaseURI + "/search"

func (c *Client) SearchMovie(querySlice []string) (*SearchMovieResult, error) {
	if querySlice == nil {
		return nil, ErrNoQueryParam
	}

	query := sumQuery(querySlice)

	searchURI += "/movie"

	var options = map[string]string{
		"query": url.QueryEscape(query),
	}

	v, err := c.constructURI(searchURI, options)
	if err != nil {
		return nil, err
	}

	s, err := c.request(*v, "GET")
	if err != nil {
		return nil, err
	}

	data := SearchMovieResult{}
	json.Unmarshal([]byte(s), &data)

	return &data, nil
}

func (c *Client) SearchCompany(query string) (*SearchCompanyResult, error) {
	if query == "" {
		return nil, ErrNoQueryParam
	}

	searchURI += "/company"

	var options = map[string]string{
		"query": url.QueryEscape(query),
	}

	v, err := c.constructURI(searchURI, options)
	if err != nil {
		return nil, err
	}

	s, err := c.request(*v, "GET")
	if err != nil {
		return nil, err
	}

	data := SearchCompanyResult{}
	json.Unmarshal([]byte(s), &data)

	return &data, nil
}

func (c *Client) SearchPerson(querySlice []string) (*SearchPersonResult, error) {
	if querySlice == nil {
		return nil, ErrNoQueryParam
	}

	query := sumQuery(querySlice)

	searchURI += "/person"

	var options = map[string]string{
		"query": url.QueryEscape(query),
	}

	v, err := c.constructURI(searchURI, options)
	if err != nil {
		return nil, err
	}

	s, err := c.request(*v, "GET")
	if err != nil {
		return nil, err
	}

	data := SearchPersonResult{}
	json.Unmarshal([]byte(s), &data)

	return &data, nil
}

func (c *Client) SearchCollection(querySlice []string) (*SearchCollectionResult, error) {
	if querySlice == nil {
		return nil, ErrNoQueryParam
	}

	query := sumQuery(querySlice)
	searchURI += "/collection"
	var options = map[string]string{
		"query": url.QueryEscape(query),
	}

	v, err := c.constructURI(searchURI, options)
	if err != nil {
		return nil, err
	}

	s, err := c.request(*v, "GET")
	if err != nil {
		return nil, err
	}

	data := SearchCollectionResult{}
	json.Unmarshal([]byte(s), &data)

	return &data, nil
}

func (c *Client) SearchKeyword(querySlice []string) (*SearchKeywordResult, error) {
	if querySlice == nil {
		return nil, ErrNoQueryParam
	}
	query := sumQuery(querySlice)
	searchURI += "/keyword"
	var options = map[string]string{
		"query": url.QueryEscape(query),
	}

	v, err := c.constructURI(searchURI, options)
	if err != nil {
		return nil, err
	}

	s, err := c.request(*v, "GET")
	if err != nil {
		return nil, err
	}

	data := SearchKeywordResult{}
	json.Unmarshal([]byte(s), &data)

	return &data, nil
}

func (c *Client) SearchTv(querySlice []string) (*SearchTvResult, error) {
	if querySlice == nil {
		return nil, ErrNoQueryParam
	}
	query := sumQuery(querySlice)
	searchURI += "/tv"
	var options = map[string]string{
		"query": url.QueryEscape(query),
	}

	v, err := c.constructURI(searchURI, options)
	if err != nil {
		return nil, err
	}

	s, err := c.request(*v, "GET")
	if err != nil {
		return nil, err
	}

	data := SearchTvResult{}
	json.Unmarshal([]byte(s), &data)

	return &data, nil
}
