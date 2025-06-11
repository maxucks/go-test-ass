package com

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Details    any    `json:"details"`
	AppCode    int    `json:"code"`
}

func (e HttpError) Error() string {
	return e.Message
}

func New(status int, message string, code int, extensions ...ExtensionFunc) HttpError {
	err := HttpError{StatusCode: status, Message: message, AppCode: code}
	for _, ext := range extensions {
		ext(&err)
	}
	return err
}

func (e HttpError) apply(extension ExtensionFunc) { extension(&e) }

type ExtensionFunc func(*HttpError)

func WithDetails(details any) ExtensionFunc {
	return func(err *HttpError) {
		err.Details = details
	}
}

func JSON[T any](w http.ResponseWriter, value T) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		return ToHttpError(err)
	}
	return nil
}

func Empty(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

func internal(extensions ...ExtensionFunc) HttpError {
	return New(404, "errors.common.internal", 1, extensions...)
}

func badRequest(extensions ...ExtensionFunc) HttpError {
	return New(400, "errors.common.badRequest", 2, extensions...)
}

func notFound(extensions ...ExtensionFunc) HttpError {
	return New(404, "errors.common.notFound", 3, extensions...)
}

func Internal(w http.ResponseWriter, extensions ...ExtensionFunc) {
	Error(w, internal(extensions...))
}

func BadRequest(w http.ResponseWriter, extensions ...ExtensionFunc) {
	Error(w, badRequest(extensions...))
}

func NotFound(w http.ResponseWriter, extensions ...ExtensionFunc) {
	Error(w, notFound(extensions...))
}

func ToHttpError(err error) HttpError {
	msg := err.Error()

	switch err {
	case io.EOF:
		msg = fmt.Sprintf("EOF reading HTTP request body: %v", msg)
		return badRequest(WithDetails(msg))
	case sql.ErrNoRows:
		msg = fmt.Sprintf("not found: %v", msg)
		return notFound(WithDetails(msg))
	}

	msg = fmt.Sprintf("internal server error: %v", msg)
	return internal(WithDetails(msg))
}

func Error(w http.ResponseWriter, err error) {
	h := w.Header()

	h.Del("Content-Length")
	h.Set("Content-Type", "application/json")
	h.Set("X-Content-Type-Options", "nosniff")

	var httpErr HttpError

	switch err := err.(type) {
	case HttpError:
		httpErr = err
	default:
		httpErr = ToHttpError(err)
	}

	w.WriteHeader(httpErr.StatusCode)
	JSON(w, httpErr)
}
