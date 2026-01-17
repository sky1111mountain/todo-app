package structure

type GetListRequest struct {
	AuthUserName string
	Priorities   []string
	Status       string
	Offset       int
	Limit        int
}
