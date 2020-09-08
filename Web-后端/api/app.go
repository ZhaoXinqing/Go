package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"unicode"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

const appPath = "/rtfdata/app"

func ListApp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var apps []string
	files, _ := ioutil.ReadDir(appPath)
	for _, f := range files {
		if f.IsDir() {
			apps = append(apps, f.Name())
		}
	}
	writeResponse(w, apps)
}

func GetApp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var v interface{}
	var buf bytes.Buffer
	name := params.ByName("name")
	if !appNameValid(name) {
		logrus.Errorf("GetApp appNameValid %s", "failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("%s/%s", appPath, name)

	files, _ := ioutil.ReadDir(path)
	buf.Write([]byte(`[`))
	for _, f := range files {
		if !f.IsDir() {
			name := fmt.Sprintf("%s/%s", path, f.Name())
			fp, err := os.Open(name)
			if err != nil {
				logrus.Errorf("GetApp open file %s err %s", name, err.Error())
				continue
			}
			data, err := ioutil.ReadAll(fp)
			if err != nil {
				logrus.Errorf("GetApp readAll file %s err %s", name, err.Error())
				continue
			}
			buf.Write(data)
			buf.Write([]byte(`,`))
		}
	}
	bytes := buf.Bytes()
	end := len(bytes) - 1
	resetBytes := bytes[:end]
	resetBytes = append(resetBytes, []byte(`]`)...)

	err := json.Unmarshal(resetBytes, &v)
	if err != nil {
		logrus.Errorf("GetApp json unmarshal path %s err %s", path, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeResponse(w, v)
}

func CreateApp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	query := r.URL.Query()
	name := query.Get("name")
	version := query.Get("version")

	if !appNameValid(name) {
		logrus.Errorf("CreateApp nameValid %s failed", name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !versionValid(version) {
		logrus.Errorf("CreateApp versionValid %s failed", version)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("CreateApp readAll err %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dir := fmt.Sprintf("%s/%s", appPath, name)

	if !exists(dir) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			logrus.Errorf("CreateApp create dir %s err %s", dir, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	path := fmt.Sprintf("%s/%s", dir, version)
	fp, err := os.Create(path)
	if err != nil {
		logrus.Errorf("CreateApp create file %s err %s", path, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		logrus.Errorf("CreateApp write file %s data err %s", path, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func UpdateApp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	query := r.URL.Query()
	name := query.Get("name")
	version := query.Get("version")
	if !appNameValid(name) {
		logrus.Errorf("UpdateApp nameValid %s failed", name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !versionValid(version) {
		logrus.Errorf("UpdateApp versionValid %s failed", version)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("%s/%s/%s", appPath, name, version)
	fp, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		logrus.Errorf("UpdateApp open file %s err %s", path, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		logrus.Errorf("UpdateApp write file %s data err %s", path, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DeletAppVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	version := params.ByName("version")

	if !appNameValid(name) {
		logrus.Errorf("DeletAppVersion nameValid %s failed", name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !versionValid(version) {
		logrus.Errorf("DeletAppVersion versionValid %s failed", version)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path := fmt.Sprintf("%s/%s/%s", appPath, name, version)

	err := os.Remove(path)
	if err != nil {
		logrus.Errorf("DeletAppVersion remove path err %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteApp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")

	if !appNameValid(name) {
		logrus.Errorf("DeleteApp nameValid %s failed", name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path := fmt.Sprintf("%s/%s/", appPath, name)

	err := os.RemoveAll(path)
	if err != nil {
		logrus.Errorf("DeleteApp remove  %s err %s", path, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetJsonSchema(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var result interface{}
	path := fmt.Sprintf("%s/json_schema", appPath)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("GetJsonSchema read file path %s err %s", path, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		logrus.Errorf("GetJsonSchema json unmarshal path %s err %s", path, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeResponse(w, result)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func appNameValid(name string) bool {
	nameByte := []byte(name)

	valid := func(name string) bool {
		for _, n := range name {
			if !(unicode.IsDigit(n) || unicode.IsLetter(n) || string(n) == "_") {
				return false
			}
		}
		return true
	}

	if name != "" && len(nameByte) <= 32 && unicode.IsLetter([]rune(name)[0]) && valid(name) {
		return true
	} else {
		return false
	}
}