package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	adexp "github.com/florian74/assignement/adexp"
)

type Controller struct {
}

func (controller Controller) HandlePut(buf []byte, rlen int, count int) {
	fmt.Println(string(buf[0:rlen]))
	fmt.Println(count)
	pushToRedis(fromJson(string(buf[0:rlen])))
}

func (controller Controller) HandleSearch(w http.ResponseWriter, req *http.Request) {

	fmt.Println("search request detected")

	id := strings.TrimPrefix(req.URL.Path, "/flight/")

	var typeSearch string
	var value string

	if !strings.Contains(id, "/") {
		typeSearch = id
		value = ""
	} else {
		typeSearch = strings.Split(id, "/")[0]
		value = strings.Split(id, "/")[1]
	}

	fmt.Println("looking for " + typeSearch + " and value " + value)

	allValues := getAllKeys()
	var result = make([]Fpl, 0)

	for i := 0; i < len(allValues); i++ {
		if strings.HasPrefix(getField(&allValues[i], typeSearch), value) || value == "" {
			result = append(result, allValues[i])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func getField(v *adexp.Fpl, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
