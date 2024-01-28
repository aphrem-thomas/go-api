package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Name  string `json:"name"`
	Roll  int    `json:"roll"`
	Age   int    `json:"age"`
	Class int    `json:"class"`
}

type resultJson struct {
	resJson []byte
}

func (j *resultJson) WriteJson(w http.ResponseWriter) {
	fmt.Fprint(w, string(j.resJson))
}

func GetJson(w http.ResponseWriter, r *http.Request) {

	eric := student{
		Name:  "Eric",
		Roll:  10,
		Age:   12,
		Class: 7,
	}
	var stud []student
	stud = append(stud, eric)
	resJas := resultJson{}
	resJas.resJson, _ = json.MarshalIndent(stud, "", "	")
	resJas.WriteJson(w)
}
