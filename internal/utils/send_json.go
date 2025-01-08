package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	internalerrors "projetoEmail/internal/internal_errors"
)

func SendJSON(w http.ResponseWriter, err error, data interface{}) {

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}

		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
}
