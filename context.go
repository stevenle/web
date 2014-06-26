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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// A Context stores the current state of the application, including request and
// response data.
type Context struct {
	Request *http.Request
	Params  map[string]string

	response   http.ResponseWriter
	body       io.ReadWriter
	statusCode int
}

func newContext(request *http.Request, response http.ResponseWriter, params map[string]string) *Context {
	return &Context{
		Request: request,
		Params:  params,

		response:   response,
		body:       new(bytes.Buffer),
		statusCode: 200,
	}
}

// SetHeader sets the value of an HTTP response header.
func SetHeader(ctx *Context, key string, value string) {
	// TODO(stevenle): Disallow certain protected headers from being written.
	ctx.response.Header().Set(key, value)
}

// SetStatusCode sets the HTTP response status code.
func SetStatusCode(ctx *Context, statusCode int) {
	ctx.statusCode = statusCode
}

// WriteResponse writes a byte string to the body of the HTTP response.
func WriteResponse(ctx *Context, content []byte) {
	ctx.body.Write(content)
}

// WriteResponseString writes a utf-8 string to the body of the HTTP response.
func WriteResponseString(ctx *Context, content string) {
	io.WriteString(ctx.body, content)
}

// WriteResponseJson serializes an object to a JSON string and writes it to the
// body of the HTTP response.
func WriteResponseJson(ctx *Context, content interface{}) error {
	SetHeader(ctx, "Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.body)
	return encoder.Encode(content)
}

func renderResponse(ctx *Context) {
	ctx.response.WriteHeader(ctx.statusCode)
	io.Copy(ctx.response, ctx.body)
}
