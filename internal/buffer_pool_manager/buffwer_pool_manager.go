package buffer_pool_manager

import "github.com/takushi-m/go-simple-rdb/internal/disk_manager"

type BufferPoolManager struct {
	Disk      disk_manager.IDiskManager
	Pool      BufferPool
	PageTable map[disk_manager.PageID]BufferID
}
