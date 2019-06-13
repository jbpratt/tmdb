package tmdb

import (
	"encoding/json"
)

var searchURI = BaseURI + "/search"

func (c *Client) SearchMovie(querySlice []string) (*SearchMovieResult, error) {
	if querySlice == nil {
		return nil, ErrNoQueryParam
	}

	v, err := c.constructURI(searchURI + "/movie")
	if err != nil {
		return nil, err
	}

	v = setQuery(v, querySlice)
	s, err := c.request(*v, "GET")
	if err != nil {
		return nil, err
	}

	data := SearchMovieResult{}
	json.Unmarshal([]byte(s), &data)

	return &data, nil
}

func (c *Client) SearchCompany(querySlice []string) (*SearchCompanyResult, error) {
	if querySlice == nil {
		return nil, ErrNoQueryParam
	}

	v, err := c.constructURI(searchURI + "/company")
	if err != nil {
		return nil, err
	}

	v = setQuery(v, querySlice)
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

	v, err := c.constructURI(searchURI + "/person")
	if err != nil {
		return nil, err
	}

	v = setQuery(v, querySlice)
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

	v, err := c.constructURI(searchURI + "/collection")
	if err != nil {
		return nil, err
	}

	v = setQuery(v, querySlice)
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

	v, err := c.constructURI(searchURI + "/keyword")
	if err != nil {
		return nil, err
	}

	v = setQuery(v, querySlice)
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

	v, err := c.constructURI(searchURI + "/tv")
	if err != nil {
		return nil, err
	}

	v = setQuery(v, querySlice)
	s, err := c.request(*v, "GET")
	if err != nil {
		return nil, err
	}

	data := SearchTvResult{}
	json.Unmarshal([]byte(s), &data)

	return &data, nil
}
