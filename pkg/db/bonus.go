package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type BonusDBO struct {

    Id          int    `json:"id"`
    Store       string `json:"store"`
    Start_date  int    `json:"start_date"`
    End_date    int    `json:"end_date"`
    Description string `json:"description"`
    Discount    string `json:"discount"`
    Link        string `json:"link"`
    
}

type bonusPost struct {

    Bonus_description    string `json:"bonus_description"`
    Discount_description string `json:"discount_description"`
    End_date             string `json:"end_date"`
    Start_date           string `json:"start_date"`
    Link                 string `json:"link"`
    Supermarket          string `json:"supermarket"`

}

var time_mapping map[string]int = map[string]int {

    "Monday": 0,
    "Tuesday": 1,
    "Wednesday": 2,
    "Thursday": 3,
    "Friday": 4,
    "Saturday": 5,
    "Sunday": 6,

}

func Find(slct string, where ...string) ([]BonusDBO, []byte) {

    db := connect("savr") 
    defer db.Close()

    query := buildQuery(slct, where...)
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }

    retDBO, ret := processRows(rows)
    return retDBO, ret
}

func buildQuery(slct string, where ...string) string {

    query := "SELECT " + slct + " FROM bonus"
    if len(where) > 0 {

        query += " WHERE "
        for _, where_part := range where {
            query += where_part + " "
        }

    }

    return query

}

func PruneOldBonus() {

    db := connect("savr") 
    defer db.Close()

    _, err := db.Query("DELETE FROM bonus where end_date < ?", time.Now().Unix())
    if err != nil {
        log.Fatal(err)
    }

}

func PostBonus(data []byte) {

    db := connect("savr")
    defer db.Close()

    var bonusPostData []bonusPost
    err := json.Unmarshal(data, &bonusPostData)
    if err != nil {
        log.Fatal(err)
    }

    query := ""

    for _, bonusPost := range bonusPostData {

        start_date, end_date := 0, 0

        if bonusPost.Start_date == "" && bonusPost.End_date == "" {
            start_date, end_date = getDefaultStartDates()
        } else {
            start_date, end_date = convDates(bonusPost)
        }

        query += fmt.Sprintf("call InsertBonus(\"%s\", %d, %d, \"%s\", \"%s\", \"%s\");", bonusPost.Supermarket, start_date, end_date, bonusPost.Bonus_description, bonusPost.Discount_description, bonusPost.Link)
    }

    _, err = db.Query(query)
    if err != nil {
        log.Println(err)
    } else {
        log.Println("Successfully inserted rows into the table!")
    }
}

func getDefaultStartDates() (int, int) {

    now := time.Now()
    base := now.Unix() - int64(now.Minute() * 60) - int64(now.Hour() * 60 * 60) - int64(now.Second())

    today_string := now.Weekday().String()
    start_date := int(base) - (time_mapping[today_string] * 24) * 3600
    end_date := int(base) + ((7 - time_mapping[today_string]) * 24) * 3600

    return start_date, end_date

}

func convDates(bonus bonusPost) (int, int) {

        start_date, err := strconv.Atoi(bonus.Start_date)
        if err != nil {
            log.Println(err)
        }

        end_date, err := strconv.Atoi(bonus.End_date)
        if err != nil {
            log.Println(err)
        }

        return start_date, end_date
}

// processRows returns both the DBOs and the Marshalled DBOs; omit the return type you don't want
func processRows(rows *sql.Rows) ([]BonusDBO, []byte) {

    var ret_array []BonusDBO

    for rows.Next() {

       var b BonusDBO
       err := rows.Scan(&b.Id, &b.Store, &b.Start_date, &b.End_date, &b.Description, &b.Discount, &b.Link)
       if err != nil {
           log.Fatal(err)
       }
       ret_array = append(ret_array, b)

    }

    ret, err := json.Marshal(ret_array)
    if err != nil {
        log.Fatal(err)
    }

    return ret_array, ret

}
