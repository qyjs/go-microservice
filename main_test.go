package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

func Test_main(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name   string
		result string
		method string
		path   string
	}{
		{name: "health", method: "GET", path: "/health", result: "healthy"},
		{name: "base64", method: "GET", path: "/qrcode/base64?content=hello", result: "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAMAAABrrFhUAAAABlBMVEX///8AAABVwtN+AAABp0lEQVR42uzdwY6rMAxAUeb/f3r2lSpix6kJOXf3Kp5azsoiRnNJkiRJkiRJkiRJkiRJkiRJ0l+yXb4PAAAAAAAAAHD/g2av//z82w1UfR8AAAAAAAAAoB7gbhAZ/Xf0/1eBAQAAAAAAAAD6AVaBAQAAAAAAAAAAAAAAAAAAAAD6D0ZGD0peezIEAAAAAAAAvAAgu7hYNfhsuykKAAAAAAAAbAyw6oHKdUoAAAAAAADAxgDZQWV0sMoOXgAAAAAAAACA3wFEFxSqb7B9QQIAAAAAAAA4ECD6A6oWKqILkwAAAAAAAACAvnODKETVIAQAAAAAAAAA6Dsgqbqx0c+z1wEAAAAAAAAA+gel7MFK1YuWAAAAAAAAAID5G5gdeGYh2hYjAAAAAAAAgAMBqh6cRB90tL84AQAAAAAAAAAo+0NJVQ88AAAAAAAAAAD7A4wOOo9bkAQAAAAAAAAATL9IGT34aHtjFAAAAAAAAACQvj46CI2CAwAAAAAAAADWA2QXJaOD0GMPRgAAAAAAAIADASRJkiRJkiRJkiRJkiRJkiRt038AAAD//3IXOAH1bEeSAAAAAElFTkSuQmCC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, tt.result, w.Body.String())
		})
	}
}
