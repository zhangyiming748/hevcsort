# HEVC Sort

一个用于检测和分类 HEVC 编码视频文件的命令行工具。

## 📋 功能特性

- 🔍 **自动检测**: 扫描源目录中的所有视频文件，自动识别 HEVC 编码
- 📁 **智能分类**: 将 HEVC 视频移动到 `dst/hevc`，非 HEVC 视频移动到 `dst/anti-hevc`
- 🔄 **保持结构**: 保持原有的文件层级目录结构不变
- ⚡ **跨平台**: 支持 Windows、Linux、macOS
- 🛡️ **安全检查**: 防止跨分区移动、目录不存在等常见问题
- 📊 **进度显示**: 实时显示处理进度
- 🎯 **冲突处理**: 自动处理文件名冲突，生成唯一文件名

## 🚀 快速开始

### 安装

```bash
# 从源码构建
go build -o hevc-sort.exe .

# 或直接下载预编译版本（通过 GitHub Releases）
```

### 使用方法

```bash
# 基本用法
hevc-sort split --src /path/to/source --dst /path/to/destination

# 简写形式
hevc-sort split -i /path/to/source -o /path/to/destination
```

### 参数说明

- `-i, --src`: 源目录路径（必需），视频文件应位于子目录中
- `-o, --dst`: 目标目录路径（必需），根目录不应有视频文件

### 使用要求

1. ✅ 源目录和目标目录不能相同
2. ✅ 源目录和目标目录必须在同一分区/磁盘上
3. ✅ 源目录根目录下不能有视频文件（视频应在子目录中）
4. ✅ 目标目录根目录下不能有视频文件
5. ✅ 源目录和目标目录必须存在且为有效目录

## 📊 输出示例

```
2026/06/08 09:14:40 找到 100 个视频文件，开始分类...
2026/06/08 09:14:40 进度: 10/100 (10.0%)
2026/06/08 09:14:40 [HEVC] 文件已移动: src/movie1/video.mp4 -> dst/hevc/movie1/video.mp4
2026/06/08 09:14:40 [非HEVC] 文件已移动: src/movie2/video.avi -> dst/anti-hevc/movie2/video.avi
2026/06/08 09:14:40 进度: 20/100 (20.0%)
...
2026/06/08 09:15:30 进度: 100/100 (100.0%)
2026/06/08 09:15:30 分类完成！
```

## 🏗️ 项目结构

```
hevcsort/
├── main.go              # 命令行入口，Cobra CLI 实现
├── core/
│   ├── core.go          # 核心分类逻辑
│   └── move_test.go     # 单元测试
├── go.mod               # Go 模块依赖
├── go.sum               # 依赖校验文件
├── .goreleaser.yml      # GoReleaser 配置
├── ISSUES.md            # 问题追踪与改进建议
├── IMPROVEMENTS.md      # 改进实施总结
└── README.md            # 项目文档
```

## 🛠️ 技术栈

- **语言**: Go 1.x
- **CLI 框架**: [Cobra](https://github.com/spf13/cobra)
- **媒体信息**: [FastMediaInfo](https://github.com/zhangyiming748/FastMediaInfo)
- **文件查找**: [finder](https://github.com/zhangyiming748/finder)
- **构建工具**: [GoReleaser](https://goreleaser.com/)
- **CI/CD**: GitHub Actions

## 📝 开发历史

本项目的完整开发历史和所有提交记录如下：

---

### Commit: d54669f
**Author:** zen <win10 on cnpc>  
**Date:** 2026-06-08  
**Message:** fix(core): 解决文件分类工具中的多个关键问题

**改进内容:**
- ✅ 实现跨分区文件移动检测，防止因跨磁盘操作导致的移动失败
- ✅ 添加目录存在性检查，验证源目录和目标目录的有效性
- ✅ 改进 PreRun 错误处理方式，统一使用标准错误输出并正确退出
- ✅ 实现文件冲突自动处理，当目标文件存在时自动生成唯一文件名
- ✅ 改进测试代码，使用临时目录提高测试可移植性
- ✅ 添加实时进度显示功能，让用户了解处理状态
- ✅ 统一日志输出格式，提升日志一致性和可读性
- ✅ 移除未使用的并发处理功能以确保兼容低性能设备

---

### Commit: 084dcc1
**Author:** zen <win10 on cnpc>  
**Date:** 2026-06-08  
**Message:** docs(project): 添加项目逻辑漏洞与改进建议文档

**文档内容:**
- 📝 记录了跨分区文件移动失败的问题及解决方案
- 📝 指出了缺少目录存在性检查的风险点
- 📝 分析了 PreRun 错误处理方式的不当之处
- 📝 提出了文件冲突处理的改进建议
- 📝 指出了测试代码硬编码路径的问题
- 📝 建议增加用户反馈和进度显示功能
- 📝 统一日志输出格式的规范化建议
- 📝 识别了并发控制缺失的性能问题
- 📝 按优先级分类了各项问题的修复顺序

---

### Commit: 99efd8d
**Author:** zen <win10 on cnpc>  
**Date:** 2026-04-21  
**Message:** fix(core): 修复移动功能并改进错误处理

**修复内容:**
- 🔧 修改测试用例中的源和目标路径配置
- 🔧 添加源目录和目标目录相同的检查避免无限递归
- 🔧 将错误输出从 fmt 改为 log.Fatalf 确保程序正确退出
- 🔧 更新根目录视频文件检测逻辑并改进错误提示
- 🔧 添加 .gitignore 文件忽略编译后的可执行文件

---

### Commit: a78d5f8
**Author:** zen <win10 on cnpc>  
**Date:** 2026-04-21  
**Message:** feat(cli): 添加命令行界面和改进视频分类功能

**新增功能:**
- ✨ 集成 cobra 库实现命令行参数解析
- ✨ 添加 src 和 dst 参数用于指定源目录和目标目录
- ✨ 实现目录验证逻辑确保源目录和目标目录根目录无视频文件
- ✨ 改进核心分类逻辑添加目录创建和文件移动功能
- ✨ 添加详细的日志输出和进度提示
- ✨ 移除旧的字符串处理依赖改用更可靠的路径操作
- ✨ 添加错误处理和异常情况的用户友好提示

---

### Commit: 9759420
**Author:** zen <win10 on cnpc>  
**Date:** 2026-04-21  
**Message:** feat(core): 实现HEVC视频分类功能

**初始实现:**
- 🎉 添加核心模块用于分离HEVC和非HEVC视频文件
- 🎉 集成FastMediaInfo库获取视频格式信息
- 🎉 实现文件遍历和分类逻辑到不同目标目录
- 🎉 添加GoReleaser配置支持多平台构建
- 🎉 配置GitHub Actions自动化发布流程
- 🎉 添加单元测试验证移动功能
- 🎉 初始化项目依赖管理文件

---

### Commit: db031d4
**Author:** zen <zhangyiming748@gmail.com>  
**Date:** 2026-04-21  
**Message:** Initial commit

**项目初始化:**
- 🌱 创建项目基础结构
- 🌱 初始化 Git 仓库

---

## 📈 项目统计

- **总提交数**: 6 次
- **主要贡献者**: zen
- **开发周期**: 2026-04-21 至 2026-06-08
- **代码行数**: ~400+ 行 Go 代码
- **测试覆盖**: 包含核心功能单元测试

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 🔗 相关链接

- [ISSUES.md](ISSUES.md) - 问题追踪与改进建议
- [IMPROVEMENTS.md](IMPROVEMENTS.md) - 改进实施详细总结
- [GitHub Actions](.github/workflows/gorelease.yml) - CI/CD 配置
