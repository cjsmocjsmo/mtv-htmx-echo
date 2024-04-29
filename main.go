package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io"
	"log"
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

type TvEpiStruct struct {
	tvid     string
	size     string
	catagory string
	name     string
	season   string
	episode  string
	path     string
	idx      string
}

type TVSeasonStruct struct {
	season string
	episodes []map[string]string
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
	e.GET("/movaction", mov_action)
	e.GET("/movarnold", mov_arnold)
	e.GET("/movbrucelee", mov_brucelee)
	e.GET("/movbrucewillis", mov_brucewillis)
	e.GET("/movbuzz", mov_buzz)
	e.GET("/movcartoons", mov_cartoons)
	e.GET("/movcharliebrown", mov_charliebrown)
	e.GET("/movchucknorris", mov_chucknorris)
	e.GET("/movcomedy", mov_comedy)
	e.GET("/movdocumentary", mov_documentary)
	e.GET("/movdrama", mov_drama)
	e.GET("/movfantasy", mov_fantasy)
	e.GET("/movgodzilla", mov_godzilla)
	e.GET("/movharrypotter", mov_harrypotter)
	e.GET("/movindianajones", mov_indianajones)
	e.GET("/movjamesbond", mov_jamesbond)
	e.GET("/movjohnwayne", mov_johnwayne)
	e.GET("/movjohnwick", mov_johnwick)
	e.GET("/movjurrassicpark", mov_jurrassicpark)
	e.GET("/movkingsman", mov_kingsman)
	e.GET("/movmeninblack", mov_meninblack)
	e.GET("/movminions", mov_minions)
	e.GET("/movmisc", mov_misc)
	e.GET("/movnicolascage", mov_nicolascage)
	e.GET("/movoldies", mov_oldies)
	e.GET("/movpirates", mov_pirates)
	e.GET("/movriddick", mov_riddick)
	e.GET("/movscifi", mov_scifi)
	e.GET("/movstalone", mov_stalone)
	e.GET("/movstartrek", mov_startrek)
	e.GET("/movstarwars", mov_starwars)
	e.GET("/movsuperheros", mov_superheros)
	e.GET("/movtinkerbell", mov_tinkerbell)
	e.GET("/movtomcruize", mov_tomcruize)
	e.GET("/movtransformers", mov_transformers)
	e.GET("/movtremors", mov_tremors)
	e.GET("/movtherock", mov_therock)
	e.GET("/movxmen", mov_xmen)
	e.GET("/tvshows", mtv_tvshows)
	e.GET("/tvaction", tv_action)
	e.GET("/shogun", tv_action_shogun_seasons)
	e.GET("/continental", tv_action_continental_seasons)
	e.GET("/tvcomedy", tv_comedy)
	e.GET("/fuubar", tv_comedy_fuubar_seasons)
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
	return c.Render(http.StatusOK, "mov_play", movid)
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

func MovInfo(cat string) []map[string]string {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, year, posteraddr, size, path, idx, movid, catagory, httpthumbpath FROM movies WHERE catagory = ? ORDER BY year DESC", cat)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.name, &movie.year, &movie.posteraddr, &movie.size, &movie.path, &movie.idx, &movie.movid, &movie.catagory, &movie.httpthumbpath); err != nil {
			log.Printf("failed to scan row: %v", err)
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
	}
	return movies
}

