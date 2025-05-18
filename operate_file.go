package main

import (
	"context"
	"os"
	osrt "runtime"
	"strings"

	"time"

	"github.com/cespare/xxhash/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type OperateFile struct {
	ctx          context.Context
	path         string       // file path
	hotTicker    *time.Ticker // hot-reload Ticker
	xxhashDigest uint64
}

func (f *OperateFile) SetupFileOperation(ctx context.Context) {
	f.SetContext(ctx)
	f.hotTicker = time.NewTicker(time.Second)
	f.hotTicker.Stop()
	go f.HotReload()
}

func (f *OperateFile) SetHotReloadTime(ms int) {
	PPrintln("set hot reload", ms, "ms")
	if ms == 0 {
		f.hotTicker.Stop()
		return
	}

	// f.hotTicker.Stop()
	f.hotTicker.Reset(time.Duration(ms) * time.Millisecond)
}

func (f *OperateFile) HotReload() {
	PPrintln("hot reload goroutine")
	defer f.hotTicker.Stop()
	for {
		select {
		case <-f.hotTicker.C:
			PPrintln("hot reload")
			runtime.WindowExecJS(f.ctx, "window.reloadMd();")
		}
	}
}

func (a *OperateFile) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func NewOperateFile() *OperateFile {
	return &OperateFile{}
}

type OpenedFile struct {
	IsModified  bool   `json:IsModified`
	FileContent string `json:"FileContent"`
	Err         bool   `json:"Err"`
}

func (f *OperateFile) GetContent() OpenedFile {
	var of OpenedFile
	PPrintln("filepath:", f.path)
	if f.path == "" {
		return of
	}
	content, err := os.ReadFile(f.path)
	if err != nil {
		of.Err = true
		return of
	}

	digest := xxhash.Sum64(content)
	// PPrintln(content)

	of.IsModified = false
	if f.xxhashDigest != digest {
		of.FileContent = string(content)
		of.IsModified = true
		f.xxhashDigest = digest
	}
	of.Err = false
	PPrintln(of)
	return of
}

func (f *OperateFile) WriteFile(filedata string) bool {
	bytedata := []byte(filedata)
	err := os.WriteFile(f.path, bytedata, 0644)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (f *OperateFile) ExistsFile() bool {
	_, err := os.Stat(f.path)
	return err == nil
}

func sanitizePath(filepath string) string {
	var sanitizedPath string
	exclude := "!\"$<>&|:?"

	if osrt.GOOS == "windows" {
		exclude += "//"
	}

	sanitizedPath = strings.Trim(filepath, exclude)
	return sanitizedPath
}

func (f *OperateFile) SetFilePathWithDialog() error {
	path, err := runtime.OpenFileDialog(f.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return err
	}
	f.path = path
	return nil
}
