package utils

import (
	"encoding/json"
	"net/http"
	internalerrors "projetoEmail/internal/internal_errors"
)

type Success struct {
	Data   interface{}
	Status int
}

func SendJSON(w http.ResponseWriter, err error, data *Success) {

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		if httpError, ok := err.(internalerrors.HttpErrorImp); ok {
			w.WriteHeader(httpError.GetStatus())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err := json.NewEncoder(w).Encode(map[string]string{"error": err.Error()}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		return
	}

	w.WriteHeader(data.Status)
	if err := json.NewEncoder(w).Encode(data.Data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
}
