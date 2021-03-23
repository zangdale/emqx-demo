package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

/*
免费的在线 MQTT 5 服务器

官网: https://www.emqx.cn/mqtt/public-mqtt5-broker

docker:  docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqx/emqx:v4.0.0
user: admin/public

Broker: broker.emqx.io
TCP 端口： 1883
Websocket 端口： 8083
TCP/TLS 端口： 8883
Websocket/TLS 端口： 8084

*/

var (
	topic  = "ss"
	broker = "127.0.0.1"

	port = 1883
)

// 全局 MQTT pub 消息处理
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("接收消息: 从话题[ %s ] 发来的内容: %s \n", msg.Topic(), msg.Payload())
}

// 连接的回调
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("连接...")
}

// 连接丢失的回调
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("链接丢失: %v", err)
}

func main() {

	opts := mqtt.NewClientOptions() // 用于设置 broker，端口，客户端 id ，用户名密码等选项
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))

	// 设置 TLS
	//tlsConfig := NewTlsConfigNoClientCert()
	//opts.SetTLSConfig(tlsConfig)

	opts.SetClientID("go_mqtt_client")
	//opts.SetUsername("emqx")
	//opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	// 遗言
	opts.SetWill("offline", "go_mqtt_client offline", 1, false)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client)

	// 断开连接将终止与服务器的连接，但是在等待指定的毫秒数以等待现有工作完成之前不会断开连接
	//TODO 设置为 0 也是可以接收消息
	client.Disconnect(0)
	client.Disconnect(250)
}

// 使用 TLS 连接
func NewTlsConfigWithClientCert() *tls.Config {
	certpool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatalln(err.Error())
	}
	certpool.AppendCertsFromPEM(ca)
	// 加载 客户端证书和私钥
	clientKeyPair, err := tls.LoadX509KeyPair("client-crt.pem", "client-key.pem")
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{clientKeyPair},
	}
}

// 不设置客户端证书
func NewTlsConfigNoClientCert() *tls.Config {
	certpool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatalln(err.Error())
	}
	certpool.AppendCertsFromPEM(ca)
	return &tls.Config{
		RootCAs: certpool,
	}
}

// 订阅
func sub(client mqtt.Client) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("接收到的话题: %s \n", topic)
}

// 发布消息
func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("> 你好 不乖 %d", i)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
