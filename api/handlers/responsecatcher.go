package handlers

import "net/http"

func newResponseCatcher(w http.ResponseWriter, r *http.Request) *responseCatcher {
	res := &responseCatcher{}
	res.w = w
	res.response.Request = r
	res.response.Proto = r.Proto
	return res
}

type responseCatcher struct {
	w        http.ResponseWriter
	response http.Response
}

func (rc *responseCatcher) Header() http.Header {
	return rc.w.Header()
}

func (rc *responseCatcher) Write(b []byte) (int, error) {
	size, err := rc.w.Write(b)
	rc.response.ContentLength += int64(size)
	return size, err
}

func (rc *responseCatcher) WriteHeader(s int) {
	rc.w.WriteHeader(s)
	rc.response.StatusCode = s
}

func (rc *responseCatcher) Flush() {
	f, ok := rc.w.(http.Flusher)
	if ok {
		f.Flush()
	}
}