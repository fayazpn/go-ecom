package product

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fayazpn/ecom/types"
	"github.com/fayazpn/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/product", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetAllProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// process the payload
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	log.Print(payload)
	// validator
	if err := utils.Validate.Struct(payload); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorMessages := make([]string, len(validationErrors))
			for i, e := range validationErrors {
				errorMessages[i] = fmt.Sprintf("Field '%s' failed validation: %s %v", e.Field(), e.Tag(), e.Value())
			}
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", strings.Join(errorMessages, "; ")))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("unexpected error during validation"))
		}
		return
	}

	// create the product
	err := h.store.CreateProduct(types.CreateProductPayload{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
