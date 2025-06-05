package Routes;

import (
        "fmt"
        "github.com/gin-gonic/gin"
        cntr "github.com/samObot19/backend/taskManagerApi/controller"
)

func StartRoute(){
        fmt.Println("Task manager API")
        router := gin.Default()
        router.GET("api/tasks", cntr.GetTasks)
        router.GET("api/tasks/:id", cntr.GetTask)
        router.DELETE("api/tasks/:id", cntr.RemoveTask)
        router.PUT("api/tasks/:id", cntr.UpdateTask)
        router.POST("api/tasks", cntr.AddTask)
        router.Run()
}
