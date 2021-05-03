package buffer_pool_manager

import "github.com/takushi-m/go-simple-rdb/internal/constant"

type Page []byte

func NewPage() Page {
	return Page(make([]byte, constant.PageSize))
}
