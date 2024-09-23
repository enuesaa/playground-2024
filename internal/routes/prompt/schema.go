package prompt

type CreateBody struct {
	Command string `json:"command"`
}

type Created struct {
	Output string `json:"output"`
}
