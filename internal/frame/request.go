package frame

import "net/http"

const (
	Get HttpMethod = iota + 1
	Post
	Head
	Put
	Patch
	Delete
	Options
)

type HttpMethod int

type HandlerFn func(request *http.Request) (*HttpResponse, error)

type Headers struct {
	setHeaderEntryMap   map[string]string
	addHeaderEntrySlice []addHeaderEntry
}

type addHeaderEntry struct {
	name  string
	value string
}

func (x *Headers) GetSetEntryMap() map[string]string {
	return x.setHeaderEntryMap
}

func (x *Headers) GetAddEntrySlice() []addHeaderEntry {
	return x.addHeaderEntrySlice
}
