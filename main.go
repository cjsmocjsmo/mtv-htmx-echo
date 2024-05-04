package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type Template struct {
	templates *template.Template
}

type MovieStruct struct {
	Name          string
	Year          string
	PosterAddr    string
	Size          string
	Path          string
	Idx           string
	MovId         string
	Catagory      string
	HttpThumbPath string
}

type TvEpiStruct struct {
	TvId     string
	Size     string
	Catagory string
	Name     string
	Season   string
	Episode  string
	Path     string
	Idx      string
}

type TVSeasonStruct struct {
	Cat string
	Sea string
	Epi []TvEpiStruct
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
	e.GET("/movcostner", mov_costner)
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
	e.GET("/Shogunsea", tv_action_shogun_seasons)
	e.GET("/TheContinentalsea", tv_action_continental_seasons)

	e.GET("/tvcomedy", tv_comedy)
	e.GET("/FuuBarsea", tv_comedy_fuubar_seasons)

	e.GET("/tvfantasy", tv_fantasy)
	e.GET("/WheelOfTimesea", tv_fantasy_wheeloftime_seasons)
	e.GET("/TheLordOfTheRingsTheRingsOfPowersea", tv_fantasy_TheLordOfTheRingsTheRingsOfPower_seasons)
	e.GET("/HouseOfTheDragonsea", tv_fantasy_houseofthedragon_seasons)

	e.GET("/tvstartrek", tv_startrek)
	e.GET("/strangenewworlds", tv_startrek_strangenewworlds_seasons)
	e.GET("/discovery", tv_startrek_discovery_seasons)
	e.GET("/picard", tv_startrek_picard_seasons)
	e.GET("/lowerdecks", tv_startrek_lowerdecks_seasons)
	e.GET("/prodigy", tv_startrek_prodigy_seasons)
	e.GET("/enterprise", tv_startrek_enterprise_seasons)
	e.GET("/voyager", tv_startrek_voyager_seasons)
	e.GET("/nextgeneration", tv_startrek_nextgeneration_seasons)
	e.GET("/sttv", tv_startrek_sttv_seasons)

	e.GET("/tvstarwars", tv_starwars)
	e.GET("/ahsoka", tv_starwars_ahsoka_seasons)
	e.GET("/andor", tv_starwars_andor_seasons)
	e.GET("/badbatch", tv_starwars_badbatch_seasons)
	e.GET("/bobafett", tv_starwars_bobafett_seasons)
	e.GET("/mandalorian", tv_starwars_mandalorian_seasons)
	e.GET("/obiwan", tv_starwars_obiwan_seasons)
	e.GET("/talesofthejedi", tv_starwars_talesofthejedi_seasons)
	e.GET("/visions", tv_starwars_visions_seasons)
	e.GET("/tvscifi", tv_scifi)
	e.GET("/threebodyproblem", tv_scifi_threebodyproblem_seasons)
	e.GET("/fallout", tv_scifi_fallout_seasons)
	e.GET("/alteredcarbon", tv_scifi_alteredcarbon_seasons)
	e.GET("/cowboybebop", tv_scifi_cowboybebop_seasons)
	e.GET("/forallmankind", tv_scifi_forallmankind_seasons)
	e.GET("/foundation", tv_scifi_foundation_seasons)
	e.GET("/halo", tv_scifi_halo_seasons)
	e.GET("/thelastofus", tv_scifi_thelastofus_seasons)
	e.GET("/lostinspace", tv_scifi_lostinspace_seasons)
	e.GET("/monarchlegacyofmonsters", tv_scifi_monarchlegacyofmonsters_seasons)
	e.GET("/nightsky", tv_scifi_nightsky_seasons)
	e.GET("/orville", tv_scifi_orville_seasons)
	e.GET("/raisedbywolves", tv_scifi_raisedbywolves_seasons)
	e.GET("/silo", tv_scifi_silo_seasons)
	e.GET("/tvscience", tv_science)
	e.GET("/prehistoricplanet", tv_science_prehistoricplanet_seasons)
	e.GET("/tvmcu", tv_mcu)
	e.GET("/falconwintersoldier", tv_mcu_falconwintersoldier_Seasons)
	e.GET("/iamgroot", tv_mcu_iamgroot_seasons)
	e.GET("/loki", tv_mcu_loki_Seasons)
	e.GET("/moonknight", tv_mcu_moonknight_seasons)
	
	e.GET("/shehulk", tv_mcu_shehulk_seasons)
	e.GET("/whatif", tv_mcu_whatif_seasons)
	e.GET("/wandavision", tv_mcu_wandavision_seasons)
	e.GET("/hawkeye", tv_mcu_hawkeye_seasons)
	e.GET("/secretInvasion", tv_mcu_secretinvasion_seasons)
	e.GET("/tvwestern", tv_western)
	e.GET("/hford1923", tv_western_1923_seasons)
	e.GET("/admin", mtv_admin)
	e.GET("/playmovie/:MovId", playmovie)
	e.Static("/assets", "assets")
	e.Logger.Fatal(e.Start(":8080"))

}

