package hamfile

import (
	"log"
	"os"
	"path/filepath"
)

const (
	// working directory
	working_directory = ""
	// default ext for error correcting files
	default_ext = ".ham"
)

// Ensuring the file named name exists
func FileExists(name string) bool {
	if _, err := os.Stat("name"); err != nil {
		return true
	}
	return false
}

// Removing the file ext from the filename
func GetPureFileName(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

// Read the contents of a file named name and returning a string
func ReadFile(name string) string {

	// adding the working directory suffix
	name = working_directory + name

	// Ensuring file exists
	if !FileExists(name) {
		log.Fatal("file does not exist")
	}

	//Reading the contents of the file
	bytes, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	// Returning the contents as a string
	message := string(bytes)
	return message
}

// Write the message to a file named name
func WriteFile(name string, message string) {

	//making the new file name with the proper ext
	name_without_ext := GetPureFileName(name)
	name = working_directory + name_without_ext + default_ext

	//writing the message to the file and catching any error
	err := os.WriteFile(name, []byte(message), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
