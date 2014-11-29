package main

// porting from martini to gin: work in progress!

import (
	"github.com/gin-gonic/gin"
	"github.com/martini-contrib/render"
)

type People []Person
type Person struct {
	Name string
}

func main() {
	// Apparently to disable logging we can't use ClassicMartini(?), so do
	// everything it would normally do and leave out logging, then make a
	// ClassicMartini struct anyhow so we can use it the way all the tutorials
	// say. There is almost certainly a better way of doing this.
	g := gin.New()
	g.Use(gin.Logger())

	// use: logger, recovery and static/public "public" directory

	//mn.MapTo(r, (*martini.Routes)(nil))
	//mn.Action(r.Handle)
	//m := &martini.ClassicMartini{mn, r}

	// note to people checking out this code: in a normal typical use case, you
	// would not need any of the above, and would instead only use the line below.
	// m := martini.Classic()

	// set a default layout for templates
	g.Use(render.Renderer(render.Options{Layout: "layout"}))

	// HTTP action controller
	g.Get("/:title", func(c *gin.Context) {
	//params martini.Params, r render.Render) {
		title := params["title"]
		members := People{
			Person{Name: "Chris McCord"},
			Person{Name: "Matt Sears"},
			Person{Name: "David Stump"},
			Person{Name: "Ricardo Thompson"},
		}

		// use an anonymous struct to pass complex template data
		// see http://talks.golang.org/2012/10things.slide#2
		context := struct {
			Title   string
			Members People
		}{
			title,
			members,
		}

		r.HTML(200, "index", c)
	})

	m.Run()
}
