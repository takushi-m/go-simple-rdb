package buffer_pool_manager

import "github.com/takushi-m/go-simple-rdb/internal/disk_manager"

type BufferID uint64

type Buffer struct {
	PageID  disk_manager.PageID
	Page    Page
	IsDirty bool
}

type Frame struct {
	UsageCount uint64
	Buffer     *Buffer
}

type BufferPool struct {
	Buffers      []Frame
	NextVictimID BufferID
}

func NewBufferPool(size uint64) *BufferPool {
	pool := BufferPool{
		Buffers:      make([]Frame, size),
		NextVictimID: 0,
	}

	return &pool
}

func (pool *BufferPool) Size() uint64 {
	return uint64(len(pool.Buffers))
}

func (pool *BufferPool) Evict() *BufferID {
	size := pool.Size()
	consecutiveCount := 0

	var evictBufferID BufferID
	for {
		nextVictimID := pool.NextVictimID
		frame := pool.Buffers[nextVictimID]

		// 利用されていないバッファを見つけた
		if frame.UsageCount == 0 {
			evictBufferID = nextVictimID
			break
		}

	}

	return &evictBufferID
}
