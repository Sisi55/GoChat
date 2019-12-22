package main

const (
	// 어플리케이션에서 사용할 세션의 키 정보
	sessionKey = "simple_cat_session"
	sessionSecret = "simple_chat_session_secret"

	socketBufferSize = 1024
)

var (
	mongoSession *mgo.Session
	renderer = render.New()
	upgrader = &websocket.Upgrader{
		ReadBufferSize: socketBufferSize,
		WriteBufferSize: socketBufferSize,
	}
)

func main(){

}