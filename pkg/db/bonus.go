package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
)

type bonusDB struct {

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

func GetAllBonus() []byte {

    db := connect("savr") 
    defer db.Close()

    rows, err := db.Query("SELECT * FROM bonus")
    if err != nil {
        log.Fatal(err)
    }

    return processRows(rows)
}

func GetBonusByStore(store string) []byte {

    db := connect("savr") 
    defer db.Close()

    rows, err := db.Query("SELECT * FROM bonus WHERE store = ?", store)
    if err != nil {
        log.Fatal(err)
    }

    return processRows(rows)
}

func PostBonus(data []byte) {

    db := connect("savr")
    defer db.Close()

    var bonusPostData []bonusPost
    err := json.Unmarshal(data, &bonusPostData)
    if err != nil {
        log.Fatal(err)
    }

    for _, bonusPost := range bonusPostData {

        log.Println("testing")

        start_date, end_date := getDates(bonusPost)

        // TODO: fix the 'too many open connections' issue. Put all the values in one insert statement instead of an insert statement for each bonusPost
        _, err = db.Query("INSERT INTO bonus (store, start_date, end_date, description, discount, link) VALUES (?, ?, ?, ?, ?, ?)", bonusPost.Supermarket, start_date, end_date, bonusPost.Bonus_description, bonusPost.Discount_description, bonusPost.Link)

        if err != nil {
            log.Println(err)
        } else {
            log.Println("Successfully inserted new row into the table!")
        }

    }
}

func getDates(bonus bonusPost) (int, int) {

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

// TODO: Fix this function to only return false when the row does not exist
func rowExists(bonus bonusPost) bool {

    db := connect("savr")
    defer db.Close()

    start_date, end_date := getDates(bonus)

    row := db.QueryRow("SELECT * from bonus where store = ? and start_date = ? and description = ? and discount = ? and link = ?", bonus.Supermarket, start_date, end_date, bonus.Bonus_description, bonus.Discount_description, bonus.Link) 

    var b bonusDB
    err := row.Scan(&b.Id, &b.Store, &b.Start_date, &b.End_date, &b.Description, &b.Discount, &b.Link)

    if err == sql.ErrNoRows {
        return false
    }

    return true
}

func processRows(rows *sql.Rows) []byte {

    var ret_array []bonusDB

    for rows.Next() {

       var b bonusDB
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

    return ret

}
