package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web-programing-susu/models"
	_ "web-programing-susu/models"
)

var (
	limitDb  = 3
	connStr  = "user=postgres password=postgres dbname=WebProgramingSusu sslmode=disable"
	database *sql.DB
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	curPage := vars["page"]
	curLim := r.URL.Query().Get("limit")
	if curLim != "" {
		limitDb, _ = strconv.Atoi(curLim)
	}
	fmt.Println(curLim)
	res, err := strconv.Atoi(curPage)
	var prevPage int
	var nextPage int
	if curPage == "" || err != nil {
		curPage = "0"
		prevPage = 0
		nextPage = 1
	} else {
		if res == 0 {
			prevPage = 0
			nextPage = 1
		} else {
			prevPage = res - 1
			nextPage = res + 1
		}
	}
	var sqlText = "SELECT * FROM" + " books ORDER BY ID LIMIT " + strconv.Itoa(limitDb) + " OFFSET " + strconv.Itoa(res*limitDb)
	rows, err := database.Query(sqlText)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Author, &p.Price, &p.Rating)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	if len(products) != 0 {
		data := models.IndexViewData{RealPage: strconv.Itoa(res), CurPage: strconv.Itoa(res + 1), BooksData: products, PrevPage: strconv.Itoa(prevPage), NextPage: strconv.Itoa(nextPage)}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", 301)
	}

}
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		name := r.FormValue("name")
		author := r.FormValue("author")
		price := r.FormValue("price")
		rating := r.FormValue("rating")

		sqlText := "INSERT INTO" + " books (firstname, author, price, rating) VALUES ('" + name + "', '" + author + "', " + price + "," + rating + ")"
		_, err = database.Exec(sqlText)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "templates/create.html")
	}
}

// Возвращаем позьзователю страницу с полями объекта с указаным id
func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sqlText := "SELECT * FROM" + " books " + "WHERE id = " + id

	row := database.QueryRow(sqlText)
	prod := models.Product{}
	err := row.Scan(&prod.Id, &prod.Name, &prod.Author, &prod.Price, &prod.Rating)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("templates/edit.html")
		tmpl.Execute(w, prod)
	}
}
func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	author := r.FormValue("author")
	price := r.FormValue("price")
	rating := r.FormValue("rating")
	sqlText :=
		"UPDATE" + " books SET firstname = '" + name + "', author = '" + author + "', price = " + price + ", rating = " + rating + " WHERE id = " + id

	_, err = database.Exec(sqlText)

	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 301)

}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sqlText := "DELETE FROM" + " books WHERE id = " + id

	_, err := database.Exec(sqlText)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
func RedirectToView(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/0", 301)
}
func main() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	database = db
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", RedirectToView)
	router.HandleFunc("/view/{page:[0-9]+}", IndexHandler)
	router.HandleFunc("/create", CreateHandler)
	router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", DeleteHandler)
	//http.HandleFunc("/create", UsersHandler)

	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}
