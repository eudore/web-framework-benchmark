package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/astaxie/beego"
	beegocontext "github.com/astaxie/beego/context"
	"github.com/devfeel/dotweb"
	"github.com/dinever/golf"
	"github.com/eudore/erouter"
	"github.com/eudore/eudore"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/kataras/iris"
	iriscontext "github.com/kataras/iris/context"
	"github.com/labstack/echo"
	"github.com/twiglab/twig"
)

type (
	Route struct {
		Method string
		Path   string
	}
)

var (
	num    = 0
	static = []*Route{
		{"GET", "/"},
		{"GET", "/cmd.html"},
		{"GET", "/code.html"},
		{"GET", "/contrib.html"},
		{"GET", "/contribute.html"},
		{"GET", "/debugging_with_gdb.html"},
		{"GET", "/docs.html"},
		{"GET", "/effective_go.html"},
		{"GET", "/files.log"},
		{"GET", "/gccgo_contribute.html"},
		{"GET", "/gccgo_install.html"},
		{"GET", "/go-logo-black.png"},
		{"GET", "/go-logo-blue.png"},
		{"GET", "/go-logo-white.png"},
		{"GET", "/go1.1.html"},
		{"GET", "/go1.2.html"},
		{"GET", "/go1.html"},
		{"GET", "/go1compat.html"},
		{"GET", "/go_faq.html"},
		{"GET", "/go_mem.html"},
		{"GET", "/go_spec.html"},
		{"GET", "/help.html"},
		{"GET", "/ie.css"},
		{"GET", "/install-source.html"},
		{"GET", "/install.html"},
		{"GET", "/logo-153x55.png"},
		{"GET", "/Makefile"},
		{"GET", "/root.html"},
		{"GET", "/share.png"},
		{"GET", "/sieve.gif"},
		{"GET", "/tos.html"},
		{"GET", "/articles/"},
		{"GET", "/articles/go_command.html"},
		{"GET", "/articles/index.html"},
		{"GET", "/articles/wiki/"},
		{"GET", "/articles/wiki/edit.html"},
		{"GET", "/articles/wiki/final-noclosure.go"},
		{"GET", "/articles/wiki/final-noerror.go"},
		{"GET", "/articles/wiki/final-parsetemplate.go"},
		{"GET", "/articles/wiki/final-template.go"},
		{"GET", "/articles/wiki/final.go"},
		{"GET", "/articles/wiki/get.go"},
		{"GET", "/articles/wiki/http-sample.go"},
		{"GET", "/articles/wiki/index.html"},
		{"GET", "/articles/wiki/Makefile"},
		{"GET", "/articles/wiki/notemplate.go"},
		{"GET", "/articles/wiki/part1-noerror.go"},
		{"GET", "/articles/wiki/part1.go"},
		{"GET", "/articles/wiki/part2.go"},
		{"GET", "/articles/wiki/part3-errorhandling.go"},
		{"GET", "/articles/wiki/part3.go"},
		{"GET", "/articles/wiki/test.bash"},
		{"GET", "/articles/wiki/test_edit.good"},
		{"GET", "/articles/wiki/test_Test.txt.good"},
		{"GET", "/articles/wiki/test_view.good"},
		{"GET", "/articles/wiki/view.html"},
		{"GET", "/codewalk/"},
		{"GET", "/codewalk/codewalk.css"},
		{"GET", "/codewalk/codewalk.js"},
		{"GET", "/codewalk/codewalk.xml"},
		{"GET", "/codewalk/functions.xml"},
		{"GET", "/codewalk/markov.go"},
		{"GET", "/codewalk/markov.xml"},
		{"GET", "/codewalk/pig.go"},
		{"GET", "/codewalk/popout.png"},
		{"GET", "/codewalk/run"},
		{"GET", "/codewalk/sharemem.xml"},
		{"GET", "/codewalk/urlpoll.go"},
		{"GET", "/devel/"},
		{"GET", "/devel/release.html"},
		{"GET", "/devel/weekly.html"},
		{"GET", "/gopher/"},
		{"GET", "/gopher/appenginegopher.jpg"},
		{"GET", "/gopher/appenginegophercolor.jpg"},
		{"GET", "/gopher/appenginelogo.gif"},
		{"GET", "/gopher/bumper.png"},
		{"GET", "/gopher/bumper192x108.png"},
		{"GET", "/gopher/bumper320x180.png"},
		{"GET", "/gopher/bumper480x270.png"},
		{"GET", "/gopher/bumper640x360.png"},
		{"GET", "/gopher/doc.png"},
		{"GET", "/gopher/frontpage.png"},
		{"GET", "/gopher/gopherbw.png"},
		{"GET", "/gopher/gophercolor.png"},
		{"GET", "/gopher/gophercolor16x16.png"},
		{"GET", "/gopher/help.png"},
		{"GET", "/gopher/pkg.png"},
		{"GET", "/gopher/project.png"},
		{"GET", "/gopher/ref.png"},
		{"GET", "/gopher/run.png"},
		{"GET", "/gopher/talks.png"},
		{"GET", "/gopher/pencil/"},
		{"GET", "/gopher/pencil/gopherhat.jpg"},
		{"GET", "/gopher/pencil/gopherhelmet.jpg"},
		{"GET", "/gopher/pencil/gophermega.jpg"},
		{"GET", "/gopher/pencil/gopherrunning.jpg"},
		{"GET", "/gopher/pencil/gopherswim.jpg"},
		{"GET", "/gopher/pencil/gopherswrench.jpg"},
		{"GET", "/play/"},
		{"GET", "/play/fib.go"},
		{"GET", "/play/hello.go"},
		{"GET", "/play/life.go"},
		{"GET", "/play/peano.go"},
		{"GET", "/play/pi.go"},
		{"GET", "/play/sieve.go"},
		{"GET", "/play/solitaire.go"},
		{"GET", "/play/tree.go"},
		{"GET", "/progs/"},
		{"GET", "/progs/cgo1.go"},
		{"GET", "/progs/cgo2.go"},
		{"GET", "/progs/cgo3.go"},
		{"GET", "/progs/cgo4.go"},
		{"GET", "/progs/defer.go"},
		{"GET", "/progs/defer.out"},
		{"GET", "/progs/defer2.go"},
		{"GET", "/progs/defer2.out"},
		{"GET", "/progs/eff_bytesize.go"},
		{"GET", "/progs/eff_bytesize.out"},
		{"GET", "/progs/eff_qr.go"},
		{"GET", "/progs/eff_sequence.go"},
		{"GET", "/progs/eff_sequence.out"},
		{"GET", "/progs/eff_unused1.go"},
		{"GET", "/progs/eff_unused2.go"},
		{"GET", "/progs/error.go"},
		{"GET", "/progs/error2.go"},
		{"GET", "/progs/error3.go"},
		{"GET", "/progs/error4.go"},
		{"GET", "/progs/go1.go"},
		{"GET", "/progs/gobs1.go"},
		{"GET", "/progs/gobs2.go"},
		{"GET", "/progs/image_draw.go"},
		{"GET", "/progs/image_package1.go"},
		{"GET", "/progs/image_package1.out"},
		{"GET", "/progs/image_package2.go"},
		{"GET", "/progs/image_package2.out"},
		{"GET", "/progs/image_package3.go"},
		{"GET", "/progs/image_package3.out"},
		{"GET", "/progs/image_package4.go"},
		{"GET", "/progs/image_package4.out"},
		{"GET", "/progs/image_package5.go"},
		{"GET", "/progs/image_package5.out"},
		{"GET", "/progs/image_package6.go"},
		{"GET", "/progs/image_package6.out"},
		{"GET", "/progs/interface.go"},
		{"GET", "/progs/interface2.go"},
		{"GET", "/progs/interface2.out"},
		{"GET", "/progs/json1.go"},
		{"GET", "/progs/json2.go"},
		{"GET", "/progs/json2.out"},
		{"GET", "/progs/json3.go"},
		{"GET", "/progs/json4.go"},
		{"GET", "/progs/json5.go"},
		{"GET", "/progs/run"},
		{"GET", "/progs/slices.go"},
		{"GET", "/progs/timeout1.go"},
		{"GET", "/progs/timeout2.go"},
		{"GET", "/progs/update.bash"},
	}

	githubAPI = []*Route{
		// OAuth Authorizations
		{"GET", "/authorizations"},
		{"GET", "/authorizations/:id"},
		{"POST", "/authorizations"},
		//{"PUT", "/authorizations/clients/:client_id"},
		//{"PATCH", "/authorizations/:id"},
		{"DELETE", "/authorizations/:id"},
		{"GET", "/applications/:client_id/tokens/:access_token"},
		{"DELETE", "/applications/:client_id/tokens"},
		{"DELETE", "/applications/:client_id/tokens/:access_token"},

		// Activity
		{"GET", "/events"},
		{"GET", "/repos/:owner/:repo/events"},
		{"GET", "/networks/:owner/:repo/events"},
		{"GET", "/orgs/:org/events"},
		{"GET", "/users/:user/received_events"},
		{"GET", "/users/:user/received_events/public"},
		{"GET", "/users/:user/events"},
		{"GET", "/users/:user/events/public"},
		{"GET", "/users/:user/events/orgs/:org"},
		{"GET", "/feeds"},
		{"GET", "/notifications"},
		{"GET", "/repos/:owner/:repo/notifications"},
		{"PUT", "/notifications"},
		{"PUT", "/repos/:owner/:repo/notifications"},
		{"GET", "/notifications/threads/:id"},
		//{"PATCH", "/notifications/threads/:id"},
		{"GET", "/notifications/threads/:id/subscription"},
		{"PUT", "/notifications/threads/:id/subscription"},
		{"DELETE", "/notifications/threads/:id/subscription"},
		{"GET", "/repos/:owner/:repo/stargazers"},
		{"GET", "/users/:user/starred"},
		{"GET", "/user/starred"},
		{"GET", "/user/starred/:owner/:repo"},
		{"PUT", "/user/starred/:owner/:repo"},
		{"DELETE", "/user/starred/:owner/:repo"},
		{"GET", "/repos/:owner/:repo/subscribers"},
		{"GET", "/users/:user/subscriptions"},
		{"GET", "/user/subscriptions"},
		{"GET", "/repos/:owner/:repo/subscription"},
		{"PUT", "/repos/:owner/:repo/subscription"},
		{"DELETE", "/repos/:owner/:repo/subscription"},
		{"GET", "/user/subscriptions/:owner/:repo"},
		{"PUT", "/user/subscriptions/:owner/:repo"},
		{"DELETE", "/user/subscriptions/:owner/:repo"},

		// Gists
		{"GET", "/users/:user/gists"},
		{"GET", "/gists"},
		//{"GET", "/gists/public"},
		//{"GET", "/gists/starred"},
		{"GET", "/gists/:id"},
		{"POST", "/gists"},
		//{"PATCH", "/gists/:id"},
		{"PUT", "/gists/:id/star"},
		{"DELETE", "/gists/:id/star"},
		{"GET", "/gists/:id/star"},
		{"POST", "/gists/:id/forks"},
		{"DELETE", "/gists/:id"},

		// Git Data
		{"GET", "/repos/:owner/:repo/git/blobs/:sha"},
		{"POST", "/repos/:owner/:repo/git/blobs"},
		{"GET", "/repos/:owner/:repo/git/commits/:sha"},
		{"POST", "/repos/:owner/:repo/git/commits"},
		//{"GET", "/repos/:owner/:repo/git/refs/*ref"},
		{"GET", "/repos/:owner/:repo/git/refs"},
		{"POST", "/repos/:owner/:repo/git/refs"},
		//{"PATCH", "/repos/:owner/:repo/git/refs/*ref"},
		//{"DELETE", "/repos/:owner/:repo/git/refs/*ref"},
		{"GET", "/repos/:owner/:repo/git/tags/:sha"},
		{"POST", "/repos/:owner/:repo/git/tags"},
		{"GET", "/repos/:owner/:repo/git/trees/:sha"},
		{"POST", "/repos/:owner/:repo/git/trees"},

		// Issues
		{"GET", "/issues"},
		{"GET", "/user/issues"},
		{"GET", "/orgs/:org/issues"},
		{"GET", "/repos/:owner/:repo/issues"},
		{"GET", "/repos/:owner/:repo/issues/:number"},
		{"POST", "/repos/:owner/:repo/issues"},
		//{"PATCH", "/repos/:owner/:repo/issues/:number"},
		{"GET", "/repos/:owner/:repo/assignees"},
		{"GET", "/repos/:owner/:repo/assignees/:assignee"},
		{"GET", "/repos/:owner/:repo/issues/:number/comments"},
		//{"GET", "/repos/:owner/:repo/issues/comments"},
		//{"GET", "/repos/:owner/:repo/issues/comments/:id"},
		{"POST", "/repos/:owner/:repo/issues/:number/comments"},
		//{"PATCH", "/repos/:owner/:repo/issues/comments/:id"},
		//{"DELETE", "/repos/:owner/:repo/issues/comments/:id"},
		{"GET", "/repos/:owner/:repo/issues/:number/events"},
		//{"GET", "/repos/:owner/:repo/issues/events"},
		//{"GET", "/repos/:owner/:repo/issues/events/:id"},
		{"GET", "/repos/:owner/:repo/labels"},
		{"GET", "/repos/:owner/:repo/labels/:name"},
		{"POST", "/repos/:owner/:repo/labels"},
		//{"PATCH", "/repos/:owner/:repo/labels/:name"},
		{"DELETE", "/repos/:owner/:repo/labels/:name"},
		{"GET", "/repos/:owner/:repo/issues/:number/labels"},
		{"POST", "/repos/:owner/:repo/issues/:number/labels"},
		{"DELETE", "/repos/:owner/:repo/issues/:number/labels/:name"},
		{"PUT", "/repos/:owner/:repo/issues/:number/labels"},
		{"DELETE", "/repos/:owner/:repo/issues/:number/labels"},
		{"GET", "/repos/:owner/:repo/milestones/:number/labels"},
		{"GET", "/repos/:owner/:repo/milestones"},
		{"GET", "/repos/:owner/:repo/milestones/:number"},
		{"POST", "/repos/:owner/:repo/milestones"},
		//{"PATCH", "/repos/:owner/:repo/milestones/:number"},
		{"DELETE", "/repos/:owner/:repo/milestones/:number"},

		// Miscellaneous
		{"GET", "/emojis"},
		{"GET", "/gitignore/templates"},
		{"GET", "/gitignore/templates/:name"},
		{"POST", "/markdown"},
		{"POST", "/markdown/raw"},
		{"GET", "/meta"},
		{"GET", "/rate_limit"},

		// Organizations
		{"GET", "/users/:user/orgs"},
		{"GET", "/user/orgs"},
		{"GET", "/orgs/:org"},
		//{"PATCH", "/orgs/:org"},
		{"GET", "/orgs/:org/members"},
		{"GET", "/orgs/:org/members/:user"},
		{"DELETE", "/orgs/:org/members/:user"},
		{"GET", "/orgs/:org/public_members"},
		{"GET", "/orgs/:org/public_members/:user"},
		{"PUT", "/orgs/:org/public_members/:user"},
		{"DELETE", "/orgs/:org/public_members/:user"},
		{"GET", "/orgs/:org/teams"},
		{"GET", "/teams/:id"},
		{"POST", "/orgs/:org/teams"},
		//{"PATCH", "/teams/:id"},
		{"DELETE", "/teams/:id"},
		{"GET", "/teams/:id/members"},
		{"GET", "/teams/:id/members/:user"},
		{"PUT", "/teams/:id/members/:user"},
		{"DELETE", "/teams/:id/members/:user"},
		{"GET", "/teams/:id/repos"},
		{"GET", "/teams/:id/repos/:owner/:repo"},
		{"PUT", "/teams/:id/repos/:owner/:repo"},
		{"DELETE", "/teams/:id/repos/:owner/:repo"},
		{"GET", "/user/teams"},

		// Pull Requests
		{"GET", "/repos/:owner/:repo/pulls"},
		{"GET", "/repos/:owner/:repo/pulls/:number"},
		{"POST", "/repos/:owner/:repo/pulls"},
		//{"PATCH", "/repos/:owner/:repo/pulls/:number"},
		{"GET", "/repos/:owner/:repo/pulls/:number/commits"},
		{"GET", "/repos/:owner/:repo/pulls/:number/files"},
		{"GET", "/repos/:owner/:repo/pulls/:number/merge"},
		{"PUT", "/repos/:owner/:repo/pulls/:number/merge"},
		{"GET", "/repos/:owner/:repo/pulls/:number/comments"},
		//{"GET", "/repos/:owner/:repo/pulls/comments"},
		//{"GET", "/repos/:owner/:repo/pulls/comments/:number"},
		{"PUT", "/repos/:owner/:repo/pulls/:number/comments"},
		//{"PATCH", "/repos/:owner/:repo/pulls/comments/:number"},
		//{"DELETE", "/repos/:owner/:repo/pulls/comments/:number"},

		// Repositories
		{"GET", "/user/repos"},
		{"GET", "/users/:user/repos"},
		{"GET", "/orgs/:org/repos"},
		{"GET", "/repositories"},
		{"POST", "/user/repos"},
		{"POST", "/orgs/:org/repos"},
		{"GET", "/repos/:owner/:repo"},
		//{"PATCH", "/repos/:owner/:repo"},
		{"GET", "/repos/:owner/:repo/contributors"},
		{"GET", "/repos/:owner/:repo/languages"},
		{"GET", "/repos/:owner/:repo/teams"},
		{"GET", "/repos/:owner/:repo/tags"},
		{"GET", "/repos/:owner/:repo/branches"},
		{"GET", "/repos/:owner/:repo/branches/:branch"},
		{"DELETE", "/repos/:owner/:repo"},
		{"GET", "/repos/:owner/:repo/collaborators"},
		{"GET", "/repos/:owner/:repo/collaborators/:user"},
		{"PUT", "/repos/:owner/:repo/collaborators/:user"},
		{"DELETE", "/repos/:owner/:repo/collaborators/:user"},
		{"GET", "/repos/:owner/:repo/comments"},
		{"GET", "/repos/:owner/:repo/commits/:sha/comments"},
		{"POST", "/repos/:owner/:repo/commits/:sha/comments"},
		{"GET", "/repos/:owner/:repo/comments/:id"},
		//{"PATCH", "/repos/:owner/:repo/comments/:id"},
		{"DELETE", "/repos/:owner/:repo/comments/:id"},
		{"GET", "/repos/:owner/:repo/commits"},
		{"GET", "/repos/:owner/:repo/commits/:sha"},
		{"GET", "/repos/:owner/:repo/readme"},
		//{"GET", "/repos/:owner/:repo/contents/*path"},
		//{"PUT", "/repos/:owner/:repo/contents/*path"},
		//{"DELETE", "/repos/:owner/:repo/contents/*path"},
		//{"GET", "/repos/:owner/:repo/:archive_format/:ref"},
		{"GET", "/repos/:owner/:repo/keys"},
		{"GET", "/repos/:owner/:repo/keys/:id"},
		{"POST", "/repos/:owner/:repo/keys"},
		//{"PATCH", "/repos/:owner/:repo/keys/:id"},
		{"DELETE", "/repos/:owner/:repo/keys/:id"},
		{"GET", "/repos/:owner/:repo/downloads"},
		{"GET", "/repos/:owner/:repo/downloads/:id"},
		{"DELETE", "/repos/:owner/:repo/downloads/:id"},
		{"GET", "/repos/:owner/:repo/forks"},
		{"POST", "/repos/:owner/:repo/forks"},
		{"GET", "/repos/:owner/:repo/hooks"},
		{"GET", "/repos/:owner/:repo/hooks/:id"},
		{"POST", "/repos/:owner/:repo/hooks"},
		//{"PATCH", "/repos/:owner/:repo/hooks/:id"},
		{"POST", "/repos/:owner/:repo/hooks/:id/tests"},
		{"DELETE", "/repos/:owner/:repo/hooks/:id"},
		{"POST", "/repos/:owner/:repo/merges"},
		{"GET", "/repos/:owner/:repo/releases"},
		{"GET", "/repos/:owner/:repo/releases/:id"},
		{"POST", "/repos/:owner/:repo/releases"},
		//{"PATCH", "/repos/:owner/:repo/releases/:id"},
		{"DELETE", "/repos/:owner/:repo/releases/:id"},
		{"GET", "/repos/:owner/:repo/releases/:id/assets"},
		{"GET", "/repos/:owner/:repo/stats/contributors"},
		{"GET", "/repos/:owner/:repo/stats/commit_activity"},
		{"GET", "/repos/:owner/:repo/stats/code_frequency"},
		{"GET", "/repos/:owner/:repo/stats/participation"},
		{"GET", "/repos/:owner/:repo/stats/punch_card"},
		{"GET", "/repos/:owner/:repo/statuses/:ref"},
		{"POST", "/repos/:owner/:repo/statuses/:ref"},

		// Search
		{"GET", "/search/repositories"},
		{"GET", "/search/code"},
		{"GET", "/search/issues"},
		{"GET", "/search/users"},
		{"GET", "/legacy/issues/search/:owner/:repository/:state/:keyword"},
		{"GET", "/legacy/repos/search/:keyword"},
		{"GET", "/legacy/user/search/:keyword"},
		{"GET", "/legacy/user/email/:email"},

		// Users
		{"GET", "/users/:user"},
		{"GET", "/user"},
		//{"PATCH", "/user"},
		{"GET", "/users"},
		{"GET", "/user/emails"},
		{"POST", "/user/emails"},
		{"DELETE", "/user/emails"},
		{"GET", "/users/:user/followers"},
		{"GET", "/user/followers"},
		{"GET", "/users/:user/following"},
		{"GET", "/user/following"},
		{"GET", "/user/following/:user"},
		{"GET", "/users/:user/following/:target_user"},
		{"PUT", "/user/following/:user"},
		{"DELETE", "/user/following/:user"},
		{"GET", "/users/:user/keys"},
		{"GET", "/user/keys"},
		{"GET", "/user/keys/:id"},
		{"POST", "/user/keys"},
		//{"PATCH", "/user/keys/:id"},
		{"DELETE", "/user/keys/:id"},
	}

	gplusAPI = []*Route{
		// People
		{"GET", "/people/:userId"},
		{"GET", "/people"},
		{"GET", "/activities/:activityId/people/:collection"},
		{"GET", "/people/:userId/people/:collection"},
		{"GET", "/people/:userId/openIdConnect"},

		// Activities
		{"GET", "/people/:userId/activities/:collection"},
		{"GET", "/activities/:activityId"},
		{"GET", "/activities"},

		// Comments
		{"GET", "/activities/:activityId/comments"},
		{"GET", "/comments/:commentId"},

		// Moments
		{"POST", "/people/:userId/moments/:collection"},
		{"GET", "/people/:userId/moments/:collection"},
		{"DELETE", "/moments/:id"},
	}

	parseAPI = []*Route{
		// Objects
		{"POST", "/1/classes/:className"},
		{"GET", "/1/classes/:className/:objectId"},
		{"PUT", "/1/classes/:className/:objectId"},
		{"GET", "/1/classes/:className"},
		{"DELETE", "/1/classes/:className/:objectId"},

		// Users
		{"POST", "/1/users"},
		{"GET", "/1/login"},
		{"GET", "/1/users/:objectId"},
		{"PUT", "/1/users/:objectId"},
		{"GET", "/1/users"},
		{"DELETE", "/1/users/:objectId"},
		{"POST", "/1/requestPasswordReset"},

		// Roles
		{"POST", "/1/roles"},
		{"GET", "/1/roles/:objectId"},
		{"PUT", "/1/roles/:objectId"},
		{"GET", "/1/roles"},
		{"DELETE", "/1/roles/:objectId"},

		// Files
		{"POST", "/1/files/:fileName"},

		// Analytics
		{"POST", "/1/events/:eventName"},

		// Push Notifications
		{"POST", "/1/push"},

		// Installations
		{"POST", "/1/installations"},
		{"GET", "/1/installations/:objectId"},
		{"PUT", "/1/installations/:objectId"},
		{"GET", "/1/installations"},
		{"DELETE", "/1/installations/:objectId"},

		// Cloud Functions
		{"POST", "/1/functions"},
	}

	apis = [][]*Route{githubAPI, gplusAPI, parseAPI}
)

