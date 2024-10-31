package web

import (
	"html/template"
	"net/http"
	"strconv"
)

type Server struct {
	Port          int
	Analytics     Analytics
	Pages         map[string]Page
	PageTemplates map[string]*template.Template
}

// Create a new instance of a web server
func New(port int) *Server {
	return &Server{
		Analytics: Analytics{
			Enabled: false,
			Url:     "",
			Id:      "",
		},
		Port:          port,
		Pages:         make(map[string]Page),
		PageTemplates: make(map[string]*template.Template),
	}
}

func (s *Server) SetAnalytics(analytics Analytics) *Server {
	s.Analytics = analytics
	return s
}

func (s *Server) SetPort(port int) *Server {
	s.Port = port
	return s
}

func (s Server) Listen() {
	Logger.Info("Starting server!", "Port", s.Port)
	Logger.Debug("Web Server", "Server", s)
	Logger.Fatal(http.ListenAndServe(":"+s.getPortString(), nil))
}

func (s Server) getPortString() string {
	return strconv.Itoa(s.Port)
}
