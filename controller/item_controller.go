package controller

import (
	"log"
	"net/http"
)

func GetItemList(w http.ResponseWriter, r *http.Request) {
	tpl := createTemplateWithHeaderAndFooter("template/item/list.html")

	attr := struct {
		PageTitle string
		Greet     string
	}{
		PageTitle: "Home",
		Greet:     "Go",
	}

	if err := tpl.ExecuteTemplate(w, "list.html", attr); err != nil {
		log.Fatal(err)
	}
}
