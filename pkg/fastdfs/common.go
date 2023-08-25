package fastDFS

const (
	ApiStatusSuccess  = "ok"
	ApiReload         = "/reload"
	ApiStat           = "/stat"
	ApiUpload         = "/upload"
	ApiDelete         = "/delete"
	ApiGetFileInfo    = "/get_file_info"
	ApiListDir        = "/list_dir"
	ApiRepairStat     = "/repair_stat"
	ApiRemoveEmptyDir = "/remove_empty_dir"
	ApiBackup         = "/backup"
	ApiRepair         = "/repair"
	ApiRepairFileInfo = "/repair_fileinfo"
	ApiStatus         = "/status"
	ApiBigUpload      = "/big/upload/"

	ContentTypeForm = "application/x-www-form-urlencoded"
	ContentTypeText = "text/plain"
	ContentTypeHtml = "text/html"
)

const (
	FileKey     = "file"   // FileKey 表示文件键的名称
	SceneKey    = "scene"  // SceneKey 表示场景键的名称
	PathKey     = "path"   // PathKey 表示路径键的名称
	OutputKey   = "output" // OutputKey 表示输出键的名称
	SceneValue  = "video"  // SceneValue 表示场景值为视频
	PathValue   = "video"  // PathValue 表示路径值为视频
	OutputValue = "json"   // OutputValue 表示输出值为 JSON 格式
)
