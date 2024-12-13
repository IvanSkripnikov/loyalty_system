package helpers

import (
	"fmt"
	"net/http"

	"loyalty_system/logger"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	writeSuccess(w, "{\"status\": \"OK\"}")
}

// -------------PRIVATE----------------------

func writeSuccess(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprint(w, message)
	if err != nil {
		logger.Error(fmt.Sprintf("write success error %s", err.Error()))

		return
	}
}
