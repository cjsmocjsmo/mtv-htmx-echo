package main 

import (
	"database/sql"
	"fmt"
	// "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	// "html/template"
	// "io"
	"log"
	"net/http"
	"os"
)

func mov_action(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Action")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_arnold(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Arnold")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_brucelee(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "BruceLee")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_brucewillis(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "BruceWillis")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_buzz(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Buzz")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_cartoons(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Cartoons")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_charliebrown(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "CharlieBrown")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_chucknorris(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "ChuckNorris")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_comedy(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Comedy")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_documentary(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Documentary")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_drama(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Drama")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_fantasy(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Fantasy")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_godzilla(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Godzilla")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_harrypotter(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "HarryPotter")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_indianajones(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "IndianaJones")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_jamesbond(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "JamesBond")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_johnwayne(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "JohnWayne")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_johnwick(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "JohnWick")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_jurrassicpark(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "JurrassicPark")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_kingsman(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "KingsMan")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_meninblack(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "MenInBlack")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_minions(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Minions")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_misc(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Misc")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_nicolascage(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "NicolasCage")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_oldies(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Oldies")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_pirates(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Pirates")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_riddick(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Riddick")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_scifi(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "SciFi")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_stalone(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Stalone")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_startrek(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "StarTrek")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_starwars(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "StarWars")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}





func mov_superheros(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "SuperHeros")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_tinkerbell(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "TinkerBell")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_tomcruize(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "TomCruize")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_transformers(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Transformers")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_tremors(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "Tremors")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_therock(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "TheRock")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_xmen(c echo.Context) error {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ?", "XMen")
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
			return fmt.Errorf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"name":          movie.name,
			"year":          movie.year,
			"posteraddr":    movie.posteraddr,
			"size":          movie.size,
			"path":          movie.path,
			"idx":           movie.idx,
			"movid":         movie.movid,
			"catagory":      movie.catagory,
			"httpthumbpath": movie.httpthumbpath,
		}
		log.Printf("movie: %v", movinfo)
		movies = append(movies, movinfo)
	}

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
		return fmt.Errorf("rows iteration error: %v", err)
	}

	return c.Render(http.StatusOK, "mov_movie", movies)
}