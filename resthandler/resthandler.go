package resthandler

import "net/http"

type RESTHandler struct {
	Index  http.Handler
	Create http.Handler
}

func (rh *RESTHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rh.Index.ServeHTTP(w, r)
	case "POST":
		rh.Create.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}