func init() {
	str := os.Getenv("NUM")
	if len(str) > 0 {
		n, err := strconv.Atoi(str)
		if err == nil {
			num = n
		}
	}
}

func benchmarkRoutes(b *testing.B, router http.Handler, routes []*Route) {
	b.ReportAllocs()
	r := httptest.NewRequest("GET", "/", nil)
	u := r.URL
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.Method
			u.Path = route.Path
			r.URL.Path = route.Path
			router.ServeHTTP(w, r)
		}
	}
}

// golf
func loadGolfRoutes(app *golf.Application, routes []*Route) {
	for i := 0; i < num; i++ {
		app.Use(func(h golf.HandlerFunc) golf.HandlerFunc {
			return func(ctx *golf.Context) {
				h(ctx)
			}
		})
	}
	for _, r := range routes {
		switch r.Method {
		case "GET":
			app.Get(r.Path, golgHandler(r.Method, r.Path))
		case "POST":
			app.Post(r.Path, golgHandler(r.Method, r.Path))
		case "PATCH":
			app.Patch(r.Path, golgHandler(r.Method, r.Path))
		case "PUT":
			app.Put(r.Path, golgHandler(r.Method, r.Path))
		case "DELETE":
			app.Delete(r.Path, golgHandler(r.Method, r.Path))
		}
	}
}

