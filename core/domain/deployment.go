package domain

import (
	"time"

	"github.com/sk000f/metrix/internal/cicd"
)

type Deployment struct {
	ID               int                   `json:"id" bson:"id"`
	Status           cicd.DeploymentStatus `json:"status" bson:"status"`
	EnvironmentName  cicd.EnvironmentName  `json:"environmentName" bson:"environment_name"`
	ProjectID        int                   `json:"projectId" bson:"project_id"`
	ProjectName      string                `json:"projectName" bson:"project_name"`
	ProjectPath      string                `json:"projectPath" bson:"project_path"`
	ProjectNamespace string                `json:"projectNamespace" bson:"project_namespace"`
	PipelineID       int                   `json:"pipelineId" bson:"pipeline_id"`
	FinishedAt       time.Time             `json:"finishedAt" bson:"finished_at"`
	Duration         int                   `json:"duration" bson:"duration"`
}
