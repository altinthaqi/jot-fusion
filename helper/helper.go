package helper

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func isEmpty(v any) bool {
	return v == nil || v == ""
}

func GetQueryInt(r *http.Request, key string) (int, error) {
	v := mux.Vars(r)[key]

	if isEmpty(v) {
		return 0, fmt.Errorf("missing required field %s", key)
	}

	i, err := strconv.Atoi(v)

	if err != nil || i < 1 {
		return 0, fmt.Errorf("invalid %s", key)
	}

	return i, nil
}
