package controller

import "net/http"

func Routing() {
	// css, js
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("resources/css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("resources/js/"))))

	http.HandleFunc("/", GetItemList)
}
