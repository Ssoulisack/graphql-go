// api/graph/fiber_adapter.go

package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// FiberHandler creates a Fiber handler from a GraphQL handler
func FiberHandler(h *handler.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Adapter for converting fiber.Ctx to http.Request and Response
		var req http.Request
		if err := fasthttpadaptor.ConvertRequest(c.Context(), &req, true); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error converting request",
				"error":   err.Error(),
			})
		}

		// Create a response writer wrapper
		resWriter := NewResponseWriter(c)

		// Serve the GraphQL request
		h.ServeHTTP(resWriter, &req)

		return nil
	}
}

// ResponseWriter is a wrapper that adapts Fiber's Context to http.ResponseWriter
type ResponseWriter struct {
	ctx *fiber.Ctx
}

func NewResponseWriter(ctx *fiber.Ctx) *ResponseWriter {
	return &ResponseWriter{ctx: ctx}
}

func (rw *ResponseWriter) Header() http.Header {
	h := http.Header{}
	rw.ctx.Response().Header.VisitAll(func(key, value []byte) {
		h.Add(string(key), string(value))
	})
	return h
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	return rw.ctx.Write(b)
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ctx.Status(statusCode)
}
