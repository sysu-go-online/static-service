package controller

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/h2non/filetype.v1"

	"github.com/gorilla/mux"
	"github.com/sysu-go-online/public-service/tools"
)

// ImageFileHandler returns image file
func ImageFileHandler(w http.ResponseWriter, r *http.Request) error {
	// parse token
	token := r.Header.Get("Authorization")
	if valid, err := tools.CheckJWT(token, AuthRedisClient); !(err == nil && valid) {
		w.WriteHeader(401)
		return nil
	}
	ok, username := tools.GetUserNameFromToken(token, AuthRedisClient)
	if !ok || username != mux.Vars(r)["username"] {
		w.WriteHeader(401)
		return nil
	}

	path := filepath.Join("/home/", username, "projects", mux.Vars(r)["filepath"])
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
			return err
		}
		w.Header().Set("Content-Type", "image/jpeg")
		io.Copy(w, file)
	}
	w.WriteHeader(400)
	return nil
}
