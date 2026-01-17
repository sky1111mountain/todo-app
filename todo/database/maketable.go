package database

import (
	"database/sql"
)

func MakeTable(db *sql.DB) error {
	query := `
	create table if not exists todolist(
    	id integer unsigned auto_increment primary key,
   		task varchar(100) not null,
    	priority varchar(100) not null,
    	status varchar(100) not null,
    	username varchar(100) not null,
    	created_at datetime
	);`

	_, err := db.Exec(query)
	return err
}
