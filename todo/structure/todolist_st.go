package structure

import (
	"database/sql"
)

type Todo struct {
	ID        int          `json:"id"`
	Task      string       `json:"task"`
	Priority  string       `json:"priority"`
	Status    string       `json:"status"`
	UserName  string       `json:"user_name"`
	CreatedAt sql.NullTime `json:"created_at"`
}
