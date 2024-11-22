package internal

import (
	"context"
	"fmt"

	r "github.com/wailsapp/wails/v2/pkg/runtime"
)

type FileService struct {
	ctx context.Context
}

func NewFileService() *FileService {
	return &FileService{
		ctx: context.Background(),
	}
}

func (s *FileService) OpenDialog() {
	dir, err := r.OpenDirectoryDialog(s.ctx, r.OpenDialogOptions{})
	if err != nil {
		fmt.Println("something")
	}
}
