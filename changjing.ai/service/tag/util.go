package tag

import "context"

// DeleteBoardPermission ...
func (h *Handler) DeleteBoardPermission(ctx *context.Context) {
	var (
		boardID         = ctx.UserValue[guard.BoardIDKey].(int64) // 从 url 中获取参数
		userIDParam     = guard.GetUserIDParam(ctx)               // 从 body 中获取参数
		boardPermission = models.BoardPermissionsModelWithDriver(h.conn)
	)
	boardPermission.DeleteDB(`"BoardID" = ? and "UserID" = ?`, boardID, userIDParam.UserID)
	response.OkWithData(ctx, map[string]interface{}{})
}
