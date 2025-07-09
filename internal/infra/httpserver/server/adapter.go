package server

import (
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

func adaptRoute(ctrl httpserver.IController) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		request := httpserver.Request{
			Host:    ctx.Hostname(),
			Path:    ctx.Path(),
			Method:  ctx.Method(),
			Headers: getHeaders(ctx.GetReqHeaders()),
			Body:    getBody(ctx.Body()),
			Params:  getParams(ctx),
			Query:   getQuery(ctx),
		}
		slog.Info("Request receveid:["+request.Host+request.Path+"]", slog.Any("request", string(request.Body)))
		response := ctrl.Handle(ctx.UserContext(), request)

		return ctx.Status(response.Code).JSON(response.Body)
	}
}

func adaptRouteFile(ctrl httpserver.IController) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("video")
		if err != nil {
			slog.Error("Error receiving file", slog.Any("error", err))
		}

		slog.Info("File receveid:[" + file.Filename + "]")

		request := httpserver.Request{
			Host:    ctx.Hostname(),
			Path:    ctx.Path(),
			Method:  ctx.Method(),
			Headers: getHeaders(ctx.GetReqHeaders()),
			Body:    getBody(ctx.Body()),
			Params:  getParams(ctx),
			Query:   getQuery(ctx),
			File:    file,
		}
		// return ctx.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

		slog.Info("Request receveid:["+request.Host+request.Path+"]", slog.Any("file", file.Filename))
		response := ctrl.Handle(ctx.UserContext(), request)

		return ctx.Status(response.Code).JSON(response.Body)
	}
}

func getHeaders(headers map[string][]string) map[string]string {
	newHeaders := map[string]string{}
	for k, v := range headers {
		newHeaders[strings.ToLower(k)] = v[0]
	}
	return newHeaders
}

func getParams(c *fiber.Ctx) map[string]string {
	values := map[string]string{}
	args := c.Route().Params
	for _, v := range args {
		values[v] = c.Params(v)
	}
	return values
}

func getQuery(c *fiber.Ctx) map[string]string {
	values := map[string]string{}
	args := c.Context().QueryArgs()

	args.VisitAll(func(key, value []byte) {
		k := string(key)
		v := string(value)

		values[k] = v
	})

	return values
}

func getBody(mutableBytes []byte) []byte {
	buffer := make([]byte, len(mutableBytes))
	copy(buffer, mutableBytes)
	return buffer
}