func mov_action(c echo.Context) error {
	movies := MovInfo("Action")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_arnold(c echo.Context) error {
	movies := MovInfo("Arnold")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_brucelee(c echo.Context) error {
	movies := MovInfo("BruceLee")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_brucewillis(c echo.Context) error {
	movies := MovInfo("BruceWillis")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_buzz(c echo.Context) error {
	movies := MovInfo("Buzz")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_cartoons(c echo.Context) error {
	movies := MovInfo("Cartoons")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_charliebrown(c echo.Context) error {
	movies := MovInfo("CharlieBrown")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_chucknorris(c echo.Context) error {
	movies := MovInfo("ChuckNorris")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_comedy(c echo.Context) error {
	movies := MovInfo("Comedy")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_documentary(c echo.Context) error {
	movies := MovInfo("Documentary")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_drama(c echo.Context) error {
	movies := MovInfo("Drama")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_fantasy(c echo.Context) error {
	movies := MovInfo("Fantasy")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_godzilla(c echo.Context) error {
	movies := MovInfo("Godzilla")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_harrypotter(c echo.Context) error {
	movies := MovInfo("HarryPotter")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_indianajones(c echo.Context) error {
	movies := MovInfo("IndianaJones")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_jamesbond(c echo.Context) error {
	movies := MovInfo("JamesBond")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_johnwayne(c echo.Context) error {
	movies := MovInfo("JohnWayne")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_johnwick(c echo.Context) error {
	movies := MovInfo("JohnWick")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_jurrassicpark(c echo.Context) error {
	movies := MovInfo("JurrassicPark")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_kingsman(c echo.Context) error {
	movies := MovInfo("Kingsman")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_meninblack(c echo.Context) error {
	movies := MovInfo("MenInBlack")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_minions(c echo.Context) error {
	movies := MovInfo("Minions")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_misc(c echo.Context) error {
	movies := MovInfo("Misc")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_nicolascage(c echo.Context) error {
	movies := MovInfo("NicolasCage")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_oldies(c echo.Context) error {
	movies := MovInfo("Oldies")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_pirates(c echo.Context) error {
	movies := MovInfo("Pirates")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_riddick(c echo.Context) error {
	movies := MovInfo("Riddick")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_scifi(c echo.Context) error {
	movies := MovInfo("SciFi")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_stalone(c echo.Context) error {
	movies := MovInfo("Stalone")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_startrek(c echo.Context) error {
	movies := MovInfo("StarTrek")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_starwars(c echo.Context) error {
	movies := MovInfo("StarWars")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_superheros(c echo.Context) error {
	movies := MovInfo("SuperHeros")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_tinkerbell(c echo.Context) error {
	movies := MovInfo("Tinkerbell")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_tomcruize(c echo.Context) error {
	movies := MovInfo("TomCruize")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_transformers(c echo.Context) error {
	movies := MovInfo("Transformers")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_tremors(c echo.Context) error {
	movies := MovInfo("Tremors")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_therock(c echo.Context) error {
	movies := MovInfo("TheRock")
	return c.Render(http.StatusOK, "mov_movie", movies)
}

func mov_xmen(c echo.Context) error {
	movies := MovInfo("XMen")
	return c.Render(http.StatusOK, "mov_movie", movies)
}





func TVInfo(cat string, sea string) TVSeasonStruct {
	dbpath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT tvid, size, catagory, name, season, episode, path, idx FROM tvshows WHERE catagory = ? AND season = ?", cat, sea)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var sea1EpiInfo []map[string]string
	for rows.Next() {
		var tv TvEpiStruct
		if err := rows.Scan(&tv.tvid, &tv.size, &tv.catagory, &tv.name, &tv.season, &tv.episode, &tv.path, &tv.idx); err != nil {
			log.Printf("failed to scan row: %v", err)
		}
		epiInfo := map[string]string{
			"tvid"   : tv.tvid,
			"size"   : tv.size,
			"catagory" : tv.catagory,
			"name"   : tv.name,
			"season" : tv.season,
			"episode": tv.episode,
			"path"   : tv.path,
			"idx"    : tv.idx,
		}
		// log.Printf("epiInfo: %v", epiInfo)
		sea1EpiInfo = append(sea1EpiInfo, epiInfo)
	}

	data := TVSeasonStruct{
		season: "01",
		episodes: sea1EpiInfo,
	}
	log.Printf("data: %v", data)

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
	}
	return data
}


func tv_action(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_action", "WORKED")
}
func tv_action_shogun_seasons(c echo.Context) error {
	var data []TVSeasonStruct
	sea1 := TVInfo("Shogun", "01")
	data = append(data, sea1)
	return c.Render(http.StatusOK, "tvshowsseasons", data)
}
func tv_action_continental_seasons(c echo.Context) error {
	var data []TVSeasonStruct
	sea1 := TVInfo("TheContinental", "01")
	data = append(data, sea1)
	return c.Render(http.StatusOK, "tv_action_contenental", data)
}


func tv_comedy(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_comedy", "WORKED")
}
func tv_comedy_fuubar_seasons(c echo.Context) error {
	var data []TVSeasonStruct
	sea1 := TVInfo("FuuBar", "01")
	data = append(data, sea1)
	return c.Render(http.StatusOK, "tvshowsseasons", data)
}


func tv_fantasy(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_fantasy", "WORKED")
}
func tv_fantasy_wheeloftime_seasons(c echo.Context) error {
	var data []TVSeasonStruct
	sea1 := TVInfo("WheelOfTime", "01")
	data = append(data, sea1)
	sea2 := TVInfo("WheelOfTime", "02")
	data = append(data, sea2)
	return c.Render(http.StatusOK, "tvshowsseasons", data)
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

