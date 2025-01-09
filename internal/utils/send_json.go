package utils

import (
	"encoding/json"
	"net/http"
	internalerrors "projetoEmail/internal/internal_errors"
)

func SendJSON(w http.ResponseWriter, err error, data interface{}) {

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		switch err.(type) {
		case internalerrors.ErrEntityNotFound:
			w.WriteHeader(http.StatusNotFound)
		case internalerrors.ErrCampaignNotPending:
			w.WriteHeader(http.StatusBadRequest)
		case internalerrors.ErrInternal:
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
}
