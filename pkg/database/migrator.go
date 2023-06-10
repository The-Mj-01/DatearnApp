package database

import (
	"Datearn/pkg/advancedError"
	"fmt"
	"gorm.io/gorm"
	"os"
	"strings"
)

const migrationDirEnvKey string = "MIGRATIONS_DIR"

var migrationsDir string

// init initializes function for first use
func init() {
	migrationsDir = getMigrationDir()
}

// Migrate runs all sqls from migration directory
func Migrate() error {
	files, err := os.ReadDir(migrationsDir)

	if err != nil {
		return advancedError.New(err, "Cannot migrate")
	}

	dbConn, err := Conn()
	if err != nil {
		return advancedError.New(err, "Cannot migrate")
	}

	dbConn.Begin()
	for _, file := range files {
		fmt.Println(file.Name())
		err = runMigrationCmd(dbConn, file)
		if err != nil {
			return advancedError.New(err, "Cannot migrate")
		}
	}

	fmt.Println("All migrations ran successfully")
	return nil
}

// runMigrationCmd from given sql file
func runMigrationCmd(db *gorm.DB, file os.DirEntry) error {
	fullDir := generateFullFileDir(file.Name())
	sqlCmd, err := os.ReadFile(fullDir)

	if err != nil {
		return advancedError.New(err, "cannot migrate")
	}

	requests := strings.Split(string(sqlCmd), ";")

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, cmd := range requests {
		if len(cmd) == 0 || cmd == "" {
			break
		}

		result := tx.Exec(cmd)
		if result.Error != nil {
			tx.Rollback()
			return advancedError.New(err, "Cannot migrate")
		}
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return advancedError.New(err, "Cannot migrate")
	}

	return nil

}

// getMigrationDir gets and returns sql migrations directories
func getMigrationDir() string {
	if os.Getenv(migrationDirEnvKey) != "" {
		migrationsDir = os.Getenv(migrationDirEnvKey)
	}

	if migrationsDir == "" {
		migrationsDir = getDefaultDir()
	}

	return migrationsDir
}

// getDefaultDir get and returns default dir for migrations directory
func getDefaultDir() string {
	currentDir, _ := os.Getwd()
	currentDir = strings.Replace(currentDir, "server", "", 1)
	currentDir = strings.Replace(currentDir, "cmd", "", 1)
	currentDir = appendDefaultDir(currentDir, "scripts", "migrations")
	return currentDir
}

// appendDefaultDir for sql functions and returns it
func appendDefaultDir(dir string, appendables ...string) string {
	for _, appendable := range appendables {
		if !strings.Contains(dir, appendable) {
			dir = fmt.Sprintf("%s/%s", dir, appendable)
		}
	}

	return dir
}

// generateFullFileDir generates full function dir and returns it
func generateFullFileDir(flName string) string {
	return fmt.Sprintf("%s/%s", migrationsDir, flName)
}
