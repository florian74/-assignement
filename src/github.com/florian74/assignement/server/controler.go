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

	id := strings.TrimPrefix(req.URL.Path, "/provisions/")

	if !strings.Contains(id, "/") {
		panic("unexpected url")
	}

	var typeSearch = strings.Split(id, "/")[0]
	var value = strings.Split(id, "/")[1]

	allValues := getAllKeys()
	var result = make([]Fpl, 0)

	for i := 0; i < len(allValues); i++ {
		if strings.Contains(getField(&allValues[i], typeSearch), value) {
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
