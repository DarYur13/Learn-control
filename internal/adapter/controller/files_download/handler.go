package files_download

import (
	"database/sql"
	"errors"
	"io"
	"net/http"

	"github.com/DarYur13/learn-control/internal/logger"
	"github.com/DarYur13/learn-control/internal/service"
	"github.com/google/uuid"
)

type Handler struct {
	svc service.Servicer
}

func New(svc service.Servicer) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.URL.Query().Get("token")
	if tokenStr == "" {
		http.Error(w, "token is required", http.StatusBadRequest)
		return
	}

	token, err := uuid.Parse(tokenStr)
	if err != nil {
		http.Error(w, "invalid token format", http.StatusBadRequest)
		return
	}

	file, err := h.svc.GetFileByToken(r.Context(), token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "file not found or token expired", http.StatusNotFound)
			return
		}

		logger.Errorf("failed to get file: %s", err)
		http.Error(w, "failed to load file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Disposition", `attachment; filename="лист_регистрации_инструктажа.docx"`)
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, file)
	if err != nil {
		logger.Errorf("failed to copy file to response writer: %s", err)
		http.Error(w, "failed to load file", http.StatusInternalServerError)
	}
}
