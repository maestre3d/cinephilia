package actor

import "github.com/neutrinocorp/ddderr"

var (
	IdRequired          = ddderr.NewRequired("actor_id")
	DisplayNameRequired = ddderr.NewRequired("actor_display_name")
	UserRequired        = ddderr.NewRequired("actor_user_id")
	AlreadyExists       = ddderr.NewAlreadyExists(nil, "actor")
	NotFound            = ddderr.NewNotFound(nil, "actor")
)
