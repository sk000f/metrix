package cicd

type EnvironmentName string

const (
	Production    EnvironmentName = "production"
	PreProduction EnvironmentName = "preproduction"
	Staging       EnvironmentName = "staging"
	Development   EnvironmentName = "development"
)

type DeploymentStatus string

const (
	Success DeploymentStatus = "success"
	Failure DeploymentStatus = "failure"
)
