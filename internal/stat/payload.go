package stat

type GetStatResponse struct {
	Period string `json:"period"`
	Total  int    `json:"total"`
}
