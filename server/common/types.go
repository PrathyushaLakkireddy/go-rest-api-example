package common

import "net/http"

// Route structure
type Route struct {
	Method  string
	Path    string
	Name    string
	Handler http.HandlerFunc
}

// Error structre
type Error struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// Response structure
type Response struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message,omitempty"`
	User    interface{} `json:"user,omitempty"`
}
