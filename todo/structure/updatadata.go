package structure

type UpdateData struct {
	TaskID        int
	AuthUserName  string
	UpdateRequest UpdateTaskRequest
}

type UpdateTaskRequest struct {
	Task     *string `json:"task,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Status   *string `json:"status,omitempty"`
}
