package handlers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web-programing-susu/pkg/models"
)

var (
	limitDb = 3
)

func (h *Handler) redirectToIndexHandler(c *gin.Context) {
	c.Redirect(301, "/books/0")
}
func (h *Handler) indexHandler(c *gin.Context) {
	vars := c.Request.URL.Query()
	curPage := c.Params.ByName("page")
	curLim := vars.Get("limit")
	data, err := h.services.GetBooksPage(curLim, curPage)
	if err != nil {
		c.Redirect(301, "/")
	}
	if len(data.BooksData) != 0 {
		tmpl, _ := template.ParseFiles("ui/html/index.html")
		tmpl.Execute(c.Writer, data)
	} else {
		c.Redirect(301, "/")
	}
}
func (h *Handler) createBookHandler(c *gin.Context) {
	r := c.Request
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	name := r.FormValue("name")
	author := r.FormValue("author")
	price, _ := strconv.Atoi(r.FormValue("price"))
	rating, _ := strconv.Atoi(r.FormValue("rating"))

	if err := h.services.CreateBook(&models.Product{Name: name, Author: author, Price: price, Rating: rating}); err != nil {
		log.Println(err)
	}
	c.Redirect(301, "/")
}
func (h *Handler) getPageToCreatHandler(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "ui/html/create.html")
}
func (h *Handler) editPage(c *gin.Context) {
	id := c.Params.ByName("id")

	prod, err := h.services.EditBookPage(id)
	if err != nil {
		log.Println(err)
		c.Error(err)
	} else {
		tmpl, _ := template.ParseFiles("ui/html/edit.html")
		tmpl.Execute(c.Writer, prod)
	}
}
func (h *Handler) editHandler(c *gin.Context) {
	r := c.Request
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	author := r.FormValue("author")
	price := r.FormValue("price")
	rating := r.FormValue("rating")
	if err := h.services.EditBookPost(id, name, author, price, rating); err != nil {
		log.Println(err)
	}
	http.Redirect(c.Writer, r, "/", 301)
}
func (h *Handler) deleteHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := h.services.DeleteBook(id); err != nil {
		log.Println(err)
	}
	c.Redirect(301, "/")
}
