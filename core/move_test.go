package core

import (
	"os"
	"path/filepath"
	"testing"
)

// 运行测试: go test -v -timeout 10m -run TestMove
func TestMove(t *testing.T) {
	// 使用临时目录进行测试
	tempDir := t.TempDir()
	src := filepath.Join(tempDir, "src")
	dst := filepath.Join(tempDir, "dst")

	// 创建源目录
	if err := os.MkdirAll(src, 0755); err != nil {
		t.Fatalf("创建源目录失败: %v", err)
	}

	// 创建目标目录
	if err := os.MkdirAll(dst, 0755); err != nil {
		t.Fatalf("创建目标目录失败: %v", err)
	}

	// 注意：这是一个集成测试，需要实际的媒体文件才能完整测试
	// 如果没有媒体文件，Split 函数会返回"未找到任何视频文件"
	// 这里主要测试目录结构和基本流程
	Split(src, dst)
}
