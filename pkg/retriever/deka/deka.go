package deka

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type Deka struct{}

const aanbiedingen_url string = "https://www.dekamarkt.nl/aanbiedingen"

var ret_array []map[string]string

func Initialize() Deka {
    return Deka{}
}

func (d Deka) Run() []byte {

    resp, err := http.Get(aanbiedingen_url)

    if err != nil {
        log.Println(err)
    }

    body, _ := io.ReadAll(resp.Body)
    reader := bytes.NewReader(body)
    ht := html.NewTokenizer(reader)

    var description string
    var discount string

    for {

        tt := ht.Next()
        if tt == html.ErrorToken {
            break
        }

        token := ht.Token()

        for _, attr := range token.Attr {

            if attr.Key == "class" && attr.Val == "title" {

                ht.Next()
                description = string(ht.Text())

            }

            if attr.Key == "class" && attr.Val == "chip" {

                ht.Next()
                discount = string(ht.Text())

                entry := map[string]string {
                    "supermarket": "deka",
                    "bonus_description": description,
                    "start_date": "",
                    "end_date": "",
                    "discount_description": discount,
                    "link": aanbiedingen_url,
                }

                ret_array = append(ret_array, entry)

            }

        }

    }

    ret, err := json.Marshal(ret_array)
    if err != nil {
        log.Fatal(err)
    }

    return ret
}
