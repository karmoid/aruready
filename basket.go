package main

import (
	. "github.com/karmoid/aruready/karmostate"
	"strings"
)

// Basket stores a list of products which are available for purchase
type Basket struct {
	entries []BasketEntry
}

type BasketEntry struct {
	project  string
	activity string
	weight   string
}

// NewBasket instantiates a new Basket object
func NewBasket() *Basket {
	return &Basket{
		entries: make([]BasketEntry, 0, 10),
	}
}

func (b *Basket) AddItem(activityName string, weightValue string) {
	codes := strings.Split(activityName, "/")
	// Si la longueur est 1, nous n'avons pas d'activité. Seulement projet
	if len(codes) == 1 {
		b.entries = append(b.entries, BasketEntry{
			project:  activityName,
			activity: "",
			weight:   weightValue,
		})
	} else {
		b.entries = append(b.entries, BasketEntry{
			project:  codes[0],
			activity: codes[1],
			weight:   weightValue,
		})
	}
	return
}

// GetBasketSize returns size of Basket
func (b *Basket) GetBasketSize() int {
	return len(b.entries)
}

// GetBasketTotal calculates total of basket in % of day
func (b *Basket) GetBasketTotal() float64 {
	basketTotal := 0.00

	for _, value := range b.entries {
		basketTotal += GetWeightValue(value.weight)
	}

	if basketTotal > 8 {
		return 1
	}
	return basketTotal / 8
}
