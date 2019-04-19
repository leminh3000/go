package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
	CreatedOn JSONTime `json:"CreatedOn"`
}
type JSONTime time.Time

func (t JSONTime)MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%d\"", time.Time(t).Unix())
	return []byte(stamp), nil
}

type Articles []Article

func allArticles(writer http.ResponseWriter, request *http.Request) {
	articles := Articles{Article{Content:"jfasld kljsakd", Desc:"desc `1" , Title:"tilesdkjfkls 2222", CreatedOn:JSONTime(time.Now())}, Article{Content:"jfasld kljsakd", Desc:"desc `2" , Title:"tilesdkjfkls 222", CreatedOn:JSONTime(time.Now())}	}
	fmt.Println("Enpoint hit: all articles")
	json.NewEncoder(writer).Encode(articles)
}


func main() {
	handleRequest()
}