func playmovie(c echo.Context) error {
	MovId := c.Param("MovId")
	fmt.Printf("MovId: %s\n", MovId)
	return c.Render(http.StatusOK, "mov_play", MovId)
}

func (t *Template) Render(w io.Writer, Name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, Name, data)
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

func admin_stats() map[string]string {
	dbPath := os.Getenv("MTV_DB_PATH")
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Printf("failed to open database: %v", err)
    }
    defer db.Close()

    var mov_count int
    err = db.QueryRow("SELECT COUNT(*) FROM movies").Scan(&mov_count)
    if err != nil {
        log.Printf("failed to execute query: %v", err)
    }

	var tv_count int
	err = db.QueryRow("SELECT COUNT(*) FROM tvshows").Scan(&tv_count)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}

	var movTotalSize int64
	err = db.QueryRow("SELECT SUM(Size) FROM movies").Scan(&movTotalSize)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}

	var tvTotalSize int64
	err = db.QueryRow("SELECT SUM(Size) FROM tvshows").Scan(&tvTotalSize)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}

	movTotalSizeGB := float64(movTotalSize) / (1024 * 1024 * 1024)
	tvTotalSizeGB := float64(tvTotalSize) / (1024 * 1024 * 1024)

	movCountStr := strconv.Itoa(mov_count)
	tvCountStr := strconv.Itoa(tv_count)
	movTotalSizeGBStr := fmt.Sprintf("%.2f", movTotalSizeGB)
	tvTotalSizeGBStr := fmt.Sprintf("%.2f", tvTotalSizeGB)

    data := map[string]string{
		"Mov_count": movCountStr,
		"Tv_count": tvCountStr,
		"MovTotalSize": movTotalSizeGBStr,
		"TvTotalSize": tvTotalSizeGBStr,
	}

	return data
}

func mtv_admin(c echo.Context) error {
	data := admin_stats()
	return c.Render(http.StatusOK, "mtv_admin", data)
}

