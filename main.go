package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"os"
)


func getRoot(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got / request\n")
	io.WriteString(w, "This works")
}

func getLogTime(w http.ResponseWriter, r *http.Request){
	fmt.Printf("got /log request\n")
	result := logTimeToServer()
	io.WriteString(w, result)
}

func formattedTimeEntry() string {
	time := time.Now().Format("02-Jan-2006 15:04:05\n")
	return time
}

func createFile(f string, time string) error {
	file, err := os.OpenFile(f, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	data := time
	_, err = file.WriteString(data)

	if err != nil {
		return err
	}

	fmt.Println("file created")
	return nil
}

func logTime(filename string) error {
	entry := formattedTimeEntry()
	err := createFile(filename, entry)
	if err != nil {
		return err
	}
	return nil
}

func logTimeToServer() string {
	err := logTime("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	return "entry logged"
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/log", getLogTime)
	

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}