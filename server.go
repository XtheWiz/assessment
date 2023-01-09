package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xthewiz/assessment/expense"

	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Stating application... ")
	db := expense.InitDB()
	server := setupRoute(db)
	start(server)
}

func start(server *echo.Echo) {
	if err := server.Start(":2565"); err != nil && err != http.ErrServerClosed {
		server.Logger.Fatal("Shutting down server...")
	}

	server.Logger.Info("Server started ..")
}

func setupRoute(db *sql.DB) *echo.Echo {
	serv := echo.New()
	serv.Use(middleware.Logger())
	serv.Use(middleware.Recover())

	expenseHl := expense.InitHandler(db)
	serv.POST("/expenses", expenseHl.CreateExpenseHandler)
	serv.GET("/expenses", expenseHl.GetExpensesHandler)
	serv.GET("/expenses/:id", expenseHl.GetExpensesByIDHandler)
	serv.PUT("/expenses/:id", expenseHl.UpdateExpenseHandler)

	return serv
}
