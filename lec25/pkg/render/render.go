package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//create a template cache

	tc, err := createTemplateCache() // tc = Template Cache

	if err != nil {
		log.Fatal(err)
	}

	//get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template) The other way to implement is like following
	myCache := map[string]*template.Template{} //Empty Slice of Pointers declared

	//get all the files named *page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//range through all files ending with *page.tmpl
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
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl") // Mixes the Base Layout with templates
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
		//ts = Template Set
	}
	// fmt.Println(myCache)
	return myCache, nil

}
