package usecases

import (
    "context"
    "task-manager/Domain"
    "task-manager/Repositories"
)

type TaskUsecase struct {
    taskRepo *repositories.TaskRepository
}

func NewTaskUsecase(taskRepo *repositories.TaskRepository) *TaskUsecase {
    return &TaskUsecase{
        taskRepo: taskRepo,
    }
}

func (u *TaskUsecase) CreateTask(ctx context.Context, task *domain.Task) error {
    return u.taskRepo.CreateTask(ctx, task)
}

func (u *TaskUsecase) GetTasksByUserID(ctx context.Context, userID string) ([]*domain.Task, error) {
    return u.taskRepo.GetTasksByUserID(ctx, userID)
}