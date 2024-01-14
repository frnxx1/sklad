package main

import (
	ind "sklad/app/indexview"
	logv "sklad/app/loginview"
	"github.com/gin-gonic/gin"
)

func handlerFunc() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("index.html",
		"templates/login.html",
		"templates/auction.html",
		"templates/information.html",
		"templates/products.html",
		"templates/promotions.html",
		"templates/sign.html")

	r.Static("/css", "./css/.")
	r.Static("/img", "./img/.")
	r.Static("/js", "./js/.")

	r.GET("/", ind.IndexView)
	r.GET("/login", logv.LoginView)

	return r
}

func main() {
	router := handlerFunc()
	router.Run("localhost:8080")
}