func golgHandler(method, path string) golf.HandlerFunc {
	return func(ctx *golf.Context) {
		ctx.Send("OK")
	}
}

func BenchmarkGolfStatic(b *testing.B) {
	app := golf.New()
	loadGolfRoutes(app, static)
	benchmarkRoutes(b, app, static)
}

func BenchmarkGolfGitHubAPI(b *testing.B) {
	app := golf.New()
	loadGolfRoutes(app, githubAPI)
	benchmarkRoutes(b, app, githubAPI)
}

func BenchmarkGolfGplusAPI(b *testing.B) {
	app := golf.New()
	loadGolfRoutes(app, gplusAPI)
	benchmarkRoutes(b, app, gplusAPI)
}

func BenchmarkGolfParseAPI(b *testing.B) {
	app := golf.New()
	loadGolfRoutes(app, parseAPI)
	benchmarkRoutes(b, app, parseAPI)
}

// echo
func loadEchoRoutes(e *echo.Echo, routes []*Route) {
	for i := 0; i < num; i++ {
		e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(ctx echo.Context) error {
				return h(ctx)
			}
		})
	}
	for _, r := range routes {
		switch r.Method {
		case "GET":
			e.GET(r.Path, echoHandler(r.Method, r.Path))
		case "POST":
			e.POST(r.Path, echoHandler(r.Method, r.Path))
		case "PATCH":
			e.PATCH(r.Path, echoHandler(r.Method, r.Path))
		case "PUT":
			e.PUT(r.Path, echoHandler(r.Method, r.Path))
		case "DELETE":
			e.DELETE(r.Path, echoHandler(r.Method, r.Path))
		}
	}
}

