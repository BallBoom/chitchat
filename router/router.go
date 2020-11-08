package router

import "github.com/gorilla/mux"

// 返回一个 mux.Router 类型指针，从而可以当作处理器使用
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range webRoutes {
		router.Methods(route.Methods).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}
	return router
}
