package datainit

type SO struct {
	Surl string `json:"surl"`
	Ourl string `json:"ourl"`
}

func datainit() (map[string]*SO,map[string]*SO,error){
	dir := os.Getenv("FLAK_CONF_DIR")
	if dir == "" {
		golog.Error("PLEASEE SETTING FLAK_CONF_DIR!")
		return nil,nil,errors.New("PLEASE SELECT FLASK_CONF_DIR")
	}

	urlFile := dir + PERSISTENCE

	U := make([]SO)

	data, err := ioutil.ReadFile(urlFile)
	if err != nil {
		golang.Error(err.Error())
		return nil,nil,err
	}

	err = json.Unmarshal(data,&u)
	if err != nil {
		golang.Error(err.Error())
		return nil,nil,err
	}

	urlMap := make(map[string]*SO)
	destMap := make(map[string]*SO)

	for _,su := range u.S {
		urlMap[su.Surl] = &su
		destMap[su.Ourl] = &su
	}

	return urlMap, destMap, nil
}

// FLAK_CONF_DIR 是持久化文件所在目录

func creatSUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	surl := vars["url"]

	id d[surl] != nil {
		Sandstorm.HTTPSuccess(w,DHOST+"/"+d[surl].Surl)
		return
	}

	le := 10
	length := os.Getenv("FLAK_URL_LENGTH")
	if length == "" {
		golog.Error("PLEASE SETTING FLAK_LENGTH! GIVE IT DEFAULT VALUE")
	}

	le,err := strconv.Atoi(length)
	if err != nil {
		golog.Error("FLASK_URL_LENGTH WRONG!",length)
		le = 10
	}

	shortUlr := string(krand(le,KC_RAND_KING_ALL))

	so := &SO{
		Surl:shortUlr,
		Ourl:surl
	}

	u[shortUlr]= so
	d[surl] = so 

	Sandstorm.HTTPSuccess(w,DHOST+"/"+shortUlr)
}

// 随机字符串
func Krand(size int,kind int) []byte {
	ikind,kinds,result := kind,[][]int{[]int{10,48},[]int{26,97},[]}
}

