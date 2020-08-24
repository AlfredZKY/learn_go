package main

import (
	"context"
	"fmt"
	"time"
)

const DB_ADDRESS = "db_address"
const CALCULATE_VALUE = "calculate_value"

func readDB(ctx context.Context, cost time.Duration) {
	fmt.Println("db address is", ctx.Value(DB_ADDRESS))
	select {
	case <-time.After(cost):
		fmt.Println("read data from db")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func calculate(ctx context.Context, cost time.Duration) {
	fmt.Println("calcuaate value is", ctx.Value(CALCULATE_VALUE))
	select {
	case <-time.After(cost):
		fmt.Println("calculate finish")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, DB_ADDRESS, "localhost:10086")
	ctx = context.WithValue(ctx, CALCULATE_VALUE, 1234)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	go readDB(ctx, time.Second*4)
	go calculate(ctx, time.Second*4)

	time.Sleep(5 * time.Second)
}
