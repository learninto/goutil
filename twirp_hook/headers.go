package twirp_hook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/learninto/goutil/memdb"
	"net/http"
	"time"

	"github.com/learninto/goutil/conf"
	"github.com/learninto/goutil/ctxkit"
	"github.com/learninto/goutil/log"
	"github.com/learninto/goutil/twirp"
	"github.com/learninto/goutil/xhttp"
	"github.com/learninto/goutil/xjwt"
)

func NewInternalHeaders() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			req, ok := twirp.HttpRequest(ctx)
			if !ok {
				return ctx, nil
			}
			sign := req.Header.Get("Sign")

			ctx = ctxkit.WithSignKey(ctx, sign)                        // 注入签名
			ctx = ctxkit.WithDevice(ctx, req.Header.Get("Device"))     // 注入 用户设备  iso、android、web
			ctx = ctxkit.WithMobiApp(ctx, req.Header.Get("MobiApp"))   // 注入 APP 标识
			ctx = ctxkit.WithVersion(ctx, req.Header.Get("Version"))   // 注入 版本 标识
			ctx = ctxkit.WithPlatform(ctx, req.Header.Get("Platform")) // 注入 平台 标识
			//ctx = ctxkit.WithUserIP(ctx, req.RemoteAddr)               // TODO 注入 客户端IP 标识  目前貌似不准确待测试

			/* ------ 用户信息 ------ */
			c, err := xjwt.CustomClaims{}.ParseToken(ctx, sign)
			if err != nil {
				return ctx, nil
			}

			user := struct {
				ID        int64 `json:"id"`
				CompanyId int64 `json:"company_id"`
			}{}
			_ = json.Unmarshal(c.Data, &user)
			ctx = ctxkit.WithUserID(ctx, user.ID)           // 注入用户id
			ctx = ctxkit.WithCompanyID(ctx, user.CompanyId) // 注入公司id

			return ctx, nil
		},
	}
}

// NewHeaders headers 拦截
func NewHeaders() *twirp.ServerHooks {
	type User struct {
		// Comment: 企业id
		CompanyID int64 `json:"company_id"`
		// Comment: 唯一标识
		ID int64 `json:"id"`
		// Comment: 是否是系统默认角色,系统默认角色不允许删除  0:非默认  1:默认
		IsSystem int64 `json:"is_system"`
		// Comment：部门id
		DepartmentID int64 `json:"department_id"`
		// Comment: 角色id数组 英文逗号隔开
		PartIds string `json:"part_ids"`
		// Comment: 角色name数组 英文逗号隔开
		PartNames string `json:"part_names"`
		// Comment: 部门id数组 英文逗号隔开
		DepartmentIds string `json:"department_ids"`
		// Comment: 管辖用户id数组 英文逗号隔开
		ManagerUserIds string `json:"manager_user_ids"`
		// Comment: 用户昵称
		NickName string `json:"nick_name"`
		// Comment: 用户登录账号
		UserName string `json:"user_name"`
		// Comment: 密码
		PassWord string `json:"pass_word"`
		// Comment: 权限编码  多个，隔开
		RolesCodes string `json:"roles_codes"`
		// Comment: 是否生效。0：未生效；100：已生效
		// Default: 100
		Status int8 `json:"status"`
		// Comment: 公司
		Company struct {
			// Comment: 是否生效。0：未生效；100：已生效
			// Default: 100
			Status int `json:"status"`
			// Comment: 过期时间
			ExpiryTime int64 `json:"expiry_time"`
		} `json:"company"`
	}
	return &twirp.ServerHooks{
		RequestRouted: func(ctx context.Context) (context.Context, error) {
			req, ok := twirp.HttpRequest(ctx)
			if !ok {
				return ctx, nil
			}
			u := User{}
			token := req.Header.Get("Sign")

			/* ------ 注入标识 ------ */
			ctx = ctxkit.WithSignKey(ctx, token)                       // 注入签名
			ctx = ctxkit.WithDevice(ctx, req.Header.Get("Device"))     // 注入 用户设备  iso、android、web
			ctx = ctxkit.WithMobiApp(ctx, req.Header.Get("MobiApp"))   // 注入 APP 标识
			ctx = ctxkit.WithVersion(ctx, req.Header.Get("Version"))   // 注入 版本 标识
			ctx = ctxkit.WithPlatform(ctx, req.Header.Get("Platform")) // 注入 平台 标识
			ctx = ctxkit.WithUserIP(ctx, req.RemoteAddr)               // TODO 注入 客户端IP 标识  目前貌似不准确待测试

			/* ------ 解析jwt token ------ */
			c, err := xjwt.CustomClaims{}.ParseToken(ctx, token)
			if err != nil {
				log.Error(ctx, "token解析失败原因：", c)
			} else {
				_ = json.Unmarshal(c.Data, &u)
				log.Info(ctx, "token 反编译：", u)
			}

			/* ------ 验证用户权限 ------ */
			resp, err := getUser(ctx, token)
			if err != nil {
				return ctx, twirp.NewError(twirp.Unauthenticated, "请先登录")
			}
			if err = json.Unmarshal(resp, &u); err != nil {
				return ctx, twirp.NewError(twirp.Unauthenticated, "请先登录")
			}
			if u.ID == 0 {
				return ctx, twirp.NewError(twirp.Unauthenticated, "请先登录")
			}
			if u.Status != 100 {
				return ctx, twirp.NewError(twirp.Unauthenticated, "抱歉您的账号已经被禁用")
			}
			if u.Company.Status != 100 {
				return ctx, twirp.NewError(twirp.Unauthenticated, "抱歉您所在的企业已经被禁用")
			}
			if u.Company.ExpiryTime <= time.Now().Unix() {
				return ctx, twirp.NewError(twirp.Unauthenticated, "抱歉您所在的企业已经过期了")
			}

			/* ------ 验证是否单一设备登录 ------ */
			if ctx, err = checkOneDevice(ctx, u.ID, token); err != nil {
				return ctx, err
			}

			/* ------ 将用户信息写入context ------ */
			ctx = ctxkit.WithUserID(ctx, u.ID)                    // 注入用户id
			ctx = ctxkit.WithUserName(ctx, u.UserName)            // 注入用户登录账号
			ctx = ctxkit.WithNickName(ctx, u.NickName)            // 注入用户昵称
			ctx = ctxkit.WithCompanyID(ctx, u.CompanyID)          // 注入公司id
			ctx = ctxkit.WithDepartmentID(ctx, u.DepartmentID)    // 注入管辖部门id
			ctx = ctxkit.WithPartIds(ctx, u.PartIds)              // 注入角色ids
			ctx = ctxkit.WithPartNames(ctx, u.PartNames)          // 注入角色names
			ctx = ctxkit.WithDepartmentIds(ctx, u.DepartmentIds)  // 注入部门ids
			ctx = ctxkit.WithManageUserIds(ctx, u.ManagerUserIds) // 注入管辖用户ids
			ctx = ctxkit.WithRolesCodes(ctx, u.RolesCodes)        // 注入权限编码

			/* ------ 将用户信息写入ResponseHeader ------ */
			_ = twirp.SetHTTPResponseHeader(ctx, "roles-codes", u.RolesCodes) // 注入权限编码
			return ctx, nil
		},
	}
}

