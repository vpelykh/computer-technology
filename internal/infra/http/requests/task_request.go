package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline"`
}

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	var deadline *time.Time
	if r.Deadline != 0 {
		dl := time.Unix(r.Deadline, 0)
		deadline = &dl
	}
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Deadline:    deadline,
	}, nil
}

func ParseTaskId(r *http.Request) (uint64, error) {
	taskIdStr := chi.URLParam(r, "taskId")
	taskId, err := strconv.ParseUint(taskIdStr, 10, 64)
	if err != nil {
	  return 0, err
	}
	return taskId, nil
  }
  type UpdateTaskRequest struct {
	Title    *string            json:"title"
	Status   *domain.TaskStatus json:"status"
	Deadline *uint64            json:"deadline"
  }
