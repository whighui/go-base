package idl

import (
	"code.byted.org/gopkg/logs"
	"context"
	"encoding/json"
)

func Serialize(ctx context.Context, root *TreeNode) string {
	if root == nil {
		return ""
	}
	bytes, err := json.Marshal(root)
	if err != nil {
		logs.CtxError(ctx, "json.Marshal(root) err:[%v]", err)
	}
	return string(bytes)
}

func Deserializes(ctx context.Context, val string) *TreeNode {
	if val == "" {
		return nil
	}
	treeNode := TreeNode{}
	err := json.Unmarshal([]byte(val), &treeNode)
	if err != nil {
		logs.CtxError(ctx, "sonic.Unmarshal([]byte(val), &treeNode) err:[%v]", err)
	}
	return &treeNode
}
