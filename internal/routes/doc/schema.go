package doc

type Detail struct {
	Slides []string `json:"slides"`
}

type UpdateRequestBody struct {
	Slides []string `json:"slides"`
}

type Creation struct {
	Ok bool `json:"ok"`
}
