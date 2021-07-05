package error_controller

import (
	"github.com/MichalH95/exampleREST/model/error_response"
	"github.com/gofiber/fiber"
)

func ServerErrorResponse(ctx *fiber.Ctx, errorMsg string) {
	sendErrorResponse(ctx, errorMsg, 500)
}

func ClientErrorResponse(ctx *fiber.Ctx, errorMsg string) {
	sendErrorResponse(ctx, errorMsg, 400)
}

func sendErrorResponse(ctx *fiber.Ctx, errorMsg string, errorCode int) {
	err := error_response.ErrorResponse{ErrorMsg: errorMsg}
	ctx.Status(errorCode).Send(err)
}
