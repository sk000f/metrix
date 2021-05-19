package gitlab

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sk000f/metrix/core/domain"
	"github.com/sk000f/metrix/internal/cicd"
	gl "github.com/xanzy/go-gitlab"
)

type GitLab struct {
	Token  string
	URL    string
	Client *gl.Client
}

func New(t, URL string) (*GitLab, error) {

	g := &GitLab{}

	err := g.SetupClient(t, URL)
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("gitlab.GetAllDeployments")
		return nil, err
	}

	return g, nil
}

func (g *GitLab) GetAllDeployments() ([]domain.Deployment, error) {

	d := []domain.Deployment{}

	p, err := g.getAllProjects()
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("gitlab.GetAllDeployments")
		return nil, err
	}

	opt := getDeploymentListOptions()

	for _, proj := range p {
		for {

			deployments, resp, err := g.Client.Deployments.ListProjectDeployments(proj.ID, opt)
			if err != nil {
				log.Error().Stack().Err(err).
					Msg("gitlab.GetAllDeployments")
				return nil, err
			}

			for _, dep := range deployments {

				if dep.Environment.Name == "production" &&
					(dep.Status == "success" || dep.Status == "failed") {

					status := convertStatus(dep.Status)
					envName := convertEnvName(dep.Environment.Name)

					d = append(d, domain.Deployment{
						ID:               dep.ID,
						Status:           status,
						EnvironmentName:  envName,
						ProjectID:        proj.ID,
						ProjectName:      proj.Name,
						ProjectPath:      proj.Path,
						ProjectNamespace: proj.Namespace,
						PipelineID:       dep.Deployable.Pipeline.ID,
						FinishedAt:       *dep.Deployable.FinishedAt,
						Duration:         int(dep.Deployable.Duration),
					})
				}

			}

			if resp.CurrentPage >= resp.TotalPages {
				opt.Page = 1
				break
			}

			opt.Page = resp.NextPage
		}
	}

	return d, nil
}

func (g *GitLab) GetDeploymentsByProject(proj string) ([]domain.Deployment, error) {
	return nil, nil
}

func (g *GitLab) GetDeploymentsByDateRange(start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (g *GitLab) GetDeploymentsByProjectAndDateRange(proj string, start time.Time, end time.Time) ([]domain.Deployment, error) {
	return nil, nil
}

func (g *GitLab) getAllProjects() ([]Project, error) {

	p := []Project{}

	opt := getProjectListOptions()

	for {

		projects, resp, err := g.Client.Projects.ListProjects(opt)
		if err != nil {
			log.Error().Stack().Err(err).
				Msg("gitlab.getAllProjects")
			return nil, err
		}

		for _, pr := range projects {
			p = append(p, Project{
				ID:                pr.ID,
				Name:              pr.Name,
				Path:              pr.Path,
				PathWithNamespace: pr.PathWithNamespace,
				Namespace:         pr.Namespace.FullPath,
				WebURL:            pr.WebURL,
			})
		}

		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		opt.Page = resp.NextPage
	}

	return p, nil
}

func (g *GitLab) SetupClient(token, baseURL string) error {
	client, err := gl.NewClient(token, gl.WithBaseURL(baseURL))
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("gitlab.SetupClient")
		return err
	}

	g.Client = client
	return nil
}

func getProjectListOptions() *gl.ListProjectsOptions {
	return &gl.ListProjectsOptions{
		ListOptions: gl.ListOptions{Page: 1, PerPage: 20},
		Simple:      gl.Bool(false),
	}
}

func getDeploymentListOptions() *gl.ListProjectDeploymentsOptions {
	return &gl.ListProjectDeploymentsOptions{
		ListOptions: gl.ListOptions{Page: 1, PerPage: 20},
		Environment: gl.String("production"),
	}
}

func convertStatus(s string) cicd.DeploymentStatus {
	switch s {
	case string(cicd.Success):
		return cicd.Success
	case string(cicd.Failure):
		return cicd.Failure
	default:
		return ""
	}
}

func convertEnvName(e string) cicd.EnvironmentName {
	switch e {
	case string(cicd.Development):
		return cicd.Development
	case string(cicd.Staging):
		return cicd.Staging
	case string(cicd.PreProduction):
		return cicd.PreProduction
	case string(cicd.Production):
		return cicd.Production
	default:
		return ""
	}
}

// Project represents a GitLab Project object
type Project struct {
	ID                int
	Name              string
	Path              string
	PathWithNamespace string
	Namespace         string
	WebURL            string
}
