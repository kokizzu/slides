// Code generated by qtc from "layout.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line layout.html:1
package templates

//line layout.html:1
import (
  "net/http"
)

//line layout.html:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line layout.html:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line layout.html:4
type Page interface {
//line layout.html:4
	Name() string
//line layout.html:4
	StreamName(qw422016 *qt422016.Writer)
//line layout.html:4
	WriteName(qq422016 qtio422016.Writer)
//line layout.html:4
	Title() string
//line layout.html:4
	StreamTitle(qw422016 *qt422016.Writer)
//line layout.html:4
	WriteTitle(qq422016 qtio422016.Writer)
//line layout.html:4
	Content() string
//line layout.html:4
	StreamContent(qw422016 *qt422016.Writer)
//line layout.html:4
	WriteContent(qq422016 qtio422016.Writer)
//line layout.html:4
}
//line layout.html:9
func StreamLayoutPage(qw422016 *qt422016.Writer, p Page, req *http.Request) {
//line layout.html:9
	qw422016.N().S(`<!DOCTYPE html><html lang=en><head><meta charset=utf-8><meta http-equiv=x-ua-compatible content="ie=edge"><meta name=viewport content="width=device-width,initial-scale=1"><meta name=author content=GoJakarta><meta name=language content=en><meta http-equiv=content-language content=en><title>`)
//line layout.html:9
	qw422016.N().S(p.Title())
//line layout.html:9
	qw422016.N().S(`</title><link rel=stylesheet href="`)
//line layout.html:9
	qw422016.N().S(Asset(`css/app.css`))
//line layout.html:9
	qw422016.N().S(`"></head><body><nav class="navbar navbar-expand-lg navbar-dark bg-primary"><a class=navbar-brand href=/ >Search IMDb</a> <button class=navbar-toggler type=button data-toggle=collapse data-target=#navbarColor01 aria-controls=navbarColor01 aria-expanded=false aria-label="Toggle navigation"><span class=navbar-toggler-icon></span></button><div class="collapse navbar-collapse" id=navbarColor01><ul class="navbar-nav mr-auto"><li class="nav-item active"><a class=nav-link href=/ >Search <span class=sr-only>(current)</span></a></li></ul><ul class="nav navbar-nav ml-auto"><li class=nav-item><a class=nav-link href=https://gophers.id/slides target=_blank>GitHub</a></li></ul></div></nav>`)
//line layout.html:9
	p.StreamContent(qw422016)
//line layout.html:9
	qw422016.N().S(`<footer class="footer text-muted"><div class=container><p class=float-right></p><p>&copy;`)
//line layout.html:9
	qw422016.N().S(year)
//line layout.html:9
	qw422016.N().S(`&nbsp; GoJakarta</p></div></footer><script src="`)
//line layout.html:9
	qw422016.N().S(Asset(`js/app.js`))
//line layout.html:9
	qw422016.N().S(`"></script></body></html>`)
//line layout.html:9
}

//line layout.html:9
func WriteLayoutPage(qq422016 qtio422016.Writer, p Page, req *http.Request) {
//line layout.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line layout.html:9
	StreamLayoutPage(qw422016, p, req)
//line layout.html:9
	qt422016.ReleaseWriter(qw422016)
//line layout.html:9
}

//line layout.html:9
func LayoutPage(p Page, req *http.Request) string {
//line layout.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line layout.html:9
	WriteLayoutPage(qb422016, p, req)
//line layout.html:9
	qs422016 := string(qb422016.B)
//line layout.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line layout.html:9
	return qs422016
//line layout.html:9
}
