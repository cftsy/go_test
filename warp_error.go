package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func getData() error {
	err := errors.Wrap(sql.ErrNoRows, "get xxx sql_no_rows")
	return errors.WithMessage(err, "get_data_failed")
}

func main() {
	err := getData()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("%+v\n", err)
	} else if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println("get_data_succ")
	}
}