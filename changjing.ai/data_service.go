package changjing_ai

//func ExportExcel(dataset *models.DatasetsModel) (string, error) {
//	feature,fieldMap := GetFeatures(dataset.ID, dataset.GetConn())
//	var field []*models.DatasetFieldsModel
//
//	for _, v := range fieldMap {
//		fields = append(fields, v)
//	}
//	sort.SliceStable(fields,func(i,j int) bool {
//		return fields[i].OrderValue < fields[j].OrderValue
//	})
//
//	var records [][]interface{}
//	var recordTitle []interface{}
//	for _, field := range field
//}

//func main() {
//	dat, err := ioutil.ReadFile("test.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	r := csv.NewReader(strings.NewReader(string(dat[:])))
//	for {
//		record, err := r.Read()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(record)
//	}
//}
