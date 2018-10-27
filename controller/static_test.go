package controller

import (
	"net/http"
	"testing"
)

func TestImageFileHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ImageFileHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("ImageFileHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
