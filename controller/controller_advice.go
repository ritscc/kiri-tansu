package controller

import "text/template"

func createTemplateWithHeaderAndFooter(pageTemplate string) template.Template {
	tpl := template.Must(template.ParseFiles("template/item/list.html",
		"template/component/head.html", "template/component/header.html",
		"template/component/footer.html"))

	return *tpl
}
