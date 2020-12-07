package models

type PagedRequest struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

func (r *PagedRequest) GetLimit() (offset int, limit int) {
	offset = (r.PageIndex - 1) * r.PageSize
	limit = r.PageSize
	return
}

type PagedResponse struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