func echoHandler(method, path string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}
}

func BenchmarkEchoStatic(b *testing.B) {
	e := echo.New()
	loadEchoRoutes(e, static)
	benchmarkRoutes(b, e, static)
}

func BenchmarkEchoGitHubAPI(b *testing.B) {
	e := echo.New()
	loadEchoRoutes(e, githubAPI)
	benchmarkRoutes(b, e, githubAPI)
}

func BenchmarkEchoGplusAPI(b *testing.B) {
	e := echo.New()
	loadEchoRoutes(e, gplusAPI)
	benchmarkRoutes(b, e, gplusAPI)
}

func BenchmarkEchoParseAPI(b *testing.B) {
	e := echo.New()
	loadEchoRoutes(e, parseAPI)
	benchmarkRoutes(b, e, parseAPI)
}

// gin
func loadGinRoutes(g *gin.Engine, routes []*Route) {
	for i := 0; i < num; i++ {
		g.Use(func(*gin.Context) {})
	}
	if os.Getenv("SP") != "" {
		g.Use(func(ctx *gin.Context) {
			fmt.Println("gin path", ctx.FullPath())
		})
	}
	for _, r := range routes {
		switch r.Method {
		case "GET":
			g.GET(r.Path, ginHandler(r.Method, r.Path))
		case "POST":
			g.POST(r.Path, ginHandler(r.Method, r.Path))
		case "PATCH":
			g.PATCH(r.Path, ginHandler(r.Method, r.Path))
		case "PUT":
			g.PUT(r.Path, ginHandler(r.Method, r.Path))
		case "DELETE":
			g.DELETE(r.Path, ginHandler(r.Method, r.Path))
		}
	}
}

