package recipes

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name string `json:"name"`
}
