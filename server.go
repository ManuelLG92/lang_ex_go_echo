package main

import (
	"api.go.com/echo/chat"
	"api.go.com/echo/config"
	"api.go.com/echo/routes"
	_ "api.go.com/echo/user"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	server := socketio.NewServer(nil)

	go chat.SetUpChat(server)

	go func() {
		err := server.Serve()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("connected")
	}()
	defer func(server *socketio.Server) {
		err := server.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(server)

	e := echo.New()

	e.Static("/chat", "public")

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	/*	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))*/
	e.Use(middleware.CORS())

	//e.Use(middleware.Recover())

	e.Any("/socket.io/", func(context echo.Context) error {
		server.ServeHTTP(context.Response(), context.Request())
		return nil
	})
	config.InitializeDB()
	routes.PrivateRoutes(e)
	routes.PublicRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

/*	var Users []int64
	server.OnConnect("/", func(s socketio.Conn) error {
		//s.SetContext("")
		log.Println("connected:", s.ID())

		return nil

	})

	server.OnConnect("/chat", func(s socketio.Conn) error {
		//s.SetContext("")
		log.Println("connected:", s.ID())
		server.BroadcastToNamespace(s.Namespace(), "welcome", "Se ha unido.")
		return nil

	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)

		s.Emit("reply", "notice message "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg interface{}) {
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

			//fmt.Println(k, "=>", v)
		}

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

		server.BroadcastToRoom(s.Namespace(), room, "reply-test", msgInterface)

	})

	server.OnEvent("/", "echo", func(s socketio.Conn, msg interface{}) {
		fmt.Println("beforer echo event")
		fmt.Println(msg)
		fmt.Println("after echo wriye")
		fmt.Println(reflect.TypeOf(msg))
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
			log.Fatalf("couldnt parse room to int 16 %s", err)
		}
		findAndDelete(Users, socketIdInt)
		log.Println("closed", reason)
	})
*/
