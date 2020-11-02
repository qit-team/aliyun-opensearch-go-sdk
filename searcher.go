package aliyun_opensearch_go_sdk

import (
	"aliyun-opensearch-go-sdk/util"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type OpenSearcher interface {
	Name() string
	// todo: to supply params
	Search(url string, params map[string]interface{}, body interface{}, method string) []byte
	Update()
	Delete()
}

const METHOD_GET = "GET"
const METHOD_POST = "POST"

const API_VERSION = "3"
const API_TYPE = "openapi"

type AliOpenSearcher struct {
	client OpenSearchClient
	TableName       string
}

func NewAliOpenSearcher(name string, client OpenSearchClient) OpenSearcher {
	if name == "" || len(name) == 0 {
		panic(PanicInfoPrefix + "table name could not be empty")
	}

	queue := new(AliOpenSearcher)
	queue.client = client
	queue.TableName = name
	return queue
}
type StrUtil struct {
}
func (u *StrUtil) Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (p *AliOpenSearcher) Name() string {
	return p.TableName
}

func (client *AliOpenSearcher) Search(url string, params map[string]interface{}, body interface{}, method string) []byte {
	if params == nil {
		params = make(map[string]interface{})
	}
	path := "/v" + API_VERSION + "/" + API_TYPE + "/apps" + url
	url = client.client.endPoint + path

	items := map[string]interface{}{}
	items["method"] = method
	items["request_path"] = path
	items["content_type"] = "application/json"
	items["accept_language"] = "zh-cn"
	items["date"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	items["opensearch_headers"] = map[string]interface{}{}
	items["content_md5"] = ""

	rand.Seed(time.Now().Unix())
	nonce := rand.Int63n(89999) + 10000
	items["opensearch_headers"].(map[string]interface{})["X-Opensearch-Nonce"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10) + strconv.FormatInt(nonce, 10)

	if method != METHOD_GET {
		if body != nil {
			strUtil := StrUtil{}
			bodyJson, err := json.Marshal(body)
			if err != nil {
				panic(err)
			}
			items["content_md5"] = strUtil.Md5(string(bodyJson))
			items["body_json"] = bodyJson
		}
	}
	items["query_params"] = params
	signature, _ := client.client.credential.Signature(items)
	items["authorization"] = "OPENSEARCH " + client.client.accessKeyId + ":" + signature

	headers := client.getHeaders(items)
	httpClient := util.HttpClient{}
	if strings.ToLower(method) == "get" {
		respond, err := httpClient.GetWithFullData(url, params, headers)
		if err != nil {
			panic(err)
		}
		fmt.Printf("23333333%v", string(respond))
		return respond
	} else if strings.ToLower(method) == "post" {
		respond, err := httpClient.PostWithFullData(url, body, headers)
		if err != nil {
			panic(err)
		}
		return respond
	} else {
		panic(method + " does not support yet!")
	}
}

func (client *AliOpenSearcher) getHeaders(items map[string]interface{}) (headers map[string]interface{}) {
	headers = make(map[string]interface{})
	headers["Content-Type"] = items["content_type"].(string)
	headers["Date"] = items["date"].(string)
	headers["Accept-Language"] = items["accept_language"].(string)
	headers["Content-Md5"] = items["content_md5"].(string)
	headers["Authorization"] = items["authorization"].(string)
	if opensearchHeaders, ok := items["opensearch_headers"]; ok {
		for k, v := range opensearchHeaders.(map[string]interface{}) {
			headers[k] = v.(string)
		}
	}
	return
}

func (p *AliOpenSearcher) Update() {
	//tempHeader := map[string]string{
	//	"test": "1",
	//}
	//p.client.Send(GET, tempHeader, "test", "test")
}

func (p *AliOpenSearcher) Delete() {
	//tempHeader := map[string]string{
	//	"test": "1",
	//}
	//p.client.Send(GET, tempHeader, "test", "test")
}
