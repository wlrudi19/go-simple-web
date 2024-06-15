package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/wlrudi19/go-simple-web/app/product/model"
	"github.com/wlrudi19/go-simple-web/app/product/service"
	httputils "github.com/wlrudi19/go-simple-web/helper/http"
	"github.com/wlrudi19/go-simple-web/infrastructure/middlewares"
)

type ProductHandler interface {
	CreateProductHandler(writer http.ResponseWriter, req *http.Request)
	FindProductHandler(writer http.ResponseWriter, req *http.Request)
	//FindProductAllHandler(writer http.ResponseWriter, req *http.Request)
	DeleteProductHandler(writer http.ResponseWriter, req *http.Request)
	UpdateProductHandler(writer http.ResponseWriter, req *http.Request)

	//use
	FindProductAllHandler(writer http.ResponseWriter, req *http.Request)
	OrderHandler(writer http.ResponseWriter, req *http.Request)
	FindOrderConditionLogic(writer http.ResponseWriter, req *http.Request)
	BulkUpdateOrder(writer http.ResponseWriter, req *http.Request)
	OrderSummaryLogic(writer http.ResponseWriter, req *http.Request)
}

type producthandler struct {
	ProductLogic service.ProductLogic
}

func NewProductHandler(productLogic service.ProductLogic) ProductHandler {
	return &producthandler{
		ProductLogic: productLogic,
	}
}

func (h *producthandler) CreateProductHandler(writer http.ResponseWriter, req *http.Request) {
	var jsonReq model.CreateProductRequest

	err := json.NewDecoder(req.Body).Decode(&jsonReq)

	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Permintaan tidak valid. Format JSON tidak sesuai",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	err = h.ProductLogic.CreateProductLogic(req.Context(), jsonReq)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 201,
		Message:   "Product created successfully",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Errors: nil,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusCreated

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}

func (h *producthandler) OrderHandler(writer http.ResponseWriter, req *http.Request) {
	var jsonReq model.Order
	err := json.NewDecoder(req.Body).Decode(&jsonReq)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Permintaan tidak valid. Format JSON tidak sesuai",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	userId := req.Context().Value(middlewares.ContextKeyUserId).(int)
	jsonReq.UserID = userId

	err = h.ProductLogic.OrderLogic(req.Context(), jsonReq)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 201,
		Message:   "Order success",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Errors: nil,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusCreated

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}

func (h *producthandler) FindProductHandler(writer http.ResponseWriter, req *http.Request) {
	var jsonReq model.ProductRequest

	err := json.NewDecoder(req.Body).Decode(&jsonReq)

	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Permintaan tidak valid. Format JSON tidak sesuai",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	var product = model.FindProductResponse{}
	product, err = h.ProductLogic.FindProductLogic(req.Context(), jsonReq.Id)

	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			respon := []httputils.StandardError{
				{
					Code:   "404",
					Title:  "Not found",
					Detail: "Product not found",
					Object: httputils.ErrorObject{},
				},
			}
			httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
			return
		}

		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 200,
		Message:   "Product finding successfully",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Data:   &product,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusOK

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}

func (h *producthandler) FindOrderConditionLogic(writer http.ResponseWriter, req *http.Request) {
	var jsonReq model.Order
	userId := req.Context().Value(middlewares.ContextKeyUserId).(int)

	err := json.NewDecoder(req.Body).Decode(&jsonReq)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Permintaan tidak valid. Format JSON tidak sesuai",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	var orders []model.Order
	orders, err = h.ProductLogic.FindOrderConditionLogic(req.Context(), userId, jsonReq)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			respon := []httputils.StandardError{
				{
					Code:   "404",
					Title:  "Not found",
					Detail: "Product not found",
					Object: httputils.ErrorObject{},
				},
			}
			httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
			return
		}

		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 200,
		Message:   "Order finding successfully",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Data:   &orders,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusOK

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}

