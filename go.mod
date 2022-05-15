module github.com/xsuners/machine

go 1.18

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/nats-io/nats.go v1.15.0
	github.com/xsuners/mo v0.1.3
	github.com/xsuners/msql.v2 v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.10.0
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/nats-io/nats-server/v2 v2.8.2 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	github.com/xsuners/mo => ../mo
	github.com/xsuners/msql.v2 => ../msql.v2
)
