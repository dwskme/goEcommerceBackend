package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dwskme/goEcommerceBackend/service/cart"
	"github.com/dwskme/goEcommerceBackend/service/order"
	"github.com/dwskme/goEcommerceBackend/service/product"
	"github.com/dwskme/goEcommerceBackend/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	prodcutHandler := product.NewHandler(productStore, userStore)
	prodcutHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Listening on:", s.addr)
	return http.ListenAndServe(s.addr, router)
}
