package handlers

import (
	"github.com/SamPariatIL/rrqd/services"
	"github.com/gin-gonic/gin"
)

type IssueHandler interface {
	GetIssues(ctx *gin.Context)
}

type issueHandler struct {
	issueService services.IssueService
}

func NewIssueHandler(is services.IssueService) IssueHandler {
	return &issueHandler{
		issueService: is,
	}
}

func (ih *issueHandler) GetIssues(ctx *gin.Context) {}
