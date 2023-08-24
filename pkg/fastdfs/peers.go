package fastdfs

// Peers peers table
type Peers struct {
	GroupName     string `gorm:"not null;type:varchar(64)"`
	ServerAddress string `gorm:"not null;type:varchar(256)"`
	ShowAddress   string `gorm:"type:varchar(256)"`
}

func NewPeers() *Peers {
	return &Peers{
		GroupName:     GroupName,
		ServerAddress: ServerAddress,
		ShowAddress:   ShowAddress,
	}
}

type PeerUser struct {
	Account         string `gorm:"not null;type:varchar(64)"`
	Password        string `gorm:"not null;type:varchar(64)"`
	Name            string `gorm:"not null;type:varchar(64)"`
	CredentialsSalt string `gorm:"not null;type:varchar(64)"`
}

func NewPeerUser() *PeerUser {
	return &PeerUser{
		Account:         Account,
		Password:        Password,
		Name:            Name,
		CredentialsSalt: CredentialsSalt,
	}
}
