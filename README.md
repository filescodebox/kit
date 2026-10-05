# FilesCodeBox kit

[![CI](https://github.com/filescodebox/kit/actions/workflows/ci.yml/badge.svg)](https://github.com/filescodebox/kit/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/tag/filescodebox/kit?label=release)](https://github.com/filescodebox/kit/tags)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](./LICENSE)
[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](https://go.dev)

FilesCodeBox 生态的共享 Go 工具库：28 个零生态依赖的通用包（重试、并发编排、限流、优雅停机、工作流等），供 core / server / fnos / p2p 及任意 Go 项目复用。

**Shared Go toolkit for the FilesCodeBox ecosystem** — 28 general-purpose packages with zero ecosystem dependencies.

## 设计原则

- **零生态依赖（叶子仓）**：不 import 任何 `github.com/filescodebox/*` 兄弟模块，CI 强制守卫；被任何模块依赖都不会成环。
- **stdlib 优先**：绝大多数包纯标准库实现；仅 3 个包带轻量外部依赖（`ratelimit` → `golang.org/x/time`，`uidgen` → `google/uuid`，`validator` → `go-playground/validator`）。
- **带测试发布**：全部包附 `-race` 通过的单元测试；lint 棘轮从 0 issues 起步只收紧不放宽。

## 包总览

### 并发与重试

| 包 | 说明 |
|---|------|
| [async](./async/) | 安全 goroutine 启动工具（GoSafe 等 panic 恢复原语，进程兜底层） |
| [group](./group/) | Actor 组：多长驻 goroutine 并发执行，任一返回即中断其余 |
| [singleflight](./singleflight/) | 并发请求合并（防击穿）：同 key 并发调用只执行一次、共享结果 |
| [syncx](./syncx/) | 并发编排原语：按键在飞闸门 + 按序扇出收集 |
| [retry](./retry/) | 可取消的重试与轮询：带退避的动作重试 + 到 deadline 的条件轮询 |
| [ratelimit](./ratelimit/) | 通用限流器（令牌桶 + 按键限流管理器，空闲键自动回收） |
| [cache](./cache/) | 进程内检查缓存：TTL 缓存 + 刷新进行中合一 |
| [progress](./progress/) | 多阶段工作流进度追踪（有序阶段生命周期 + 总进度） |

### 错误与校验

| 包 | 说明 |
|---|------|
| [errors](./errors/) | 线程安全的多错误聚合 |
| [validator](./validator/) | 统一参数验证（go-playground/validator 封装） |
| [enumx](./enumx/) | 枚举管理 |

### 网络与数据

| 包 | 说明 |
|---|------|
| [httpjson](./httpjson/) | HTTP+JSON 调用纪律单源：ctx 感知构造请求、重试与 Retry-After 尊重 |
| [jsonrepair](./jsonrepair/) | LLM 输出宽容 JSON 修复：栅栏剥离 → 散文抽对象 → 逐层修复 |
| [encoding](./encoding/) | 编码与哈希工具 |
| [pagination](./pagination/) | 框架无关的分页参数定义 |
| [uidgen](./uidgen/) | 分布式唯一 ID / 随机串生成 |

### 基础类型工具

| 包 | 说明 |
|---|------|
| [convert](./convert/) | 常用类型转换 |
| [slicex](./slicex/) | 泛型 Slice 工具 |
| [stringx](./stringx/) | 字符串操作（截断等） |
| [mathx](./mathx/) | 数学工具（统计等） |
| [timex](./timex/) | 时间操作 + cron 子包 |
| [pathx](./pathx/) | 文件路径匹配 |
| [ptrx](./ptrx/) | 指针泛型工具 |

### 构建信息

| 包 | 说明 |
|---|------|
| [version](./version/) | 编译时版本注入（-ldflags）+ 运行时查询（Metadata()） |

### 服务治理与网络

| 包 | 说明 |
|---|------|
| [shutdown](./shutdown/) | 优雅停机管理器：信号接入 + 引用计数请求阻断 + 资源逆序释放（日志走 stdlib slog） |
| [workflow](./workflow/) | 流程编排：Pipe 串行 / Parallel 并行 / Async+Future / Switch 条件分支，可任意嵌套，Hook AOP 递归传播 |
| [streamx](./streamx/) | 跨协议流式事件抽象（Event/Source/Sink/Hub）：一次写 Event，SSE/WS/Chunked 多协议扇出，零框架依赖 |
| [wsutil](./wsutil/) | WebSocket 房间广播器：Conn 接口解耦不绑具体 ws 库，空房间自动清理、容错广播 |

> 日志约定：以上包的运行日志统一走 stdlib `log/slog`（`slog.Default()`），消费方可用自定义 handler 接管输出格式。

## 快速开始

```bash
go get github.com/filescodebox/kit/retry@latest
```

```go
// 带退避的重试
err := retry.Do(ctx, retry.Config{Attempts: 3, Delay: 100 * time.Millisecond},
    func(attempt int) error { return doSomething() })

// 按键在飞闸门：同 key 并发上限（如按分片 ID 限在飞）
gate := syncx.NewKeyedGate(8) // cap<=0 = 不限
release, err := gate.Acquire(ctx, "chunk-42")
if err != nil { /* ctx 取消/排队失败 */ }
defer release()
```

各包自带 doc 注释与 `_test.go` 用例，行为以测试为准。

## 本仓库在生态中的位置

```
server / fnos / frontend ──► core ──► contracts ──► thrift runtime
p2p（叶子仓）────────────────► kit（本仓，零生态依赖地基层；core 亦按需消费）
```

纯库仓：无二进制、无镜像、无 chart；以 git tag（`v*`）发布，`go get` 直接消费。

## License

[Apache-2.0](./LICENSE)。其中通用工具包移植自 [enjoydream/ekit](https://git.enjoye.top/enjoydream/ekit)（MIT License, Copyright (c) 2024-2026 enjoydream），归属声明见 [NOTICE](./NOTICE)。
