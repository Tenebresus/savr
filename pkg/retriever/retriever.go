package retriever

import (
	"fmt"

	"github.com/tenebresus/savr/pkg/retriever/ah"
)

type Runner interface {
    Run() []byte
}

func Run() {

    var runners []Runner
    runners = append(runners, ah.Initialize())

    for _, runner := range runners {
        json := runner.Run()
        fmt.Println(string(json))
    }

}

