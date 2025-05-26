package tests

import (
	"context"
	"testing"

	"github.com/Masum-Osman/lex-scope/modules/text/repository"
	"github.com/Masum-Osman/lex-scope/modules/text/usecase"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestTextCRUDIntegration(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	repo := repository.NewTextRepository(mt.DB)
	service := usecase.NewTextService(repo)

	ctx := context.Background()

	// CREATE
	id, err := service.Create(ctx, "Hello world. Another line.")
	assert.NoError(t, err)

	// READ
	text, err := service.Get(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, 2, text.SentenceCount)

	// UPDATE
	err = service.Update(ctx, id, "Updated text with more sentences. Another one!")
	assert.NoError(t, err)

	// DELETE
	err = service.Delete(ctx, id)
	assert.NoError(t, err)
}
