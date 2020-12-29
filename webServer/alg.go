package webServer

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"unicode"
)

const algPath = "/rtfdata/alg"
const md5Path = "/rtfdata/md5sum"

type Alg struct {
	Name    string             `json:"name"`
	Version map[string]Version `json:"version"`
}

type Version struct {
	Attr  map[string]string `json:"attr"`
	Mtime int64             `json:"mtime"`
	Size  int64             `json:"size"`
}

type FileInfo struct {
	Info map[string]map[string]Version
}

// CreateAlg ...
func CreateAlg(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	version := query.Get("version")

	r.ParseMultipartForm(32 << 20)

	if !nameValid(name) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !versionValid(version) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 根据字段名获取表单文件
	formFile, _, err := r.FormFile("uploadfile")
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer formFile.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, formFile); err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	md5Value := getMd5(buf.Bytes())

	// 创建md5目录
	dir := fmt.Sprintf("%s/%s/%s", md5Path, md5Value[:2], md5Value[2:4])

	// 创建保存文件

	path := fmt.Sprintf("%s/%s", dir, md5Value)

	// 读取表单文件，写入保存文件
	// _, err = io.Copy(destFile, formFile)
	// if err != nil {

	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	for k, v := range r.MultipartForm.Value {
		if k == "name" || k == "version" || len(v) == 0 {
			continue
		}

		if !attrKeyValid(k) || !attrValueValid(v[0]) {

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_ = fmt.Sprintf("user.%s-%s-%s", name, version, k)

	}

	algDir := fmt.Sprintf("%s/%s/%s", algPath, name[:1], name)

	newLink := fmt.Sprintf("%s/%s", algDir, version)
	err = os.Link(path, newLink)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdateAlg ...
func UpdateAlg(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	version := query.Get("version")

	if !nameValid(name) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !versionValid(version) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// path := fmt.Sprintf("%s/%s/%s/%s", algPath, name[:1], name, version)

	for k, v := range query {
		if k == "name" || k == "version" || len(v) == 0 {
			continue
		}
		if !attrKeyValid(k) || !attrValueValid(v[0]) {

			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteAlg ...
func DeleteAlg(w http.ResponseWriter, r *http.Request) {
	var err error

	name := "name"
	version := "version"

	if !nameValid(name) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !versionValid(version) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path := fmt.Sprintf("%s/%s/%s/%s", algPath, name[:1], name, version)

	err = os.Remove(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DownloadAlg ...
func DownloadAlg(w http.ResponseWriter, r *http.Request) {
	name := "name"
	version := "version"

	if !nameValid(name) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !versionValid(version) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path := fmt.Sprintf("%s/%s/%s/%s", algPath, name[:1], name, version)
	file, err := os.Open(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	defer file.Close()
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s:%s", name, version))
	w.Header().Add("Content-Type", "application/octet-stream")
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func getMd5(buf []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(buf)
	md5InBytes := md5Ctx.Sum(nil)
	md5Value := hex.EncodeToString(md5InBytes[:])
	return md5Value
}

// Writes the response as a standard JSON response with StatusOK
func writeResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(errorMsg)
}

func versionValid(version string) bool {
	versionByte := []byte(version)

	valid := func(version string) bool {
		for _, n := range version {
			if !(unicode.IsDigit(n) || string(n) == ".") {
				return false
			}
		}
		return true
	}

	if version != "" && len(versionByte) <= 8 && valid(version) {
		return true
	}
	return false

}

// nameValid ...
func nameValid(name string) bool {
	nameByte := []byte(name)

	valid := func(name string) bool {
		for _, n := range name {
			if !(unicode.IsDigit(n) || unicode.IsLower(n) || string(n) == "_") {
				return false
			}
		}
		return true
	}

	if name != "" && len(nameByte) <= 32 && unicode.IsLower([]rune(name)[0]) && valid(name) {
		return true
	}
	return false
}

// attrKeyValid ...
func attrKeyValid(name string) bool {
	nameByte := []byte(name)

	valid := func(name string) bool {
		for _, n := range name {
			if !(unicode.IsDigit(n) || unicode.IsLetter(n) || string(n) == "_" || string(n) == ".") {
				return false
			}
		}
		return true
	}

	if name != "" && len(nameByte) <= 32 && valid(name) {
		return true
	}
	return false

}

// attrValueValid ...
func attrValueValid(val string) bool {
	valByte := []byte(val)

	if len(valByte) <= 1024 {
		return true
	}
	return false
}
