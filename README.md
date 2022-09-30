### 生成账户
`go test -timeout 3000s -run ^TestGenAccounts$ github.com/wangdayong228/benchmark_conflux_transfer`
### 批量转账
`go test -v -timeout 3000s -run ^TestDispatchTreasure$ github.com/wangdayong228/benchmark_conflux_transfer`
### 启动压测
`go run .`