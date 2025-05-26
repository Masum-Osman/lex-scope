package di

import (
	"context"
	"testing"

	"github.com/Masum-Osman/lex-scope/modules/text/usecase"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.uber.org/fx"
)

func TestDIWiring(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	var service usecase.TextService

	app := fx.New(
		fx.Provide(
			func() *mongo.Database {
				return mt.DB
			},
		),
		Module,
		fx.Populate(&service),
	)

	err := app.Start(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, service)
	_ = app.Stop(context.Background())

}
