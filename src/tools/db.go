package tools

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type SqlCron struct {
	UserID     int
	Count      int
	EntryID    int
	Expression string
	Message    string
}

func DBInit() {
	var DBError error
	db, DBError = sql.Open("sqlite3", "./sql.db")
	CheckErr(DBError)
}

func DBUpdateEntryID(old int, new int) {
	sql_ := "UPDATE cron set entry_id = ? where entry_id = ?"
	stmt, err := db.Prepare(sql_)
	CheckErr(err)
	_, err = stmt.Exec(new, old)
	CheckErr(err)
}

func DBUpdateCount(message string, userID string, exps string) {
	sql_ := "SELECT * FROM cron WHERE message = '" + message + "' AND user_id = '" + userID + "' AND expression = '" + exps + "'"
	dbQuery := DBQuery(sql_)
	count := dbQuery[0].Count - 1
	query := "UPDATE cron set count = ? where message = ? and user_id = ? and expression = ?"
	if count == 0 {
		query = "DELETE from cron where count = ? and message = ? and user_id = ? and expression = ?"
	}

	stmt, err := db.Prepare(query)
	CheckErr(err)
	_, err = stmt.Exec(count, message, userID, exps)
	CheckErr(err)
}

func DBInsert(sqlCron SqlCron) {
	query := "INSERT INTO cron(user_id, count, entry_id, expression, message) values (?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	CheckErr(err)
	_, err = stmt.Exec(sqlCron.UserID, sqlCron.Count, sqlCron.EntryID, sqlCron.Expression, sqlCron.Message)
	CheckErr(err)
}

func DBQuery(sql string) []SqlCron {
	var sqlCrons []SqlCron
	rows, _ := db.Query(sql)
	for rows.Next() {
		sqlCron := new(SqlCron)
		err := rows.Scan(&sqlCron.UserID, &sqlCron.Count, &sqlCron.EntryID, &sqlCron.Expression, &sqlCron.Message)
		sqlCrons = append(sqlCrons, *sqlCron)
		CheckErr(err)
	}
	return sqlCrons
}

func DBOrderByEntryID() {
	sql_ := "select * from cron order by entry_id"
	sqlCrons := DBQuery(sql_)
	for i, cron := range sqlCrons {
		i = i + 1
		if cron.EntryID != i {
			DBUpdateEntryID(cron.EntryID, i)
		}
	}
}

func DBDelete(entryID string) {
	stmt, err := db.Prepare("Delete from cron where entry_id=?")
	CheckErr(err)
	_, err = stmt.Exec(entryID)
	CheckErr(err)
}
