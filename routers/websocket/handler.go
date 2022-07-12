package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func init() {
	// Start listening for incoming chat messages
	go handleMessages()
}

// Message 定义一个对象来管理消息，反引号包含的文本是 Go 在对象和 JSON 之间进行序列化和反序列化时需要的元数据。
type Message struct {
	Type       string `json:"type"`
	StatusCode int    `json:"status_code"` //参考http状态码： https://zh.wikipedia.org/zh-hans/HTTP%E7%8A%B6%E6%80%81%E7%A0%81
	UserID     string `json:"user_id"`
	token      string `json:"token"` //认证用
	Detail     string `json:"detail"`
	DataString string `json:"data_string"` //附加的json字符串数据，服务器根据情况解析
}

//创建一个 upGrader 的实例。这只是一个对象，它具备一些方法，这些方法可以获取一个普通 HTTP 链接然后将其升级成一个 WebSocket
var upGrader = websocket.Upgrader{
	//ReadBufferSize:  4096,//读缓存区大小 单位是 bytes，依需求設定（设为 0，则不限制大小）
	//WriteBufferSize: 1024,// 写缓存区大小 同上
	// use default options
	// 检测请求来源 //检查是否跨域
	CheckOrigin: func(r *http.Request) bool {
		////验证方法，只支持Get的话这样写
		//if r.Method != "GET" {
		//	fmt.Println("method is not GET")
		//	return false
		//}
		//验证路径
		if r.URL.Path != "/api/ws" {
			fmt.Println("path error")
			return false
		}
		return true
	},
}
var clientID = 0

//map 映射，其键对应是一个指向 WebSocket 的指针，
//其值是一个int值。我们实际上并不需要这个值，但使用的映射数据结构需要有一个映射值，这样做更容易添加和删除单项。
var clients = make(map[*websocket.Conn]int) // connected clients
//用于由客户端发送消息的队列，扮演通道的角色。后面定义了一个 goroutine 来从这个通道读取新消息，然后将它们发送给其它连接到服务器的客户端。
var broadcast = make(chan Message) // broadcast channel

// WsHandler
//路由是 "/ws",即 ws://127.0.0.1:1234/api/ws
func WsHandler(c *gin.Context) {
	//Upgrade 函数将 http get请求升级到 WebSocket 协议。
	//   responseHeader包含在对客户端升级请求的响应中。
	//// 使用responseHeader指定cookie（Set-Cookie）和应用程序协商的子协议（Sec-WebSocket-Protocol）。
	//// 如果升级失败，则升级将使用HTTP错误响应回复客户端
	//// 返回一个 Conn 指针(wsConn)，拿到他后，可使用 Conn 读写数据与客户端通信。
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("Error during connection up gradation:", err) //连接升级出错
		return
	}
	// 把新的客户端添加到全局的 "clients" 映射表中进行注册
	clients[wsConn] = clientID
	clientID++
	//比 defer wsConn.Close() 更好?
	//通知 Go 在函数返回的时候关闭 WebSocket。
	defer func() {
		closeSocketErr := wsConn.Close()
		if closeSocketErr != nil {
			fmt.Println(err)
		}
	}()

	// 无限循环，等待要写入 WebSocket 的新消息，将其从 JSON 反序列化为 Detail 对象然后送入广播频道。
	for {
		////读取ws中的数据,反序列为json（序列化：将对象转化成字节序列的过程。 反序列化：就是讲字节序列转化成对象的过程。）
		var msg Message // Read in a new message as JSON and map it to a Detail object
		err = wsConn.ReadJSON(&msg)
		if err != nil {
			//fmt.Println()
			log.Printf("websocket服务器错误: %v", err)
			//如果从 socket 中读取数据有误，我们假设客户端已经因为某种原因断开。我们记录错误并从全局的 “clients” 映射表里删除该客户端，这样一来，我们不会继续尝试与其通信。
			delete(clients, wsConn)
			break
		} else {
			if WsDebug {
				fmt.Printf("websocket服务器收到: %v\n", msg)
			}
			// Send the newly received message to the broadcast channel
			broadcast <- msg
		}
	}
}

//一个简单循环，从“broadcast”中连续读取数据，然后通过各自的 WebSocket 连接将消息传播到客户端。
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast //广播频道
		// Send it out to every client that is currently connected
		for client := range clients {
			switch msg.Type {
			case "ping", "ping!", "乒":
				handPingMessage(client, msg)
			case "heartbeat":
				handHeartbeatMessage(client, msg)
			case "sync_page":
				handSyncPageMessage(client, msg)
			default:
				handDefaultMessage(client, msg)
			}
		}
	}
}
