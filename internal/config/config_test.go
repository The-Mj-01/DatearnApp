package config_test

import (
	"Datearn/internal/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const testFileName string = ".env"

const testFileContext string = "MYSQL_USER=\"MJ\"\n" +
	"MYSQL_PASSWORD=\"AZ\""

// TestReadConfig tests the ReadConfig function of the config package. and checks whether it functions correct or not
func TestReadConfig(t *testing.T) {
	err := config.ReadConfig()
	assert.Error(t, err, "At first line read config file shout throw an error because there is no .env file")

	err = prepareTestEnv()
	assert.NoError(t, err, "Cannot create .env file for running test please check os permissions")

	err = config.ReadConfig()
	assert.NoError(t, err, "Env file exists and it should not throws an error")

	err = deleteTestFiles()
	assert.NoError(t, err, "Could not delete created files")
}

// prepareTestEnv prepares Env file for testing
func prepareTestEnv() error {
	file, err := os.Create(testFileName)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(testFileContext))
	return err
}

// deleteEnvFile deletes .env file name
func deleteTestFiles() error {
	return os.Remove(testFileName)
}
