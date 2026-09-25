# 概念
一个自动化的工作流程

比如 Go 工作流可能会自动：

下载你的代码；
安装 Go 环境；
编译项目；
运行测试；
报告有没有出错。

## yml/yaml
例子：
 
name: Go test

on: [push]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
      - run: go test ./...

AML 可以理解为一种结构化配置语言，但它不是 Python、Go 那种能独立完成复杂计算的编程语言。
更准确地说：
- 它有明确的语法规则，所以不只是随便写的格式；
- 它主要用来表达“参数、层级和选项”；
- 和 JSON 类似，但通常更适合人阅读和编辑；
- .yml / .yaml 是它常见的文件后缀。

