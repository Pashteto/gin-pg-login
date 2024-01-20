package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	ginhandlers "gin-pg-login/internal/gin-handlers"
)

type Servo struct {
	r    *gin.Engine
	gh   *ginhandlers.GinHandlers
	port string // ":9999"
	Serv *http.Server
}

func NewServo(gh *ginhandlers.GinHandlers, port string) Servo {
	if os.Getenv("LOCAL_DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
		port = ":443"
	}
	r := gin.Default()
	r.LoadHTMLGlob("./templates/**/*")
	r.Static("/static", "./static")
	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	return Servo{
		r:    r,
		gh:   gh,
		port: port,
		Serv: server,
	}
}

func (s Servo) InitRoutes() {
	//==========================================================================
	// BIO
	//==========================================================================
	s.r.GET("/", s.gh.Bio)

	s.r.GET("/getBio", s.gh.GetBio)

	//==========================================================================
	// HELPING
	//==========================================================================
	s.r.GET("/helping-elena", s.gh.GetElena)

	s.r.GET("/helping-nikita", s.gh.GetNikita)

	s.r.GET("/helping-mikhail", s.gh.GetMikhail)

	//==========================================================================
	// PAGES
	//==========================================================================
	/*	s.r.GET("/bot1", s.gh.Entry)

		s.r.GET("/signup", s.gh.SignUp)

		s.r.GET("/locations", s.gh.Locations)

		s.r.GET("/location/:id", s.gh.LocationId)

		s.r.GET("/logout", s.gh.Logout)

		s.r.GET("/profile_created", s.gh.ProfileCreated)

		s.r.GET("/profile_page", s.gh.ProfilePage)

		//==========================================================================
		// ACTIONS
		//==========================================================================

		s.r.POST("/actions/signup", s.gh.SignUpPost)

		s.r.POST("/actions/login", s.gh.LoginPost)

		s.r.POST("/actions/location", s.gh.LocationPost)
	*/
}

func (s Servo) Run() {
	if os.Getenv("LOCAL_DEBUG") != "true" {
		s.Serv.Addr = ":443"
		err := s.Serv.ListenAndServeTLS(
			"/etc/letsencrypt/live/pashteto.com/fullchain.pem",
			"/etc/letsencrypt/live/pashteto.com/privkey.pem",
		)
		if err != nil {
			log.Fatalln("ListenAndServeTLS: ", err.Error())
			return
		}
	} else {
		err := s.Serv.ListenAndServe()
		if err != nil {
			log.Fatalln("ListenAndServe: ", err.Error())
			return
		}
	}
	log.Printf("server initiated: listen on port: %s", s.Serv.Addr)
}

func (s Servo) KillServer() {
	// The context is used to give a timeout for the graceful server shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.Serv.Shutdown(ctx); err != nil {
		log.Fatalln("Server forced to shutdown:", err.Error())
	}

	log.Println("Server exiting")
}
