// Package ctxkit 操作请求 ctx 信息
package ctxkit

import (
	"context"
)

type key int

const (
	// TraceIDKey 请求唯一标识，类型：string
	TraceIDKey key = iota
	// StartTimeKey 请求开始时间，类型：time.Time
	StartTimeKey
	// UserIDKey 用户 ID，未登录则为 0，类型：int64
	UserIDKey
	// UserNameKey 用户 名称，类型：string
	UserNameKey
	// NickNameKey 用户 昵称，类型：string
	NickNameKey
	// CompanyIDKey 公司ID，未登录则为 0，类型：int64
	CompanyIDKey
	// DepartmentIDKey 获取当前管辖部门 ID，类型：int64
	DepartmentIDKey
	// DepartmentIdsKey 获取当前部门 IDs 英文逗号隔开，类型：string
	DepartmentIdsKey
	// PartIdsKey 获取当前角色 IDs 英文逗号隔开，类型：string
	PartIdsKey
	// UserIPKey 用户 IP，类型：string
	UserIPKey
	// PlatformKey 用户使用平台，ios, android, pc
	PlatformKey
	// BuildKey 客户端构建版本号
	BuildKey
	// VersionKey 客户端版本号
	VersionKey
	// AccessKey 移动端支付令牌
	AccessKey
	// DeviceKey 移动 app 设备标识，ios, android, phone, pad
	DeviceKey
	// MobiAppKey 移动 app 标识，ios, android, phone, pad
	MobiAppKey
	// UserPortKey 用户端口
	UserPortKey
	// ManageUserKey 管理后台用户名
	ManageUserKey
	// BuvidKey 非登录用户标识
	BuvidKey
	// CookieKey web 用户登录令牌
	CookieKey
	// CompanyAppKeyKey
	CompanyAppKeyKey
	// AppKeyKey 接口签名标识
	AppKeyKey
	// TsKey 时间戳
	TSKey
	// SignKey 签名
	SignKey
	// IsValidSignKeyKey 签名正确则置为 true
	IsValidSignKeyKey
)

// GetUserID 获取当前登录用户 ID
func GetUserID(ctx context.Context) int64 {
	uid, _ := ctx.Value(UserIDKey).(int64)
	return uid
}

// WithUserID 注入当前登录用户 ID
func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetUserName 获取当前登录用户 Name
func GetUserName(ctx context.Context) (uName string) {

	if uName, _ = ctx.Value(UserNameKey).(string); uName == "" {
		uName = "sys"
	}

	return uName
}

// WithUserName 注入当前登录用户 Name
func WithUserName(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, UserNameKey, userName)
}

// GetNickName 获取当前登录用户 昵称
func GetNickName(ctx context.Context) (nickName string) {

	if nickName, _ = ctx.Value(NickNameKey).(string); nickName == "" {
		nickName = "sys"
	}

	return nickName
}

// WithNickName 注入当前登录用户 昵称
func WithNickName(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, NickNameKey, userName)
}

// GetCompanyID 获取当前公司 ID
func GetCompanyID(ctx context.Context) int64 {
	uid, _ := ctx.Value(CompanyIDKey).(int64)
	return uid
}

// WithCompanyID 注入当前公司 ID
func WithCompanyID(ctx context.Context, companyID int64) context.Context {
	return context.WithValue(ctx, CompanyIDKey, companyID)
}

// GetDepartmentID 获取当前管辖部门 ID
func GetDepartmentID(ctx context.Context) int64 {
	departmentID, _ := ctx.Value(DepartmentIDKey).(int64)
	return departmentID
}

// WithDepartmentID 注入当前管辖部门 ID
func WithDepartmentID(ctx context.Context, departmentID int64) context.Context {
	return context.WithValue(ctx, DepartmentIDKey, departmentID)
}

// GetDepartmentIds 获取当前部门 ID 英文逗号隔开
func GetDepartmentIds(ctx context.Context) string {
	uid, _ := ctx.Value(DepartmentIdsKey).(string)
	return uid
}

// WithDepartmentIdS 注入当前管辖部门 ID 英文逗号隔开
func WithDepartmentIds(ctx context.Context, departmentID string) context.Context {
	return context.WithValue(ctx, DepartmentIdsKey, departmentID)
}

// GetPartIds 获取当前角色 ID 英文逗号隔开
func GetPartIds(ctx context.Context) string {
	uid, _ := ctx.Value(PartIdsKey).(string)
	return uid
}

