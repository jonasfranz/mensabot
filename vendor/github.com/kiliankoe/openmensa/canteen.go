package openmensa

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// A Canteen object
type Canteen struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	City       string      `json:"city"`
	Address    string      `json:"address"`
	Coordinate *Coordinate `json:"coordinates"`
}

func (c Canteen) String() string {
	return fmt.Sprintf("%s", c.Name)
}

// GetAllCanteens returns a list of all known canteens
// Since OpenMensa uses pagination for the results, this will result in multiple requests.
func GetAllCanteens() (canteens []*Canteen, err error) {
	// FIXME: Concurrency anyone?

	url, _ := url.Parse(BaseURL + "canteens")
	page := 1

	for {
		params := url.Query()
		params.Set("page", strconv.Itoa(page))
		url.RawQuery = params.Encode()

		resp, err := get(url.String())

		var currentCanteens []*Canteen
		err = json.Unmarshal(resp, &currentCanteens)
		if err != nil {
			return nil, err
		}

		if len(currentCanteens) == 0 {
			// Reached the last page
			break
		}

		canteens = append(canteens, currentCanteens...)

		page++
	}

	return
}

// GetCanteens returns a list of canteen objects for the given ids
func GetCanteens(ids ...int) (canteens []*Canteen, err error) {
	url, _ := url.Parse(fmt.Sprintf("%s/canteens", BaseURL))

	stringIDs := []string{}
	for _, id := range ids {
		stringIDs = append(stringIDs, strconv.Itoa(id))
	}

	params := url.Query()
	params.Set("ids", strings.Join(stringIDs, ","))
	url.RawQuery = params.Encode()

	resp, err := get(url.String())

	err = json.Unmarshal(resp, &canteens)

	return
}
