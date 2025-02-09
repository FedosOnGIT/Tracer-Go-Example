package handler

import (
	"github.com/FedosOnGIT/TracerLib/uploadBatch"
	"net/http"
)

const (
	RequestIDHeaderName = "X-Request-Id"
)

type TracerHandler struct {
	logger *uploadBatch.Logger
}

func New(logger *uploadBatch.Logger) *TracerHandler {
	return &TracerHandler{
		logger: logger,
	}
}

func (handler *TracerHandler) HandleWithWarn(w http.ResponseWriter, r *http.Request) {
	logger := handler.logger
	requestID := r.Header.Get(RequestIDHeaderName)
	if requestID != "" {
		logger = logger.WithRequestID(requestID)
	} else {
		handler.logger.Warn("No request ID")
	}

	logger = logger.WithTag("worker", "handler")

	logger.Warn("Warning")

	w.WriteHeader(http.StatusOK)
}

func (handler *TracerHandler) HandleWithWarnf(w http.ResponseWriter, r *http.Request) {
	logger := handler.logger
	requestID := r.Header.Get(RequestIDHeaderName)
	if requestID != "" {
		logger = logger.WithRequestID(requestID)
	} else {
		logger.Warn("No request ID")
	}

	logger = logger.WithTag("worker", "handler")

	queryParam := r.URL.Query().Get("param")
	if queryParam == "" {
		logger.Error("Missing query parameter: param")
		http.Error(w, "Missing query parameter: param", http.StatusBadRequest)
		return
	}

	logger.Warnf("Warning. Parameter: %s", queryParam)

	w.WriteHeader(http.StatusOK)
}

func (handler *TracerHandler) HandleWithError(w http.ResponseWriter, r *http.Request) {
	logger := handler.logger
	requestID := r.Header.Get(RequestIDHeaderName)
	if requestID != "" {
		logger = logger.WithRequestID(requestID)
	} else {
		logger.Warn("No request ID")
	}

	logger = logger.WithTag("worker", "handler")
	logger.Error("Error")

	w.WriteHeader(http.StatusInternalServerError)
}

func (handler *TracerHandler) HandleWithErrorf(w http.ResponseWriter, r *http.Request) {
	logger := handler.logger
	requestID := r.Header.Get(RequestIDHeaderName)
	if requestID != "" {
		logger = logger.WithRequestID(requestID)
	} else {
		handler.logger.Warn("No request ID")
	}

	logger = logger.WithTag("worker", "handler")

	queryParam := r.URL.Query().Get("param")
	if queryParam == "" {
		logger.Error("Missing query parameter: param")
		http.Error(w, "Missing query parameter: param", http.StatusBadRequest)
		return
	}

	logger.Errorf("Error. Parameter: %s", queryParam)

	w.WriteHeader(http.StatusInternalServerError)
}
