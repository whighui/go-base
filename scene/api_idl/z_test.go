package idl

import (
	"code.byted.org/toutiao/api_risk_detected_dal/model/sql"
	"code.byted.org/toutiao/api_risk_detected_dal/mysql"
	"context"
	"fmt"
	"testing"
)

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func TestJsonConvertIdl(t *testing.T) {
	query := "{\"pagesize\":[\"200\"],\"__token\":[\"361ff20e1ca767e0f1950cbbe689c53b\"],\"page\":[\"0\"],\"search_type\":[\"1\"]}"
	reqBody := "{\n  \"code\": 0,\n  \"data\": [\n    {\n      \"a\": \"\",\n      \"b\": \"\"\n    }\n  ],\n  \"list\": [\n    1,\n    2,\n    3\n  ],\n  \"ext\": {},\n  \"message\": \"success\"\n}"
	respBody := "{\n  \"code\": 0,\n  \"data\": [\n    {\n      \"a\": \"\",\n      \"b\": \"\"\n    }\n  ],\n  \"list\": [\n    [\n      {\n        \"a\": \"\",\n        \"b\": \"\"\n      }\n    ]\n  ],\n  \"ext\": {},\n  \"message\": \"success\"\n}"
	tree := JsonConvIdl(context.Background(), query, reqBody, respBody)

	query1 := "{\"pagesiz1e\":[\"200\"],\"__token\":[\"361ff20e1ca767e0f1950cbbe689c53b\"],\"page\":[\"0\"],\"search_type\":[\"1\"]}"
	reqBody1 := "{\n  \"code\": 0,\n  \"data\": [\n    {\n      \"a\": \"\",\n      \"b\": \"\"\n    }\n  ],\n  \"list\": [\n    1,\n    2,\n    3\n  ],\n  \"ext\": {},\n  \"message\": \"success\"\n}"
	respBody1 := "{\n  \"code\": 0,\n  \"data\": [\n    {\n      \"a\": \"\",\n      \"b\": \"\"\n    }\n  ],\n  \"list\": [\n    [\n      {\n        \"a\": \"\",\n        \"b\": \"\"\n      }\n    ]\n  ],\n  \"ext\": {},\n  \"message\": \"success\"\n}"
	tree1 := JsonConvIdl(context.Background(), query1, reqBody1, respBody1)
	fmt.Println(DiffIdlAnalyse(tree, tree1))
}

func TestUpdateIdl(t *testing.T) {
	mysql.Init()
	sql.Init()
	apiIdl := sql.WriteDB.BytetreeAPIIdl
	str := "abc"
	apiIdl.WithContext(context.Background()).Where(apiIdl.PathID.Eq(10)).Update(apiIdl.Idl, str)
}

func TestSliceZero(t *testing.T) {
	var id []int
	fmt.Println(id == nil)
}

func TestDiff(t *testing.T) {

	baseResp := "{\"code\":0,\"msg\":\"成功\",\"data\":{\"docs\":[{\"title\":\"首页功能使用说明\",\"content\":\"简要概述新版首页功能的功能，QIC基地的首页暂时迁移至越库管理模块下方\",\"url\":\"https://bytedance.feishu.cn/docx/BNNFdhpUVo2Q9bxjgRNcrPMxnnh\",\"status\":1,\"id\":510132,\"createTime\":1670937933,\"createBy\":\"wangtao\",\"creatorName\":\"汪涛\",\"version\":1,\"updateTime\":1670937933,\"updateBy\":\"\",\"updateName\":\"\"}],\"total\":1,\"page\":1,\"pageSize\":4}}"

	detectedResp := "{\n  \"code\": 0,\n  \"msg\": \"成功\",\n  \"data\": {\n    \"total\": 0,\n    \"page\": 1,\n    \"pageSize\": 4\n  },\n  \"a\": \"\"\n}"

	baseIdl := JsonConvIdl(context.Background(), "", "", baseResp)
	detectedIdl := JsonConvIdl(context.Background(), "", "", detectedResp)
	fmt.Println(DiffIdlAnalyse(detectedIdl, baseIdl))
}
