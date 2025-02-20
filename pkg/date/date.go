package date

import (
	"log"
	"time"
)

func Parse(layout string, value string) int64 {

    t, err := time.Parse(layout, value)
    if err != nil {
        log.Fatal(err)
    }

    return t.Unix()

}
