package conn

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var sysDb *sqlx.DB = nil

type SqlDb struct {
	sysDb *sqlx.DB
}

func NewSqlDb() *SqlDb {
	if sysDb == nil {
		var err error
		// sysPath := GetPath()
		// sysDb, err = sqlx.Connect("sqlite3", sysPath)
		sysDb, err = Connect()
		if err != nil {
			fmt.Println("error here in db connection")
		}
	}
	return &SqlDb{sysDb: sysDb}
}

func Connect() (*sqlx.DB, error) {
	dsn := "root:Nilesh@045@tcp(127.0.0.1:3306)/trackit?parseTime=true"
	var err error
	DB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("DB connection failed:", err)
	}
	log.Println("✅ Connected to MySQL database TrakIt")
	return DB, err
}

func GetPath() string {
	exec, _ := os.Executable()
	basepath := filepath.Dir(exec)
	// basepath := "D:/GoProjects/project1"
	syspath := filepath.Join(basepath, "../../databases/project_db.data")
	fmt.Println("syspath")
	return syspath
}

func (sqlDb *SqlDb) QuerySql(query string, params ...any) QueryResult {
	return sqlDb.querySqlData(sqlDb.sysDb, query, params...)
}

func (sqlDb *SqlDb) querySqlData(db *sqlx.DB, query string, params ...any) QueryResult {
	rows, err := db.Query(query, params...)
	if err != nil {
		_ = fmt.Errorf("error while executing query: %v", err)
	}
	return QueryResult{Rows: rows, Error: err}
}

func (sqlDb *SqlDb) ExecSql(query string, params ...any) ExecResult {
	return sqlDb.execSqlQry(sqlDb.sysDb, query, params...)
}

func (sqlDb *SqlDb) execSqlQry(db *sqlx.DB, query string, params ...any) ExecResult {
	result, err := db.Exec(query, params...)
	execResult := ExecResult{Error: err}
	if err != nil {
		_ = fmt.Errorf("error while executing query: %v", err)
	} else {
		execResult.RowsAffected, _ = result.RowsAffected()
		execResult.InsertId, _ = result.LastInsertId()
		_ = fmt.Errorf("execSql rowsAffected: %d, lastInsertId: %d,",
			execResult.RowsAffected, execResult.InsertId)
	}
	return execResult
}
