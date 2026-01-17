package structure

type InsertData struct {
	Task         string `json:"task"`
	Priority     string `json:"priority"`
	Status       string `json:"status"`
	AuthUserName string
}
