package docs

type Item struct {
	Path    string `json:"path"`
	DirName string `json:"dirName"`
}

type ItemWithContent struct {
	Path    string `json:"path"`
	DirName string `json:"dirName"`
	Content string `json:"content"`
}

type CreateBody struct {
	DirName string `json:"dirName"`
}

type Created struct {
	Ok bool `json:"ok"`
}

type UpdateBody struct {
	Content string `json:"content"`
}
