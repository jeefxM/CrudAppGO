package main

import (
	"context"

	"goLang/db"

	"github.com/gofiber/fiber/v2"
)

// Global variable to store the createdPost


func main() {
    client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer client.Prisma.Disconnect()

	app := fiber.New()
    
    app.Post("/action", func(c *fiber.Ctx) error {

        todoInput := new(struct{
            Todo string `json:"Todo"`
            Done bool `json:"done"`
        })
       	if err := c.BodyParser(todoInput); err != nil {
			return c.Status(400).SendString("Failed to parse input")
		}

        createdToDo, err := client.TodoList.CreateOne(
            db.TodoList.Todo.Set(todoInput.Todo),
            db.TodoList.Done.Set(todoInput.Done),
        ).Exec(context.Background())
        if err != nil {
            return c.Status(500).SendString("Failed to create todo")
        }
        return c.JSON(createdToDo)

    })
    app.Get("/todo/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		post, err := client.TodoList.FindUnique(
			db.TodoList.ID.Equals(id),
		).Exec(context.Background())
		if err != nil {
			return c.Status(404).SendString("Post not found")
		}

		return c.JSON(post)
	})

    app.Delete("/todo/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

        todo, err := client.TodoList.FindUnique(
            db.TodoList.ID.Equals(id),
        ).Delete().Exec(context.Background())
        if err != nil {
            return c.Status(404).SendString("Post not found")
        }
        return c.JSON(todo)
        })

    app.Put("/todo/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")

        

        newTodoInput := new(struct{
            Todo string `json:"Todo"`
            Done bool `json:"done"`
        })

        if err := c.BodyParser(newTodoInput); err != nil {
            return c.Status(400).SendString("Failed to parse input")
        }

        newTodo, err := client.TodoList.FindUnique(
            db.TodoList.ID.Equals(id),
        ).Update(
            db.TodoList.Todo.Set(newTodoInput.Todo),
            db.TodoList.Done.Set(newTodoInput.Done),
        ).Exec(context.Background())
        if err != nil {
            return c.Status(404).SendString("Post not found")
        }

        return c.JSON(newTodo)
    })
        

	app.Listen(":3000")
}


