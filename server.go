package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"./computation"
)

var arrGlobal []int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/getNumbers", getNumbers)
	http.HandleFunc("/getResult", getResult)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getNumbers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := r.PostForm["data"]
	arrT := strings.Split(data[0], ",")
	for _, v := range arrT {
		i, _ := strconv.Atoi(strings.TrimSpace(v))
		arrGlobal = append(arrGlobal, i)
	}
	result, err := computation.GetRepeatNumber(arrGlobal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", result)
}

func getResult(w http.ResponseWriter, r *http.Request) {
	result, err := computation.GetRepeatNumber(arrGlobal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Finished array: %s Result: %s", strings.Trim(strings.Replace(fmt.Sprint(arrGlobal), " ", ",", -1), "[]"), result)
}
