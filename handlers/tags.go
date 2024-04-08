package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jay-bhogayata/blogapi/database"
	"github.com/jay-bhogayata/blogapi/internal/helper"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Tag struct {
	Name string `json:"name"`
}

func (h *Handlers) GetAllTags(w http.ResponseWriter, r *http.Request) {
	Tags, err := h.query.GetAllTags(r.Context())
	if err != nil {
		h.logger.Error("error while fetching Tags: ", "error", err.Error())
		h.respondWithError(w, http.StatusInternalServerError, "error while fetching Tags")
		return
	}

	h.respondWithJSON(w, http.StatusOK, Tags)
}

func (h *Handlers) GetTagByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.ParseID(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid Tag id", "error", err.Error())
		h.respondWithError(w, http.StatusBadRequest, "invalid Tag id")
		return
	}

	Tag, err := h.query.GetTagById(r.Context(), id)
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.logger.Error("no Tag found with given id", "error", err.Error())
			h.respondWithJSON(w, http.StatusNotFound, &ErrorResponse{Message: "No Tag found with given id"})
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.respondWithJSON(w, http.StatusOK, Tag)
}

func (h *Handlers) CreateTag(w http.ResponseWriter, r *http.Request) {
	var Tag Tag

	err := helper.DecodeJSONBody(w, r, &Tag)
	if err != nil {
		h.logger.Error("error while decoding request body", "error", err.Error())
		h.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.query.CreateTag(r.Context(), Tag.Name)
	if err != nil {
		h.logger.Error("error while creating Tag", "error", err.Error())
		h.respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.respondWithJSON(w, http.StatusCreated, id)
}

func (h *Handlers) UpdateTag(w http.ResponseWriter, r *http.Request) {
	var Tag database.UpdateTagParams
	err := helper.DecodeJSONBody(w, r, &Tag)
	if err != nil {
		h.logger.Error("error while decoding request body", "error", err.Error())
		h.respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedTag, err := h.query.UpdateTag(r.Context(), Tag)
	if err != nil {
		h.logger.Error("error while updating Tag", "error", err.Error())
		h.respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.respondWithJSON(w, http.StatusOK, updatedTag)
}

func (h *Handlers) DeleteTag(w http.ResponseWriter, r *http.Request) {
	id, err := h.ParseID(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("invalid Tag id", "error", err.Error())
		h.respondWithError(w, http.StatusBadRequest, "invalid Tag id")
		return
	}

	type res struct {
		Message string `json:"message"`
	}

	if _, err := h.query.GetTagById(r.Context(), id); err != nil {
		if err.Error() == "no rows in result set" {
			h.logger.Error("no Tag found with given id", "error", err.Error())
			h.respondWithJSON(w, http.StatusNotFound, &ErrorResponse{Message: "No Tag found with given id"})
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.query.DeleteTag(r.Context(), id); err != nil {
		h.logger.Error("error while deleting Tag", "error", err.Error())
		h.respondWithError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	h.respondWithJSON(w, http.StatusOK, &res{Message: "Tag deleted successfully"})
}

func (h *Handlers) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJSON(w, code, &ErrorResponse{Message: message})
}

func (h *Handlers) respondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		h.logger.Error(err.Error())
		h.respondWithError(w, http.StatusInternalServerError, "error while marshalling response")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (h *Handlers) ParseID(idStr string) (int32, error) {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.logger.Error(err.Error())
		return 0, err
	}
	return int32(id), nil
}