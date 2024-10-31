package main

import "go.baylor.dev/web"

func main() {

	//web.EnableDebug()

	server := web.New(3000)

	server.RegisterPage(web.Page{
		Route:      "/",
		PageFile:   "index",
		LayoutFile: "default",
	},
		server.HandleUnprotectedPage("index", web.PageData{
			Head: web.PageHead{
				Title: "Home",
			},
		}),
	)

	server.Listen()

}
