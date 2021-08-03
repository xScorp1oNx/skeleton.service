package routing

import (
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HandlesChain struct {
	head *HandleChainLeaf
	tail *HandleChainLeaf
}

func (hg *HandlesChain) append(handle HttpHandler) {
	leaf := HandleChainLeaf{handler: handle}

	if hg.head == nil && hg.tail == nil {
		hg.head = &leaf
		hg.tail = &leaf
	} else {
		hg.tail.next = &leaf
		hg.tail = &leaf
	}
}

type HandleChainLeaf struct {
	handler HttpHandler
	next    *HandleChainLeaf
}

func (hg *HandleChainLeaf) Handle(w http.ResponseWriter, r *http.Request) {
	hg.handler(w, r)
	if hg.next != nil {
		hg.next.Handle(w, r)
	}
}

func MakeHandlesChain(handlers ...HttpHandler) http.HandlerFunc {
	chain := HandlesChain{}

	for _, h := range handlers {
		chain.append(h)
	}

	return chain.head.Handle
}
