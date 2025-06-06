package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Models "github.com/samObot19/backend/taskManagerApi/model"
	service "github.com/samObot19/backend/taskManagerApi/services"
)

func AddTask(ctx *gin.Context) {
	var newTask Models.Task

	err := ctx.ShouldBindJSON(&newTask)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.AddTask(&newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	val, err := service.GetTask(&id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	} else {
		ctx.JSON(http.StatusOK, val)
	}

}

func GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"tasks": service.GetTasks()})

}

func RemoveTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := service.RemoveTask(&id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
	}

}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := service.UpdateTask(&id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
	}

}
