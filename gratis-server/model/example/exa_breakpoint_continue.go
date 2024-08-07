package example

import (
	"github/stable-diffusion-go/server/global"
)

// ExaFile file struct, 文件结构体
type ExaFile struct {
	global.Model
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// ExaFileChunk file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.Model
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
