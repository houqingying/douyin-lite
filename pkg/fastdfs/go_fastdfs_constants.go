package fastdfs

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
	GroupName     = "tiktok"
	ServerAddress = "http://192.168.20.56:8085"
	ShowAddress   = ""
)

const (
	Account         = "root"
	Password        = "d06d6575eb571f01e15ff3e077098ae1"
	Name            = "root"
	CredentialsSalt = "f40bc3eccfa4e9985e5298be1254001"
)

const (
	Scene = "video"
	Path  = "video"
)
