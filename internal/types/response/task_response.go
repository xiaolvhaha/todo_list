package response

import "todolist/internal/types"

type TaskInfoResponse struct {
	Deadline string `json:"deadline"`
	Info     types.TaskDomain
}

type TaskListResponse struct {
	List []*types.TaskDomain `json:"list"`
}
