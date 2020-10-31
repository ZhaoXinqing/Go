package account


// Login ...
func (h *Handle) Login(ctx *context.Context) {
	signParam, err := common.GetSignParam(ctx)
	if err != nil {
		response.Error(ctx, response.WrongSignParams)
		return
	}
	if signParam.Password == "" || signParam.UserName == "" {
		response.Error(ctx, response.WrongUserNameOrPassword)
		return
	}
	
	data, err := h.accountClient.Sign(signParam.UserName， signParam.Password)
	if err != nil {
		logger.Error(err)
		response.Error(ctx,response.LoginFailed)
		return
	}

	roleModels, _ := h.adminClient.GetUserRoles(data.User.ID)
	menuIds, _ := h.adminClient.GetRoleMenus(roleModels)

	log := &guard.ConnectionInfo{Info:&models.LoginLog{
		IP: ctx.LocalIP(),
		UserID:data.User.ID,
		UserName:data.User.UserName,
		Token:data.Token,
	},State:0}

	SaveLogin(h.conn, log)
	response.OkWithData(ctx,map[string]interface{}{"token":data.Token,"user":data.User})
}