package repository

import (
	"testing"

	trade "github.com/ShatAlex/trading-app"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestTradeCreate(t *testing.T) {
	db, mock, _ := sqlmock.Newx()
	defer db.Close()

	r := NewTradePostgres(db)

	type input struct {
		userId int
		trade  trade.Trade
	}

	tests := []struct {
		name   string
		input  input
		mock   func(input input, id int)
		wantId int
		retErr bool
	}{
		{
			name: "OK",
			input: input{
				userId: 1,
				trade: trade.Trade{
					Ticker: "AAPL",
					TypeId: 3,
					Price:  400,
					Amount: 2,
				},
			},
			mock: func(input input, id int) {

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO trades").WithArgs(
					input.trade.Ticker, input.userId, input.trade.TypeId, input.trade.Price,
					input.trade.Amount, input.trade.TypeId).WillReturnRows(rows)
			},
			wantId: 1,
			retErr: false,
		},
		{
			name: "Empty Ticker",
			input: input{
				userId: 1,
				trade: trade.Trade{
					TypeId: 3,
					Price:  400,
					Amount: 2,
				},
			},
			mock: func(input input, id int) {

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO trades").WithArgs(
					input.userId, input.trade.TypeId, input.trade.Price,
					input.trade.Amount, input.trade.TypeId).WillReturnRows(rows)
			},
			retErr: true,
		},
		{
			name: "Empty TypeId",
			input: input{
				userId: 1,
				trade: trade.Trade{
					Ticker: "AAPL",
					Price:  400,
					Amount: 2,
				},
			},
			mock: func(input input, id int) {

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO trades").WithArgs(
					input.trade.Ticker, input.userId, input.trade.Price,
					input.trade.Amount, input.trade.TypeId).WillReturnRows(rows)
			},
			retErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input, testCase.wantId)

			got, err := r.Create(testCase.input.userId, testCase.input.trade)
			if testCase.retErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.wantId, got)
			}
		})
	}
}

func TestTradeGetAll(t *testing.T) {
	db, mock, _ := sqlmock.Newx()
	defer db.Close()

	r := NewTradePostgres(db)

	type input struct {
		userId int
	}

	tests := []struct {
		name   string
		input  input
		mock   func()
		want   []trade.Trade
		retErr bool
	}{
		{
			name: "OK",
			input: input{
				userId: 1,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "ticker", "user_id", "type_id", "price", "amount"}).
					AddRow(1, "ticker1", 1, 1, 100, 5).
					AddRow(2, "ticker2", 1, 1, 100, 5).
					AddRow(3, "ticker3", 1, 1, 100, 5)

				mock.ExpectQuery("SELECT (.+) FROM trades").WithArgs(1).WillReturnRows(rows)
			},
			want: []trade.Trade{
				{1, "ticker1", 1, 1, 100, 5},
				{2, "ticker2", 1, 1, 100, 5},
				{3, "ticker3", 1, 1, 100, 5},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()
		})

		got, err := r.GetAll(testCase.input.userId)
		if testCase.retErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.want, got)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	}
}

// func TestTradeGetId(t *testing.T) {
// 	db, mock, _ := sqlmock.Newx()
// 	defer db.Close()

// 	r := NewTradePostgres(db)

// 	type input struct {
// 		userId  int
// 		tradeId int
// 	}

// 	tests := []struct {
// 		name   string
// 		input  input
// 		mock   func()
// 		want   trade.Trade
// 		retErr bool
// 	}{
// 		{
// 			name: "OK",
// 			input: input{
// 				userId:  1,
// 				tradeId: 1,
// 			},
// 			mock: func() {
// 				rows := sqlmock.NewRows([]string{"id", "ticker", "user_id", "type_id", "price", "amount"}).
// 					AddRow(1, "ticker1", 1, 1, 100, 5)

// 				mock.ExpectQuery("SELECT (.+) FROM trades").WithArgs(1).WillReturnRows(rows)
// 			},
// 			want: trade.Trade{
// 				{1, "ticker1", 1, 1, 100, 5},
// 			},
// 		},
// 	}

// 	for _, testCase := range tests {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			testCase.mock()
// 		})

// 		got, err := r.GetTradeById(testCase.input.userId, testCase.input.tradeId)
// 		if testCase.retErr {
// 			assert.Error(t, err)
// 		} else {
// 			assert.NoError(t, err)
// 			assert.Equal(t, testCase.want, got)
// 		}
// 		assert.NoError(t, mock.ExpectationsWereMet())
// 	}
// }
