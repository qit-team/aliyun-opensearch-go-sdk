package aliyun_opensearch_go_sdk

import (
	"github.com/gogap/errors"
)

const (
	ALIYUN_OPENSEARCH_ERR_NS = "OPENSEARCH"

	ALIYUN_OPENSEARCH_ERR_TEMPLATE = "aliyun_opensearch response status error,code: {{.resp.Code}}, message: {{.resp.Message}}, resource: {{.resource}}"
)

var (
	ERR_MNS_ACCESS_DENIED                = errors.TN(ALIYUN_OPENSEARCH_ERR_NS, 100, ALIYUN_OPENSEARCH_ERR_TEMPLATE)
)
