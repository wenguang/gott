package main

import (
	"crypto/sha1"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"os"
	//"reflect"
	"sort"
	"strings"
	//"unicode/utf8"
)

func main() {
	http.HandleFunc("/", apiParamsSort)
	http.ListenAndServe(":9000", nil)
}

func apiParamsSort(w http.ResponseWriter, r *http.Request) {

	//	var authorization string
	//	authorization = r.Header.Get("Authorization")
	//	fmt.Fprintln(w, reflect.TypeOf(authorization))
	//	if authorization == nil {
	//		fmt.Fprintln(w, "Authorization invavid")
	//	}

	var params map[string][]string
	if r.Method == http.MethodGet {
		params = r.URL.Query()
	} else if r.Method == http.MethodPost {
		params = r.Form
	} else {
		fmt.Fprintln(w, "hello")
		return
	}

	if params == nil || len(params) == 0 {
		fmt.Fprintln(w, "hello")
		return
	}

	i := 0
	paramKeys := make([]string, len(params))
	for k, _ := range params {
		paramKeys[i] = k
		i++
	}
	sort.Strings(paramKeys)
	paramVals := make([]string, len(paramKeys))
	for j := 0; j < len(paramKeys); j++ {
		value, ok := params[paramKeys[j]]
		if ok {
			valueStr := strings.Join(value, "")
			paramVals[j] = valueStr
		}
	}
	paramsValueString := strings.Join(paramVals, "")
	fmt.Fprintln(w, paramsValueString)

	apiSecret := "abc123"
	strs := []string{apiSecret, paramsValueString}
	secretStr := strings.Join(strs, "")
	fmt.Fprintln(w, secretStr)
}
