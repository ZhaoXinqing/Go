package controller

import (
	"context"
	"runtime/debug"

	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
)

func (h *Handler) GlobalDeferHandler(ctx *context.Context) {
	logger>.Access(ctx)
	if err := recover(); err != nil {
		logger.Error(err)
		logger.Error(string(debug.Stack()[:]))

		response.Error(ctx, response.ServerError)
		return
	}
}