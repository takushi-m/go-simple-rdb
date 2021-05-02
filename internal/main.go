package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/takushi-m/go-simple-rdb/internal/disk_manager"
)

func main() {
	path := filepath.Join(os.TempDir(), "table.dat")
	d, err := disk_manager.New(path)
	if err != nil {
		panic(err)
	}

	ID, err := d.AllocatePage()
	if err != nil {
		panic(err)
	}

	data := []byte("hello world\n")
	if err := d.WritePage(ID, data); err != nil {
		panic(err)
	}

	buf := make([]byte, len(data))
	if err := d.ReadPage(ID, buf); err != nil {
		panic(err)
	}

	fmt.Print(string(buf))
}
