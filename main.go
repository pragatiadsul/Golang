// Package classification Account API.
//
// this is to show how to write RESTful APIs in golang.
// that is to provide a detailed overview of the language specs
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:8080
//     Version: 1.0.0
//     Contact: Supun Muthutantri<mydocs@example.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
// swagger:meta
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/articles/id", createNewArticle)
	//	http.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 3", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
func main() {
	// articles := []Article{
	// 	Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	// 	Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	// }

	// swagger:operation GET /hello users
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: successful operation
	handleRequests()

}
