package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	defaultEnvDir string = ".env"    // defaultEnvDir is the default directory to look for the .env file
	envDirEnvKey  string = "ENV_DIR" // envDirEnvKey is the environment variable that

)

// ReadConfig reads the configuration variables from the .env file
func ReadConfig() error {
	err := readFromEnv()
	return err
}

func readFromEnv() error {
	dir := defaultEnvDir              // set the directory to the default directory
	var shouldThrowError bool = false // shouldThrowError flag is used to determine whether an error should be thrown if the file is not found

	if os.Getenv(envDirEnvKey) != "" { // check if the ENV_DIR environment variable is set
		dir = os.Getenv(envDirEnvKey) // if it is set, use that directory instead of the default
	}

LoadenvOp:
	err := godotenv.Load(dir)           // try to load the .env file from the specified directory
	if err != nil && shouldThrowError { // if an error is encountered and the shouldThrowError flag is set
		return err // exit the program with the error message
	}

	if err != nil { // if an error is encountered
		shouldThrowError = true // set the shouldThrowError flag to true
		dir = "../../.env"      // change the directory to the backup directory
		goto LoadenvOp          // retry loading the .env file
	}
	return nil // return nil to indicate success
}
