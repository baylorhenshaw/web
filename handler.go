package web

import (
	"net/http"
	"time"
)

type PageData struct {
	Head PageHead
	Data any
}

type PageHead struct {
	Title     string
	Analytics Analytics
}

func (s *Server) HandleUnprotectedPage(page string, data PageData) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		Logger.Debug("Received a new request for page:" + page + " on route " + r.URL.Path)

		// TODO: Inject Page Data
		// injectors.InjectUser(&data, r)
		// injectors.InjectAnalytics(&data)

		// Serve Page
		err := s.PageTemplates[page].ExecuteTemplate(w, s.Pages[page].LayoutFile, data)
		if err != nil {
			Logger.Error("An error occured on page "+page+"!", "error", err)
		} else {
			Logger.Debug("Served page " + page + " in " + time.Since(start).String())
		}

	}
}
