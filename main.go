package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	hello = "hello world..."
	howdi = "howdi partner..."
)

type student struct {
	Name  string `json:"name"`
	Roll  int    `json:"roll"`
	Age   int    `json:"age"`
	Class int    `json:"class"`
}

func main() {

	eric := student{
		Name:  "Eric",
		Roll:  10,
		Age:   12,
		Class: 7,
	}
	var stud []student
	stud = append(stud, eric)
	out, err := json.MarshalIndent(stud, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(out))
	})

	http.ListenAndServe(":8080", nil)
}
