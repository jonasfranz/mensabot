package openmensa

import (
	"encoding/json"
	"fmt"
	"time"
)

// A Meal object
type Meal struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Prices   struct {
		Students  *float32 `json:"students"`
		Employees *float32 `json:"employees"`
		Pupils    *float32 `json:"pupils"`
		Others    *float32 `json:"others"`
	} `json:"prices"`
	Notes []string `json:"notes"`
}

func (m Meal) String() string {
	return m.Name
}

// TodaysMeals is a helper wrapper for Meals() for today's menu
func (c Canteen) TodaysMeals() (meals []*Meal, err error) {
	now := time.Now()
	return c.GetMeals(now)
}

// GetMeals returns all meals for a given date
func (c Canteen) GetMeals(date time.Time) (meals []*Meal, err error) {
	url := fmt.Sprintf("%s/canteens/%d/days/%s/meals", BaseURL, c.ID, date.Format("2006-01-02"))

	resp, err := get(url)

	err = json.Unmarshal(resp, &meals)

	return
}
