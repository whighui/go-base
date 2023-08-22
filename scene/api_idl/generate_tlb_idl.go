package idl

import (
	"code.byted.org/aweme_privacy/privacy_common/util/metric"
	"code.byted.org/gopkg/env"
	"code.byted.org/gopkg/logs"
	"code.byted.org/toutiao/api_risk_detected_dal/model/bo"
	"code.byted.org/toutiao/api_risk_detected_dal/model/do"
	"code.byted.org/toutiao/api_risk_detected_dal/model/sql"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

//对tlb流量生成idl
//如果是没有生成idl 进行第一次生成并进行metric打点
//				  如果生成idl 并存在diff 进行metric打点

func GenerateTlbApiIdl(ctx context.Context, detectedApi *bo.DetectAPI) {
	if detectedApi == nil {
		return
	}
	apiIdlDB := sql.WriteDB.BytetreeAPIIdl
	psm, path, pathID, query, payload, respBody := detectedApi.Psm, detectedApi.Path, detectedApi.Info.PathId, detectedApi.Query, detectedApi.Payload, detectedApi.Content
	idlData, _ := apiIdlDB.WithContext(ctx).Where(apiIdlDB.PathID.Eq(pathID)).Find()                      //根据PathId来判断是否存在Idl
	metric.EmitCounter(env.PSM(), "FilterTLBDataIDLTraffic", map[string]string{"psm": psm, "path": path}) //计算流量QPS
	if len(idlData) == 0 {                                                                                //没有生成idl  进行生成并存储 打点
		idl := Serialize(ctx, JsonConvIdl(ctx, mapConvString(query), payload, respBody))
		apiIdlDB.WithContext(ctx).Create(&do.BytetreeAPIIdl{
			PathID:     pathID,
			Idl:        &idl,
			CreateTime: time.Now(),
		})
		logs.CtxInfo(ctx, "入库idl->pathID:", pathID)
		metric.EmitCounter(env.PSM(), "FilterTLBDataIDLGen", map[string]string{"psm": psm, "path": path, "pathID": strconv.Itoa(int(pathID))})
	} else { //已经生成idl  判断是否存在diff
		idl := idlData[0].Idl
		if idl == nil {
			return
		}
		baseIdl := Deserializes(ctx, *idl)
		detectedIdl := JsonConvIdl(ctx, mapConvString(query), payload, respBody)
		if DiffIdlAnalyse(baseIdl, detectedIdl) {
			//如果存在diff打点 并进行merge
			mergeIdl := MergeHistoryIdl(baseIdl, detectedIdl)
			mergeIdlStr := Serialize(ctx, mergeIdl)
			apiIdlDB.WithContext(ctx).Where(apiIdlDB.PathID.Eq(pathID)).Update(apiIdlDB.Idl, mergeIdlStr)       //进行更新操作
			apiIdlDB.WithContext(ctx).Where(apiIdlDB.PathID.Eq(pathID)).Update(apiIdlDB.Version, "0.01")        //如果判断diff成功进行更新一次版本
			apiIdlDB.WithContext(ctx).Where(apiIdlDB.PathID.Eq(pathID)).Update(apiIdlDB.CreateTime, time.Now()) //如果存在merge进行更新create_time
			metric.EmitCounter(env.PSM(), "FilterTLBDataIDLDiff", map[string]string{"psm": psm, "path": path, "pathID": strconv.Itoa(int(pathID))})
		}
	}
}

func mapConvString(m map[string][]string) string {
	data, _ := json.Marshal(m)
	return string(data)
}
