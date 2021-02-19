免费的在线 MQTT 5 服务器

官网: https://www.emqx.cn/mqtt/public-mqtt5-broker

- Broker: broker.emqx.io
- TCP 端口： 1883
- Websocket 端口： 8083
- TCP/TLS 端口： 8883
- Websocket/TLS 端口： 8084

> golang 的 demo 案例

```go
go run main.go

连接...
接收到的话题: topic/buguai 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 0 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 1 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 2 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 3 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 4 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 5 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 6 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 7 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 8 
接收消息: 从话题[ topic/buguai ] 发来的内容: > 你好 不乖 9 

Process finished with exit code 0
```