func MovInfo(cat string) []map[string]string {
	dbPath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT Name, Year, PosterAddr, Size, Path, Idx, MovId, Catagory, HttpThumbPath FROM movies WHERE Catagory = ? ORDER BY Year DESC", cat)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var movies []map[string]string
	for rows.Next() {
		var movie MovieStruct
		if err := rows.Scan(&movie.Name, &movie.Year, &movie.PosterAddr, &movie.Size, &movie.Path, &movie.Idx, &movie.MovId, &movie.Catagory, &movie.HttpThumbPath); err != nil {
			log.Printf("failed to scan row: %v", err)
		}
		var movinfo = map[string]string{
			"Name":          movie.Name,
			"Year":          movie.Year,
			"PosterAddr":    movie.PosterAddr,
			"Size":          movie.Size,
			"Path":          movie.Path,
			"Idx":           movie.Idx,
			"MovId":         movie.MovId,
			"Catagory":      movie.Catagory,
			"HttpThumbPath": movie.HttpThumbPath,
		}
		// log.Printf("movie: %v", movinfo)
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

func mov_costner(c echo.Context) error {
	movies := MovInfo("Costner")
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
	movies := MovInfo("TinkerBell")
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

func TVSeasonInfo2(cat string, sea string) TVSeasonStruct {
	dbPath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT TvId, Size, Catagory, Name, Season, Episode, Path, Idx FROM tvshows WHERE Catagory = ? AND Season = ? ORDER BY Episode ASC", cat, sea)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var sea1EpiInfo []TvEpiStruct
	for rows.Next() {
		var tv TvEpiStruct
		if err := rows.Scan(&tv.TvId, &tv.Size, &tv.Catagory, &tv.Name, &tv.Season, &tv.Episode, &tv.Path, &tv.Idx); err != nil {
			log.Printf("failed to scan row: %v", err)
		}
		sea1EpiInfo = append(sea1EpiInfo, tv)

	}
	

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
	}

	seaInfo := TVSeasonStruct{
		Cat: cat,
		Sea: sea,
		Epi: sea1EpiInfo,
	}
	// log.Printf("data: %v", seaInfo)

	return seaInfo
}

func TVSeasonInfo(cat string) []map[string]string {
	dbPath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT Season FROM tvshows WHERE Catagory = ? ORDER BY Season ASC", cat)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var result []map[string]string
	for rows.Next() {
		var season string
		if err := rows.Scan(&season); err != nil {
			log.Printf("failed to scan row: %v", err)
		}
		info := map[string]string{
			"Catagory": cat,
			"Season":   season,
		}
		result = append(result, info)
	}

	log.Printf("data: %v", result)

	if err := rows.Err(); err != nil {
		log.Printf("rows iteration error: %v", err)
	}
	return result
}

// func TVSeasonInfo2(cat string, sea string) []map[string]string {
// 	dbPath := os.Getenv("MTV_DB_PATH")
// 	db, err := sql.Open("sqlite3", dbPath)
// 	if err != nil {
// 		log.Printf("failed to open database: %v", err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT TvId, Size, Catagory, Name, Season, Episode, Path, Idx FROM tvshows WHERE Catagory = ? AND Season = ? ORDER BY Episode ASC", cat, sea)
// 	if err != nil {
// 		log.Printf("failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	var sea1EpiInfo []map[string]string
// 	for rows.Next() {
// 		var tv TvEpiStruct
// 		if err := rows.Scan(&tv.TvId, &tv.Size, &tv.Catagory, &tv.Name, &tv.Season, &tv.Episode, &tv.Path, &tv.Idx); err != nil {
// 			log.Printf("failed to scan row: %v", err)
// 		}
// 		epiInfo := map[string]string{
// 			"TvId":     tv.TvId,
// 			"Size":     tv.Size,
// 			"Catagory": tv.Catagory,
// 			"Name":     tv.Name,
// 			"Season":   tv.Season,
// 			"Episode":  tv.Episode,
// 			"Path":     tv.Path,
// 			"Idx":      tv.Idx,
// 		}
// 		sea1EpiInfo = append(sea1EpiInfo, epiInfo)

// 	}
// 	// log.Printf("data: %v", sea1EpiInfo)

// 	if err := rows.Err(); err != nil {
// 		log.Printf("rows iteration error: %v", err)
// 	}
// 	return sea1EpiInfo
// }

//////////////////////////////// ACTION TV SHOWS //////////////////////////////////////

func tv_action(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_action", "WORKED")
}
func tv_action_shogun_seasons(c echo.Context) error {
	result := TVSeasonInfo2("Shogun", "01")
	return c.Render(http.StatusOK, "tv_test", result)
}
func tv_action_continental_seasons(c echo.Context) error {
	result := TVSeasonInfo2("TheContinental", "01")
	return c.Render(http.StatusOK, "tv_test", result)
}

//////////////////////////////// COMEDY TV SHOWS //////////////////////////////////////

func tv_comedy(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_comedy", "WORKED")
}
func tv_comedy_fuubar_seasons(c echo.Context) error {
	result := TVSeasonInfo2("FuuBar", "01")
	return c.Render(http.StatusOK, "tv_test", result)
}

//////////////////////////////// FANTASY TV SHOWS //////////////////////////////////////

func tv_fantasy(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_fantasy", "WORKED")
}
func tv_fantasy_wheeloftime_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("WheelOfTime", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("WheelOfTime", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_fantasy_TheLordOfTheRingsTheRingsOfPower_seasons(c echo.Context) error {
	result := TVSeasonInfo2("RingsOfPower", "01")
	return c.Render(http.StatusOK, "tv_test", result)
}
func tv_fantasy_houseofthedragon_seasons(c echo.Context) error {
	result := TVSeasonInfo2("HouseOfTheDragon", "01")
	return c.Render(http.StatusOK, "tv_test", result)
}

func tv_startrek(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_startrek", "WORKED")
}
func tv_startrek_strangenewworlds_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("StrangeNewWorlds", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("StrangeNewWorlds", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_startrek_discovery_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Discovery", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Discovery", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("Discovery", "03")
	season_list = append(season_list, sea3)
	sea4 := TVSeasonInfo2("Discovery", "04")
	season_list = append(season_list, sea4)
	sea5 := TVSeasonInfo2("Discovery", "05")
	season_list = append(season_list, sea5)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_startrek_picard_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Picard", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Picard", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_startrek_lowerdecks_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("LowerDecks", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("LowerDecks", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("LowerDecks", "03")
	season_list = append(season_list, sea3)
	sea4 := TVSeasonInfo2("LowerDecks", "04")
	season_list = append(season_list, sea4)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_startrek_prodigy_seasons(c echo.Context) error {
	result := TVSeasonInfo2("Prodigy", "01")
	return c.Render(http.StatusOK, "tv_test", result)
}
func tv_startrek_enterprise_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Enterprise", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Enterprise", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("Enterprise", "03")
	season_list = append(season_list, sea3)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_startrek_voyager_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Voyager", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Voyager", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("Voyager", "03")
	season_list = append(season_list, sea3)
	sea4 := TVSeasonInfo2("Voyager", "04")
	season_list = append(season_list, sea4)
	sea5 := TVSeasonInfo2("Voyager", "05")
	season_list = append(season_list, sea5)
	sea6 := TVSeasonInfo2("Voyager", "06")
	season_list = append(season_list, sea6)
	sea7 := TVSeasonInfo2("Voyager", "07")
	season_list = append(season_list, sea7)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_startrek_nextgeneration_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("TNG", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("TNG", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("TNG", "03")
	season_list = append(season_list, sea3)
	sea4 := TVSeasonInfo2("TNG", "04")
	season_list = append(season_list, sea4)
	sea5 := TVSeasonInfo2("TNG", "05")
	season_list = append(season_list, sea5)
	sea6 := TVSeasonInfo2("TNG", "06")
	season_list = append(season_list, sea6)
	sea7 := TVSeasonInfo2("TNG", "07")
	season_list = append(season_list, sea7)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_startrek_sttv_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("STTV", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("STTV", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("STTV", "03")
	season_list = append(season_list, sea3)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}

func tv_starwars(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_starwars", "WORKED")
}
func tv_starwars_ahsoka_seasons(c echo.Context) error {
	result := TVSeasonInfo2("Ahsoka", "01")
	return c.Render(http.StatusOK, "tv_test", result)
}
func tv_starwars_andor_seasons(c echo.Context) error {
	result := TVSeasonInfo2("Andor", "01")
	return c.Render(http.StatusOK, "tv_test", result)
}
func tv_starwars_badbatch_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("BadBatch", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("BadBatch", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("BadBatch", "03")
	season_list = append(season_list, sea3)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_starwars_bobafett_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("BobaFett", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_starwars_mandalorian_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Mandalorian", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Mandalorian", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("Mandalorian", "03")
	season_list = append(season_list, sea3)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_starwars_obiwan_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("ObiWan", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_starwars_talesofthejedi_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("TalesOfTheJedi", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_starwars_visions_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Visions", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Visions", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}

func tv_scifi(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_scifi", "WORKED")
}
func tv_scifi_threebodyproblem_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("ThreeBodyProblem", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_scifi_fallout_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("Fallout", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_scifi_alteredcarbon_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("AlteredCarbon", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("AlteredCarbon", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_scifi_cowboybebop_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("CowboyBebop", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_scifi_forallmankind_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("ForAllMankind", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("ForAllMankind", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("ForAllMankind", "03")
	season_list = append(season_list, sea3)
	sea4 := TVSeasonInfo2("ForAllMankind", "04")
	season_list = append(season_list, sea4)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_scifi_foundation_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Foundation", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Foundation", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_scifi_halo_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Halo", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Halo", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_scifi_thelastofus_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("TheLastOfUs", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_scifi_lostinspace_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("LostInSpace", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("LostInSpace", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("LostInSpace", "03")
	season_list = append(season_list, sea3)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_scifi_monarchlegacyofmonsters_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("MonarchLegacyOfMonsters", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_scifi_nightsky_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("NightSky", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_scifi_orville_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Orville", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Orville", "02")
	season_list = append(season_list, sea2)
	sea3 := TVSeasonInfo2("Orville", "03")
	season_list = append(season_list, sea3)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_scifi_raisedbywolves_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("RaisedByWolves", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("RaisedByWolves", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_scifi_silo_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("Silo", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}

func tv_science(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_science", "WORKED")
}
func tv_science_prehistoricplanet_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("PrehistoricPlanet", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("PrehistoricPlanet", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}

func tv_mcu(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_mcu", "WORKED")
}
func tv_mcu_falconwintersoldier_Seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("FalconWinterSoldier", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_mcu_iamgroot_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("IAmGroot", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("IAmGroot", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_mcu_loki_Seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("Loki", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("Loki", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_mcu_moonknight_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("MoonKnight", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_mcu_shehulk_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("SheHulk", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_mcu_wandavision_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("WandaVision", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_mcu_whatif_seasons(c echo.Context) error {
	season_list := []TVSeasonStruct{}
	sea1 := TVSeasonInfo2("WhatIf", "01")
	season_list = append(season_list, sea1)
	sea2 := TVSeasonInfo2("WhatIf", "02")
	season_list = append(season_list, sea2)
	return c.Render(http.StatusOK, "tv_test2", season_list)
}
func tv_mcu_hawkeye_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("Hawkeye", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
func tv_mcu_secretinvasion_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("SecretInvasion", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}

func tv_western(c echo.Context) error {
	return c.Render(http.StatusOK, "tv_western", "WORKED")
}
func tv_western_1923_seasons(c echo.Context) error {
	sea1 := TVSeasonInfo2("HFord1923", "01")
	return c.Render(http.StatusOK, "tv_test", sea1)
}
