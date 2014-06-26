// Copyright 2014 Steven Le. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"net/http"

	"github.com/stevenle/routetrie"
)

type Handler interface {
	Handle(*Context)
}

type HandlerFunc func(*Context)

type Router struct {
	routes *routetrie.RouteTrie
}

// NewRouter returns a new instance of a router.
func NewRouter() *Router {
	return &Router{
		routes: routetrie.NewRouteTrie(),
	}
}

func (router *Router) Handle(pattern string, handler Handler) {
	router.routes.Add(pattern, handler)
}

func (router *Router) HandleFunc(pattern string, handlerFunc HandlerFunc) {
	router.routes.Add(pattern, handlerFunc)
}

// ServeHTTP implements the interface for http.Handler.
func (router *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	result, params := router.routes.Get(request.URL.Path)
	ctx := newContext(request, response, params)

	if result != nil {
		switch h := result.(type) {
		case Handler:
			h.Handle(ctx)
		case HandlerFunc:
			h(ctx)
		default:
			// This should never really happen, but in case it does...
			SetStatusCode(ctx, http.StatusNotFound)
		}
	} else {
		SetStatusCode(ctx, http.StatusNotFound)
	}

	renderResponse(ctx)
}