func ginHandler(method, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}

func BenchmarkGinStatic(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	loadGinRoutes(g, static)
	benchmarkRoutes(b, g, static)
}

func BenchmarkGinGitHubAPI(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	loadGinRoutes(g, githubAPI)
	benchmarkRoutes(b, g, githubAPI)
}

func BenchmarkGinGplusAPI(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	loadGinRoutes(g, gplusAPI)
	benchmarkRoutes(b, g, gplusAPI)
}

func BenchmarkGinParseAPI(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	loadGinRoutes(g, parseAPI)
	benchmarkRoutes(b, g, parseAPI)
}

// dotweb
type dotwebMiddlware struct {
	dotweb.BaseMiddlware
}

func (m *dotwebMiddlware) Handle(ctx dotweb.Context) error {
	return m.Next(ctx)
}

func loadDotwebRoute(app *dotweb.DotWeb, routes []*Route) {
	for i := 0; i < num; i++ {
		app.Use(&dotwebMiddlware{})
	}
	router := app.HttpServer.Router()
	for _, r := range routes {
		router.RegisterRoute(r.Method, r.Path, dotwebHandler(r.Method, r.Path))
	}
}

func dotwebHandler(method, path string) dotweb.HttpHandle {
	return func(ctx dotweb.Context) error {
		return ctx.WriteString("OK")
	}
}

