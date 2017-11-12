package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	log.Println("Initializing...")
	// log.Println("")
	// log.Println("The following variables must be set:")
	// log.Println("LOGSTASH_HOST: IP where logstash runs")
	// log.Println("LOGSTASH_PORT: Port where logstash runs")
	// log.Println("INPUT_DIR: Directory where to find the files")
	// log.Println("FILEENDINGS: Fileendings of files to send. Separated by comma.")
	// log.Println("SLEEP: Seconds to sleep between sendings")

	logstashHost := os.Getenv("LOGSTASH_HOST")
	logstashPort := os.Getenv("LOGSTASH_PORT")
	inputFolder := os.Getenv("INPUT_DIR")
	fileendings := os.Getenv("FILEENDINGS")
	sleep, _ := strconv.Atoi(os.Getenv("SLEEP"))

	log.Println("Waiting for files...")
	for true {
		// Sync
		var wg sync.WaitGroup

		// Parent dir
		files, _ := ioutil.ReadDir(inputFolder)
		wg.Add(len(files))

		for _, file := range files {
			// Parallel execution
			go func(fileTmp os.FileInfo) {
				defer wg.Done()

				// Checking filetype
				isValid := false
				for _, fileending := range strings.Split(fileendings, ",") {
					isValid = isValid || strings.HasSuffix((fileTmp).Name(), fileending)
				}
				// Checking file / dir
				isValid = isValid && !(fileTmp).IsDir()

				// Send file
				if isValid {
					log.Println("File found: ", (fileTmp).Name())
					sendErr := SendFile(logstashHost, logstashPort, inputFolder, &fileTmp)
					if sendErr == nil {
						// DELETE fileTmp
						err := os.Remove(inputFolder + string(os.PathSeparator) + (fileTmp).Name())
						if err == nil {
							log.Println("File deleted: ", (fileTmp).Name())
						} else {
							log.Println("File could not be deleted: ", (fileTmp).Name())
						}
					} else {
						log.Println("File could not be send: ", (fileTmp).Name(), " ", sendErr.Error())
					}
				} else {
					// log.Println("File ignored: " + (fileTmp).Name())
				}
			}(file)
		}

		// Sleep
		time.Sleep(time.Duration(sleep) * time.Second)
		// Wait for threads
		wg.Wait()
	}
}

/**
* Sends file to host:port.
**/
func SendFile(logstashHost string, logstashPort string, inputFolder string, file *os.FileInfo) error {
	// Connection
	conn, errConn := net.Dial("tcp", logstashHost+":"+logstashPort)
	if errConn != nil {
		return errConn
	}
	defer conn.Close()

	// File
	f, errFile := os.Open(inputFolder + string(os.PathSeparator) + (*file).Name())
	if errFile != nil {
		return errFile
	}
	defer f.Close()

	// Read file, write socket.
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(f)
	_, errCopy := io.Copy(writer, reader)
	if errCopy != nil {
		return errCopy
	}
	writer.WriteByte('\n')
	errFlush := writer.Flush()
	if errFlush != nil {
		return errFlush
	}
	return nil
}
