if dataset.Load(id); dataset.ID == 0 || dataset.IsDelete {
	response.Error(ctx, response.WrongDatasetID)
	return
}

// UpdateTagStatus ...
func (h *Handler) UpdataTagStatus(ctx *context.Context) {
	var (
		id = ctx.UserValue[guard.DatasetIDKey].(int64)
		datasetTagStatusParam = guard.GetDatasetTagStatusParam(ctx)
		dataset = models.DatasedsModelWithDriver(h.conn)
	)
	if dataset.Load(id); dataset.ID == 0 || dataset.IsDelete {
		response.Error(ctx,response.WrongDatasetID)
		return
	}
	var notEditFields = []string{
		"Title", "FeatureType", "FeatureCount", "CRS", "IsDelete",
		"HasParsed", "ShareType", "CreatorID", "TitleFieldID", "FolderID",
	}
	dataset.IsTagHide = datasetTagStatusParam.IsTagHide
	dataset.FieldNotAllowEdit(notEditFields...)
	dataset.Save()
	response.OkWithData(ctx,map[string]interface{}{
		"Dataset":domain.BuildDatasetResWithModel(dataset),
	})
}

if fir.NextID != 0 {
	nums := field.Select(`"OrderValue" between ? and ? and "DatasetID" = ? and "IsDelete"=?`, prevOrderValue, nextOrderValue, datasetID, false)
	logger.Debug(nums)
	if len(num) > 2 {
		response.Error(ctx, response.WrongNotNeighbour)
		return
	}

type DatasetTagStatusParam struct {
	IsTagHide bool
}


// UpdataTagStatus ...
func (g *Guard) UpdateTagStatus(ctx *context.Context) {
	if VerifyReqID(DatasetIDKey,ctx) != nil {
		Alter(ctx, response.WrongDatasetID)
		return
	}
	datasetTagStatusParam := &DatasetTagStatusParam{}
	if err := odysseyUtils.ParseRequestJSON(ctx, datasetTagStatusParam)
	ctx.Next()
}

// GetDatasetTagStatusParam ...
func GetDatasetTagStatusParam(ctx *context.Context) *DatasetTagStatusParam {
	return ctx.UserValue[getDatasetTagStatusParamKey].(*DatasetTagStatusParam)
}

route.PUT("/datasets/:"+guard.DatasetIDKey+"/tag-status",bagel.guardian.UpdataTagStatus,bagel.handler.UpdataTagStatus)


if searchKey != "" {
	filteredRequests = searchFeatures(filteredResults, searchKey)
}

// searchFeature 通过关键字进行搜索
func searchFeatures(results []map[string]interface{}, searchKey string) (filteredResults
 []map[string]interface{}){
	 for _, 
 }
