package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jay-bhogayata/blogapi/database"
	"github.com/jay-bhogayata/blogapi/logger"
)

func (h *Handlers) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	query := database.New(h.DB)

	cate, err := query.GetAllCategories(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log.Error(err.Error())
		return
	}

	res, err := json.Marshal(cate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *Handlers) GetCategoryById(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Log.Error(err.Error())
		return
	}

	query := database.New(h.DB)

	cate, err := query.GetCategoryById(r.Context(), int32(id))
	if err != nil {
		if err.Error() == "no rows in result set" {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			res, _ := json.Marshal(&ErrorResponse{Message: "No category found with given id"})
			w.Write(res)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log.Error(err.Error())
		return
	}

	res, err := json.Marshal(cate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Log.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
