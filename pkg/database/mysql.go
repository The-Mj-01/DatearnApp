package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

const (
	dbEnvUserKey   string = "MYSQL_USER" // environment variable for database username
	dbEnvPassKey   string = "MYSQL_PASS" // environment variable for database password
	dbEnvAddrKey   string = "MYSQL_ADDR" // environment variable for database address
	dbEnvPortKey   string = "MYSQL_PORT" // environment variable for database port
	dbEnvDbNameKey string = "MYSQL_DB"   // environment variable for database name
	dbEnvCharKey   string = "MYSQL_CHAR" // environment variable for database character set

	defaultDBAddress   = "localhost"  // default database address
	defaultDBPort      = "3306"       // default database port
	defaultDBChar      = "utf8mb4"    // default database character set
	productionStageKey = "production" //Define key production stage of application
)

// Conn returns a pointer to a GORM database instance.
func Conn() (*gorm.DB, error) {
	if os.Getenv("APPLICATION_STAGE") == productionStageKey {
		return productionStage()
	}
	return developmentStage()
}

// productionStage makes connection that is appropriate for production stage
func productionStage() (*gorm.DB, error) {
	dsn := makeConnDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

// developmentStage configures database for development stage in order to test an
func developmentStage() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	return db, err
}

// makeConnDsn generates a database connection string using environment variables.
func makeConnDsn() string {
	user := os.Getenv(dbEnvUserKey) // get database username from environment variable
	pass := os.Getenv(dbEnvPassKey) // get database password from environment variable
	db := os.Getenv(dbEnvDbNameKey) // get database name from environment variable
	addr := getDBAddress()          // get database address (default or from environment variable)
	port := getDbPort()             // get database port (default or from environment variable)
	char := getCharKey()            // get database character set (default or from environment variable)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, pass, addr, port, db, char)
	return dsn
}

// getDBAddress returns the database address from an environment variable or the default.
func getDBAddress() string {
	dbAddress := os.Getenv(dbEnvAddrKey)

	if dbAddress == "" {
		dbAddress = defaultDBAddress
	}

	return dbAddress
}

// getDbPort returns the database port from an environment variable or the default.
func getDbPort() string {
	port := os.Getenv(dbEnvPortKey)

	if port == "" {
		port = defaultDBPort
	}

	return port
}

// getCharKey returns the database character set from an environment variable or the default.
func getCharKey() string {
	char := os.Getenv(dbEnvCharKey)

	if char == "" {
		char = defaultDBChar
	}

	return char
}
