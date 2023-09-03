package lru

import "container/list"

type Lru struct {
	list *list.List
	m    map[int]struct{}
}
