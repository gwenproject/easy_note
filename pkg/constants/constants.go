package constants

const (
	NoteTableName           = "note"
	UserTableName           = "user"
	SecretKey               = "d9s8fjDISF"
	IdentityKey             = "id"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	ApiServiceName          = "apiservice"
	NoteServiceName         = "noteservice"
	UserServiceName         = "userservice"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
)

var (
	MySQLDefaultDSN = "root:root@tcp(" + GetIp("MysqlIp") + ":3306)/easy_note?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = GetIp("EtcdIp") + ":2379"
)
