package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewToken(t *testing.T) {
	id := uuid.New()
	expiredAt := time.Minute

	payload, err := NewPayload(id, expiredAt)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.WithinDuration(t, time.Now(), payload.IssuedAt, time.Second)
	require.WithinDuration(t, time.Now().Add(expiredAt), payload.ExpiredAt, time.Second)

}
