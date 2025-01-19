package admin

import (
	"context"
	v1 "demo/api/admin/v1"
	"demo/internal/dao"
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/errors/gerror"
)

type cAdmin struct{}

var Admin cAdmin

func (c *cAdmin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	adminInfo := entity.Admin{}
	err = dao.Admin.Ctx(ctx).Where(dao.Admin.Columns().Name).Scan(&adminInfo)
	if err != nil {
		return nil, gerror.New("用户名或密码错误")
	}
	if adminInfo.Password != req.Password {
		return nil, gerror.New("用户名或密码错误")
	}
	return &v1.LoginRes{
		Id:   adminInfo.Id,
		Name: adminInfo.Name,
	}, nil
}
