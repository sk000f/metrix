package domain

import "time"

type Deployment struct {
	ID               int        `json:"id"`
	Status           string     `json:"status"`
	EnvironmentName  string     `json:"environmentName"`
	ProjectID        int        `json:"projectId"`
	ProjectName      string     `json:"projectName"`
	ProjectPath      string     `json:"projectPath"`
	ProjectNamespace string     `json:"projectNamespace"`
	PipelineID       int        `json:"pipelineId"`
	FinishedAt       *time.Time `json:"finishedAt"`
	Duration         float64    `json:"duration"`
}
