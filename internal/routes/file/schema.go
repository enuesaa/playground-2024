package file

type Detail struct {
	Path string `json:"path"`
	Content string `json:"content"`
}

type CreateRequestBody struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type Creation struct {
	Ok bool `json:"ok"`
}
