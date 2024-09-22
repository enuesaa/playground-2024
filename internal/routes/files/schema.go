package routes

type Item struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
	IsDir    bool   `json:"isDir"`
}
