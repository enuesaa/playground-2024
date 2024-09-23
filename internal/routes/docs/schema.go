package docs

type Item struct {
	Path    string `json:"path"`
	DirName string `json:"dirName"`
}

type Detail struct {
	Path    string `json:"path"`
	DirName string `json:"dirName"`
	Content string `json:"content"`
}

type CreateRequestBody struct {
	DirName string `json:"dirName"`
}

type Creation struct {
	Ok bool `json:"ok"`
}

type UpdateRequestBody struct {
	Content string `json:"content"`
}
