package director

import "github.com/neutrinocorp/ddderr"

var (
	IdRequired          = ddderr.NewRequired("director_id")
	UserRequired        = ddderr.NewRequired("director_user_id")
	DisplayNameRequired = ddderr.NewRequired("director_display_name")
)
