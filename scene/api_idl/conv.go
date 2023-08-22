package idl

import (
	"context"
	"encoding/json"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"unicode"
)

const (
	RootTag     = "root"
	QueryTag    = "query"
	ReqBodyTag  = "req_body"
	RespBodyTag = "resp_body"

	RootLayer       = 0
	QueryFirstLayer = 1
	BodyFirstLayer  = 1

	PrefixArray = "[]"
	UnknownType = "Unknown"

	UserId       = "user_id" //判断特殊情况的key 全部由数字组成
	UserIdLength = 16
	UserIdType   = "int64"

	CNType = "HAN"
)

var paramLayerMap map[string]int //存放当前IDL所有参数 层级

func JsonConvIdl(ctx context.Context, query, reqBody, respBody string) *TreeNode {
	//if !preJudgeBody(reqBody) && !preJudgeBody(respBody) {
	//	return nil
	//} 如果这个条件限制了 那么就是会返回nil 后续不好判断
	root := &TreeNode{
		Info: &NodeInfo{
			Tag:   RootTag,
			Layer: RootLayer,
		},
		Child: make([]*TreeNode, 3, 4),
	}
	root.Child[0] = parseQuery(ctx, query, QueryTag)
	root.Child[1] = parseBody(ctx, reqBody, ReqBodyTag)
	root.Child[2] = parseBody(ctx, respBody, RespBodyTag)
	return root
}

// 这里边如果body穿空回怎么样
func parseQuery(ctx context.Context, query, queryTag string) *TreeNode {
	queryRoot := &TreeNode{
		Info: &NodeInfo{
			Tag:   queryTag,
			Layer: QueryFirstLayer,
		},
		Child: make([]*TreeNode, 0),
	}
	if query == "" || query == "{}" {
		return queryRoot
	}

	kvs := make(map[string]interface{})
	err := json.Unmarshal([]byte(query), &kvs)
	if err != nil {
		return queryRoot //这里即便解析json 也是要返回节点
	}

	//list := make([]string, 0, len(kvs))  //后边进行排序 在这里就是先不用进行排序
	//for key, _ := range kvs {
	//	list = append(list, key)
	//}
	//sort.Strings(list)

	queryRoot.Child = make([]*TreeNode, 0, len(kvs))
	childs := make([]*TreeNode, 0, len(kvs))
	for paramKey, _ := range kvs {

		paramType := paramTypeJudge(kvs[paramKey])
		if isUserIdCase(paramKey) {
			paramKey = UserId
			paramType = UserIdType
		}
		if isChineseCase(paramKey) {
			paramKey = CNType
		}
		if paramType == "slice" { //数组形式进行额外判断
			value := kvs[paramKey]
			val := value.([]interface{})
			if len(val) != 0 {
				//判断数组第一个位置 进行递归下去 返回paramType以及map
				paramType, _ = judgeSliceType(val[0], PrefixArray)
			} else {
				paramType = PrefixArray + UnknownType
			}
		}
		childs = append(childs, &TreeNode{
			Info: &NodeInfo{
				Tag:       queryTag,
				ParamKey:  paramKey,
				ParamType: paramType,
				Layer:     queryRoot.Info.Layer + 1,
			},
		})
	}
	queryRoot.Child = SortSliceNodeInfo(childs) //对一层key 进行去重和排序

	return queryRoot
}

func parseBody(ctx context.Context, body, bodyTag string) *TreeNode {
	bodyRoot := &TreeNode{
		Info: &NodeInfo{
			Tag:   bodyTag,
			Layer: BodyFirstLayer,
		},
		Child: make([]*TreeNode, 0), //这里边child先默认有值  否则后续再进行merge判断的时候不好进行判断
	}

	//前置判断
	if body == "" || body == "{}" {
		return bodyRoot
	}

	kvs := make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &kvs)
	if err != nil {
		return bodyRoot
	}
	recursiveParseBody(ctx, bodyRoot, kvs, bodyTag)
	return bodyRoot
}

