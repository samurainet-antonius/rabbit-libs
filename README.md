# Rabbit Lib
Manage Rabbit connection, publish, consume, etc

## How to change existing code
1. add import `import "github.com/samurainet-antonius/rabbit-libs/rabbitmq"`
2. Replace `amqp.Connection` with `rabbitmq.Connection` and `amqp.Channel` with `rabbitmq.Channel`

## Example
### Auto reconnect consumer
> go run go run example/consumer/demo.go

### Auto reconnect publisher
> go run go run example/publisher/demo.go

### Standart Aino Consumer
> go run example/std-consumer/demo.go

### Standart Aino Publisher
> go run example/std-publisher/demo.go