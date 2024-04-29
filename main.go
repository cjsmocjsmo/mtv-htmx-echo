package main

import (
	// "database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io"
	// "log"
	"net/http"
	"os"
)

type Template struct {
	templates *template.Template
}

type MovieStruct struct {
	name          string
	year          string
	posteraddr    string
	size          string
	path          string
	idx           string
	movid         string
	catagory      string
	httpthumbpath string
}

type TvShowStruct struct {
	tvid     string
	size     string
	catagory string
	name     string
	season   string
	episode  string
	path     string
	idx      string
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
	checkDBExists()
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
	e.GET("/movaction", Mov_action)
	e.GET("/movarnold", Mov_arnold)
	e.GET("/movbrucelee", Mov_brucelee)
	e.GET("/movbrucewillis", Mov_brucewillis)
	e.GET("/movbuzz", Mov_buzz)
	e.GET("/movcartoons", Mov_cartoons)
	e.GET("/movcharliebrown", Mov_charliebrown)
	e.GET("/movchucknorris", Mov_chucknorris)
	e.GET("/movcomedy", Mov_comedy)
	e.GET("/movdocumentary", Mov_documentary)
	e.GET("/movdrama", Mov_drama)
	e.GET("/movfantasy", Mov_fantasy)
	e.GET("/movgodzilla", Mov_godzilla)
	e.GET("/movharrypotter", Mov_harrypotter)
	e.GET("/movindianajones", Mov_indianajones)
	e.GET("/movjamesbond", Mov_jamesbond)
	e.GET("/movjohnwayne", Mov_johnwayne)
	e.GET("/movjohnwick", Mov_johnwick)
	e.GET("/movjurrassicpark", Mov_jurrassicpark)
	e.GET("/movkingsman", Mov_kingsman)
	e.GET("/movmeninblack", Mov_meninblack)
	e.GET("/movminions", Mov_minions)
	e.GET("/movmisc", Mov_misc)
	e.GET("/movnicolascage", Mov_nicolascage)
	e.GET("/movoldies", Mov_oldies)
	e.GET("/movpirates", Mov_pirates)
	e.GET("/movriddick", Mov_riddick)
	e.GET("/movscifi", Mov_scifi)
	e.GET("/movstalone", Mov_stalone)
	e.GET("/movstartrek", Mov_startrek)
	e.GET("/movstarwars", Mov_starwars)
	e.GET("/movsuperheros", Mov_superheros)
	e.GET("/movtinkerbell", Mov_tinkerbell)
	e.GET("/movtomcruize", Mov_tomcruize)
	e.GET("/movtransformers", Mov_transformers)
	e.GET("/movtremors", Mov_tremors)
	e.GET("/movtherock", Mov_therock)
	e.GET("/movxmen", Mov_xmen)
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
	e.GET("/playmovie/:movid", playmovie)
	e.Static("/assets", "assets")
	e.Logger.Fatal(e.Start(":8080"))

}

func playmovie(c echo.Context) error {
	movid := c.Param("movid")
	fmt.Printf("movid: %s\n", movid)
	return c.Render(http.StatusOK, "Mov_play", movid)
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

// func Mov_action(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Action")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_arnold(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Arnold")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_brucelee(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "BruceLee")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_brucewillis(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "BruceWillis")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_buzz(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Buzz")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_cartoons(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Cartoons")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_charliebrown(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "CharlieBrown")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_chucknorris(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "ChuckNorris")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_comedy(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Comedy")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_documentary(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Documentary")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_drama(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Drama")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_fantasy(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Fantasy")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_godzilla(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Godzilla")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_harrypotter(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "HarryPotter")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_indianajones(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "IndianaJones")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_jamesbond(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "JamesBond")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_johnwayne(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "JohnWayne")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_johnwick(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "JohnWick")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_jurrassicpark(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "JurrassicPark")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_kingsman(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "KingsMan")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_meninblack(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "MenInBlack")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_minions(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Minions")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_misc(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Misc")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_nicolascage(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "NicolasCage")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_oldies(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Oldies")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_pirates(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Pirates")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_riddick(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Riddick")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_scifi(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "SciFi")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_stalone(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Stalone")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_startrek(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "StarTrek")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_starwars(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "StarWars")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_superheros(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "SuperHeros")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_tinkerbell(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "TinkerBell")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_tomcruize(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "TomCruize")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_transformers(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Transformers")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_tremors(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Tremors")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_therock(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "TheRock")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }

// func Mov_xmen(c echo.Context) error {
// 	dbpath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbpath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 		return fmt.Errorf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "XMen")
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 		return fmt.Errorf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var movies []map[string]string
// 	for rows.Next() {
// 		var movie MovieStruct
// 		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 			return fmt.Errorf("failed to scan row: %v", err)
// 		}
// 		var movinfo = map[string]string{
// 			"name":          movie.name,
// 			"year":          movie.year,
// 			"posteraddr":    movie.posteraddr,
// 			"size":          movie.size,
// 			"path":          movie.path,
// 			"idx":           movie.idx,
// 			"movid":         movie.movid,
// 			"catagory":      movie.catagory,
// 			"httpthumbpath": movie.httpthumbpath,
// 		}
// 		log.Printf("movie: %v", movinfo)
// 		movies = append(movies, movinfo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 		return fmt.Errorf("rows iteration error: %v", err)
// 	}

// 	return c.Render(http.StatusOK, "Mov_movie", movies)
// }









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

func tv_shogun_seasons(c echo.Context) error {
	//get season info from db
	return c.Render(http.StatusOK, "tvshowsseasons", "WORKED")
}