func recursiveParseBody(ctx context.Context, root *TreeNode, kvs map[string]interface{}, tag string) {
	if len(kvs) == 0 {
		return
	}
	root.Child = make([]*TreeNode, len(kvs), len(kvs))
	index := -1
	for paramKey, _ := range kvs {
		index++
		value := kvs[paramKey]
		paramType := paramTypeJudge(value)

		if isUserIdCase(paramKey) { //加上特殊判断
			paramKey = UserId
			paramType = UserIdType
		}
		if isChineseCase(paramKey) {
			paramKey = CNType
		}
		root.Child[index] = &TreeNode{
			Info: &NodeInfo{
				Tag:       tag,
				ParamKey:  paramKey,
				ParamType: paramType,
				Layer:     root.Info.Layer + 1,
			},
		}
		if paramType == "map" { //map类型继续进行递归下去
			val := value.(map[string]interface{})
			if len(val) != 0 {
				recursiveParseBody(ctx, root.Child[index], val, tag)
			}
		} else if paramType == "slice" { //slice类型判断有没有map
			val := value.([]interface{})
			if len(val) != 0 {
				//判断数组第一个位置 进行递归下去 返回paramType以及map
				paramTypeSlice, m := judgeSliceType(val[0], PrefixArray)
				root.Child[index].Info.ParamType = paramTypeSlice
				if m != nil {
					recursiveParseBody(ctx, root.Child[index], m, tag)
				}
			} else {
				root.Child[index].Info.ParamType = PrefixArray + UnknownType
			}
		}
	}
	root.Child = SortSliceNodeInfo(root.Child) //进行排序和去重
}

// Kind的分类只有少数的几种：基础类型Bool、String以及各种数字类型；聚合类型Array和Struct；引用类型Chan、Func、Ptr、Slice和Map、接口类型Interface；
// 最后还有Invalid类型，表示它们还没有任何值。
func paramTypeJudge(val interface{}) string {
	if val == nil {
		return UnknownType
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.Int:
		return "int"

	case reflect.Int8:
		return "int8"

	case reflect.Int16:
		return "int16"

	case reflect.Int32:
		return "int32"

	case reflect.Int64:
		return "int64"

	case reflect.Uint:
		return "uint"

	case reflect.Uint8:
		return "uint8"

	case reflect.Uint16:
		return "uint16"

	case reflect.Uint32:
		return "uint32"

	case reflect.Uint64:
		return "uint64"

	case reflect.Float32:
		return "float32"

	case reflect.Float64:
		return "float64"

	case reflect.Complex64:
		return "complex64"

	case reflect.Complex128:
		return "complex128"

	case reflect.Map:
		return "map"

	case reflect.Slice:
		return "slice"

	case reflect.String:
		return "string"
	}
	return UnknownType
}

func preJudgeBody(body string) bool {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &m)
	if err != nil {
		return false
	}
	return true
}

func sortMapKey(m map[string]interface{}) []string {
	list := make([]string, 0, len(m))
	for key, _ := range m {
		list = append(list, key)
	}
	sort.Strings(list)
	return list
}

//judgeSliceType 只判断数组第一个位置  如果一直是数组返回paramType

func judgeSliceType(val interface{}, prefix string) (paramType string, kvs map[string]interface{}) {

	if val == nil {
		return UnknownType, nil
	}

	switch reflect.TypeOf(val).Kind() {

	case reflect.Map:
		return prefix + "map", val.(map[string]interface{})

	case reflect.Slice:
		valSlice := val.([]interface{})
		if len(valSlice) != 0 {
			return judgeSliceType(valSlice[0], prefix+PrefixArray)
		} else {
			return prefix + UnknownType, nil
		}

	case reflect.String:
		return prefix + "string", nil

	case reflect.Int:
		return prefix + "int", nil

	case reflect.Int8:
		return prefix + "int8", nil

	case reflect.Int16:
		return prefix + "int16", nil

	case reflect.Int32:
		return prefix + "int32", nil

	case reflect.Int64:
		return prefix + "int64", nil

	case reflect.Uint:
		return prefix + "uint", nil

	case reflect.Uint8:
		return prefix + "uint8", nil

	case reflect.Uint16:
		return prefix + "uint16", nil

	case reflect.Uint32:
		return prefix + "uint32", nil

	case reflect.Uint64:
		return prefix + "uint64", nil

	case reflect.Float32:
		return prefix + "float32", nil

	case reflect.Float64:
		return prefix + "float64", nil

	case reflect.Complex64:
		return prefix + "complex64", nil

	case reflect.Complex128:
		return prefix + "complex128", nil

	default: //防止数组里边没有任何值 走到这里来判断
		return UnknownType, nil
	}
}

func isUserIdCase(paramKey string) bool {
	_, err := strconv.ParseFloat(paramKey, 64)
	if err == nil && len(paramKey) == UserIdLength {
		return true
	}
	return false
}

// isChineseCase  判断参数是不是包含中文字符串
func isChineseCase(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
