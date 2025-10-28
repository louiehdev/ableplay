package frontend

import (
	"net/http"
)

func (f *frontendConfig) handlerHome(w http.ResponseWriter, r *http.Request) {
	f.templates.ExecuteTemplate(w, "home.html", nil)
}
