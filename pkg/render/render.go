package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates using "html/template"
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}

}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if the template is already in the cache
	_, inMap := tc[t]
	if !inMap {
		// create the template if not in the cache
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)

		if err != nil {
			log.Println(err)
		}
	} else {
		// template in cache already
		log.Println("using template in cache")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}
	// parse the template
	// ... takes each entry in a slice and inserts them as individual strings
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// add template to the cache
	tc[t] = tmpl

	return nil
}
