package wineMap

import (
	"encoding/json"
	"github.com/grugrut/wine-map-plot/src/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/add", add)
	http.HandleFunc("/api/", api)
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/base.html", "template/main.html"))
	tmpl.Execute(w, nil)
}

func add(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "method: %v, path: %v", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		q := r.URL.Query()
		if q.Get("lat") == "" {
			log.Warningf(ctx, "Lat is null, %v", q.Get("lat"))
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		lat, err := strconv.ParseFloat(q.Get("lat"), 32)
		if err != nil {
			log.Errorf(ctx, err.Error())
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		if q.Get("lng") == "" {
			log.Warningf(ctx, "Lng is null, %v", q.Get("lng"))
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		lng, err := strconv.ParseFloat(q.Get("lng"), 32)
		if err != nil {
			log.Errorf(ctx, err.Error())
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
			log.Errorf(ctx, "name:%v, nameJa:%v, latS:%v, lngS:%v", name, nameJa, latS, lngS)
			http.Redirect(w, r, "/", http.StatusOK)
			return
		}
		lat, err := strconv.ParseFloat(latS, 64)
		if err != nil {
			log.Errorf(ctx, err.Error())
			http.Redirect(w, r, "/", http.StatusOK)
			return
		}
		lng, err := strconv.ParseFloat(lngS, 64)
		if err != nil {
			log.Errorf(ctx, err.Error())
			http.Redirect(w, r, "/", http.StatusOK)
			return
		}

		winery := model.Winery{
			Name:      name,
			NameJa:    nameJa,
			Latitude:  lat,
			Longitude: lng,
		}
		_, err = model.AddWinery(ctx, winery)
		if err != nil {
			log.Errorf(ctx, err.Error())
		}
		http.Redirect(w, r, "/", http.StatusOK)
	default:
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "method: %v, path: %v", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/api/fetch":
			wineries, err := model.FetchAllWinery(ctx)
			if err != nil {
				log.Errorf(ctx, err.Error())
				http.Redirect(w, r, "/", http.StatusBadRequest)
			}
			result, err := json.Marshal(wineries)
			if err != nil {
				log.Errorf(ctx, err.Error())
				http.Redirect(w, r, "/", http.StatusBadRequest)
			}
			w.Write(result)
		}
	default:
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
	}
}
