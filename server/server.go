package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"../computation"
)

type MyRequest struct {
	Data string `json:"data"`
}

var arrGlobal []int

func checkInputData(input string) (bool, string) {
	var result string
	var resultBool bool
	pattern := ``
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(input, "")
	matched, _ := regexp.Match(pattern, []byte(replace))
	if matched {
		result = fmt.Sprint(replace)
		resultBool = true
	} else {
		result = fmt.Sprintf("'%s' is not a valid data\n", replace)
	}

	return resultBool, result
}

func insertInGlobal(input string) {
	arrT := strings.Split(input, ",")
	for _, v := range arrT {
		i, _ := strconv.Atoi(v)
		if i != 0 {
			arrGlobal = append(arrGlobal, i)
		}
	}
}

// Up server localhost:8000.
func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/numbers", numbers)
	http.HandleFunc("/result", result)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func numbers(w http.ResponseWriter, r *http.Request) {
	var result string
	decoder := json.NewDecoder(r.Body)
	var myReq MyRequest
	err := decoder.Decode(&myReq)
	if err != nil {
		log.Println("Error decode")
	}
	log.Printf("Get input data - %s", myReq.Data)
	chB, ch := checkInputData(myReq.Data)
	if chB {
		insertInGlobal(ch)
		res, _ := computation.GetRepeatNumber(arrGlobal)
		result = res
		fmt.Printf("%v\nRequired number: %s\n", arrGlobal, result)
	}
	message := map[string]string{
		"data": fmt.Sprintf("%s", result),
	}
	js, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func result(w http.ResponseWriter, r *http.Request) {
	result, err := computation.GetRepeatNumber(arrGlobal)
	if err != nil {
		log.Println("Error get repeat number")
	}
	m := fmt.Sprintf("Finished array: %v \nRequired number: %s", arrGlobal, result)
	message := map[string]string{
		"data": m,
	}
	js, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func test(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var myReq MyRequest
	err := decoder.Decode(&myReq)
	if err != nil {
		log.Println("Error decode")
	}
	log.Printf("Client send - %s", myReq.Data)

	message := map[string]string{
		"data": "Message from server",
	}
	js, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
