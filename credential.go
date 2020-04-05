package aliyun_opensearch_go_sdk

import (
	"aliyun-opensearch-go-sdk/util"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"strings"
)

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
	Signature(items map[string]interface{}) (signature string, err error)
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
func (p *AliOsCredential) Signature(items map[string]interface{}) (signature string, err error) {
	params := map[string]interface{}{}
	if _, ok := items["query_params"]; ok {
		params = items["query_params"].(map[string]interface{})
	}
	// 定义加密内容
	signContent := ""
	signContent += strings.ToUpper(items["method"].(string)) + "\n"
	signContent += items["content_md5"].(string) + "\n"
	signContent += items["content_type"].(string) + "\n"
	signContent += items["date"].(string) + "\n"

	// 构建专有header内容
	xHeaders := p.filter(items["opensearch_headers"].(map[string]interface{}))
	if len(xHeaders) == 0 {
		panic("CanonicalizedOpenSearchHeaders is empty!")
	}
	for k, v := range xHeaders {
		signContent += strings.ToLower(k) + ":" + v.(string) + "\n"
	}

	// 构建canonicalizedResource方法
	resource := url.QueryEscape(items["request_path"].(string))
	resource = strings.ReplaceAll(resource, "%2F", "/")
	sortParams := p.filter(params)
	queryString := p.buildQuery(sortParams)
	canonicalizedResource := resource
	if queryString != "" {
		canonicalizedResource += "?" + queryString
	}
	signContent += canonicalizedResource

	// 执行加密
	key := p.accessKeySecret
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(signContent))
	signResult := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return signResult, nil
}

func (p *AliOsCredential)filter(params map[string]interface{}) map[string]interface{} {
	newParams := map[string]interface{}{}

	if params != nil && len(params) > 0 {
		for k, v := range params {
			if k == "Signature" || v == "" || v == nil {
				continue
			} else {
				newParams[k] = v
			}
		}
		util := util.Util{}
		return util.KSort(newParams)
	} else {
		return newParams
	}
}

func (p *AliOsCredential) buildQuery(params map[string]interface{}) string {
	values := url.Values{}
	for k, v := range params {
		values.Add(k, v.(string))
	}
	return strings.ReplaceAll(values.Encode(), "+", "%20")
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
