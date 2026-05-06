package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/pedrfelip/go-gateway/gateway/internal/config"
	"github.com/pedrfelip/go-gateway/gateway/internal/proxy"
)

func matchRoute(reqPath string, routePath string) (bool, string) {
	if routePath == "/" {
		return true, ""
	}

	if strings.HasSuffix(routePath, "/*") {
		base := strings.TrimSuffix(routePath, "/*")
		if strings.HasPrefix(reqPath, base) {
			remainder := strings.TrimPrefix(reqPath, base)
			return true, remainder
		}
		return false, ""
	}

	return reqPath == routePath, ""
}

func main() {
	configPath := os.Args[1]

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.LoadConfig(data)
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + strconv.Itoa(cfg.Server.Port)

	client := &http.Client{}

		handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		reqPath := path.Clean(req.URL.Path)

		for _, route := range cfg.Routes {
			matched, _ := matchRoute(reqPath, route.Path)
			if matched {
				targetPath := strings.TrimSuffix(route.Target, "/") + reqPath

				p := proxy.NewProxy(targetPath, client)
				p.ServeHTTP(w, req)
				return
			}
		}

		http.NotFound(w, req)
	})

	log.Printf("gateway listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
