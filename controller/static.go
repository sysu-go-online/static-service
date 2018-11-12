package controller

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	pModel "github.com/sysu-go-online/project-service/model"

	"gopkg.in/h2non/filetype.v1"

	"github.com/gorilla/mux"
	"github.com/sysu-go-online/public-service/tools"
	userModel "github.com/sysu-go-online/service-end/model"
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

	vars := mux.Vars(r)
	projectName := vars["projectname"]
	filePath := vars["filepath"]
	givenUsername := vars[username]

	if givenUsername != username {
		w.WriteHeader(401)
		return nil
	}

	// Get project information
	session := MysqlEngine.NewSession()
	u := userModel.User{Username: username}
	ok, err := u.GetWithUsername(session)
	if !ok {
		w.WriteHeader(400)
		return nil
	}
	if err != nil {
		return err
	}
	p := pModel.Project{Name: projectName, UserID: u.ID}
	has, err := p.GetWithUserIDAndName(session)
	if !has {
		w.WriteHeader(204)
		return nil
	}
	if err != nil {
		return err
	}

	path := filepath.Join("/home/", username, "projects", p.Path, filePath)
	f, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Println(err)
		w.WriteHeader(204)
		return nil
	}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		w.WriteHeader(204)
		return nil
	}
	// check file size
	// get the size
	size := f.Size()
	if size > 4*1024*1024 {
		w.WriteHeader(400)
		return nil
	}
	if filetype.IsImage(buf) {
		file, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return err
		}
		w.Header().Set("Content-Type", "image/jpeg")
		io.Copy(w, file)
	}
	w.WriteHeader(400)
	return nil
}
