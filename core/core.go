package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/zhangyiming748/FastMediaInfo"
	"github.com/zhangyiming748/finder"
)

/*
src为源文件根目录
dst为目标文件根目录
*/
func Split(src, dst string) {
	videos := finder.FindAllVideos(src)
	if len(videos) == 0 {
		log.Println("未找到任何视频文件")
		return
	}

	hevc := filepath.Join(dst, "hevc")
	antiHevc := filepath.Join(dst, "anti-hevc")

	// 创建目标目录
	if err := os.MkdirAll(hevc, 0755); err != nil {
		log.Fatalf("创建目录 %s 失败: %v", hevc, err)
	}
	if err := os.MkdirAll(antiHevc, 0755); err != nil {
		log.Fatalf("创建目录 %s 失败: %v", antiHevc, err)
	}

	log.Printf("找到 %d 个视频文件，开始分类...", len(videos))

	// 进度显示
	total := len(videos)
	for i, video := range videos {
		// 每处理10个文件或最后一个文件时显示进度
		if (i+1)%10 == 0 || i == total-1 {
			log.Printf("进度: %d/%d (%.1f%%)", i+1, total, float64(i+1)/float64(total)*100)
		}

		mi := FastMediaInfo.GetStandMediaInfo(video)
		vInfo := mi.Video
		var targetDir string
		var targetType string
		if vInfo.Format == "HEVC" {
			targetDir = hevc
			targetType = "HEVC"
		} else {
			targetDir = antiHevc
			targetType = "非HEVC"
		}

		// 计算目标文件路径
		relPath, err := filepath.Rel(src, video)
		if err != nil {
			log.Printf("计算相对路径失败 %s: %v", video, err)
			continue
		}
		target := filepath.Join(targetDir, relPath)

		// 创建目标文件的父目录
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			log.Printf("创建目录 %s 失败: %v", filepath.Dir(target), err)
			continue
		}

		// 检查目标文件是否已存在，如果存在则自动重命名
		finalTarget := target
		if _, err := os.Stat(target); err == nil {
			// 文件已存在，生成新的文件名
			finalTarget = generateUniqueFilename(target)
			log.Printf("⚠️  文件已存在，自动重命名: %s -> %s", target, finalTarget)
		}

		// 移动文件
		if err := os.Rename(video, finalTarget); err != nil {
			log.Printf("移动文件 %s 到 %s 失败: %v", video, finalTarget, err)
			continue
		}

		log.Printf("[%s] 文件已移动: %s -> %s", targetType, video, finalTarget)
	}

	log.Println("分类完成！")
}

// generateUniqueFilename 生成唯一的文件名，避免冲突
func generateUniqueFilename(originalPath string) string {
	dir := filepath.Dir(originalPath)
	base := filepath.Base(originalPath)
	ext := filepath.Ext(base)
	nameWithoutExt := base[:len(base)-len(ext)]

	// 尝试添加时间戳
	timestamp := fmt.Sprintf("_%d", int(time.Now().Unix()))
	newName := nameWithoutExt + timestamp + ext
	newPath := filepath.Join(dir, newName)

	// 如果还是冲突，添加序号
	counter := 1
	for {
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}
		newName = fmt.Sprintf("%s_%d%s", nameWithoutExt, counter, ext)
		newPath = filepath.Join(dir, newName)
		counter++
		if counter > 1000 { // 防止无限循环
			break
		}
	}

	// 最后的备选方案
	return filepath.Join(dir, fmt.Sprintf("%s_conflict_%d%s", nameWithoutExt, time.Now().UnixNano(), ext))
}
