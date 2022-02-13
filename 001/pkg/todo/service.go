package todo

type TodoService struct {
	inflation float32
}

func NewTodoService() (*TodoService, error) {
	s := &TodoService{
		inflation: 1.5,
	}

	return s, nil
}
