package postmi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gcfg.v1"
	"net/http"
)

type Service struct {
	Store Store

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

	store, e := Open(cfg.DataBase.Driver, cfg.DataBase.Setting)
	if e != nil {
		return nil, e
	}
	s.Store = store

	return s, nil
}

func (s *Service) Run() error {
	eng := gin.Default()
	eng.LoadHTMLGlob(s.Config.App.TemplatesPath + "/*")
	eng.Static("/assets", s.Config.App.AssetPath)

	prefix := "/"
	if s.Config.App.Prefix != "" {
		prefix = s.Config.App.Prefix
	}

	g := eng.Group(prefix, gin.BasicAuth(gin.Accounts{
		s.Config.Admin.Login: s.Config.Admin.Password,
	}))

	g.GET("/", func(c *gin.Context) {
		posts, e := s.Store.GetSlice(10, 0)
		res := gin.H{}
		if e != nil {
			fmt.Println(e)
			res["error"] = e
		}
		res["posts"] = posts
		c.HTML(200, "index.html", res)
	})

	g.Any("/create", func(c *gin.Context) {
		rm := gin.H{}
		if c.Request.Method == "POST" {
			rm["message"] = "posted"

			title := c.PostForm("title")
			content := c.PostForm("content")

			p := new(Post)
			p.Title = title
			p.Body = content

			e := s.Store.Save(p)
			if e != nil {
				fmt.Println(e)
				rm["error"] = e
			}
		}
		c.HTML(200, "edit.html", rm)
	})

	g.GET("/p/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.HTML(200, "show.html", gin.H{"da": id})
	})

	g.POST("/p/:id", func(c *gin.Context) {

	})

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Config.App.Port), eng)
}
