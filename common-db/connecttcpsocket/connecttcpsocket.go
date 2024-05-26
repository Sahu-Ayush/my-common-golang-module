package connecttcpsocket

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// connectTCPSocket initializes a TCP connection pool for a Cloud SQL
// instance of MySQL.
func ConnectTCPSocket() (*sql.DB, error) {

	fmt.Println("Using connectTCPSocket")

	// mustGetenv := func(k string) string {
	// 	v := os.Getenv(k)
	// 	if v == "" {
	// 		log.Fatalf("Fatal Error in connect_tcp.go: %s environment variable not set.", k)
	// 	}
	// 	return v
	// }

	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep secrets safe.
	var (
		dbUser    = os.Getenv("DB_USER")       // e.g. 'my-db-user'
		dbPwd     = os.Getenv("DB_PASS")       // e.g. 'my-db-password'
		dbName    = os.Getenv("DB_NAME")       // e.g. 'my-database'
		dbPort    = os.Getenv("DB_PORT")       // e.g. '3306'
		dbTCPHost = os.Getenv("INSTANCE_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	fmt.Println("dbURI: ", dbURI)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	// ...

	return dbPool, nil
}
