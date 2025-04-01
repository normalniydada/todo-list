package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todo-list/internal/api/dto"
	"todo-list/internal/domain/service"
)

type TaskHandler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type taskHandlerImpl struct {
	service service.TaskService
}

func NewTaskHandler(s service.TaskService) TaskHandler {
	return &taskHandlerImpl{service: s}
}

func (h *taskHandlerImpl) Create(c echo.Context) error {
	var req dto.TaskRequestDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	createdTask, err := h.service.CreateTask(req.Title, req.Content)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.ToTaskResponseDTO(createdTask))

}
func (h *taskHandlerImpl) List(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}
func (h *taskHandlerImpl) Get(c echo.Context) error {
	id := c.Param("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Task is not found"})
	}

	return c.JSON(http.StatusOK, dto.ToTaskResponseDTO(task))
}
func (h *taskHandlerImpl) Update(c echo.Context) error {
	id := c.Param("id")
	var req dto.TaskRequestDTO

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	task, err := h.service.UpdateTask(id, req.Title, req.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error id / Title or Content is empty"})
	}

	return c.JSON(http.StatusOK, dto.ToTaskResponseDTO(task))
}
func (h *taskHandlerImpl) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Task is not found"})
	}

	return c.NoContent(http.StatusNoContent)
}
