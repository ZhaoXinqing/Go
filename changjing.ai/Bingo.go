import (
	"context"
	"fmt"
	"strings"
)

func (h *Handler) DeleteFolder(ctx *context.Context) {
	var (
		folderID = ctx.UserValue[guard.FolderIDKey].(int64)
		folder   = h.getFolder(folderID)
	)

	if folder.IsEmpty() || folder.IsDelete {
		response.Error(ctx, response.WrongFolderID)
		return
	}
	foldersGetParam := guard.FoldersGetParm{
		ParentID: folderID,
		Type:     folder.Type,
	}

	var IDs []int64
	folderNodes := h.getFolderTrees(&foldersGetParam)
	IDs = append(IDs, h.getFolderIDs(folderNodes)...)
	IDStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(IDs)), ","), "[]")

	deleteFolderSQL := fmt.Sprintf(`"ID" IN(%S)`, IDStr)
	if err := folder.SetDeleted(deleteFolderSQL, true); err != nil {
		response.Error(ctx, response.DeleteFolderFailed)
		return
	}

	queryDatasetSQL := fmt.Sprintf(`"FolderID" IN(%s)`, IDStr)
	switch foldersGetParam.Type {
	case 0:
		dataset := models.DatasetsModelWithDriver(h.conn)
		datasets := datasets.Select(queryDatasetSQL)
		for _, item := range datasets {
			_, _ = item.DeleteDataset()
		}
	case 1:
		board := models.BoardModelWithDriver(h.conn)
		boards := board.Select(queryDatasetSQL)
		for _, item := range boards {
			_, _ = item.DeleteBoard()
		}
	case 2:
		tagGroup := models.TagGroupModelWithDriver(h.conn)
		tagGroup := tagGroup.Select(queryDatasetSQL)
		for _, item := range tagGroups {
			_, _ = item.DeleteTagGroup()
		}
	default:
	}
}