// 验证单设备登录
func checkOneDevice(ctx context.Context, id int64, token string) (context.Context, error) {
	dType := ctxkit.GetDevice(ctx)
	if dType == "ios" || dType == "android" {
		dType = "app"
	}
	deviceCacheKey := fmt.Sprintf("user_login_%s_%d", dType, id)

	// 验证原token
	ctx, cache := memdb.Get(ctx, "DEFAULT")
	if oldToken, err := cache.Get(ctx, deviceCacheKey).Result(); err != redis.Nil && oldToken != token {
		_ = cache.Del(ctx, oldToken) // 删除老的key
		return ctx, twirp.NewError(twirp.Unauthenticated, "抱歉您的账号已在其他设备登录！")
	}

	return ctx, nil
}

// 获取用户信息
func getUser(ctx context.Context, sign string) (b []byte, err error) {
	ctx, cache := memdb.Get(ctx, "DEFAULT")

	/* ------ 获取缓存中的用户信息 ------ */
	userBody, err := cache.Get(ctx, sign).Bytes()
	if err != nil && userBody != nil {
		return userBody, nil
	}

	/* ------ 刷新用户权限，重新写入缓存 ------ */
	req, _ := http.NewRequest(
		http.MethodPost,
		conf.Get("FRAME_ADDR")+conf.Get("FRAME_REFRESH_USER_URI"),
		bytes.NewReader([]byte("")),
	)
	req.Header.Set("SIGN", sign)
	req.Header.Set("Device", ctxkit.GetDevice(ctx))     // 注入 用户设备  iso、android、web
	req.Header.Set("MobiApp", ctxkit.GetMobiApp(ctx))   // 注入 APP 标识
	req.Header.Set("Version", ctxkit.GetVersion(ctx))   // 注入 版本 标识
	req.Header.Set("Platform", ctxkit.GetPlatform(ctx)) // 注入 平台 标识

	timeout := 2 * time.Second
	if d := conf.GetDuration("INTERNAL_API_TIMEOUT"); d > 0 {
		timeout = d * time.Millisecond
	}
	resp, err := xhttp.NewClient(timeout).Do(ctx, req)
	if err != nil {
		log.Error(ctx, "请求FRAME_REFRESH_USER_URI失败：", err)
		return b, nil
	}
	if resp.StatusCode != 200 {
		return b, twirp.NewError(twirp.Unauthenticated, "请先登录")
	}
	userBody, err = cache.Get(ctx, sign).Bytes()
	return userBody, err
}
