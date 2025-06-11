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

func Internal(extensions ...ExtensionFunc) HttpError {
	return New(404, "errors.common.internal", 1)
}

func BadRequest(extensions ...ExtensionFunc) HttpError {
	return New(400, "errors.common.badRequest", 2)
}

func NotFound(extensions ...ExtensionFunc) HttpError {
	return New(404, "errors.common.notFound", 3)
}

func JSON[T any](w http.ResponseWriter, value T) error {
	if err := json.NewEncoder(w).Encode(value); err != nil {
		return ToHttpError(err)
	}
	return nil
}

func Empty(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

func ToHttpError(err error) HttpError {
	msg := err.Error()

	switch err {
	case io.EOF:
		msg = fmt.Sprintf("EOF reading HTTP request body: %v", msg)
		return BadRequest(WithDetails(msg))
	case sql.ErrNoRows:
		msg = fmt.Sprintf("not found: %v", msg)
		return NotFound(WithDetails(msg))
	}

	msg = fmt.Sprintf("internal server error: %v", msg)
	return Internal(WithDetails(msg))
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
