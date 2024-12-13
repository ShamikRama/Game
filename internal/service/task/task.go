package task

import (
	"Game/internal/repository"
	"fmt"
	"strings"
)

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

func (r *TaskService) EnterRefCode(userID int, referrerID int) error {
	pointsForEnterCode := 6
	pointsForReferral := 13

	if userID == referrerID {
		return fmt.Errorf("can not enter own ref_code")
	}

	exist, err := r.repo.UserExists(referrerID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("referrer by ID not found")
	}

	err = r.repo.CompleteRef(userID, referrerID)
	if err != nil {
		if strings.Contains(err.Error(), "twice ref code") {
			return fmt.Errorf("unique constraint")
		}
		return err
	}

	err = r.repo.UpdatePoints(userID, pointsForEnterCode)
	if err != nil {
		return err
	}

	err = r.repo.UpdatePoints(referrerID, pointsForReferral)
	if err != nil {
		return err
	}

	return nil
}
