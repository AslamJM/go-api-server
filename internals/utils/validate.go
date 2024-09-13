package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorValidateHandler(err error, w http.ResponseWriter) {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errs := make(map[string]string)

		for _, e := range validationErrs {
			errs[e.Field()] = e.Tag()
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)

	} else {
		log.Println("Non-validation error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
