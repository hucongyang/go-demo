package cache_service

import (
	"github.com/hucongyang/go-demo/pkg/cache"
	"strconv"
	"strings"
)

type Tag struct {
	ID    int
	Name  string
	State int

	PageNum  int
	PageSize int
}

// 获取tag的cache key
func (tag *Tag) GetTagKey() string {
	keys := []string{
		cache.Tag,
		"LIST",
	}

	if tag.Name != "" {
		keys = append(keys, tag.Name)
	}
	if tag.State >= 0 {
		keys = append(keys, strconv.Itoa(tag.State))
	}
	if tag.PageNum > 0 {
		keys = append(keys, strconv.Itoa(tag.PageNum))
	}
	if tag.PageSize > 0 {
		keys = append(keys, strconv.Itoa(tag.PageSize))
	}
	return strings.Join(keys, "_")
}
