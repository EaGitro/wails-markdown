package main

import (
	"context"
	"fmt"
	"os"
	osrt "runtime"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type OperateFile struct {
	path string
	ctx  context.Context
}

func (a *OperateFile) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func NewOperateFile() *OperateFile {
	return &OperateFile{}
}

type OpenedFile struct {
	FileContent string `json:"FileContent"`
	Err         bool   `json:"Err"`
}

func (f *OperateFile) GetContent() OpenedFile {
	var of OpenedFile
	fmt.Println("filepath:", f.path)
	content, err := os.ReadFile(f.path)
	if err != nil {
		of.Err = true
		return of
	}

	fmt.Println(content)

	of.Err = false
	of.FileContent = string(content)
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
