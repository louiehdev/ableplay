package web

import (
	"net/http"
)

func (f *frontendConfig) handlerHome(w http.ResponseWriter, r *http.Request) {
	f.Templates.ExecuteTemplate(w, "home.html", nil)
}