func BenchmarkDotwebStatic(b *testing.B) {
	app := dotweb.New()
	loadDotwebRoute(app, static)
	benchmarkRoutes(b, app.HttpServer, static)
}

func BenchmarkDotwebGitHubAPI(b *testing.B) {
	app := dotweb.New()
	loadDotwebRoute(app, githubAPI)
	benchmarkRoutes(b, app.HttpServer, githubAPI)
}

func BenchmarkDotwebGplusAPI(b *testing.B) {
	app := dotweb.New()
	loadDotwebRoute(app, gplusAPI)
	benchmarkRoutes(b, app.HttpServer, gplusAPI)
}

func BenchmarkDotwebParseAPI(b *testing.B) {
	app := dotweb.New()
	loadDotwebRoute(app, parseAPI)
	benchmarkRoutes(b, app.HttpServer, parseAPI)
}

// iris
func loadIrisRoutes(app *iris.Application, routes []*Route) {
	for _, r := range routes {
		app.Handle(r.Method, r.Path, irisHandler(r.Method, r.Path))
	}
}

func irisHandler(method, path string) iriscontext.Handler {
	return func(ctx iriscontext.Context) {
		ctx.Text("OK")
	}
}

/*
func BenchmarkIrisStatic(b *testing.B) {
	app := iris.Default()
	loadIrisRoutes(app, static)
	benchmarkRoutes(b, app, static)
}

func BenchmarkIrisGitHubAPI(b *testing.B) {
	app := iris.Default()
	loadIrisRoutes(app, githubAPI)
	benchmarkRoutes(b, app, githubAPI)
}

func BenchmarkIrisGplusAPI(b *testing.B) {
	app := iris.Default()
	loadIrisRoutes(app, gplusAPI)
	benchmarkRoutes(b, app, gplusAPI)
}

func BenchmarkIrisParseAPI(b *testing.B) {
	app := iris.Default()
	loadIrisRoutes(app, parseAPI)
	benchmarkRoutes(b, app, parseAPI)
}*/

