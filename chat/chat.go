package chat

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"strconv"
)

func SetUpChat(server *socketio.Server) {

	var Users []int64

	server.OnConnect("/", func(s socketio.Conn) error {
		//s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnConnect("/chat", func(s socketio.Conn) error {
		//s.SetContext("")
		log.Println("connected:", s.ID())
		guess := "Se ha unido " + s.ID()
		server.BroadcastToNamespace(s.Namespace(), "welcome", guess)
		return nil

	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)

		s.Emit("reply", "notice message "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg interface{}) {

		log.Println("init")
		log.Println(msg)
		log.Println("after interface")
		var room = ""
		var user = ""
		var msgInterface = ""
		m, ok := msg.(map[string]interface{})
		if !ok {
			fmt.Println("error")
		}
		for k, v := range m {
			if k == "room" {
				user = v.(string)
			}
			if k == "user" {
				user = v.(string)
			}
			if k == "msg" {
				msgInterface = v.(string)
			}

			fmt.Println(k, "=>", v)
		}

		fmt.Println("before socket id print")
		fmt.Println(s.ID())
		fmt.Println(user)

		userId, err := strconv.ParseInt(s.ID(), 10, 64)
		if err != nil {
			log.Fatalf("couldnt parse room to int 16 %s", err)
		}

		userFound := findUser(Users, userId)
		fmt.Println(userFound)
		if !userFound {
			fmt.Println("inside !userFound")
			Users = append(Users, userId)
			fmt.Println("After assign to user array")
			s.Join(room)

		}

		s.SetContext(msg)
		fmt.Println("connect id", s.ID())
		/*resp, _ :=  json.Marshal(chatResponse{user: s.ID(),message: msgInterface})
		fmt.Println("response", resp)*/

		server.BroadcastToRoom(s.Namespace(), room, "reply-test", s.ID(), msgInterface)

	})

	server.OnEvent("/", "echo", func(s socketio.Conn, msg interface{}) {

		m, ok := msg.(map[string]interface{})
		if !ok {
			fmt.Println("error")
		}
		for k, v := range m {
			fmt.Println(k, "=>", v)
		}

	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		socketIdInt, err := strconv.ParseInt(s.ID(), 10, 64)
		if err != nil {
			log.Fatalf("couldnt parse room to int 16 %s", err)
		}
		findAndDelete(Users, socketIdInt)
		err = s.Close()
		if err != nil {
			return ""
		}
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		socketIdInt, err := strconv.ParseInt(s.ID(), 10, 64)
		if err != nil {
			log.Fatalf("couldnt desconection parse room to int 16 %s", err)
		}
		findAndDelete(Users, socketIdInt)
		log.Println("closed", reason)
	})

}

func findAndDelete(s []int64, item int64) []int64 {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}

func findUser(users []int64, item int64) bool {
	userFound := false
	for _, v := range users {
		if v == item {
			userFound = true
		}
	}
	return userFound
}
