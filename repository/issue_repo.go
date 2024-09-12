package repository

import "go.mongodb.org/mongo-driver/mongo"

type IssueRepository interface {
	GetIssues()
}

type issueRepository struct {
	db *mongo.Database
}

func NewIssueRepository(db *mongo.Database) IssueRepository {
	return &issueRepository{db: db}
}

func (ir *issueRepository) GetIssues() {}
