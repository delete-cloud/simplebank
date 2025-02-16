package gapi

import (
	"testing"
	"time"

	db "github.com/delete-cloud/simplebank/db/sqlc"
	"github.com/delete-cloud/simplebank/util"
	"github.com/delete-cloud/simplebank/worker"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
