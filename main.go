package main

import (
	// "golang.org/x/crypto/acme"
	// "database/sql"
	// "flag"
	"fmt"
	"html/template"
	"io"
	// "log"
	
	"net/http"
	
	"os"
	// "regexp"
	// "strconv"
	// "time"
	// "unicode"
	// "strings"

	// "github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// _ "github.com/mattn/go-sqlite3"
)

type Template struct {
	templates *template.Template
}

type MovieStruct struct {
	Name 			string
	Year 			string
	PosterAddr 		string
	Size 			string
	Path 			string
	Idx 			string
	MovId 			string
	Catagory 		string
	HttpThumbPath 	string
}

type TvShowStruct struct {
	TvId 			string
	Size			string
	Catagory		string
	Name 			string
	Season 			string
	Episode 		string
	Path 			string
	Idx 			string
}

func checkDBExists() {
	mtvDBPath := os.Getenv("MTV_DB_PATH")
	if _, err := os.Stat(mtvDBPath); os.IsNotExist(err) {
		// file does not exist
		fmt.Println("Database file does not exist\n Please run mtvsetup.")
		os.Exit(1)
	} else if err != nil {
		// other error
		fmt.Println("Error checking for database file: ", err)
		os.Exit(1)
	}
	// file exists
	fmt.Println("Database file exists.")
}

func init() {

	godotenv.Load("mtvhtmxecho.env")
	// checkDBExists()

	

	

	// filePath := os.Getenv("ATS_DB_PATH")
	// _, err3 := os.OpenFile(filePath, os.O_CREATE|os.O_EXCL, 0755)
	// if err3 != nil {
	// 	if os.IsExist(err3) {
	// 		fmt.Println("file exists")
	// 	} else {
	// 		fmt.Println(err3)
	// 		fmt.Print("unable to create db file")
	// 	}

	// }
}

func main() {

	
		e := echo.New()
		e.Use(middleware.CORS())
		e.Use(middleware.Gzip())
		// e.Use(middleware.Recover())
		t := &Template{
			templates: template.Must(template.ParseGlob("MtvTemplates/*")),
		}
		e.Renderer = t

		e.GET("/", mtv_index)
		e.GET("/movies", mtv_movies)
		e.GET("/movaction", mov_action)
		e.GET("/movarnold", mov_arnold)
		e.GET("/movbrucelee", mov_brucelee)
		e.GET("/movbrucewillis", mov_brucewillis)

		e.GET("/tvshows", mtv_tvshows)
		e.GET("/tvaction", tv_action)
		e.GET("/tvcomedy", tv_comedy)
		e.GET("/tvfantasy", tv_fantasy)
		e.GET("/tvstartrek", tv_startrek)
		e.GET("/tvstarwars", tv_starwars)
		e.GET("/tvscifi", tv_scifi)
		e.GET("/tvscience", tv_science)
		e.GET("/tvmcu", tv_mcu)
		e.GET("/tvwestern", tv_western)
		e.GET("/admin", mtv_admin)
		e.Static("/assets", "assets")
		e.Logger.Fatal(e.Start(":8080"))
	
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func mov_action(c echo.Context) error {
	return c.Render(http.StatusOK, "mov_action", "WORKED")
}

func mov_arnold(c echo.Context) error {
	return c.Render(http.StatusOK, "mov_arnold", "WORKED")
}

func mov_brucelee(c echo.Context) error {
	return c.Render(http.StatusOK, "mov_brucelee", "WORKED")
}

func mov_brucewillis(c echo.Context) error {
	return c.Render(http.StatusOK, "mov_brucewillis", "WORKED")
}













func mtv_index(c echo.Context) error {
	return c.Render(http.StatusOK, "mtv_index", "WORKED")
}

func mtv_movies(c echo.Context) error {
	return c.Render(http.StatusOK, "mtv_movies", "WORKED")
}

func mtv_tvshows(c echo.Context) error {
	return c.Render(http.StatusOK, "mtv_tvshows", "WORKED")
}

func mtv_admin(c echo.Context) error {
	return c.Render(http.StatusOK, "mtv_admin", "WORKED")
}

func tv_action(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_action", "WORKED")
}

func tv_comedy(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_comedy", "WORKED")
}

func tv_fantasy(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_fantasy", "WORKED")
}

func tv_startrek(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_startrek", "WORKED")
}

func tv_starwars(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_starwars", "WORKED")
}

func tv_scifi(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_scifi", "WORKED")
}

func tv_science(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_science", "WORKED")
}

func tv_mcu(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_mcu", "WORKED")
}

func tv_western(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_western", "WORKED")
}

