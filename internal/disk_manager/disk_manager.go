package disk_manager

import (
	"os"

	"golang.org/x/xerrors"

	"github.com/takushi-m/go-simple-rdb/internal/constant"
)

type PageID int64

type DiskManager struct {
	file       *os.File
	nextPageID PageID
}

type IDiskManager interface {
	AllocatePage() (PageID, error)
	ReadPage(ID PageID, data []byte) error
	WritePage(ID PageID, data []byte) error
}

func New(path string) (*DiskManager, error) {
	d := DiskManager{}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, xerrors.Errorf("%+w", err)
	}
	d.file = file

	info, err := d.file.Stat()
	if err != nil {
		return nil, xerrors.Errorf("%+w", err)
	}

	fileSize := info.Size()
	d.nextPageID = PageID(fileSize / constant.PageSize)

	return &d, nil
}

func (d *DiskManager) AllocatePage() (PageID, error) {
	pageID := d.nextPageID
	d.nextPageID += 1
	return pageID, nil
}

func (d *DiskManager) ReadPage(ID PageID, data []byte) error {
	offset := constant.PageSize * int64(ID)
	_, err := d.file.ReadAt(data, offset)
	if err != nil {
		return xerrors.Errorf("%+w", err)
	}
	return nil
}

func (d *DiskManager) WritePage(ID PageID, data []byte) error {
	offset := constant.PageSize * int64(ID)
	_, err := d.file.WriteAt(data, offset)
	if err != nil {
		return xerrors.Errorf("%+w", err)
	}
	return nil
}
