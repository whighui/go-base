package idl

import (
	"context"
	"encoding/json"
)

func Serialize(ctx context.Context, root *TreeNode) string {
	if root == nil {
		return ""
	}
	bytes, err := json.Marshal(root)
	if err != nil {
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
	}
	return &treeNode
}
