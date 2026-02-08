package main

import (
	"context"
	"fmt"
	simpleconnection "study/feature_postgres/simple_connection"
	simplesql "study/feature_postgres/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := simpleconnection.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simplesql.CreateTable(conn, ctx); err != nil {
		panic(err)
	}

	if err := simplesql.InsertRow(ctx, conn, "Об", "пойти в школу пообедать", false, time.Now()); err != nil {
		panic(err)
	}

	tasks, err := simplesql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 5 {
			task.Title = "Выгул Собаки"
			task.Description = "погулять с собакой"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simplesql.UpdateTask(ctx, *conn, task); err != nil {
				panic(err)
			}
			break
		}
	}
	fmt.Println("sucseed ")

}
