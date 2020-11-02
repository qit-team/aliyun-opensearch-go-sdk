package aliyun_opensearch_go_sdk


type ErrorEntity struct {
	Code      string    `json:"code,omitempty"`
	Message   string    `json:"message,omitempty"`
}

type FacetEntity struct {
	Key   string `json:"key,omitempty"`
	items []ChildItem `json:"key,omitempty"`
}

type ChildItem struct {
	Value   string `json:"value,omitempty"`
	Count   string `json:"count,omitempty"`
}

type item struct {
	Fields map[string] interface{} `json:"fields,omitempty"`
	Property map[string] interface{} `json:"property,omitempty"`
	Attribute map[string] interface{} `json:"attribute,omitempty"`
	VariableValue map[string] interface{} `json:"variable_value,omitempty"`
	SortExprValues []string `json:"sort_expr_values,omitempty"`
}

type SearchResult struct {
	Searchtime     float64    `json:"searchtime,omitempty"`
	Total          int64    `json:"total,omitempty"`
	Num            int64    `json:"num,omitempty"`
	Viewtotal      int64    `json:"viewtotal,omitempty"`
	Items          []item    `json:"items,omitempty"`
	Facet          []FacetEntity    `json:"facet,omitempty"`

	Message   string    `json:"message,omitempty"`
}

type SearchResponse struct {
	Status    string         `json:"code,omitempty"`
	RequestId string         `json:"request_id,omitempty"`
	Result    SearchResult         `json:"result,omitempty"`
	HostId    string         `json:"host_id,omitempty"`
	Errors    []ErrorEntity  `json:"errors,omitempty"`
}