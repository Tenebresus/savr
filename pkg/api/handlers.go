package api

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/tenebresus/savr/pkg/db"
)

func Run() {

    http.HandleFunc("GET /app", app)
    http.HandleFunc("GET /api/v1/bonus", getAllBonus)
    http.HandleFunc("POST /api/v1/bonus", postBonus)
    http.HandleFunc("GET /api/v1/bonus/{store}", getBonusByStore)
    http.ListenAndServe("0.0.0.0:8080", nil)

}

func getAllBonus(w http.ResponseWriter, req *http.Request) {

    bonus := db.GetAllBonus()
    io.WriteString(w, string(bonus))

}

func postBonus(w http.ResponseWriter, req *http.Request) {

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

    tp, err := template.ParseFiles("templates/index.html")
    if err != nil {
        log.Fatal(err)
    }

    bonusDBOs := db.GetAllBonusDBO()

    tp.Execute(w, bonusDBOs)

}

