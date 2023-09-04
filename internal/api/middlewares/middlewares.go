// error_handling_middleware.go
package middleware

import (
	"net/http"
)

func WrapHandlerWithErr(handlerFunc func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if err := handlerFunc(w, r); err != nil {
            http.Error(w, "Erro interno", http.StatusInternalServerError)
            // Ou lide com o erro de acordo com suas necessidades
        }
    }
}
