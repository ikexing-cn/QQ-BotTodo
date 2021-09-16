package utils

import (
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type SqlCron struct {
	UserID     int
	Count      int
	EntityID   int
	Expression string
	Message    string
}

func DBInit() {
	var DBError error
	db, DBError = sql.Open("sqlite3", "./sql.db")
	CheckErr(DBError)
}

func DBInsert(sqlCron SqlCron) {
	query := "INSERT INTO cron(user_id, count, entry_id, expression, message) values (?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	CheckErr(err)
	_, err = stmt.Exec(sqlCron.UserID, sqlCron.Count, sqlCron.EntityID, sqlCron.Expression, sqlCron.Message)
	CheckErr(err)
}

func DBQuery(sql string) []SqlCron {
	var sqlCrons []SqlCron
	rows, _ := db.Query(sql)
	for rows.Next() {
		sqlCron := new(SqlCron)
		err := rows.Scan(&sqlCron.UserID, &sqlCron.Count, &sqlCron.EntityID, &sqlCron.Expression, &sqlCron.Message)
		sqlCrons = append(sqlCrons, *sqlCron)
		CheckErr(err)
	}
	return sqlCrons
}
