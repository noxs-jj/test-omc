package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/noxs-jj/test-omc/001/pkg/todo"
)

const (
	EXIT_SUCCESS = "EXIT_SUCCESS"
	EXIT_FAILED  = "EXIT_FAILED"
)

// # missing

type AppServices struct {
	todoService TodoService
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(EXIT_FAILED)
		os.Exit(-1)
	}

	/*
		# missing
	*/

	fmt.Fprintf(os.Stderr, "income[%v] | months[%v]\n", income, months)

	ctx := context.TODO()
	todoSvc, err := todo.NewTodoService()
	if err != nil {
		fmt.Printf("%s: %s\n", EXIT_FAILED, err.Error())
		os.Exit(-1)
	}

	app := AppServices{
		todoService: todoSvc,
	}

	result, err := app.todoService.AddMonthsInflation(ctx, float32(income), months)
	if err != nil {
		fmt.Printf("%s: %s\n", EXIT_FAILED, err.Error())
		os.Exit(-1)
	}

	fmt.Printf("with income[%d] and months:[%d], result is [%.2f]\n", int(income), months, result)

	fmt.Println(EXIT_SUCCESS)
	os.Exit(0)
}
