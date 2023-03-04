package main

type Recipe struct {
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
}

var recipes []Recipe = []Recipe{
	{"Ceasar Salad", []string{"Creamy", "Green"}, []string{"Lettuce", "Dressing"}},
	{"Spring Green Salad", []string{"Sour", "Low Lat"}, []string{"Spring Green", "Vinerette"}},
}
