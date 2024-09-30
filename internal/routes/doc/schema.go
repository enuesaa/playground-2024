package doc

type Slide struct {
	Content string `json:"content"`
	Drawing string `json:"drawing"`
}

type Detail struct {
	Slides []Slide `json:"slides"`
}

type UpdateRequestBody struct {
	Slides []Slide `json:"slides"`
}

type Creation struct {
	Ok bool `json:"ok"`
}