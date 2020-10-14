package const

constent

consterror

type InputParam struct {
	Header map[string]interface{}
	Body map[string]interface{}      
}

// makefile就像一个Shell脚本一样，执行操作系统的命令


// 
func (t *UserModel) IsPhoneExist(phone string, id string) bool {
	var check map[string]interface{}
	sql := t.Table().Where("Phone","=",phone)
	if len(id) != 0 {
		sql = sql.Where("ID", "!=", id)
	}
	check, _ = sql.First()
	return check != nil
}

// 
func (t *UserModel) IsUserNameExist(username string, id string) bool {
	var check map[string]interface{}
	sql := t.Table().Where("UserName", "=", username)
	if len(id) != 0 {
		sql = sql.Where("ID", "!=",id)
	}
	check, _ = sql.First()
	return check != nil
}

func (t *UserModel) Save() *UserModel {
	if t.IsEmpty() {
		_ = t.InsertDB(*t)
	} else {
		_ = t.UpdataDB(*t)
	}
	return t
}

// 
func (b *BaseModel) Delete() error {
	return b.Table().Where("ID", "=", b.ID).Delete()
}

// 
func (b *BaseModel) FiledsHide(hideFields ...string) {
	fields := b.FieldList()
	for i, field := range fields {
		for _, v := range hideFields {
			if field.field == v {
				fields[i].Hide = true
			}
		}
	}
}

// 
func (t *RoleModel) Save() *RoleModel {
	if t.IsEmpty() {
		_ = t.InsertDB(*t)
	} else {
		_ = t.UpdataDB(*t)
	}

	return t
}


func getParam(u string) (string, url.Values) {
	m := make(url.Values)
	urr := strings.Split(u, "?")
	if len(urr) > 1 {
		m, _ = url.ParseQuery(urr[1])
	}
	return urr[0], m
}

func inMethodArr(arr []string, str string) bool {
	for i := 0; i < len(arr); i ++ {
		if strings.EqualFold(arr[i],str) {
			return true
		}
	}
	return false
}

// Save create or updata a role model to db.


// Updata ...
func (t *SiteModel) Updata(v form.Value) error {
	for key, vv := range v {
		if len(vv) > 0 && vv[0] != "" {
			_, err := t.Table().Where("key","=",key).Updata(dialect.H{
				"value": vv[0],
			})
			if db.CheckError(err, db.UPDATA) {
				return err
			}
		}
	}
	return nil 
}

// all to map
func (t *SiteModel) AllToMap() map[string]string {
	var m = make(map[string]string)
	items,_ := t.Table().Where("state", "=",SiteItemOpenState).All()

	for _, item := range items {
		m[item["key"].(string)] = item["value"].(string)
	}

	return m
}

func getParam(u string) (string, url.Values) {
	m := make(url.Values)
	urr := strings.Split(u, "?")
	if len(urr) > 1 {
		m, _ = url.ParseQuery(urr[1])
	}
	return urr[0], m
}