package Admin

import (
	c "github.com/GoAdminGroup/go-admin/modules/config"
    "github.com/GoAdminGroup/go-admin/modules/service"
    "github.com/GoAdminGroup/go-admin/plugins"
)

type Example struct {
	*plugin.Base
}

func NewExample() *Example {
	return &Example{

	}
}


type Plugin interface {
	
}