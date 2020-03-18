package aliyun_opensearch_go_sdk

import (
	"testing"
	"time"
	"math/rand"
	"strconv"
	"fmt"
)

func TestSignature(t *testing.T)  {
	fmt.Println("begin")
	params := make(map[string]interface{})
	path := "/v3/openapi/apps/qudian/search"
	items := map[string]interface{}{}
	items["method"] = "get"
	items["request_path"] = path
	items["content_type"] = "application/json"
	items["accept_language"] = "zh-cn"
	items["date"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	items["opensearch_headers"] = map[string]interface{}{}
	items["content_md5"] = ""

	rand.Seed(time.Now().Unix())
	nonce := rand.Int63n(89999) + 10000
	items["opensearch_headers"].(map[string]interface{})["X-Opensearch-Nonce"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10) + strconv.FormatInt(nonce, 10)

	//if method != METHOD_GET {
	//	if body != nil {
	//		strUtil := StrUtil{}
	//		bodyJson, err := json.Marshal(body)
	//		if err != nil {
	//			panic(err)
	//		}
	//		items["content_md5"] = strUtil.Md5(string(bodyJson))
	//		items["body_json"] = bodyJson
	//	}
	//}
	items["query_params"] = params
	service := NewAliOsCredential("", "", "")
	signature, err := service.Signature(items)
	fmt.Println("err: ", err)
	fmt.Println("result: ", signature)
}