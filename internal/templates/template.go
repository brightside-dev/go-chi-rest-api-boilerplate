package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

type TemplateData struct {
	CurrentYear int
	Form        any
	Flash       string
}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func NewTemplateCache() (map[string]*template.Template, error) {
	// Initialise a new ma to act as the cache
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() function to get a slice of all filepaths that
	// match the pattern "./ui/html/pages/*.tmpl". This will essentially gives
	// us a slice of all the filepaths for our application 'page' templates
	// like: [ui/html/pages/home.tmpl ui/html/pages/view.tmpl]
	pages, err := filepath.Glob("./views/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.tmpl') from the full filepath
		// and assign it to the name variable
		name := filepath.Base(page)

		// Parse the base template file into a template set.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./views/html/base.html")
		if err != nil {
			return nil, err
		}

		// Parse all the 'partial' templates into the set.
		ts, err = ts.ParseGlob("./views/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Parse the page template file into the set.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key.
		cache[name] = ts
	}

	// Return the map
	return cache, nil
}

func Render(w http.ResponseWriter, r *http.Request, tmpl string, data *TemplateData, templateCache map[string]*template.Template) {
	// Retrieve the template set from the cache based on the page name
	ts, ok := templateCache[tmpl]
	if !ok {
		// print error
		fmt.Println("Error getting template from cache")
		return
	}

	// Initialize a new buffer
	buf := new(bytes.Buffer)

	// Write the template to the buffer, instead of straight to the
	// http.ResponseWriter
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		fmt.Println("Error executing template")
		return
	}

	// Write the contents of the buffer to the http.ResponseWriter
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser")
		return
	}
}
