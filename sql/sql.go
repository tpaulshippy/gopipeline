package sql

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

// ConnectionInfo is what is needed to connect to database
type ConnectionInfo struct {
	debug    bool
	password string
	port     int
	server   string
	user     string
	database string
}

func openDatabase(connInfo *ConnectionInfo) (*sql.DB, error) {
	if connInfo == nil {
		log.Fatal("connInfo is nil")
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", connInfo.server, connInfo.user, connInfo.password, connInfo.port, connInfo.database)
	if connInfo.debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)

	return conn, err
}
