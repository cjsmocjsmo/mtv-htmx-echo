package main

import (
	// "crypto/tls"
	// "golang.org/x/crypto/acme"
	// "database/sql"
	// "flag"
	// "fmt"
	"html/template"
	"io"
	// "log"
	// "mime/multipart"
	"net/http"
	// "net/mail"
	// "os"
	// "regexp"
	// "strconv"
	// "time"
	// "unicode"
	// "strings"

	// "github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/mailjet/mailjet-apiv3-go/v4"
	// _ "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/acme/autocert"
)

type Template struct {
	templates *template.Template
}

// func createAccountsDB(db_path string) {
// 	db, err := sql.Open("sqlite3", db_path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	sqlStmt := `
// 	CREATE TABLE IF NOT EXISTS accounts (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		acctid TEXT NOT NULL UNIQUE,
// 		email TEXT NOT NULL UNIQUE,
// 		date TEXT NOT NULL
// 	);
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createCommentsDB(db_path string) {
// 	db, err := sql.Open("sqlite3", db_path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	sqlStmt := `
// 	CREATE TABLE IF NOT EXISTS comments (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		acctid TEXT NOT NULL,
// 		comid TEXT NOT NULL,
// 		name TEXT NOT NULL,
// 		email TEXT NOT NULL,
// 		rating TEXT NOT NULL,
// 		comment TEXT NOT NULL,
// 		date TEXT NOT NULL,
// 		media TEXT NOT NULL,
// 		status TEXT NOT NULL
// 	);
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createEstimatesDB(db_path string) {
// 	db, err := sql.Open("sqlite3", db_path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	sqlStmt := `
// 	CREATE TABLE IF NOT EXISTS estimates (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		acctid TEXT NOT NULL,
// 		estid TEXT NOT NULL,
// 		name TEXT NOT NULL,
// 		address TEXT NOT NULL,
// 		city TEXT NOT NULL,
// 		phone TEXT NOT NULL,
// 		email TEXT NOT NULL,
// 		servdate TEXT NOT NULL,
// 		recdate TEXT NOT NULL,
// 		comment TEXT NOT NULL,
// 		media TEXT NOT NULL,
// 		status TEXT NOT NULL
// 	);
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createUploadsDir(uploads_path string) {
// 	_, err := os.Stat(uploads_path)
// 	if os.IsNotExist(err) {
// 		err := os.MkdirAll(uploads_path, 0755)
// 		if err != nil {
// 			fmt.Println(err)
// 			fmt.Print("unable to create uploads dir")
// 		}
// 	}
// }

// func dbFileExists(db_path string) bool {
// 	_, err := os.Stat(db_path)
// 	return os.IsNotExist(err) 
// }

// func createCertDir(cert_path string) {
// 	err := os.MkdirAll(cert_path, 0755)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Print("unable to create cert dir")
// 	}
// }

func init() {

	// godotenv.Load("atshtmxecho.env")
	godotenv.Load("test.env")

	// if !dbFileExists(os.Getenv("ATS_DB_PATH")) {
	// 	log.Println("DB file does not exist creating it")
		// dbpath := os.Getenv("ATS_DB_PATH")
		// createAccountsDB(dbpath)
		// createCommentsDB(dbpath)
		// createEstimatesDB(dbpath)
	// }

	// uploadsPath := os.Getenv("ATS_UPLOADS_PATH")
	// createUploadsDir(uploadsPath)

	// certpath := os.Getenv("ATS_CERT_PATH")
	// createCertDir(certpath)

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

