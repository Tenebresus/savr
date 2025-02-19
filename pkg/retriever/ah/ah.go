package ah

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

type AH struct {}

type authResT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

var ret_array []map[string]string
var wg sync.WaitGroup

func (a AH) Run() []byte {

    var authRes authResT
    json.Unmarshal(getAuthToken(), &authRes)

    categories := getCategories()
    promotion_types := getPromotionType()

    for _, category := range categories {
        go getBonusByGroup(authRes.AccessToken, category, false)
    }

    for _, promotion_type := range promotion_types {
        go getBonusByGroup(authRes.AccessToken, promotion_type, true)
    }

    wg.Wait()

    ret, err := json.Marshal(ret_array)
    if err != nil {
        log.Fatal(err)
    }

    return ret

}

func Initialize() AH {
    return AH{}
}

func getBonusByGroup(access_token string, group string, is_promotion bool) {

    wg.Add(1)
    defer wg.Done()

    var bonusRes bonusResT
    bonus := getBonus(access_token, group, is_promotion)
    json.Unmarshal(bonus, &bonusRes)

    for _, product := range bonusRes.BonusGroupOrProducts {

        go getBonusByProduct(product.BonusGroup)

    } 

}

func getBonusByProduct(product bonusGroupT) {

    entry := map[string]string {
        "bonus_description": product.SegmentDescription,
        "start_date": product.BonusStartDate,
        "end_date": product.BonusEndDate,
        "discount_description": product.DiscountDescription,
    }

    ret_array = append(ret_array, entry)
}

func getAuthToken() []byte {

    client := &http.Client{}
	var data = strings.NewReader(`{"clientId":"appie"}`)
	req, err := http.NewRequest("POST", "https://api.ah.nl/mobile-auth/v1/auth/token/anonymous", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Appie/8.22.3")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
    if err != nil {
		log.Fatal(err)
	}

    return bodyText

}

func getBonus(access_token string, value string, isPromotion bool) []byte {

    option := ""

    if isPromotion {
        option = "?promotionType=" + value
    } else {
        option = "?category=" + value + "&promotionType=NATIONAL"
    }

    client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.ah.nl/mobile-services/bonuspage/v2/section" + option, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Appie/8.22.3")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", "AHWEBSHOP")
	req.Header.Set("Authorization", "Bearer " + access_token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
    return bodyText
}

func getPromotionType() []string {

    return []string {
        "ETOS",
        "GALL",
        "GALLCARD",
        "AHONLINE",
    }

}

func getCategories() []string {

    return []string {
        "Aardappel%2C%20groente%2C%20fruit",
        "Salades%2C%20pizza%2C%20maaltijden",
        "Vlees%2C%20vis",
        "Vegetarisch%2C%20vegan%20en%20plantaardig",
        "Kaas%2C%20vleeswaren%2C%20tapas",
        "Zuivel%2C%20eieren%2C%20boter",
        "Bakkerij",
        "Ontbijtgranen%20en%20beleg",
        "Chips%2C%20noten%2C%20toast%2C%20popcorn",
        "Snoep%2C%20chocolade%2C%20koek",
        "Koffie%2C%20thee",
        "Frisdrank%2C%20sappen%2C%20siropen%2C%20water",
        "Wijn%20en%20bubbels",
        "Bier%20en%20aperitieven",
        "Pasta%2C%20rijst%20en%20wereldkeuken",
        "Soepen%2C%20sauzen%2C%20kruiden%2C%20olie",
        "Diepvries",
        "Drogisterij",
        "Gezondheid%2C%20sport",
        "Baby%20en%20kind",
        "Huishouden",
        "Huisdier",
        "Koken%2C%20tafelen%2C%20vrije%20tijd",
    }

}
