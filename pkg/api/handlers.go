package api

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tenebresus/savr/pkg/db"
)

func Run() {

    // Main App
    http.HandleFunc("GET /app", app)
    http.HandleFunc("GET /static/", static)

    // API
    http.HandleFunc("GET /api/v1/bonus", getAllBonus)
    http.HandleFunc("POST /api/v1/bonus", postBonus)
    http.HandleFunc("GET /api/v1/bonus/{store}", getBonusByStore)

    log.Println("started listening on port 8080...")
    http.ListenAndServe(":8080", nil)

}

func getAllBonus(w http.ResponseWriter, req *http.Request) {

    bonus := db.GetAllBonus()
    io.WriteString(w, string(bonus))

}

func postBonus(w http.ResponseWriter, req *http.Request) {

    log.Println("Received POST data from retriever...")
    postData, err := io.ReadAll(req.Body)
    if err != nil {
        log.Println(err)
    } else {
        db.PostBonus(postData)
    }

}

func getBonusByStore(w http.ResponseWriter, req *http.Request) {

    bonus := db.GetBonusByStore(req.PathValue("store"))
    io.WriteString(w, string(bonus))

}

func app(w http.ResponseWriter, req *http.Request) {

    log.Println("Received request")

    tp, err := template.ParseFiles("static/templates/index.html")
    if err != nil {
        log.Fatal(err)
    }

    search := req.URL.Query().Get("search")
    bonusDBOs := []db.BonusDBO{}

    if search != "" {
        bonusDBOs = db.GetBonusWhereDescriptionDBO(search)
    } else {
        bonusDBOs = db.GetAllBonusDBO()
    }
    tp.Execute(w, bonusDBOs)

}

func static(w http.ResponseWriter, req *http.Request) {

    path := strings.Split(req.URL.Path, "/")
    file := path[2]

    file_path := ""

    if strings.HasSuffix(file, "css") {
        file_path = "static/stylesheets/" + file
        w.Header().Set("Content-Type", "text/css")
    }

    if strings.HasSuffix(file, "png") {
        file_path = "static/images/" + file
        w.Header().Set("Content-Type", "image/png")
    }

    data, err := os.ReadFile(file_path)
    if err != nil {
        log.Println(err)
    }

    w.Write(data)

}
