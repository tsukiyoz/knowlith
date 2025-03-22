package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/enums/option"
	"github.com/parakeet-nest/parakeet/llm"
	"github.com/tsukiyoz/knowlith/internal/pkg/core"
	"github.com/tsukiyoz/knowlith/internal/pkg/errorsx"
	v1 "github.com/tsukiyoz/knowlith/pkg/api/apiserver/v1"
)

type Handler struct {
	url   string
	model string
}

func NewHandler() *Handler {
	return &Handler{
		url:   "http://localhost:11434",
		model: "llama3.2:latest",
	}
}

func (h *Handler) Prompt(c *fiber.Ctx) error {
	slog.Info("Prompt function called")

	var rq v1.PromptRequest
	if err := c.BodyParser(&rq); err != nil {
		return core.WriteResponse(c, nil, errorsx.ErrBind)
	}

	options := llm.SetOptions(map[string]interface{}{
		option.Temperature: 0.5,
	})

	question := llm.GenQuery{
		Model:   h.model,
		Prompt:  rq.Prompt,
		Options: options,
	}

	answer, err := completion.Generate(h.url, question)
	if err != nil {
		return core.WriteResponse(c, nil, err)
	}

	return core.WriteResponse(c, v1.PromptResponse{Answer: answer.Response}, nil)
}