// beego
func loadBeegoRoutes(app *beego.App, routes []*Route) {
	for _, r := range routes {
		switch r.Method {
		case "GET":
			app.Handlers.Get(r.Path, beegoHandler(r.Method, r.Path))
		case "POST":
			app.Handlers.Post(r.Path, beegoHandler(r.Method, r.Path))
		case "PATCH":
			app.Handlers.Patch(r.Path, beegoHandler(r.Method, r.Path))
		case "PUT":
			app.Handlers.Put(r.Path, beegoHandler(r.Method, r.Path))
		case "DELETE":
			app.Handlers.Delete(r.Path, beegoHandler(r.Method, r.Path))
		}
	}
}

func beegoHandler(method, path string) beego.FilterFunc {
	return func(ctx *beegocontext.Context) {
		ctx.Output.Body([]byte("OK"))
	}
}

func BenchmarkBeegoStatic(b *testing.B) {
	if num > 0 {
		return
	}
	beego.SetLevel(beego.LevelEmergency)
	app := beego.NewApp()
	loadBeegoRoutes(app, static)
	benchmarkRoutes(b, app.Handlers, static)
}

func BenchmarkBeegoGitHubAPI(b *testing.B) {
	if num > 0 {
		return
	}
	app := beego.NewApp()
	loadBeegoRoutes(app, githubAPI)
	benchmarkRoutes(b, app.Handlers, githubAPI)
}

func BenchmarkBeegoGplusAPI(b *testing.B) {
	if num > 0 {
		return
	}
	app := beego.NewApp()
	loadBeegoRoutes(app, gplusAPI)
	benchmarkRoutes(b, app.Handlers, gplusAPI)
}

func BenchmarkBeegoParseAPI(b *testing.B) {
	if num > 0 {
		return
	}
	app := beego.NewApp()
	loadBeegoRoutes(app, parseAPI)
	benchmarkRoutes(b, app.Handlers, parseAPI)
}

// twig
func loadTwigRoutes(app twig.Register, routes []*Route) {
	app.Use(func(h twig.HandlerFunc) twig.HandlerFunc {
		return func(ctx twig.Ctx) error {
			return h(ctx)
		}
	})
	for _, r := range routes {
		app.AddHandler(r.Method, r.Path, twigHandler(r.Method, r.Path))
	}
}

func twigHandler(method, path string) twig.HandlerFunc {
	return func(ctx twig.Ctx) error {
		return ctx.String(200, "OK")
	}
}

func BenchmarkTwigStatic(b *testing.B) {
	app := twig.TODO()
	loadTwigRoutes(app.Config(), static)
	benchmarkRoutes(b, app, static)
}

func BenchmarkTwigGitHubAPI(b *testing.B) {
	app := twig.TODO()
	loadTwigRoutes(app.Config(), githubAPI)
	benchmarkRoutes(b, app, githubAPI)
}

func BenchmarkTwigGplusAPI(b *testing.B) {
	app := twig.TODO()
	loadTwigRoutes(app.Config(), gplusAPI)
	benchmarkRoutes(b, app, gplusAPI)
}

func BenchmarkTwigParseAPI(b *testing.B) {
	app := twig.TODO()
	loadTwigRoutes(app.Config(), parseAPI)
	benchmarkRoutes(b, app, parseAPI)
}

// eudore radix router
func loadEuodreRoutes(app *eudore.App, routes []*Route) {
	for i := 0; i < num; i++ {
		app.AddMiddleware("ANY", "", func(eudore.Context) {})
	}
	if os.Getenv("SP") != "" {
		app.AddMiddleware("ANY", "", func(ctx eudore.Context) {
			fmt.Println("eudore path", ctx.Path())
		})
	}
	for _, r := range routes {
		app.RegisterHandler(r.Method, r.Path, eudore.HandlerFuncs{eudoreHandler(r.Method, r.Path)})
	}
}

var eudoreCtx = os.Getenv("EUDORECTX") == ""

func eudoreHandler(method, path string) eudore.HandlerFunc {
	if eudoreCtx {
		return func(ctx eudore.Context) {
			ctx.WriteString("OK")
		}
	}
	return eudore.NewHandlerFunc(func(ctx eudore.ContextData) {
		ctx.WriteString("OK")
	})
}

func BenchmarkEudoreRadixStatic(b *testing.B) {
	app := eudore.NewCore()
	loadEuodreRoutes(app.App, static)
	benchmarkRoutes(b, eudore.GetNetHttpHandler(nil, app), static)
}

func BenchmarkEudoreRadixGitHubAPI(b *testing.B) {
	app := eudore.NewCore()
	loadEuodreRoutes(app.App, githubAPI)
	benchmarkRoutes(b, eudore.GetNetHttpHandler(nil, app), githubAPI)
}

func BenchmarkEudoreRadixGplusAPI(b *testing.B) {
	app := eudore.NewCore()
	loadEuodreRoutes(app.App, gplusAPI)
	benchmarkRoutes(b, eudore.GetNetHttpHandler(nil, app), gplusAPI)
}

