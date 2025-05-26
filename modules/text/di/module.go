package di

import (
	"github.com/Masum-Osman/lex-scope/modules/text/handler"
	"github.com/Masum-Osman/lex-scope/modules/text/repository"
	"github.com/Masum-Osman/lex-scope/modules/text/usecase"
	"go.uber.org/fx"
)

var Module = fx.Module("text-module",
	fx.Provide(
		repository.NewTextRepository,
		usecase.NewTextService,
		handler.NewTextHandler,
	),
)
