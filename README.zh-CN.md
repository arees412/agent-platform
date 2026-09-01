# AgentForge AI — 企业级多 Agent AI 运营平台

> **Reference note:** Upstream and Chinese-language documentation is retained here for reference. The English [README](./README.md) is the canonical source for current AgentForge-specific functionality and project status.

> 通过可扩展的平台协调专业 AI Agent、企业知识、工具、记忆与人工审批。

[English](./README.md) | [简体中文](./README.zh-CN.md)

AgentForge AI 是基于 [`atliliw/agent-platform`](https://github.com/atliliw/agent-platform) 的独立分支项目。仓库保留上游 Git 历史、作者归属、Go 模块路径、内部服务标识和 MIT 许可。AgentForge 品牌、Business Agent Pack 与相关文档是本分支新增内容；不代表上游项目认可或背书。

## 定位

AgentForge AI 面向需要多 Agent 编排、企业知识检索、长期记忆、MCP 工具和治理控制的团队。HTTP 请求通过 Gin Gateway 进入，后端服务使用 gRPC 通信；Agent、知识、记忆、工具和 Harness 治理职责相互分离。

本仓库是软件基础，不宣称已经获得生产认证、合规认证、规模验证或客户验证。生产使用前，部署方必须完成身份与权限、租户隔离、密钥管理、数据保留、评测、监控和人工审批配置。

## 核心能力

- 多 Agent 执行、流式响应、handoff、检查点、干预和会话回放。
- BM25 + 向量检索的 RAG 知识服务，以及分层长期记忆。
- 内置和远程 MCP 工具、技能库与渐进式加载。
- Harness 护栏、审批、评测、Prompt、工作流、SLO、成本和可观测性能力。
- React 19 前端、Docker Compose 部署和 OpenTelemetry Collector。
- 新增五个可编译且已接入默认 Agent 初始化流程的 Business Agent。

## Business Agent Pack

| Agent | 领域 | 风险级别 | 人工边界 |
| --- | --- | --- | --- |
| Business Operations Director | 跨职能运营 | 高 | 高影响行动仅提供建议，需授权人员批准 |
| Revenue Intelligence Agent | 收入运营 | 高 | 不修改 CRM/财务数据、不定价、不承诺、不外发消息 |
| Customer Operations Agent | 客户运营 | 高 | 账户变更、退款、取消和客户消息均为待审批草稿 |
| Risk and Compliance Review Agent | 风险与合规 | 高 | 不作最终法律结论，不批准、申报、认证或执行 |
| Executive Briefing Agent | 高管决策支持 | 中 | 简报不代表批准，也不执行决策 |

这些 Agent 只使用仓库中已验证的 `knowledge_search`、`web_search` 和 `calculator` 工具，并要求标注证据、不确定性和安全回退。全新安装会与现有默认 Agent 一起初始化；已有 Agent 数据的安装不会被自动重置或补种。详见 [Business Agents](./docs/en/business-agents.md)。

## 快速开始

```bash
# 生成已被 gitignore 的各服务 config.yaml
bash scripts/init-config.sh sk-your-dashscope-key

# Windows PowerShell：
pwsh scripts/init-config.ps1 sk-your-dashscope-key

# 构建并启动完整栈
docker compose -f docker/docker-compose.yaml up -d --build

# 健康检查
curl http://localhost:9000/health
```

- Gateway：`http://localhost:9000`
- Frontend：`http://localhost:8888`

开发检查：

```bash
go test ./pkg/businessagents ./pkg/agent
go test ./...
make lint     # 已安装 golangci-lint 时
make fmt
```

## 安全与治理

Business Agent Pack 默认只执行分析与草拟，不包含账户写入、支付、CRM 写入、申报或消息发送工具。法律、财务、合规、客户账户和外部沟通等高影响行动必须由授权人员审批。部署方仍需根据自身风险模型配置 Harness 规则、审批、租户隔离、审计与工具权限。

## 文档

- [文档索引](./docs/README.md)
- [架构与信任边界](./docs/en/architecture.md)
- [Business Agent Pack](./docs/en/business-agents.md)
- [中文配置](./docs/zh-CN/configuration.md)
- [中文部署](./docs/zh-CN/deployment.md)
- [中文开发](./docs/zh-CN/development.md)

## 项目状态、归属与许可

AgentForge AI 正在持续开发中；本文不提供未经验证的生产、性能、客户或合规声明。上游代码、历史和作者归属仍属于 `atliliw` 与贡献者。AgentForge 的品牌、文档和 Business Agent 模块属于本分支的修改，不暗示上游背书。

项目使用 MIT License。详见 [LICENSE](./LICENSE) 与 [NOTICE.md](./NOTICE.md)。
