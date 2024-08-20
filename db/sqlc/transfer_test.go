package db

import (
	"context"
	"database/sql"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAcountID: account1.ID,
		ToAccountID:  account2.ID,
		Amount:       util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAcountID, transfer.FromAcountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestUpdateTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:           transfer.ID,
		Amount:       util.RandomMoney(),
		FromAcountID: transfer.FromAcountID,
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, arg.Amount, transfer2.Amount)
	require.Equal(t, transfer.FromAcountID, transfer2.FromAcountID)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.Equal(t, transfer.FromAcountID, transfer2.FromAcountID)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
func TestListTransfer(t *testing.T) {

	for i := 0; i < 5; i++ {
		CreateRandomTransfer(t)
	}

	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
