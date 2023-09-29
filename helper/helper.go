package helper

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func isEmpty(v any) bool {
	return v == nil || v == ""
}

func GetQueryInt(r *http.Request, key string) (int, bool) {
	v := mux.Vars(r)[key]

	if isEmpty(v) {
		return 0, false
	}

	i, err := strconv.Atoi(v)

	if err != nil || i < 1 {
		return 0, false
	}

	return i, true
}
