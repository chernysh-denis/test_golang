package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type MyRequest struct {
	Data string `json:"data"`
}

func input(in string) (number string) {
	message := map[string]string{
		"data": in,
	}
	jsonValue, _ := json.Marshal(message)

	resp, e := http.Post("http://localhost:8000/numbers", "application/json", bytes.NewBuffer(jsonValue))
	if e != nil {
		fmt.Println("Server is down")
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var myReq MyRequest
	err := decoder.Decode(&myReq)
	if err != nil {
		log.Println("Error decode")
	}

	return fmt.Sprintf("%s", myReq.Data)
}

func exit() {
	resp, e := http.Get("http://localhost:8000/result")
	if e != nil {
		fmt.Println("Server is down")
	}
	defer resp.Body.Close()
	fmt.Printf("%s\n", resp.Body)

	fmt.Println("Client program closed")
	os.Exit(1)
}

func main() {
	fmt.Println("Start client")
	for {
		fmt.Print("Enter data: ")
		myscanner := bufio.NewScanner(os.Stdin)
		myscanner.Scan()
		line := myscanner.Text()
		if line == "exit" {
			os.Exit(1)
		} else {
			r := input(line)
			fmt.Printf("Server request:  %s\n", r)
		}
	}
}
