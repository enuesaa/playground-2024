package file

type CreateBody struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type Created struct {
	Ok bool `json:"ok"`
}
