package request

type CreateTaskRequest struct {
	CategoryId int64  `json:"category_id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Priority   int64  `json:"priority"`
	Deadline   string `json:"deadline"`
}
