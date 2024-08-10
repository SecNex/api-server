package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	res "github.com/secnex/server/api/res"
	"github.com/secnex/server/api/routes"
	"github.com/secnex/server/db"
)

type ApiConfiguration struct {
	Address string
	Port    int
}

type Api struct {
	Name           string
	Address        string
	Port           int
	TrustedProxies []TrustedProxy
	Database       *db.Connection
}

type ApiRoute struct {
	Path    string
	Handler http.HandlerFunc
}

type ApiRoutes []ApiRoute

type TrustedProxy string

func NewTrustedProxies() []TrustedProxy {
	return []TrustedProxy{}
}

func (a *Api) AddTrustedProxy(proxy string) {
	a.TrustedProxies = append(a.TrustedProxies, TrustedProxy(proxy))
}

func NewApiRouteCollection() ApiRoutes {
	return []ApiRoute{}
}

func (a *ApiRoutes) AddRoute(path string, handler http.HandlerFunc) {
	*a = append(*a, ApiRoute{Path: path, Handler: handler})
}

func NewAPI(address string, port int, proxies []TrustedProxy, cnx *db.Connection, name string) *Api {
	return &Api{
		Address:        address,
		Port:           port,
		TrustedProxies: proxies,
		Database:       cnx,
		Name:           name,
	}
}

func CheckTrustedProxies(next http.Handler, proxies []TrustedProxy) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyCount := len(proxies)
		clientIP := extractIP(r)
		trusted := false
		if proxyCount > 0 {
			for _, proxy := range proxies {
				if string(proxy) == clientIP {
					trusted = true
					break
				}
			}
		} else {
			trusted = true
		}
		if !trusted {
			log.Printf("Unauthorized access from %v", r.RemoteAddr)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			result := res.ResultError{
				Code:    http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "Unauthorized access",
			}
			w.Write([]byte(result.String()))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *Api) Start(apiRoutes ApiRoutes) {
	r := http.NewServeMux()
	myRouter := routes.NewRouter(a.Database)
	r.HandleFunc("/health", myRouter.RouteHealth)
	for _, route := range apiRoutes {
		r.HandleFunc(route.Path, route.Handler)
	}
	router := CheckTrustedProxies(r, a.TrustedProxies)
	log.Printf("Starting server at %v:%v", a.Address, a.Port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", a.Address, a.Port), router)
}

func extractIP(r *http.Request) string {
	clientIP := r.RemoteAddr
	if strings.Contains(clientIP, "[") || strings.Contains(clientIP, "]") {
		clientIP = strings.Split(clientIP, "]")[0]
		clientIP = strings.Split(clientIP, "[")[1]
	} else {
		clientIP = strings.Split(clientIP, ":")[0]
	}
	return clientIP
}
