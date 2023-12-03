package common

import "github.com/valyala/fasthttp"

type Resource struct {
	Request  *fasthttp.Request
	Response *fasthttp.Response
}

func AcquireResource() *Resource {
	return &Resource{
		Request:  fasthttp.AcquireRequest(),
		Response: fasthttp.AcquireResponse(),
	}
}

func (r *Resource) Release() {
	defer fasthttp.ReleaseRequest(r.Request)
	defer fasthttp.ReleaseResponse(r.Response)
}
