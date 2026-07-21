# Todo CLI - 命令行待办事项管理

Go + Cobra 框架实现的命令行待办管理工具。数据存储在 JSON 文件中。

## 功能

- 添加待办事项（todo add）
- 列出所有待办（todo list）
- 标记完成（todo done）
- 删除待办（todo delete）

## 快速启动

```bash
go run main.go
```

## 使用示例

```bash
# 添加待办
go run . add "买牛奶"
go run . add "写作业"
go run . add "去健身房"

# 列出待办
go run . list
# 输出:
# [ ]   1. 买牛奶
# [ ]   2. 写作业
# [ ]   3. 去健身房

# 标记完成
go run . done 1

# 再次列出
go run . list
# 输出:
# [x]   1. 买牛奶
# [ ]   2. 写作业
# [ ]   3. 去健身房

# 删除
go run . delete 2
```

## 命令说明

| 命令 | 用法 | 说明 |
|------|------|------|
| add | todo add [内容] | 添加新的待办事项 |
| list | todo list | 列出所有待办 |
| done | todo done [ID] | 标记待办为已完成 |
| delete | todo delete [ID] | 删除待办事项 |

## 技术栈

- **Go** 1.22+ - 后端语言
- **Cobra** - CLI 命令行框架
- **JSON** - 数据持久化存储

## 项目结构

```
├── main.go         入口文件
├── cmd/
│   ├── root.go     根命令
│   ├── todo.go     数据模型和存取
│   ├── add.go      添加命令
│   ├── list.go     列表命令
│   ├── done.go     完成命令
│   └── delete.go   删除命令
└── todos.json      数据文件（自动生成）
```
