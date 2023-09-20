package repository

import (
	"fmt"
	"testing"

	trade "github.com/ShatAlex/trading-app"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestTypeTradeCreate(t *testing.T) {
	db, mock, _ := sqlmock.Newx()
	defer db.Close()

	r := NewTypeTradePostgres(db)

	type input struct {
		typeTrade trade.TypeTrade
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
				typeTrade: trade.TypeTrade{
					Trade_type: "test",
				},
			},
			mock: func(input input, id int) {

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO types").WithArgs(
					input.typeTrade.Trade_type).WillReturnRows(rows)
			},
			wantId: 1,
			retErr: false,
		},
		{
			name:  "empty input",
			input: input{},
			mock: func(input input, id int) {

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO types").WithArgs().WillReturnRows(rows)
			},
			retErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input, testCase.wantId)

			got, err := r.Create(testCase.input.typeTrade)
			if testCase.retErr {
				assert.Error(t, err)
			} else {
				fmt.Println(got)
				fmt.Println(err)
				assert.Equal(t, testCase.wantId, got)
			}
		})
	}
}
