package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type VerifyEmalTxParams struct {
	EmailId    int64
	SecretCode string
}

type VerifyEmalTxResult struct {
	User        User
	VerifyEmail VerifyEmail
}

func (store *SQLStore) VerifyEmalTx(ctx context.Context, arg VerifyEmalTxParams) (VerifyEmalTxResult, error) {
	var result VerifyEmalTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Username: result.VerifyEmail.Username,
			IsEmailVerified: pgtype.Bool{
				Bool:  true,
				Valid: true,
			},
		})
		return err
	})

	return result, err
}
