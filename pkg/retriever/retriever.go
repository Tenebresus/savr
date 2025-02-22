package retriever

import (
	"log"
	"net/http"
	"strings"

	"github.com/tenebresus/savr/pkg/retriever/ah"
)

type Runner interface {
    Run() []byte
}

func Run() {

    var runners []Runner
    runners = append(runners, ah.Initialize())

    for _, runner := range runners {
        bonusData := runner.Run()
        postRequest(bonusData)
    }

}

func postRequest(bonusData []byte) {

    data := strings.NewReader(string(bonusData))
    _, err := http.Post("http://127.0.0.1:8080/api/v1/bonus", "application/json", data)

    if err != nil {
        log.Println(err)
    }

}

