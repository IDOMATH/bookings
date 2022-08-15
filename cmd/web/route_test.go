package main

import (
	"fmt"
	"testing"

	"github.com/IDOMATH/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch valueType := mux.(type) {
	case *chi.Mux:
		// Do nothing
	default:
		t.Errorf(fmt.Sprintf("type is not *chi.Mux, but is %T", valueType))
	}
}
