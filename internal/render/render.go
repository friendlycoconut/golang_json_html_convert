package render

import "html/template"

// Load parses the HTML template file once at startup.
// The returned template is safe for concurrent Execute calls.
// Implements Single Responsibility
func Load(path string) (*template.Template, error) {
	return template.ParseFiles(path)
}
