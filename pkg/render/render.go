package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using "html/template"
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// old way to parse templates, replaced by what's above
	//parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	//err := parsedTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("error parsing template: ", err)
	//	return
	//}

}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files named *page.tmpl from the templates folder ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all found files ending in *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}

// a simple way to make a template cache, commented out to build out the more complex way of doing it
//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	// check to see if the template is already in the cache
//	_, inMap := tc[t]
//	if !inMap {
//		// create the template if not in the cache
//		log.Println("creating template and adding to cache")
//		err = createTemplateCache(t)
//
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		// template in cache already
//		log.Println("using template in cache")
//	}
//
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//
//	if err != nil {
//		log.Println(err)
//	}
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
//	}
//	// parse the template
//	// ... takes each entry in a slice and inserts them as individual strings
//	tmpl, err := template.ParseFiles(templates...)
//
//	if err != nil {
//		return err
//	}
//
//	// add template to the cache
//	tc[t] = tmpl
//
//	return nil
//}
