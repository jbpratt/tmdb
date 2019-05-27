package tmdb

type SearchMovieResult struct {
	Page    int `json:"page"`
	Results []struct {
		PosterPath       string  `json:"poster_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalTitle    string  `json:"original_title"`
		OriginalLanguage string  `json:"original_language"`
		Title            string  `json:"title"`
		BackdropPath     string  `json:"backdrop_path"`
		Popularity       float64 `json:"popularity"`
		VoteCount        int     `json:"vote_count"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
	} `json:"results"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
}

type SearchCompanyResult struct {
	Page    int `json:"page"`
	Results []struct {
		ID       int    `json:"id"`
		LogoPath string `json:"logo_path"`
		Name     string `json:"name"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchPersonResult struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		Popularity  float64 `json:"popularity"`
		ID          int     `json:"id"`
		ProfilePath string  `json:"profile_path"`
		Name        string  `json:"name"`
		KnownFor    []struct {
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int     `json:"vote_count"`
			ID               int     `json:"id"`
			Video            bool    `json:"video"`
			MediaType        string  `json:"media_type"`
			Title            string  `json:"title"`
			Popularity       float64 `json:"popularity"`
			PosterPath       string  `json:"poster_path"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			GenreIds         []int   `json:"genre_ids"`
			BackdropPath     string  `json:"backdrop_path"`
			Adult            bool    `json:"adult"`
			Overview         string  `json:"overview"`
			ReleaseDate      string  `json:"release_date"`
		} `json:"known_for"`
		Adult bool `json:"adult"`
	} `json:"results"`
}

type SearchKeywordResult struct {
	Page    int `json:"page"`
	Results []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchCollectionResult struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool   `json:"adult"`
		BackdropPath     string `json:"backdrop_path"`
		ID               int    `json:"id"`
		Name             string `json:"name"`
		OriginalLanguage string `json:"original_language"`
		OriginalName     string `json:"original_name"`
		Overview         string `json:"overview"`
		PosterPath       string `json:"poster_path"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type SearchTvResult struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		OriginalName     string   `json:"original_name"`
		ID               int      `json:"id"`
		Name             string   `json:"name"`
		VoteCount        int      `json:"vote_count"`
		VoteAverage      float64  `json:"vote_average"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date"`
		Popularity       float64  `json:"popularity"`
		GenreIds         []int    `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
	} `json:"results"`
}
