package doc

type Detail struct {
	Path   string   `json:"path"`
	Slides []string `json:"slides"`
}

type UpdateRequestBody struct {
	Slides []string `json:"slides"`
}

type Updated struct {
	Ok bool `json:"ok"`
}
