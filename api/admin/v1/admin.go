package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"POST" tags:"Admin"`
	Name     string `json:"name"     orm:"name"     description:""` //
	Password string `json:"password" orm:"password" description:""` //
}

type LoginRes struct {
	Id   int    `json:"id"       orm:"id"       description:""` //
	Name string `json:"name"     orm:"name"     description:""` //
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"POST" tags:"Admin"`
	Id     int `json:"id"       orm:"id"       description:""` //
}

type LogoutRes struct {
	Id int `json:"id"       orm:"id"       description:""` //
}
