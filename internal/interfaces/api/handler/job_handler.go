package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JobDTO 任务数据传输对象，对接前端表单
type JobDTO struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Cron    string `json:"cron"`
	Command string `json:"command"`
	Status  int    `json:"status"` // 0: 停止, 1: 运行中
}

type JobHandler struct {
	// 以后可以注入 Service 或 Repository
}

func NewJobHandler() *JobHandler {
	return &JobHandler{}
}

// List 获取任务列表
func (h *JobHandler) List(c *gin.Context) {
	// TODO: 从存储获取任务列表
	c.JSON(http.StatusOK, gin.H{"message": "Jobs list (decoupled)"})
}

// Save 保存或更新任务
func (h *JobHandler) Save(c *gin.Context) {
	var job JobDTO
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: 写入存储
	c.JSON(http.StatusOK, gin.H{"message": "Job saved (decoupled)", "data": job})
}

// Delete 删除任务
func (h *JobHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	// TODO: 从存储中删除
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted (decoupled)", "id": id})
}

// RunNow 立即触发一次任务
func (h *JobHandler) RunNow(c *gin.Context) {
	id := c.Param("id")
	// TODO: 调度逻辑
	c.JSON(http.StatusOK, gin.H{"message": "Job triggered (decoupled)", "id": id})
}
