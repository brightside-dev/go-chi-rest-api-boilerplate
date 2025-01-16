package templates

import (
	"html/template"
	"net/http"
	"time"
)

type TemplateData struct {
	CurrentYear int
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

func Render(w http.ResponseWriter, r *http.Request, page string, data *TemplateData) {
	files := []string{
		"./views/html/base.html",
		"./views/html/partials/nav.html",
		"./views/html/pages/home.html",
	}

	// Parse the template files...
	ts, err := template.New("").Funcs(functions).ParseFiles(files...)
	if err != nil {
		println(err.Error())
		return
	}

	// And then execute them. Notice how we are passing in the snippet
	// data (a models.Snippet struct) as the final parameter?
	// Pass in the templateData struct when executing the template.
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		println(err.Error())
		return
	}
}
