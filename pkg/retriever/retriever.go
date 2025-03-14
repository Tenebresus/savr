package retriever

import (
	"log"
	"net/http"
	"strings"

	"github.com/tenebresus/savr/pkg/os"
	"github.com/tenebresus/savr/pkg/retriever/ah"
	"github.com/tenebresus/savr/pkg/retriever/lidl"
	"github.com/tenebresus/savr/pkg/retriever/deka"
)

type Runner interface {
    Run() []byte
}

func Run() {

    var runners []Runner
    runners = append(runners, ah.Initialize())
    runners = append(runners, deka.Initialize())
    runners = append(runners, lidl.Initialize())

    for _, runner := range runners {
        bonusData := runner.Run()
        postRequest(bonusData)
    }

}

func postRequest(bonusData []byte) {

    data := strings.NewReader(string(bonusData))
    _, err := http.Post("http://" + os.GetEnv("API_HOST") + ":8080/api/v1/bonus", "application/json", data)

    if err != nil {
        log.Println(err)
    }

}

