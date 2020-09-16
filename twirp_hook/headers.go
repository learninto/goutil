package twirp_hook

import (
	"context"
	"encoding/json"

	"github.com/learninto/goutil/ctxkit"
	"github.com/learninto/goutil/jwt"
	"github.com/learninto/goutil/log"
	"github.com/learninto/goutil/twirp"
)

// NewHeaders
func NewHeaders() *twirp.ServerHooks {
	type user struct {
		// Comment: 企业id
		CompanyID int64 `json:"company_id"`
		// Comment: 唯一标识
		ID           int64 `json:"id"`
		DepartmentID int64 `json:"department_id"`
		// Comment: 角色id数组 英文逗号隔开
		PartIds string `json:"part_ids"`
		// Comment: 部门id数组 英文逗号隔开
		DepartmentIds string `json:"department_ids"`
		//Comment: 用户昵称
		NickName string `json:"nick_name"`
	}
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			req, _ := twirp.HttpRequest(ctx)

			u := user{}

			if c, err := jwt.NewJWT().ParseToken(req.Header.Get("Sign")); err == nil {
				_ = json.Unmarshal(c.Data, &u)
			} else {
				log.Get(ctx).Error("jwt解密 Error", err)
			}

			ctx = ctxkit.WithUserID(ctx, u.ID)                   // 注入用户id
			ctx = ctxkit.WithNickName(ctx, u.NickName)           // 注入用户昵称
			ctx = ctxkit.WithCompanyID(ctx, u.CompanyID)         // 注入公司id
			ctx = ctxkit.WithDepartmentID(ctx, u.DepartmentID)   // 注入管辖部门id
			ctx = ctxkit.WithPartIds(ctx, u.PartIds)             // 注入角色id
			ctx = ctxkit.WithDepartmentIds(ctx, u.DepartmentIds) // 注入部门id

			ctx = ctxkit.WithSignKey(ctx, req.Header.Get("Sign"))      // 注入签名
			ctx = ctxkit.WithDevice(ctx, req.Header.Get("Device"))     // 注入 用户设备  iso、android、web
			ctx = ctxkit.WithMobiApp(ctx, req.Header.Get("MobiApp"))   // 注入 APP 标识
			ctx = ctxkit.WithVersion(ctx, req.Header.Get("Version"))   // 注入 版本 标识
			ctx = ctxkit.WithPlatform(ctx, req.Header.Get("Platform")) // 注入 平台 标识
			ctx = ctxkit.WithUserIP(ctx, req.RemoteAddr)               // TODO 注入 客户端IP 标识  目前貌似不准确待测试

			return ctx, nil
		},
	}
}
