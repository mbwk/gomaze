package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mbwk/gomaze/algorithms"
	"github.com/mbwk/gomaze/maze"
	"github.com/mbwk/gomaze/search"
)

var m *maze.Maze

func GetOrDefaultInt(query url.Values, key string, defaultValue int) int {
	if query.Has(key) {
		intStr := query.Get(key)
		i, err := strconv.ParseInt(intStr, 10, 64)
		if err == nil {
			return int(i)
		}
	}
	return defaultValue
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	width, height := GetOrDefaultInt(query, "width", 40), GetOrDefaultInt(query, "height", 20)
	seed := int64(GetOrDefaultInt(query, "seed", 0))
	algorithm := query.Get("algorithm")

	start := strings.Split(query.Get("start"), ",")
	end := strings.Split(query.Get("end"), ",")
	startX, startY, endX, endY := width/3, 1, width-(width/3), 1
	parseThrough := func(numstr string, initial int) int {
		converted, err := strconv.ParseInt(numstr, 10, 64)
		if err != nil {
			return initial
		}
		return int(converted)
	}
	if len(start) >= 2 {
		startX = parseThrough(start[0], startX)
		startY = parseThrough(start[1], startY)
	}
	if len(end) >= 2 {
		endX = parseThrough(end[0], endX)
		endY = parseThrough(end[1], endY)
	}

	m = maze.NewMaze(maze.Dimensions{X: width, Y: height})
	m.InitializeGrid()
	mazeSeed := algorithms.GenerateMaze(m, algorithm, seed)
	search.FindPath(m, maze.Coords{startX, startY}, maze.Coords{endX, endY})
	io.WriteString(w, maze.PrintGrid(&m.G))
	io.WriteString(w, fmt.Sprintln("\nMaze seed:", mazeSeed))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", indexHandler)
	s := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		// TLSConfig:         &tls.Config{},
		// ReadHeaderTimeout: 0,
		// IdleTimeout:       0,
		// MaxHeaderBytes:    0,
		// TLSNextProto:      map[string]func( *http.Server,  *tls.Conn,  http.Handler){},
		// ConnState: func( net.Conn,  http.ConnState) {
		// },
		// ErrorLog: &log.Logger{},
		// BaseContext: func( net.Listener) context.Context {
		// },
		// ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		// },
	}
	fmt.Println("starting server")
	s.ListenAndServe()
}
