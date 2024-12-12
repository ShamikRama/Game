package task

import "Game/internal/repository"

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (r *TaskService) CompleteTaskTelegram(userID int) error {
	goal_type := "Telegram"
	points := 10

	err := r.repo.CompleteTask(userID, goal_type)
	if err != nil {
		return err
	}

	err = r.repo.UpdatePoints(userID, points)
	if err != nil {
		return err
	}

	return nil
}
