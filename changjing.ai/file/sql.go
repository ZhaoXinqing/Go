packages controller

import 

func insql([]interface)result string {
	fieldModel             := models.DatasetFieldsModelWithDriver(t.Conn)
	FieldIDStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ap.FieldIDs)), ","), "[]")
	querySQL := fmt.Sprintf(`"ID" IN (%s)`, FieldIDStr)
	fieldModels := fieldModel.Select(querySQL)
}