# simple TMDB API wrapper

## Example usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/jbpratt78/tmdb"
)

func main() {
	c := tmdb.New("API_KEY")
	s, err := c.SearchMovie([]string{"james", "bond"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
```

## TODO
- [ ] Improve documentation
- [ ] Search
  - [x] Companies
  - [x] Collections
  - [x] Keywords
  - [x] Movies
  - [ ] Multi
  - [x] People
  - [x] Tv
- [ ] Discover
  - [ ] Movie
  - [ ] Tv
- [x] Trending
