package configs

import fastDFS "douyin-lite/pkg/fastdfs"

// 错误消息常量
const (
	ErrNoFileUploaded     = "未上传文件"
	ErrCreateTmpDirectory = "创建临时目录失败"
	ErrSaveTmpFile        = "保存临时文件失败"
	ErrGetFastDFSURL      = "获取FastDFS URL失败"
	ErrUploadToDFS        = "上传到DFS失败"
	ErrDeleteTmpFile      = "删除临时文件失败"
)

const (
	TmpFileDir = "tmp" // 临时存储目录
	FrameNum   = 1     // 设置视频帧数作为封面图片
)

const (
	GroupName       = "tiktok"                           // 组名
	ServerAddress   = "http://192.168.20.56:8085"        // 服务器地址
	ShowAddress     = ""                                 // 显示地址
	Account         = "root"                             // 账户名
	Password        = "d06d6575eb571f01e15ff3e077098ae1" // 密码
	Name            = "root"                             // 名称
	CredentialsSalt = "f40bc3eccfa4e9985e5298be1254001"  // 凭据盐值
)

func FastDFSInit() {
	fastDFS.NewFDClient(
		GroupName,
		ServerAddress,
		ShowAddress,
		Account,
		Password,
		Name,
		CredentialsSalt,
	)
}
