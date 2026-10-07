// Package version 提供编译时版本信息注入和运行时查询。
//
// PigeonBox 各 Go 服务共享此包，通过 -ldflags 在编译时注入 git 信息。
//
// # 注入方式
//
// 在 build.sh 或 Makefile 中：
//
//	GIT_COMMIT=$(git rev-parse --short HEAD)
//	GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
//	BUILD_TIME=$(date -u '+%Y-%m-%dT%H:%M:%SZ')
//	MOD="github.com/pigeonbox/kit"
//	go build -ldflags "-X ${MOD}/version.BuildCommit=${GIT_COMMIT} \
//	  -X ${MOD}/version.BuildBranch=${GIT_BRANCH} \
//	  -X ${MOD}/version.BuildTime=${BUILD_TIME}" -o myapp .
//
// # 导出变量
//
//   - Version: 语义化版本号
//   - BuildCommit: Git commit hash
//   - BuildBranch: Git 分支名
//   - BuildTime: 编译时间
//   - StartTime: 进程启动时间
//
// 版本信息 HTTP 端点不内置（kit 保持零框架依赖）：需要暴露
// /buildVersion 时用任意 HTTP 框架把 Metadata() 序列化返回即可。
package version
