package webServer

//
//import (
//	"bytes"
//	"crypto/md5"
//	"encoding/hex"
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//	"path/filepath"
//	"strings"
//	"unicode"
//
//	"github.com/julienschmidt/httprouter"
//)
//
//const algPath = "/rtfdata/alg"
//const md5Path = "/rtfdata/md5sum"
//
//type Alg struct {
//	Name    string             `json:"name"`
//	Version map[string]Version `json:"version"`
//}
//
//type Version struct {
//	Attr  map[string]string `json:"attr"`
//	Mtime int64             `json:"mtime"`
//	Size  int64             `json:"size"`
//}
//
//type FileInfo struct {
//	Info map[string]map[string]Version
//}
//
//func ListAlg(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	algs := []*Alg{}
//	allAlgs := getDirFileInfo(algPath)
//
//	for k, allVersions := range allAlgs.Info {
//		versions := make(map[string]Version)
//		for key, v := range allVersions {
//			versions[key] = v
//		}
//		algs = append(algs, &Alg{Name: k, Version: versions})
//	}
//	writeResponse(w, algs)
//}
//
//func GetAlg(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	alg := Alg{}
//	name := params.ByName("name")
//	if !nameValid(name) {
//		logrus.Errorf("GetAlg nameValid %s failed", name)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	path := fmt.Sprintf("%s/%s/%s", algPath, name[:1], name)
//	allAlgs := getDirFileInfo(path)
//	for k, allVersions := range allAlgs.Info {
//		versions := make(map[string]Version)
//		for key, v := range allVersions {
//			versions[key] = v
//		}
//		alg.Name = k
//		alg.Version = versions
//	}
//
//	if alg.Name == "" {
//		logrus.Errorf("GetAlg alg name is %s", "empty")
//		w.WriteHeader(http.StatusNotFound)
//		return
//	}
//	writeResponse(w, alg)
//}
//
//func CreateAlg(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	query := r.URL.Query()
//	name := query.Get("name")
//	version := query.Get("version")
//
//	r.ParseMultipartForm(32 << 20)
//
//	if !nameValid(name) {
//		logrus.Errorf("CreateAlg nameValid %s failed", name)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	if !versionValid(version) {
//		logrus.Errorf("CreateAlg versionValid %s failed", version)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	// 根据字段名获取表单文件
//	formFile, _, err := r.FormFile("uploadfile")
//	if err != nil {
//		logrus.Errorf("CreateAlg get form file failed: %s", err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	defer formFile.Close()
//
//	buf := bytes.NewBuffer(nil)
//	if _, err := io.Copy(buf, formFile); err != nil {
//		logrus.Errorf("CreateAlg copy file buf failed: %s", err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	md5Value := getMd5(buf.Bytes())
//
//	// 创建md5目录
//
//	dir := fmt.Sprintf("%s/%s/%s", md5Path, md5Value[:2], md5Value[2:4])
//	if !exists(dir) {
//		err = os.MkdirAll(dir, os.ModePerm)
//		if err != nil {
//			logrus.Errorf("CreateAlg create dir %s err %s", dir, err.Error())
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
//
//	// 创建保存文件
//
//	path := fmt.Sprintf("%s/%s", dir, md5Value)
//
//	if !exists(path) {
//		destFile, err := os.Create(path)
//		if err != nil {
//			logrus.Errorf("CreateAlg crate file failed: %s", err.Error())
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//		defer destFile.Close()
//
//		if _, err = formFile.Seek(0, io.SeekStart); err != nil {
//			logrus.Errorf("CreateAlg seek failed: %s", err.Error())
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//
//		// 读取表单文件，写入保存文件
//		_, err = io.Copy(destFile, formFile)
//		if err != nil {
//			logrus.Errorf("CreateAlg copy file failed: %s", err.Error())
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//
//		for k, v := range r.MultipartForm.Value {
//			if k == "name" || k == "version" || len(v) == 0 {
//				continue
//			}
//
//			if !attrKeyValid(k) || !attrValueValid(v[0]) {
//				logrus.Errorf("CreateAlg attr valid failed k %s v %s", k, v)
//				w.WriteHeader(http.StatusBadRequest)
//				return
//			}
//
//			attrName := fmt.Sprintf("user.%s-%s-%s", name, version, k)
//			err = xattr.Set(path, attrName, []byte(v[0]))
//			if err != nil {
//				logrus.Errorf("CreateAlg set attr %s failed: %s", k, err.Error())
//				w.WriteHeader(http.StatusInternalServerError)
//				return
//			}
//		}
//
//	}
//
//	algDir := fmt.Sprintf("%s/%s/%s", algPath, name[:1], name)
//	if !exists(algDir) {
//		err = os.MkdirAll(algDir, os.ModePerm)
//		if err != nil {
//			logrus.Errorf("CreateAlg create algDir %s err %s", algDir, err.Error())
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
//
//	newLink := fmt.Sprintf("%s/%s", algDir, version)
//	err = os.Link(path, newLink)
//	if err != nil {
//		logrus.Errorf("CreateAlg create path %s newLink %s link failed: %s", path, newLink, err.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//}
//
//func UpdateAlg(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	query := r.URL.Query()
//	name := query.Get("name")
//	version := query.Get("version")
//
//	if !nameValid(name) {
//		logrus.Errorf("UpdateAlg nameValid %s failed", name)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	if !versionValid(version) {
//		logrus.Errorf("UpdateAlg versionValid %s failed", version)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	path := fmt.Sprintf("%s/%s/%s/%s", algPath, name[:1], name, version)
//
//	for k, v := range query {
//		if k == "name" || k == "version" || len(v) == 0 {
//			continue
//		}
//		if !attrKeyValid(k) || !attrValueValid(v[0]) {
//			logrus.Errorf("UpdateAlg attr valid failed k %s v %s", k, v)
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		attrName := fmt.Sprintf("user.%s-%s-%s", name, version, k)
//		err := xattr.Set(path, attrName, []byte(v[0]))
//		if err != nil {
//			logrus.Errorf("UpdateAlg set attr %s failed: %s", k, err.Error())
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//	}
//	w.WriteHeader(http.StatusOK)
//}
//
//func DeleteAlg(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	var err error
//
//	name := params.ByName("name")
//	version := params.ByName("version")
//
//	if !nameValid(name) {
//		logrus.Errorf("DeleteAlg nameValid %s failed", name)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	if !versionValid(version) {
//		logrus.Errorf("DeleteAlg versionValid %s failed", version)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	path := fmt.Sprintf("%s/%s/%s/%s", algPath, name[:1], name, version)
//
//	err = os.Remove(path)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
//
//func DownloadAlg(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	name := params.ByName("name")
//	version := params.ByName("version")
//
//	if !nameValid(name) {
//		logrus.Errorf("DownloadAlg nameValid %s failed", name)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	if !versionValid(version) {
//		logrus.Errorf("DownloadAlg versionValid %s failed", version)
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	path := fmt.Sprintf("%s/%s/%s/%s", algPath, name[:1], name, version)
//	file, err := os.Open(path)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logrus.Errorf("DownloadAlg unable to open and read file error %v", err)
//		return
//	}
//	defer file.Close()
//	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s:%s", name, version))
//	w.Header().Add("Content-Type", "application/octet-stream")
//	_, err = io.Copy(w, file)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		logrus.Errorf("DownloadAlg copy file error %v", err)
//		return
//	}
//}
//
//func getDirFileInfo(pwd string) FileInfo {
//	files := make(map[string]map[string]Version)
//	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
//		splits := strings.Split(path, "/")
//		if len(splits) == 6 && !info.IsDir() {
//			name := info.Name()
//			apps, ok := files[splits[4]]
//			if !ok {
//				apps = make(map[string]Version)
//			}
//			attrs, err := xattr.List(path)
//			if err != nil {
//				logrus.Errorf("getDirFileInfo path %s list attr err %s", path, err.Error())
//			}
//			logrus.Debugf("getDirFileInfo path %s attrs %v", path, attrs)
//			allAttr := make(map[string]string)
//			for _, attr := range attrs {
//				v, err := xattr.Get(path, attr)
//				if err != nil {
//					logrus.Errorf("getDirFileInfo path %s get %s attr err %s", path, attr, err.Error())
//				}
//				logrus.Debugf("getDirFileInfo arch %s", string(v[:]))
//				if len(attr) <= 5 {
//					continue
//				}
//				splits := strings.Split(attr[5:], "-")
//				if len(splits) < 3 {
//					continue
//				}
//				allAttr[splits[2]] = string(v[:])
//			}
//
//			apps[name] = Version{
//				Attr:  allAttr,
//				Mtime: info.ModTime().Unix(),
//				Size:  info.Size(),
//			}
//			files[splits[4]] = apps
//		}
//		return nil
//	})
//	info := FileInfo{}
//	info.Info = files
//	return info
//}
//
//func getMd5(buf []byte) string {
//	md5Ctx := md5.New()
//	md5Ctx.Write(buf)
//	md5InBytes := md5Ctx.Sum(nil)
//	md5Value := hex.EncodeToString(md5InBytes[:])
//	return md5Value
//}
//
//// Writes the response as a standard JSON response with StatusOK
//func writeResponse(w http.ResponseWriter, m interface{}) {
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//	if err := json.NewEncoder(w).Encode(m); err != nil {
//		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
//	}
//}
//
//// Writes the error response as a Standard API JSON response with a response code
//func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(errorCode)
//	json.NewEncoder(w).Encode(errorMsg)
//}
//
//func versionValid(version string) bool {
//	versionByte := []byte(version)
//
//	valid := func(version string) bool {
//		for _, n := range version {
//			if !(unicode.IsDigit(n) || string(n) == ".") {
//				return false
//			}
//		}
//		return true
//	}
//
//	if version != "" && len(versionByte) <= 8 && valid(version) {
//		return true
//	} else {
//		return false
//	}
//}
//
//func nameValid(name string) bool {
//	nameByte := []byte(name)
//
//	valid := func(name string) bool {
//		for _, n := range name {
//			if !(unicode.IsDigit(n) || unicode.IsLower(n) || string(n) == "_") {
//				return false
//			}
//		}
//		return true
//	}
//
//	if name != "" && len(nameByte) <= 32 && unicode.IsLower([]rune(name)[0]) && valid(name) {
//		return true
//	} else {
//		return false
//	}
//}
//
//func attrKeyValid(name string) bool {
//	nameByte := []byte(name)
//
//	valid := func(name string) bool {
//		for _, n := range name {
//			if !(unicode.IsDigit(n) || unicode.IsLetter(n) || string(n) == "_" || string(n) == ".") {
//				return false
//			}
//		}
//		return true
//	}
//
//	if name != "" && len(nameByte) <= 32 && valid(name) {
//		return true
//	} else {
//		return false
//	}
//}
//
//func attrValueValid(val string) bool {
//	valByte := []byte(val)
//
//	if len(valByte) <= 1024 {
//		return true
//	}
//	return false
//}
