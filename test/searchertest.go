package main

import (
	aliyun_opensearch_go_sdk "aliyun-opensearch-go-sdk"
	"aliyun-opensearch-go-sdk/util"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

var literalTypeField = []string{
	"user_idnum", "user_email", "user_mobile", "platform",
}

var timestampField = []string{
	"created_at", "updated_at", "next_time",
}

func main() {
	params := map[string]interface{}{
		"filter":[]interface{}{
			map[string]interface{}{
				"name":  "reponse_time",
				"op":    ">=",
				"value": "1",
			},
			map[string]interface{}{
					"name":  "reponse_time",
					"op":    "<=",
					"value": "20000",
			},
			map[string]interface{}{
					"name":  "is_offline",
					"op":    "=",
					"value": "1",
			},
		},
		"config":map[string]interface{}{
			"page":"1",
			"per_page":"10",
		},
	}


	OpenSearchClient :=  aliyun_opensearch_go_sdk.NewAliOpenSearchClient("","","","","")
	OpenSearchClient.SetTable("call_records");




	searchQuery := buildQuery(params)
	url := "/" + OpenSearchClient.AppName+ "/search"
	searchParams := map[string]interface{}{
		"query": searchQuery,
	}

	respond := OpenSearchClient.Searcher.Search(url, searchParams, "", "get")
	data := map[string]interface{}{}
	json.Unmarshal(respond, &data)
}



func buildQuery(params map[string]interface{}) string {
	start := 0
	hit := 20
	mapUtil := util.Util{}
	var searchQuery string
	config := ""
	if !mapUtil.Empty(params, "config") {
		configParams := params["config"].(map[string]interface{})
		if !mapUtil.Empty(configParams, "page") {
			perPage, err := strconv.Atoi(configParams["per_page"].(string))
			if err != nil {
				panic(err)
			}
			hit = perPage

			page, err := strconv.Atoi(configParams["page"].(string))
			if err != nil {
				panic(err)
			}
			start = (page - 1) * hit
		}

		config += "start:" + strconv.Itoa(start) + ",hit:" + strconv.Itoa(hit) + ",format:fulljson"
	} else {
		config += "hit:500" + ",format:fulljson"
	}
	searchQuery = "config=" + config

	if mapUtil.Empty(params, "keywords") && mapUtil.Empty(params, "task_id") {
		searchQuery += "&&query=default:'客服系统'"
	} else {
		// 根据任务id查询
		if !mapUtil.Empty(params, "task_id") {
			taskId := params["task_id"].(string)
			searchQuery += "&&query=id:'" + taskId + "'"
		}
		// 根据关键字查询
		if !mapUtil.Empty(params, "keywords") {
			keywords := params["keywords"].(string)
			if strings.Contains(searchQuery, "query=") {
				searchQuery += " AND default:'" + keywords + "'"
			} else {
				searchQuery += "&&query=default:'" + keywords + "'"
			}
		}
	}

	strUtil := util.StrUtil{}
	if !mapUtil.Empty(params, "index") {
		indexParams := params["index"].([]interface{})
		indexQuery := ""
		for _, v := range indexParams {
			newValue := v.(map[string]interface{})
			name := newValue["name"].(string)
			//op := newValue["op"].(string)
			value := ""
			if reflect.TypeOf(newValue["value"]).String() == "float64" {
				value = strconv.FormatFloat(newValue["value"].(float64), 'f', -1, 64)
			} else {
				value = newValue["value"].(string)
			}

			indexQuery += name + ":'" + value + "' AND "
		}
		indexQueryLength := utf8.RuneCountInString(indexQuery)
		if indexQueryLength > 0 {
			validQuery := strUtil.Substr(indexQuery, 0, indexQueryLength-5)
			searchQuery += " AND " + validQuery
		}
	}

	if !mapUtil.Empty(params, "sort") {
		sortConfig := params["sort"].(map[string]interface{})
		searchQuery += "&&sort="
		for key, value := range sortConfig {
			searchQuery += value.(string) + key + ";"
		}
		searchQuery = strings.Trim(searchQuery, ";")
	} else {
		// 默认按照创建时间倒排
		searchQuery += "&&sort=-created_at"
	}

	if !mapUtil.Empty(params, "filter") {
		filter := ""
		filterParams := params["filter"].([]interface{})
		for _, value := range filterParams {

			newValue := value.(map[string]interface{})

			if _, ok := newValue["query"]; ok {
				filter += (newValue["query"]).(string) + " AND "
				continue
			}

			op := newValue["op"].(string)
			name := newValue["name"].(string)

			if (op == "in" || op == "notin") && len(newValue["value"].([]interface{})) > 0 {
				var strValues []string
				for _, v := range newValue["value"].([]interface{}) {
					strValues = append(strValues, v.(string))
				}
				filter += op + "(" + name + ",\"" + strings.Join(strValues, "|") + "\") AND "
				continue
			}

			arrUtil := util.ArrayUtil{}
			if arrUtil.InArray(literalTypeField, newValue["name"].(string)) {
				// 数据类型兼容
				tempValue := ""
				if reflect.TypeOf(newValue["value"]).String() == "float64" {
					tempValue = strconv.FormatFloat(newValue["value"].(float64), 'f', -1, 64)
				} else {
					tempValue = newValue["value"].(string)
				}
				filter += name + op + "\"" + tempValue + "\" AND "

			} else if arrUtil.InArray(timestampField, newValue["name"].(string)) {
				timestamp := util.StrToTimestamp(newValue["value"].(string))
				filter += name + op + strconv.FormatInt(timestamp, 10) + " AND "
			} else {
				// 数据类型兼容
				tempValue := ""
				if reflect.TypeOf(newValue["value"]).String() == "float64" {
					tempValue = strconv.FormatFloat(newValue["value"].(float64), 'f', -1, 64)
				} else {
					tempValue = newValue["value"].(string)
				}
				filter += name + op + tempValue + " AND "
			}
		}
		filterLength := utf8.RuneCountInString(filter)
		if filterLength > 0 {
			filter = strUtil.Substr(filter, 0, filterLength-5)
			searchQuery += "&&filter=" + filter
		}
	}

	if !mapUtil.Empty(params, "aggregate") {
		aggregate := ""
		aggregateParams := params["aggregate"].([]interface{})
		if len(aggregateParams) > 0 {
			for _, v := range aggregateParams {
				item := v.(map[string]interface{})
				aggregate += "group_key:" + item["group_key"].(string) + ",agg_fun:" + item["agg_fun"].(string) + ";"
			}
		}
		aggregateLength := utf8.RuneCountInString(aggregate)
		if aggregateLength > 0 {
			searchQuery += "&&aggregate=" + strUtil.Substr(aggregate, 0, aggregateLength-1)
		}
	}

	return searchQuery
}