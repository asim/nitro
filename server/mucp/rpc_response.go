package mucp

import (
	"net/http"

	"github.com/micro/go-micro/v3/codec"
	"github.com/micro/go-micro/v3/transport"
)

type rpcResponse struct {
	header map[string]string
	socket transport.Socket
	codec  codec.Codec
}

func (r *rpcResponse) Codec() codec.Writer {
	return r.codec
}

func (r *rpcResponse) WriteHeader(hdr map[string]string) {
	for k, v := range hdr {
		r.header[k] = v
	}
}

func (r *rpcResponse) Write(b []byte) error {
	if _, ok := r.header["Content-Type"]; !ok {
		r.header["Content-Type"] = http.DetectContentType(b)
	}

	return r.socket.Send(&transport.Message{
		Header: r.header,
		Body:   b,
	})
}
