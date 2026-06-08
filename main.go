package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"hevcsort/core"

	"github.com/spf13/cobra"
	"github.com/zhangyiming748/finder"
)

var (
	src string
	dst string
)

// version 会在构建时通过 ldflags 注入，默认为 "dev"
var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "hevc-sort",
	Short: "HEVC视频分类工具",
	Long:  `hevc-sort 是一个用于检测和分类 HEVC 编码视频文件的工具`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Long:  `显示 hevc-sort 的当前版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hevc-sort version %s\n", version)
	},
}

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "将视频文件按 HEVC 编码分类到不同目录",
	Long:  `split 命令会扫描源目录中的所有视频文件,并根据是否为 HEVC 编码将其移动到目标目录下的 hevc 或 anti-hevc 子目录中`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// 检查源目录和目标目录是否符合要求
		if src == dst {
			fmt.Fprintf(os.Stderr, "❌ 错误: 源目录和目标目录不能相同\n")
			os.Exit(1)
		}

		// 检查目录是否存在
		if err := checkDirectoryExists(src); err != nil {
			fmt.Fprintf(os.Stderr, "❌ 错误: %v\n", err)
			os.Exit(1)
		}
		if err := checkDirectoryExists(dst); err != nil {
			fmt.Fprintf(os.Stderr, "❌ 错误: %v\n", err)
			os.Exit(1)
		}

		// 检查是否跨分区
		if isCrossPartition(src, dst) {
			fmt.Fprintf(os.Stderr, "❌ 错误: 源目录和目标目录不能在不同分区/磁盘上\n")
			fmt.Fprintf(os.Stderr, "   源目录: %s\n", src)
			fmt.Fprintf(os.Stderr, "   目标目录: %s\n", dst)
			fmt.Fprintf(os.Stderr, "   请确保两个目录在同一分区/磁盘上\n")
			os.Exit(1)
		}

		videos := finder.FindAllVideosInRoot(src)
		if len(videos) != 0 {
			fmt.Fprintf(os.Stderr, "❌ 错误: src 根目录下不应该有视频文件\n找到 %d 个视频文件: %v\n请确保视频文件都在子目录中\n", len(videos), videos)
			os.Exit(1)
		}

		videod := finder.FindAllVideosInRoot(dst)
		if len(videod) != 0 {
			fmt.Fprintf(os.Stderr, "❌ 错误: dst 根目录下不应该有视频文件\n找到 %d 个视频文件: %v\n请清空目标目录或使用其他目录\n", len(videod), videod)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		core.Split(src, dst)
	},
}

func init() {
	splitCmd.Flags().StringVarP(&src, "src", "i", "", "源目录路径")
	splitCmd.Flags().StringVarP(&dst, "dst", "o", "", "目标目录路径")
	splitCmd.MarkFlagRequired("src")
	splitCmd.MarkFlagRequired("dst")

	rootCmd.AddCommand(splitCmd)
	rootCmd.AddCommand(versionCmd)

	// 添加全局 -v/--version 标志
	rootCmd.Version = version
}

// checkDirectoryExists 检查目录是否存在且为目录
func checkDirectoryExists(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("目录不存在: %s", path)
		}
		return fmt.Errorf("无法访问目录 %s: %v", path, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("路径不是目录: %s", path)
	}
	return nil
}

// isCrossPartition 检查两个路径是否在不同分区/磁盘上
func isCrossPartition(src, dst string) bool {
	// 获取两个路径的卷标或根路径
	srcVol := getVolumePath(src)
	dstVol := getVolumePath(dst)
	return srcVol != dstVol
}

// getVolumePath 获取路径所在的卷/分区标识
func getVolumePath(path string) string {
	// 转换为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		return path
	}

	if runtime.GOOS == "windows" {
		// Windows: 提取盘符 (如 C:\)
		if len(absPath) >= 2 {
			return absPath[:2] // 返回 "C:" 这样的格式
		}
	} else {
		// Linux/macOS: 提取根目录 (如 /)
		// 对于 Unix 系统，我们检查是否在同一个挂载点
		// 简化处理：如果根目录相同则认为在同一分区
		return "/"
	}

	return absPath
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
