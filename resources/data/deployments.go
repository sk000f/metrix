package data

import (
	"time"

	"github.com/sk000f/metrix/core/domain"
)

var SampleDeployments = []domain.Deployment{
	{
		ID:               1,
		Status:           "success",
		EnvironmentName:  "production",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       1,
		FinishedAt:       time.Now(),
		Duration:         12,
	},
	{
		ID:               2,
		Status:           "success",
		EnvironmentName:  "production",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       2,
		FinishedAt:       time.Now(),
		Duration:         43,
	},
	{
		ID:               3,
		Status:           "failure",
		EnvironmentName:  "production",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       3,
		FinishedAt:       time.Now(),
		Duration:         14,
	},
	{
		ID:               4,
		Status:           "failure",
		EnvironmentName:  "staging",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       4,
		FinishedAt:       time.Now(),
		Duration:         28,
	},
	{
		ID:               5,
		Status:           "success",
		EnvironmentName:  "staging",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       5,
		FinishedAt:       time.Now(),
		Duration:         19,
	},
	{
		ID:               6,
		Status:           "success",
		EnvironmentName:  "production",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       6,
		FinishedAt:       time.Now(),
		Duration:         54,
	},
	{
		ID:               7,
		Status:           "success",
		EnvironmentName:  "production",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       7,
		FinishedAt:       time.Now(),
		Duration:         123,
	},
	{
		ID:               8,
		Status:           "success",
		EnvironmentName:  "production",
		ProjectID:        2,
		ProjectName:      "dummy-project",
		ProjectPath:      "dummy-group/dummy-project",
		ProjectNamespace: "dummy-group",
		PipelineID:       8,
		FinishedAt:       time.Now(),
		Duration:         123,
	},
	{
		ID:               9,
		Status:           "success",
		EnvironmentName:  "production",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       9,
		FinishedAt:       time.Now(),
		Duration:         43,
	},
	{
		ID:               10,
		Status:           "success",
		EnvironmentName:  "production",
		ProjectID:        1,
		ProjectName:      "test-project",
		ProjectPath:      "test-group/test-project",
		ProjectNamespace: "test-group",
		PipelineID:       10,
		FinishedAt:       time.Now(),
		Duration:         76,
	},
}
