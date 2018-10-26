## 🍴openmensa

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/kiliankoe/openmensa) [![travis](https://img.shields.io/travis/kiliankoe/openmensa.svg?style=flat-square&maxAge=2592000)](https://travis-ci.org/kiliankoe/openmensa)

This is a small golang wrapper for the [OpenMensa](https://openmensa.org) API. Very WIP.

Use it something like this:

```go
import (
	"fmt"

  	"github.com/kiliankoe/openmensa"
)

func main() {
	canteens, _ := openmensa.GetCanteens(79, 80)
	meals, _ := canteens[0].TodaysMeals()
	for _, meal := range meals {
		fmt.Println(meal)
	}
}

// Hähnchenkeule mit Geflügelsoße, dazu buntes Gemüse und Petersilienkartoffeln
// Sojageschnetzeltes mit Waldpilzen, dazu Apfelrotkohl und Klöße
// Gefülltes Pizzabrötchen
// ...
```
