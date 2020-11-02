package aliyun_opensearch_go_sdk

import (
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

type Method string

const (
	GET    Method = "GET"
	PUT           = "PUT"
	POST          = "POST"
	DELETE        = "DELETE" // can be omitted?
)

const (
	PanicInfoPrefix = "opensearch-go-sdk-panic:"
)

const (
	DefaultTimeout int64 = 60
)

const (
	version = "2019-12-05"
)

// todo check the Criterion num
const CriterionParamNumOfEndPoint = 5

//type OpenSearchClient interface {
//	Send(method Method, headers map[string]string, message interface{}, resource string) (*fasthttp.Response, error)
//}

// opensearch client基本信息
type OpenSearchClient struct {
	credential   Credential
	Searcher     OpenSearcher
	Timeout      int64
	endPoint     string
	AppName      string
}
type aliOsClient struct {
	Timeout      int64
	endPoint     *neturl.URL
	//credential   Credential
	credential   AliOsCredential
	accessKeyId  string
	client       *fasthttp.Client
	clientLocker sync.Mutex
	regionId     string
	proxyUrl     string // todo: check this one is supported or not ？
}

func NewAliOpenSearchClient(endPoint, accessKeyId, accessKeySecret, securityToken string, AppName string) OpenSearchClient {
	if endPoint == "" || len(endPoint) == 0 {
		panic(PanicInfoPrefix + "endpoint is empty")
	}

	// parse region and other info
	//pieces := strings.Split(endPoint, ".")
	//if len(pieces) != CriterionParamNumOfEndPoint {
	//	panic(PanicInfoPrefix + "endPoint is invalid")
	//}

	// todo: get regionId, to confirm the position of regionId in endpoint
	regionId := "to do"
	credential := NewAliOsCredential(accessKeyId, accessKeySecret, securityToken)
	client := *new(OpenSearchClient)
	client.accessKeyId = accessKeyId
	client.credential = *credential
	client.regionId = regionId
	client.AppName = AppName

	client.endPoint = endPoint

	client.initFastHttpClient()

	return client
}

func (p *OpenSearchClient) SetTable(name string) {
	searcher := NewAliOpenSearcher(name, *p)
	p.Searcher = searcher
}

func (p *OpenSearchClient) initFastHttpClient() {
	p.clientLocker.Lock()
	defer p.clientLocker.Unlock()

	timeoutInt := DefaultTimeout

	if p.Timeout > 0 {
		timeoutInt = p.Timeout
	}

	timeout := time.Second * time.Duration(timeoutInt)

	p.client = &fasthttp.Client{ReadTimeout: timeout, WriteTimeout: timeout}
}
//
//// todo : rewrite?
//func (p *OpenSearchClient) Send(method Method, headers map[string]string, message interface{}, resource string) (*fasthttp.Response, error) {
//
//	var xmlContent []byte
//	var err error
//
//	if message == nil {
//		xmlContent = []byte{}
//	} else {
//		switch m := message.(type) {
//		case []byte:
//			{
//				xmlContent = m
//			}
//		default:
//			if bXml, e := xml.Marshal(message); e != nil {
//				return nil, err
//			} else {
//				xmlContent = bXml
//			}
//		}
//	}
//
//	xmlMD5 := md5.Sum(xmlContent)
//	strMd5 := fmt.Sprintf("%x", xmlMD5)
//
//	if headers == nil {
//		headers = make(map[string]string)
//	}
//
//	headers[OsVersion] = version
//	headers[ContentType] = "application/xml"
//	headers[ContentMd5] = base64.StdEncoding.EncodeToString([]byte(strMd5))
//	headers[DATE] = time.Now().UTC().Format(http.TimeFormat)
//
//	if authHeader, e := p.authorization(method, headers, fmt.Sprintf("/%s", resource)); e != nil {
//		return nil, err
//	} else {
//		headers[Authorization] = authHeader
//	}
//
//	var buffer bytes.Buffer
//	buffer.WriteString(p.endPoint)
//	buffer.WriteString("/")
//	buffer.WriteString(resource)
//
//	url := buffer.String()
//
//	req := fasthttp.AcquireRequest()
//
//	req.SetRequestURI(url)
//	req.Header.SetMethod(string(method))
//	req.SetBody(xmlContent)
//
//	for header, value := range headers {
//		req.Header.Set(header, value)
//	}
//
//	resp := fasthttp.AcquireResponse()
//
//	if err = p.client.Do(req, resp); err != nil {
//		return nil, err
//	}
//
//	return resp, nil
//}
//
//func (p *OpenSearchClient) authorization(method Method, headers map[string]string, resource string) (authHeader string, err error) {
//	if signature, e := p.credential.Signature(method, headers, resource); e != nil {
//		return "", e
//	} else {
//		authHeader = fmt.Sprintf("MNS %s:%s", p.accessKeyId, signature)
//	}
//
//	return
//}