func BenchmarkEudoreRadixParseAPI(b *testing.B) {
	app := eudore.NewCore()
	loadEuodreRoutes(app.App, parseAPI)
	benchmarkRoutes(b, eudore.GetNetHttpHandler(nil, app), parseAPI)
}

// eudore full router
func BenchmarkEudoreFullStatic(b *testing.B) {
	app := eudore.NewCore()
	app.Router = eudore.NewRouterFull()
	loadEuodreRoutes(app.App, static)
	benchmarkRoutes(b, eudore.GetNetHttpHandler(nil, app), static)
}

func BenchmarkEudoreFullGitHubAPI(b *testing.B) {
	app := eudore.NewCore()
	app.Router = eudore.NewRouterFull()
	loadEuodreRoutes(app.App, githubAPI)
	benchmarkRoutes(b, eudore.GetNetHttpHandler(nil, app), githubAPI)
}

func BenchmarkEudoreFullGplusAPI(b *testing.B) {
	app := eudore.NewCore()
	app.Router = eudore.NewRouterFull()
	loadEuodreRoutes(app.App, gplusAPI)
	benchmarkRoutes(b, eudore.GetNetHttpHandler(nil, app), gplusAPI)
}

func BenchmarkEudoreFullParseAPI(b *testing.B) {
	app := eudore.NewCore()
	app.Router = eudore.NewRouterFull()
	loadEuodreRoutes(app.App, parseAPI)
	benchmarkRoutes(b, eudore.GetNetHttpHandler(nil, app), parseAPI)
}

// httprouter
func loadHttprouterRoutes(g *httprouter.Router, routes []*Route) {
	for _, r := range routes {
		switch r.Method {
		case "GET":
			g.GET(r.Path, httprouterHandler(r.Method, r.Path))
		case "POST":
			g.POST(r.Path, httprouterHandler(r.Method, r.Path))
		case "PATCH":
			g.PATCH(r.Path, httprouterHandler(r.Method, r.Path))
		case "PUT":
			g.PUT(r.Path, httprouterHandler(r.Method, r.Path))
		case "DELETE":
			g.DELETE(r.Path, httprouterHandler(r.Method, r.Path))
		}
	}
}

func httprouterHandler(method, path string) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

func BenchmarkHttprouterStatic(b *testing.B) {
	if num > 0 {
		return
	}
	router := httprouter.New()
	loadHttprouterRoutes(router, static)
	benchmarkRoutes(b, router, static)
}

func BenchmarkHttprouterGitHubAPI(b *testing.B) {
	if num > 0 {
		return
	}
	router := httprouter.New()
	loadHttprouterRoutes(router, githubAPI)
	benchmarkRoutes(b, router, githubAPI)
}

func BenchmarkHttprouterGplusAPI(b *testing.B) {
	if num > 0 {
		return
	}
	router := httprouter.New()
	loadHttprouterRoutes(router, gplusAPI)
	benchmarkRoutes(b, router, gplusAPI)
}

func BenchmarkHttprouterParseAPI(b *testing.B) {
	if num > 0 {
		return
	}
	router := httprouter.New()
	loadHttprouterRoutes(router, parseAPI)
	benchmarkRoutes(b, router, parseAPI)
}

// erouter
func loadErouterRoutes(router erouter.Router, routes []*Route) {
	for i := 0; i < num; i++ {
		router.AddMiddleware("ANY", "", func(h erouter.Handler) erouter.Handler {
			return func(w http.ResponseWriter, r *http.Request, p erouter.Params) {
				h(w, r, p)
			}
		})
	}
	for _, r := range routes {
		router.AddHandler(r.Method, r.Path, erouterHandler(r.Method, r.Path))
	}
}

func erouterHandler(method, path string) erouter.Handler {
	return func(w http.ResponseWriter, req *http.Request, p erouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

func BenchmarkErouterRadixStatic(b *testing.B) {
	router := erouter.NewRouterRadix()
	loadErouterRoutes(router, static)
	benchmarkRoutes(b, router, static)
}

func BenchmarkErouterRadixGitHubAPI(b *testing.B) {
	router := erouter.NewRouterRadix()
	loadErouterRoutes(router, githubAPI)
	benchmarkRoutes(b, router, githubAPI)
}

func BenchmarkErouterRadixGplusAPI(b *testing.B) {
	router := erouter.NewRouterRadix()
	loadErouterRoutes(router, gplusAPI)
	benchmarkRoutes(b, router, gplusAPI)
}

func BenchmarkErouterRadixParseAPI(b *testing.B) {
	router := erouter.NewRouterRadix()
	loadErouterRoutes(router, parseAPI)
	benchmarkRoutes(b, router, parseAPI)
}

func BenchmarkErouterFullStatic(b *testing.B) {
	router := erouter.NewRouterFull()
	loadErouterRoutes(router, static)
	benchmarkRoutes(b, router, static)
}

func BenchmarkErouterFullGitHubAPI(b *testing.B) {
	router := erouter.NewRouterFull()
	loadErouterRoutes(router, githubAPI)
	benchmarkRoutes(b, router, githubAPI)
}

func BenchmarkErouterFullGplusAPI(b *testing.B) {
	router := erouter.NewRouterFull()
	loadErouterRoutes(router, gplusAPI)
	benchmarkRoutes(b, router, gplusAPI)
}

func BenchmarkErouterFullParseAPI(b *testing.B) {
	router := erouter.NewRouterFull()
	loadErouterRoutes(router, parseAPI)
	benchmarkRoutes(b, router, parseAPI)
}
