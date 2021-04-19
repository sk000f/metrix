package gitlab_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/sk000f/metrix/adapters/ciserver/gitlab"
	"github.com/sk000f/metrix/core/domain"
)

func TestGitLab(t *testing.T) {
	t.Run("get all deployments from GitLab", func(t *testing.T) {
		mux, server, g := setupMockGitLabClient(t)
		defer server.Close()

		// setup mock project
		mux.HandleFunc("/api/v4/projects", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `[{
					"id": 1, 
					"name": "test",
					"path": "test",
					"path_with_namespace": "test/test", 
					"web_url": "http://test.com/test/test",
					"namespace" :{
						"full_path": "test/test"
					}
					}]`)
		})

		// setup mock deployments
		mux.HandleFunc("/api/v4/projects/1/deployments", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `[
				{
					"id": 1, 
					"status": "success",
					"environment": {
						"name": "production"
					}, 
					"deployable": { 
						"finished_at": "2020-10-06T15:30:53.355Z",
						"duration": 123,
						"pipeline": {
							"id": 1
						}
					}
				}
				]`)
		})

		timestamp, e := time.Parse(time.RFC3339, "2020-10-06T15:30:53.355Z")
		if e != nil {
			t.Errorf(e.Error())
		}

		want := []domain.Deployment{{
			ID:               1,
			Status:           "success",
			EnvironmentName:  "production",
			ProjectID:        1,
			ProjectName:      "test",
			ProjectPath:      "test",
			ProjectNamespace: "test/test",
			PipelineID:       1,
			FinishedAt:       timestamp,
			Duration:         123,
		}}

		got, err := g.GetAllDeployments()
		if err != nil {
			t.Errorf("Error getting Deployments: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v; wanted %+v", got, want)
		}

	})
}

func setupMockGitLabClient(t *testing.T) (*http.ServeMux, *httptest.Server, *gitlab.GitLab) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)

	g := &gitlab.GitLab{
		Token: "",
		URL:   server.URL,
	}

	err := g.SetupClient(g.Token, g.URL)
	if err != nil {
		server.Close()
		t.Fatalf("Error creating mock GitLab client: %v", err)
	}

	return mux, server, g
}
