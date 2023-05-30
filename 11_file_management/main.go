package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// Create directory
	err := os.Mkdir("mydir", 0755)
	checkError(err)

	// 0755
	// 7: owner has read, write, and execute permissions
	// 5: groupes have read and execute permissions
	// 5: others have read and execute permissions

	// Create directory with a parent directory
	err = os.MkdirAll("mydir/childir", 0755)
	checkError(err)

	// Read the contents of a directory
	files, err := os.ReadDir(".")
	checkError(err)

	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Create file
	new_file, err := os.Create("new_file.txt")
	checkError(err)
	defer new_file.Close()

	// Open a file for read only
	file, err := os.Open("file.txt")
	checkError(err)
	defer file.Close()

	// Open file for read and write
	file, err = os.OpenFile("file.txt", os.O_APPEND, 0644)

	// Constants for file permissions
	// O_RDONLY: read only
	// O_WRONLY: write only
	// O_RDWR: read and write
	// O_APPEND: append to end of file
	// O_CREATE: create file if it doesn't exist
	// O_EXCL: used with O_CREATE, file must not exist

	// Rename or move file
	os.Rename("file.txt", "file2.txt")

	// Copy bytes from one file to another, look at Writer & Reader interface & what File implements
	_, err = io.Copy(file, new_file)
	checkError(err)

	// Delete file or directory
	os.Remove("new_file.txt")

	// Read file
	bytes, err := ioutil.ReadFile("file2.txt")
	checkError(err)
	fmt.Println(string(bytes))

	// Read large file line by line
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		fmt.Println("Line", lines, ":", scanner.Text())
		lines++
	}

	new_file, err = os.Create("new_file.txt")
	checkError(err)
	// Write to file
	writeBuffer := bufio.NewWriter(new_file)
	for i := 0; i < 5; i++ {
		writeBuffer.WriteString(fmt.Sprintln("New Line", i, ": New text here blablabl"))
	}
	writeBuffer.Flush() // Commits changes to the disk

	// Get some file stats
	getFileStats("file2.txt")

	// TODO: Copy file
	copyFile("file2.txt", "file_copy.txt")

	// TODO: Watch a file
	watchFile("file2.txt")

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getFileStats(fileName string) {
	stats, err := os.Stat(fileName)
	checkError(err)
	fmt.Println("File Name:", stats.Name())
	fmt.Println("Size:", stats.Size())
	fmt.Println("Last Modified:", stats.ModTime())
	fmt.Println("Mode:", stats.Mode())
}

// TODO: write a function which copies a file
func copyFile(src string, dst string) {
	srcFile, err := os.Open(src)
	checkError(err)
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	checkError(err)
	defer dstFile.Close()

	// Copy bytes from source to destination.
	_, err = io.Copy(dstFile, srcFile)
	checkError(err)

	// Ensures all buffered data is written to disk.
	err = dstFile.Sync()
	checkError(err)
}

// TODO: Implement a watcher which logs the last modification time of a given file to a log file.
func watchFile(fileName string) {
	logFile, err := os.Create("logModTime.txt")
	checkError(err)
	defer logFile.Close()

	stats, err := os.Stat(fileName)
	checkError(err)
	counter := 0
	for {
		time.Sleep(1 * time.Second)

		newStats, err := os.Stat(fileName)
		checkError(err)
		lastModTime := newStats.ModTime()
		if stats.ModTime() != lastModTime {
			logFile.Write([]byte(fmt.Sprintln("Change Number", counter, ": File was modified at", lastModTime)))
			counter++
			stats = newStats
		}
	}
}
