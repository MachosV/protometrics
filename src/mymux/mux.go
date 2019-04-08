package mymux

import "net/http"

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
}

/*
GetMux returns the mux that functions register to
*/
func GetMux() *http.ServeMux {
	return mux
}
