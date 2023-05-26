package main

import (
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create directory

	// 0755
	// 7: owner has read, write, and execute permissions
	// 5: grouphas read and execute permissions
	// 5: others have read and execute permissions

	// Create directory with a parent directory

	// Read the contents of a directory

	// Open a file for read only

	// Open file for read and write

	// Constants for file permissions
	// O_RDONLY: read only
	// O_WRONLY: write only
	// O_RDWR: read and write
	// O_APPEND: append to end of file
	// O_CREATE: create file if it doesn't exist
	// O_EXCL: used with O_CREATE, file must not exist

	// Rename or move file

	// TODO: Copy file

	// Delete file or directory

	// Read file

	// Read large file line by line

	// Write to file

	// Get some file stats

	// TODO: Watch a file

}

// TODO: write a function which copies a file, hint: os.Create() & io.Copy()
func copyFile(src string, dst string) {

}

func getFileStats(fileName string) {

}

// TODO: Implement a watcher which logs the last modification time of a given file to a log file.
func watchFile(fileName string) {

}
