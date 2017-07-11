package main

import (
	"html/template"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/ONSdigital/dp-frontend-renderer/assets"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/ageSelectorList"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/ageSelectorRange"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset-filter/filterOverview"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset/finishPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset/middlePage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/dataset/startPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/datasetLandingPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/errorPage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/handlers/productPage"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/handlers/healthcheck"
	"github.com/ONSdigital/go-ns/handlers/requestID"
	"github.com/ONSdigital/go-ns/handlers/timeout"
	"github.com/ONSdigital/go-ns/log"
	"github.com/c2h5oh/datasize"
	"github.com/gorilla/pat"
	"github.com/justinas/alice"
	unrolled "github.com/unrolled/render"
)

func main() {
	bindAddr := os.Getenv("BIND_ADDR")
	if len(bindAddr) == 0 {
		bindAddr = ":20010"
	}

	var err error
	config.DebugMode, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Error(err, nil)
	}

	if config.DebugMode {
		config.PatternLibraryAssetsPath = "http://localhost:9000/dist"
	}

	log.Namespace = "dp-frontend-renderer"

	render.Renderer = unrolled.New(unrolled.Options{
		Asset:         assets.Asset,
		AssetNames:    assets.AssetNames,
		IsDevelopment: config.DebugMode,
		Layout:        "main",
		Funcs: []template.FuncMap{{
			"humanSize": func(s int) string {
				return datasize.ByteSize(s).HumanReadable()
			},
			"safeHTML": func(s string) template.HTML {
				return template.HTML(s)
			},
			"dateFormat": func(s string) template.HTML {
				t, err := time.Parse(time.RFC3339, s)
				if err != nil {
					log.Error(err, nil)
					return template.HTML(s)
				}
				return template.HTML(t.Format("02 January 2006"))
			},
			"last": func(x int, a interface{}) bool {
				return x == reflect.ValueOf(a).Len()-1
			},
			"loop": func(n, m int) []int {
				arr := make([]int, m-n)
				v := n
				for i := 0; i < m-v; i++ {
					arr[i] = n
					n++
				}
				return arr
			},
			"subtract": func(x, y int) int {
				return x - y
			},
		}},
	})

	router := pat.New()
	alice := alice.New(
		timeout.Handler(10*time.Second),
		log.Handler,
		requestID.Handler(16),
	).Then(router)

	router.Get("/healthcheck", healthcheck.Handler)
	router.Post("/homepage", homepage.Handler)
	router.Post("/dataset-landing-page-static", datasetLandingPage.StaticHandler)
	router.Post("/dataset-landing-page-filterable", datasetLandingPage.FilterHandler)
	router.Post("/productPage", productPage.Handler)
	router.Post("/error", errorPage.Handler)
	router.Post("/dataset-filter/filter-overview", filterOverview.Handler)
	router.Post("/dataset-filter/age-selector-range", ageSelectorRange.Handler)
	router.Post("/dataset-filter/age-selector-list", ageSelectorList.Handler)
	router.Post("/dataset/startpage", startPage.Handler)
	router.Post("/dataset/middlepage", middlePage.Handler)
	router.Post("/dataset/finishpage", finishPage.Handler)

	log.Debug("Starting server", log.Data{"bind_addr": bindAddr})

	server := &http.Server{
		Addr:         bindAddr,
		Handler:      alice,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}
