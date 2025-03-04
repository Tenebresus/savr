package lidl

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/html"
)

type Lidl struct {}

const aanbiedingen_url string = "https://www.lidl.nl/aanbiedingen"

var ret_array []map[string]string

func Initialize() Lidl {
    return Lidl{}
}

func (l Lidl) Run() []byte {

    resp, err := http.Get(aanbiedingen_url)

    if err != nil {
        log.Println(err)
    }

    body, _ := io.ReadAll(resp.Body)
    reader := bytes.NewReader(body)
    ht := html.NewTokenizer(reader)

    for {

        tt := ht.Next()
        if tt == html.ErrorToken {
            break
        }

        token := ht.Token()

        var lidl LidlT
        for _, attr := range token.Attr {
            if attr.Key == "data-grid-data" {

                start_date := int(time.Now().Unix())
                end_date := int(time.Now().Add(time.Hour * 24 * 7).Unix())

                json.Unmarshal([]byte(attr.Val), &lidl)
                entry := map[string]string {
                    "supermarket": "lidl",
                    "bonus_description": lidl.FullTitle,
                    "start_date": strconv.Itoa(start_date),
                    "end_date": strconv.Itoa(end_date),
                    "discount_description": lidl.Price.Discount.DiscountTitle,
                    "link": "https://www.lidl.nl" + lidl.CanonicalPath,
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
