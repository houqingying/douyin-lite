package fastDFS

import (
	"github.com/astaxie/beego/httplib"
)

var FDClient *FastDFSClient

type FastDFSClient struct {
	fastDFS     *FastDFS
	fastDFSUser *FastDFSUser
}

// FastDFS fastDFSs table
type FastDFS struct {
	GroupName     string `gorm:"not null;type:varchar(64)"`
	ServerAddress string `gorm:"not null;type:varchar(256)"`
	ShowAddress   string `gorm:"type:varchar(256)"`
}

type FastDFSUser struct {
	Account         string `gorm:"not null;type:varchar(64)"`
	Password        string `gorm:"not null;type:varchar(64)"`
	Name            string `gorm:"not null;type:varchar(64)"`
	CredentialsSalt string `gorm:"not null;type:varchar(64)"`
}

func NewFDClient(groupName, serverAddress, showAddress, account, password, name, credentialsSalt string) {
	FDClient = &FastDFSClient{
		fastDFS:     NewFastDFS(groupName, serverAddress, showAddress),
		fastDFSUser: NewFastDFSUser(account, password, name, credentialsSalt),
	}
}

func NewFastDFS(groupName, serverAddress, showAddress string) *FastDFS {
	return &FastDFS{
		GroupName:     groupName,
		ServerAddress: serverAddress,
		ShowAddress:   showAddress,
	}
}

func NewFastDFSUser(account, password, name, credentialsSalt string) *FastDFSUser {
	return &FastDFSUser{
		Account:         account,
		Password:        password,
		Name:            name,
		CredentialsSalt: credentialsSalt,
	}
}

// GetFastDFSUrl 获取FastDFSUrl
func (client *FastDFSClient) GetFastDFSUrl() (string, error) {
	if client.fastDFS.GroupName != "" {
		return client.fastDFS.ServerAddress + "/" + client.fastDFS.GroupName, nil
	}
	return client.fastDFS.ServerAddress, nil
}

// getShowUrl 获取ShowUrl
func (client *FastDFSClient) getShowUrl() string {
	showUrl := ""
	if client.fastDFS.ShowAddress == "" {
		if client.fastDFS.GroupName == "" {
			showUrl = client.fastDFS.ServerAddress
		} else {
			showUrl = client.fastDFS.ServerAddress + "/" + client.fastDFS.GroupName
		}
	} else {
		if client.fastDFS.GroupName == "" {
			showUrl = client.fastDFS.ShowAddress
		} else {
			showUrl = client.fastDFS.ShowAddress + "/" + client.fastDFS.GroupName
		}
	}
	return showUrl
}

// GetShowUrlNotGroup 获取ShowUrl
func (client *FastDFSClient) GetShowUrlNotGroup() string {
	showUrl := ""
	if client.fastDFS.ShowAddress == "" {
		showUrl = client.fastDFS.ServerAddress
	} else {
		showUrl = client.fastDFS.ShowAddress
	}
	return showUrl
}

func (client *FastDFSClient) UploadGoFastDFS(filePath string) (map[string]interface{}, error) {
	// 上传到dfs
	fastDFSUrl, err := client.GetFastDFSUrl()
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}
	req := httplib.Post(fastDFSUrl + ApiUpload)
	req.PostFile(FileKey, filePath)
	req.Param(OutputKey, OutputValue)
	req.Param(SceneKey, SceneValue)
	req.Param(PathKey, PathValue)
	err = req.ToJSON(&obj)
	if err != nil {
		return nil, err
	}

	// 构建文件的访问URL，并添加到返回的对象中
	obj["url"] = client.GetShowUrlNotGroup() + obj["path"].(string)
	return obj, nil
}
