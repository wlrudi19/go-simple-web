package api

import (
	"github.com/go-chi/chi"
	"github.com/wlrudi19/go-simple-web/infrastructure/middlewares"
)

func NewProductRouter(productHandler ProductHandler) *chi.Mux {
	router := chi.NewRouter()

	authMiddleware := middlewares.Authenticate

	router.Route("/api/products", func(r chi.Router) {
		r.With(authMiddleware).Post("/create", productHandler.CreateProductHandler)
		r.Get("/findbyId", productHandler.FindProductHandler)
		r.Get("/findall", productHandler.FindProductAllHandler)
		r.With(authMiddleware).Put("/deletebyId", productHandler.DeleteProductHandler)
		r.With(authMiddleware).Put("/update/{product_id}", productHandler.UpdateProductHandler)

		//use
		r.With(authMiddleware).Post("/order", productHandler.OrderHandler)
		r.With(authMiddleware).Get("/get-order", productHandler.FindOrderConditionLogic)
		r.With(authMiddleware).Post("/bulk-update-order", productHandler.BulkUpdateOrder)
		r.With(authMiddleware).Get("/get-order-summary", productHandler.OrderSummaryLogic)
	})

	return router
}
