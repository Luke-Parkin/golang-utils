package safeio

import (
	"bytes"
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
)

func TestWriteString(t *testing.T) {
	var buf bytes.Buffer
	text := faker.Sentence()
	n, err := WriteString(context.Background(), &buf, text)
	require.NoError(t, err)
	assert.Equal(t, len(text), n)
	assert.Equal(t, text, buf.String())
	buf.Reset()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cancel()
	n, err = WriteString(ctx, &buf, text)
	require.Error(t, err)
	assert.True(t, commonerrors.Any(err, commonerrors.ErrCancelled))
	assert.Zero(t, n)
	assert.Empty(t, buf.String())
}