func (h *producthandler) FindProductAllHandler(writer http.ResponseWriter, req *http.Request) {
	if req.ContentLength != 0 {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Tidak boleh ada inputan di body JSON",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	var products []model.Product
	products, err := h.ProductLogic.FindProductAllLogic(req.Context())
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			respon := []httputils.StandardError{
				{
					Code:   "404",
					Title:  "Not found",
					Detail: "Product not found",
					Object: httputils.ErrorObject{},
				},
			}
			httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
			return
		}

		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 200,
		Message:   "Products finding successfully",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Data:   &products,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusOK

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}

func (h *producthandler) DeleteProductHandler(writer http.ResponseWriter, req *http.Request) {
	//set context
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	var jsonReq model.ProductRequest

	err := json.NewDecoder(req.Body).Decode(&jsonReq)

	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Permintaan tidak valid. Format JSON tidak sesuai",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	err = h.ProductLogic.DeleteProductLogic(req.Context(), jsonReq.Id)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			respon := []httputils.StandardError{
				{
					Code:   "404",
					Title:  "Not found",
					Detail: "Product not found",
					Object: httputils.ErrorObject{},
				},
			}
			httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
			return
		} else if strings.Contains(err.Error(), "product has been deleted before") {
			respon := []httputils.StandardError{
				{
					Code:   "400",
					Title:  "Bad Request",
					Detail: "Product telah dihapus sebelumnya",
					Object: httputils.ErrorObject{},
				},
			}
			httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
			return
		}

		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 200,
		Message:   "Product deleted successfully",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Errors: nil,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusOK

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}

func (h *producthandler) UpdateProductHandler(writer http.ResponseWriter, req *http.Request) {
	var jsonReq model.UpdateProductRequest

	err := json.NewDecoder(req.Body).Decode(&jsonReq)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Permintaan tidak valid. Format JSON tidak sesuai",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	productId, err := strconv.Atoi(chi.URLParam(req, "product_id"))
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Permintaan tidak valid. Format JSON tidak sesuai",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	err = h.ProductLogic.UpdateProductLogic(req.Context(), productId, jsonReq)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			respon := []httputils.StandardError{
				{
					Code:   "404",
					Title:  "Not found",
					Detail: "Product not found",
					Object: httputils.ErrorObject{},
				},
			}
			httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
			return
		}

		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 200,
		Message:   "Product updated successfully",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Errors: nil,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusOK

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}

func (h *producthandler) OrderSummaryLogic(writer http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(middlewares.ContextKeyUserId).(int)

	var orders model.OrderSummary
	orders, err := h.ProductLogic.OrderSummaryLogic(req.Context(), userId)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			respon := []httputils.StandardError{
				{
					Code:   "404",
					Title:  "Not found",
					Detail: "Product not found",
					Object: httputils.ErrorObject{},
				},
			}
			httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
			return
		}

		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 200,
		Message:   "Product finding successfully",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Data:   &orders,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusOK

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}

func (h *producthandler) BulkUpdateOrder(writer http.ResponseWriter, req *http.Request) {
	var jsonReq []model.BulkUpdateOrder

	err := json.NewDecoder(req.Body).Decode(&jsonReq)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "400",
				Title:  "Bad Request",
				Detail: "Permintaan tidak valid. Format JSON tidak sesuai",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusBadRequest, respon)
		return
	}

	err = h.ProductLogic.BulkUpdateOrderLogic(req.Context(), jsonReq)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			respon := []httputils.StandardError{
				{
					Code:   "404",
					Title:  "Not found",
					Detail: "Product not found",
					Object: httputils.ErrorObject{},
				},
			}
			httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
			return
		}

		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	status := httputils.StandardStatus{
		ErrorCode: 200,
		Message:   "Order update successfully",
	}

	envelope := httputils.StandardEnvelope{
		Status: &status,
		Errors: nil,
	}

	responFix, err := json.Marshal(envelope)
	if err != nil {
		respon := []httputils.StandardError{
			{
				Code:   "500",
				Title:  "Internal server error",
				Detail: "Terjadi kesalahan internal pada server",
				Object: httputils.ErrorObject{},
			},
		}
		httputils.WriteErrorResponse(writer, http.StatusInternalServerError, respon)
		return
	}

	contentType := httputils.NewContentTypeDecorator("application/json")
	httpStatus := http.StatusOK

	httputils.WriteResponse(writer, responFix, httpStatus, contentType)
}
