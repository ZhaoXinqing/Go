func LoadFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    url:= p.ByName("url")
    file, err:= os.Open("./" + url)
    if err!= nil {
        httputil.WriteErrorResponse(w, "文件不存在")
        return
    }
    defer file.Close()
    buff, _ := ioutil.ReadAll(file)
    w.Write(buff)
}