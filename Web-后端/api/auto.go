package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

const filename = "/rtfdata/platform/config/autostart.json"

type Autoapp struct {
	Nameversion []string `json:"name"`
}

//查询
// 1、判断文件是否存在，没有则创建
// 2、读文件
// 3、写出来
func GetAutostart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// var result []string
	type Info struct {
		Name string
		Path string
		Type string
		Size int
		Data string
	}
	result := make([]string, 0)

	// 判断路径是否存在
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Println("该路径不存在", err)

		// 创建目录
		err := os.MkdirAll("/rtfdata/platform/config", os.ModePerm)
		if err != nil {
			fmt.Println("创建目录失败", err)
		}

		// 创建文件
		fp, err := os.Create("/rtfdata/platform/config/autostart.json")
		if err != nil {
			fmt.Println("文件创建失败。")
			//创建文件失败的原因有：1、路径不存在  2、权限不足  3、打开文件数量超过上限  4、磁盘空间不足等
		}
		// defer延迟调用
		defer fp.Close() //关闭文件，释放资源。
	}

	// 读json文件内容
	datas, err := ioutil.ReadFile(filename)
	fmt.Println(datas)
	if err != nil {
		logrus.Errorf("GetAutostart open file err %s", err.Error())
	}

	// 判断是否为空值
	if len(datas) == 0 {
		goto label1
	}

	// 将解码文件放入目标result，并返回
	err = json.Unmarshal(datas, &result)
	if err != nil {
		logrus.Errorf("解码文件错误 %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println("result", result)

label1:
	writeResponse(w, result)
}

// 添加
func CreateAutostart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fields := make([]string, 0)

	// 解析接受参数s
	name := params.ByName("name")
	version := params.ByName("version")
	fmt.Println(name)
	fmt.Println(version)

	// 判断路径是否存在
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Println("该路径不存在", err)

		// 创建目录
		err := os.MkdirAll("/rtfdata/platform/config", os.ModePerm)
		if err != nil {
			fmt.Println("创建目录失败", err)
		}

		// 创建文件
		fp, err := os.Create("/rtfdata/platform/config/autostart.json")
		if err != nil {
			fmt.Println("文件创建失败。")
			//创建文件失败的原因有：
			//1、路径不存在  2、权限不足  3、打开文件数量超过上限  4、磁盘空间不足等
		}
		// defer延迟调用
		defer fp.Close() //关闭文件，释放资源。
	}
	fmt.Println("路径存在")

	// 读取文件中数据，保存为map格式
	data, _ := ioutil.ReadFile("/rtfdata/platform/config/autostart.json")

	// 判断是否为空值
	if len(data) == 0 {
		goto label2
	}

	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}

label2:
	// 合并map
	fields = append(fields, name+":"+version)
	fmt.Println(fields)

	// 写入文件
	out, _ := json.Marshal(fields)
	_ = ioutil.WriteFile("/rtfdata/platform/config/autostart.json", out, 0755)
}

// 更新PUT
func UpdateAutostart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// 删除原来json文件
	err := os.Remove("/rtfdata/platform/config/autostart.json")
	if err != nil {
		fmt.Println("删除文件出错")
	}
	fmt.Println("已删除文件")

	// 获取json更新数据
	datastr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 定义类型
	var data map[string][]string

	// 解码json
	err = json.Unmarshal(datastr, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("获取body", data)

	// 创建文件
	_, err = os.Create("/rtfdata/platform/config/autostart.json")
	if err != nil {
		fmt.Println("文件创建失败。")
	}
	fmt.Println("创建文件成功")

	// 写入文件
	out, _ := json.Marshal(data["json"])
	_ = ioutil.WriteFile("/rtfdata/platform/config/autostart.json", out, 0755)

}
