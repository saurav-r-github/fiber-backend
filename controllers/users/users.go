package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

const coll = client.Database("school").Collection("students") // Database collection

func Home(c *fiber.Ctx) error {
	type Address struct {
		Street string
		City   string
		State  string
	}

	type Student struct {
		FirstName string  `bson:"first_name,omitempty"`
		LastName  string  `bson:"last_name,omitempty"`
		Address   Address `bson:"inline"`
		Age       int
	}

	address1 := Address{"1 Lakewood Way", "Elwood City", "PA"}
	student1 := Student{FirstName: "Arthur", Address: address1, Age: 8}
	createdStudent, err := coll.InsertOne(context.TODO(), student1)
	return c.SendString(createdStudent)
}
