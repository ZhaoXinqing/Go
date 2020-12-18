package changjing_ai

import (
	"context"
	"fmt"
	"strings"
)

// GetBoards 获取看板列表
func (h *Handler) GetBoards(ctx *context.Context) {
	var (
		whereSQL string
		boards   = make([]*domain.BoardRes, 0)
		// boardIDField   = modules.FilterField("ID", constsql.Delimiter)
		isDeleteField  = modules.FilterField("IsDelete", constsql.Delimiter)
		updatedAtField = modules.FilterField("UpdatedAt", constsql.Delimiter)
		user           = guard.GetCurrentUser(ctx)
		args           []interface{}
	)

	// 获取权限下看板 ID 集合
	boardPermissionsModel := models.BoardPermissionsModelWithDriver(h.conn)
	tableName0 := boardPermissionsModel.GetTableName()
	selectSQL0 := fmt.Sprintf(`SELECT "BoardID" FROM %s WHERE "UserID" = ? ORDER BY "UpdatedAt" DESC`, tableName0)
	idSlice, _ := boardPermissionsModel.SelectBySQL(selectSQL0, user.ID)

	boardModel := models.BoardsModelWithDriver(h.conn)
	tableName := boardModel.GetTableName()
	fields := boardModel.GetSelectFields()
	fieldSQL := strings.Join(fields, ",")
	for _, mapid := range idSlice {
		boardID := mapid["BoardID"]
		id := boardID.(int64)
		args = append(args, id)
	}

	if boardParam.FolderID > 0 {
		whereSQL = fmt.Sprintf("%s AND %s = ?", whereSQL, folderIDField)
		args = append(args, boardParam.FolderID)
	} else {
		whereSQL = fmt.Sprintf("%s AND %s <= 0", whereSQL, folderIDField)
	}

	whereSQL = fmt.Sprintf(`%s = False`, isDeleteField)
	selectSQL := fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s DESC", fieldSQL, tableName, whereSQL, updatedAtField)
	logger.Debug(whereSQL)
	logger.Debug(selectSQL)
	results, _ := boardModel.SelectBySQL(selectSQL, args...)
	logger.Debug(results)

	for _, result := range results {
		boards = append(boards, domain.BuildBoardResWithMap(result))
	}
	response.OkWithData(ctx, map[string]interface{}{
		"Boards": boards,
	})
}
