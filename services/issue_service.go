package services

import "github.com/SamPariatIL/rrqd/repository"

type IssueService interface {
	GetIssues()
}

type issueService struct {
	issueRepository repository.IssueRepository
}

func NewIssueService(ir repository.IssueRepository) IssueService {
	return &issueService{
		issueRepository: ir,
	}
}

func (is *issueService) GetIssues() {}
