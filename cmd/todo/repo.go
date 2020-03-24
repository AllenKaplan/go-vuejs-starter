package todo

import (
	"github.com/google/uuid"
)

func GetTodos() []Todo {
	todo1 := Todo{uuid.New().ID(), "Clean my room"}
	todo2 := Todo{uuid.New().ID(), "Do my laundry"}
	todo3 := Todo{uuid.New().ID(), "Finish my homework"}
	todo4 := Todo{uuid.New().ID(), "Meal prep"}
	todo5 := Todo{uuid.New().ID(), "Go to the gym"}

	todos := []Todo{todo1, todo2, todo3, todo4, todo5}

	return todos

}
