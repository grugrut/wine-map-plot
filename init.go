package wineMap

import (
	"html/template"
	"model"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/add", add)
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/base.html", "template/main.html"))
	tmpl.Execute(w, nil)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		q := r.URL.Query()
		if q.Get("lat") == "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		lat, err := strconv.ParseFloat(q.Get("lat"), 32)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		if q.Get("lng") == "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		lng, err := strconv.ParseFloat(q.Get("lng"), 32)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		latlng := []float64{lat, lng}
		tmpl := template.Must(template.ParseFiles("template/base.html", "template/add.html"))
		tmpl.Execute(w, latlng)
	case http.MethodPost:
		name := r.FormValue("wineryName")
		nameJa := r.FormValue("wineryNameJa")
		latS := r.FormValue("wineryLat")
		lngS := r.FormValue("wineryLng")
		if name == "" || nameJa == "" || latS == "" || lngS == "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		lat, err := strconv.ParseFloat(latS, 64)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		lng, err := strconv.ParseFloat(lngS, 64)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		winery := model.Winery{
			Name:      name,
			NameJa:    nameJa,
			Latitude:  lat,
			Longitude: lng,
		}
		model.AddWinery(r.Context(), winery)
		http.Redirect(w, r, "/", http.StatusFound)
	default:

		http.Redirect(w, r, "/", 500)
	}
}
