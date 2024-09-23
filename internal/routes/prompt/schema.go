package prompt

type CreateRequestBody struct {
	Command string `json:"command"`
}

type Creation struct {
	Output string `json:"output"`
}
