package files

type Detail struct {
	Content  string `json:"content"`
}

type UpdateRequestBody struct {
	Content  string `json:"content"`
}

type Creation struct {
	Ok bool `json:"ok"`
}
