package logger

import "go.uber.org/fx"

var Module = fx.Provide(
	func() (Logger, error) {
		return NewLogger()
	},
)
