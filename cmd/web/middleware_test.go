package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch valueType := h.(type) {
	case http.Handler:
		// Do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", valueType))
	}
}
