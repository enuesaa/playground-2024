package docs

type Item struct {
	Path    string `json:"path"`
	DirName string `json:"dirName"`
}

type CreateBody struct {
	DirName string `json:"dirName"`
}

type Created struct {
	Ok bool `json:"ok"`
}
