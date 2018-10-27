package controller

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/h2non/filetype.v1"

	"github.com/gorilla/mux"
)

// ImageFileHandler returns image file
func ImageFileHandler(w http.ResponseWriter, r *http.Request) error {
	path := filepath.Join("/home/", mux.Vars(r)["filepath"])
	f, err := os.Stat(path)
	if os.IsNotExist(err) {
		w.WriteHeader(204)
		return nil
	}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(204)
		return nil
	}
	// check file sizef
	// get the size
	size := f.Size()
	if size > 4*1024*1024 {
		w.WriteHeader(400)
		return nil
	}
	if filetype.IsImage(buf) {
		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		w.Header().Set("Content-Type", "image/jpeg")
		io.Copy(w, file)
	}
	w.WriteHeader(400)
	return nil
}
