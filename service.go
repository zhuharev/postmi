package postmi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gcfg.v1"
	"net/http"
)

type Service struct {
	Store

	Config
}

func New(confPath string) (*Service, error) {
	var cfg Config
	e := gcfg.ReadFileInto(&cfg, confPath)
	if e != nil {
		return nil, e
	}

	return NewWithConfig(cfg)
}

func NewWithConfig(cfg Config) (*Service, error) {
	s := new(Service)
	s.Config = cfg

	return s, nil
}

func (s *Service) Run() error {
	eng := gin.Default()
	eng.LoadHTMLGlob(s.Config.App.TemplatesPath + "/*")

	prefix := "/"
	if s.Config.App.Prefix != "" {
		prefix = s.Config.App.Prefix
	}

	g := eng.Group(prefix)

	g.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	g.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.HTML(200, "show.html", gin.H{"da": id})
	})

	g.POST("/:id", func(c *gin.Context) {

	})

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Config.App.Port), eng)
}
