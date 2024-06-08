package src

import (
	"context"
	"fmt"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/server"
)

type CustomHandler struct {
	server.EmptyHandler
	Pool *client.Pool
}

func (h *CustomHandler) HandleQuery(query string) (*mysql.Result, error) {

	if IsReadQuery(query) {
		fmt.Println("Received read query:", query)
	} else {
		fmt.Println("Received query:", query)
	}

	conn, err := h.Pool.GetConn(context.Background())

	if err != nil {
		return nil, err
	}

	res, err := conn.Execute(query)

	if err != nil {
		return nil, err
	}

	// Convert the result into a mysql.Resultset
	var resultSet mysql.Resultset
	resultSet.Fields = res.Fields
	resultSet.Values = res.Values
	resultSet.RowDatas = res.RowDatas

	return &mysql.Result{
		Status:       mysql.SERVER_STATUS_AUTOCOMMIT,
		AffectedRows: 0,
		InsertId:     0,
		Resultset:    &resultSet,
	}, nil
}
