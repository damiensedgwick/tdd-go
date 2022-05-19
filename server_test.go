package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		PlayerServer(res, req)

		got := res.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		res := httptest.NewRecorder()

		PlayerServer(res, req)

		got := res.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
