// Package view contains
package view

// Page is used for template generation and provides within
// Data structures the will be parsed by the templates
// Errors holds the name of the HTML input field and the errormessage
// Slug is the last part of the breadcrumb e.g. 'Products'
// Titel is the page title
type Page struct {
	Title  string            // Page title within the browser
	Slug   string            // is the plural of the entity (for breadcrumb and to mark active menu entry)
	Data   interface{}       // the data that will be handed over to the templates
	Errors map[string]string // holds the errors on a page
}

// NewPage creates a page with a title and the slug (breadcrumb)
func NewPage(title string, slug string) *Page {
	p := &Page{Title: title, Slug: slug}
	p.Errors = make(map[string]string)
	return p
}
