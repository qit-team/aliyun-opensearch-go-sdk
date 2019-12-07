package aliyun_opensearch_go_sdk

const (
	Authorization = "Authorization"
	OsVersion     = "" // todo : to supply
	ContentType   = "Content-Type"
	ContentMd5    = "Content-MD5"
	HOST          = "Host"
	DATE          = "Date"
	KeepAlive     = "Keep-Alive"
	SecurityToken = "security-token"
)

type Credential interface {
	Signature(method Method, headers map[string]string, resource string) (signature string, err error)
	SetSecretKey(accessKeySecret string)
	GetAccessKeyId() string
	GetAccessKeySecret() string
	GetSecurityToken() string
}

type AliOsCredential struct {
	accessKeyId     string
	accessKeySecret string
	securityToken   string // this param can be omitted?
}

func NewAliOsCredential(accessKeyId, accessKeySecret, securityToken string) *AliOsCredential {

	if accessKeyId == "" || len(accessKeyId) == 0 {
		panic(PanicInfoPrefix + "access key id is empty")
	}

	if accessKeySecret == "" || len(accessKeySecret) == 0 {
		panic(PanicInfoPrefix + "access key secret is empty")
	}

	aliOsCredential := new(AliOsCredential)
	aliOsCredential.accessKeyId = accessKeyId
	aliOsCredential.accessKeySecret = accessKeySecret
	aliOsCredential.securityToken = securityToken
	return aliOsCredential
}

// todo supply Signature algorithm
func (p *AliOsCredential) Signature(method Method, headers map[string]string, resource string) (signature string, err error) {
	return
}

func (p *AliOsCredential) SetSecretKey(accessKeySecret string) {
	p.accessKeySecret = accessKeySecret
}

func (p *AliOsCredential) GetAccessKeyId() string {
	return p.accessKeyId
}

func (p *AliOsCredential) GetAccessKeySecret() string {
	return p.accessKeySecret
}

func (p *AliOsCredential) GetSecurityToken() string {
	return p.securityToken
}