// WithPartIds 注入当前角色 ID 英文逗号隔开
func WithPartIds(ctx context.Context, partIds string) context.Context {
	return context.WithValue(ctx, PartIdsKey, partIds)
}

// GetUserIP 获取用户 IP
func GetUserIP(ctx context.Context) string {
	ip, _ := ctx.Value(UserIPKey).(string)
	return ip
}
func WithUserIP(ctx context.Context, userIP string) context.Context {
	return context.WithValue(ctx, UserIPKey, userIP)
}

// GetUserPort 获取用户端口
func GetUserPort(ctx context.Context) string {
	port, _ := ctx.Value(UserPortKey).(string)
	return port
}

// GetPlatform 获取用户平台
func GetPlatform(ctx context.Context) string {
	platform, _ := ctx.Value(PlatformKey).(string)
	return platform
}

func WithPlatform(ctx context.Context, platform string) context.Context {
	return context.WithValue(ctx, PlatformKey, platform)
}

// GetCompanyAppKey 获取用户平台
func GetCompanyAppKey(ctx context.Context) string {
	companyKey, _ := ctx.Value(CompanyAppKeyKey).(string)
	return companyKey
}

// IsIOSPlatform 判断是否为 IOS 平台
func IsIOSPlatform(ctx context.Context) bool {
	return GetPlatform(ctx) == "ios"
}

// GetTraceID 获取用户请求标识
func GetTraceID(ctx context.Context) string {
	id, _ := ctx.Value(TraceIDKey).(string)
	return id
}

// GetBuild 获取客户端构建版本号
func GetBuild(ctx context.Context) string {
	build, _ := ctx.Value(BuildKey).(string)
	return build
}

// GetDevice 获取用户设备，配合 GetPlatform 使用
func GetDevice(ctx context.Context) string {
	device, _ := ctx.Value(DeviceKey).(string)
	return device
}

// WithDevice 注入 Device 标识
func WithDevice(ctx context.Context, deviceKey string) context.Context {
	return context.WithValue(ctx, DeviceKey, deviceKey)
}

// GetMobiApp 获取 APP 标识
func GetMobiApp(ctx context.Context) string {
	app, _ := ctx.Value(MobiAppKey).(string)
	return app
}

// WithMobiApp 获取 APP 标识
func WithMobiApp(ctx context.Context, mobiAppKey string) context.Context {
	return context.WithValue(ctx, MobiAppKey, mobiAppKey)
}

// GetVersion 获取客户端版本
func GetVersion(ctx context.Context) string {
	version, _ := ctx.Value(VersionKey).(string)
	return version
}
func WithVersion(ctx context.Context, version string) context.Context {
	return context.WithValue(ctx, VersionKey, version)
}

// GetAccessKey 获取客户端认证令牌
func GetAccessKey(ctx context.Context) string {
	key, _ := ctx.Value(AccessKey).(string)
	return key
}

// WithAccessKey 注入客户端认证令牌
func WithAccessKey(ctx context.Context, accessKey string) context.Context {
	return context.WithValue(ctx, AccessKey, accessKey)
}

// WithTraceID 注入 trace_id
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

// GetManageUser 获取管理后台用户名
func GetManageUser(ctx context.Context) string {
	user, _ := ctx.Value(ManageUserKey).(string)
	return user
}

// GetBuvid 获取用户 buvid
func GetBuvid(ctx context.Context) string {
	buvid, _ := ctx.Value(BuvidKey).(string)
	return buvid
}

// GetCookie 获取 web cookie
func GetCookie(ctx context.Context) string {
	key, _ := ctx.Value(CookieKey).(string)
	return key
}

func GetSign(ctx context.Context) (appkey, ts, sign string) {
	appkey, _ = ctx.Value(AppKeyKey).(string)
	ts, _ = ctx.Value(TSKey).(string)
	sign, _ = ctx.Value(SignKey).(string)
	return
}

// WithAccessKey 注入签名认证令牌
func WithSignKey(ctx context.Context, signKey string) context.Context {
	return context.WithValue(ctx, SignKey, signKey)
}

// IsValidSignKey 判断业务签名是否正确
func IsValidSignKey(ctx context.Context) bool {
	valid, _ := ctx.Value(IsValidSignKeyKey).(bool)
	return valid
}
