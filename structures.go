// Copyright (c) 2018, Oracle and/or its affiliates. All rights reserved.
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"net/http"
	"time"
)

// WerckerTestDemo is used to consolidate data items used in the demo
type WerckerTestDemo struct {
	client        *http.Client
	application   string
	username      string
	werckerToken  string
	branch        string
	pipeline      string
	commitHash    string
	runID         string
	appID         string
	sourceID      string
	pipelineID    string
	triggerObject *WerckerRun
}

// WerckerApplication is the response from getting an application request
type WerckerApplication struct {
	AllowedActions []string  `json:"allowedActions"`
	BadgeKey       string    `json:"badgeKey"`
	Builds         string    `json:"builds"`
	CreatedAt      time.Time `json:"createdAt"`
	Deploys        string    `json:"deploys"`
	HasWebHook     bool      `json:"hasWebHook"`
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Owner          struct {
		Avatar struct {
			Gravatar string `json:"gravatar"`
		} `json:"avatar"`
		Meta struct {
			Type            string `json:"type"`
			Username        string `json:"username"`
			WerckerEmployee bool   `json:"werckerEmployee"`
		} `json:"meta"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"owner"`
	Scm struct {
		Domain      string `json:"domain"`
		Owner       string `json:"owner"`
		Repository  string `json:"repository"`
		ScmProvider string `json:"scmProvider"`
		Type        string `json:"type"`
	} `json:"scm"`
	Settings struct {
		IgnoredBranches []string `json:"ignoredBranches"`
		Privacy         string   `json:"privacy"`
		Stack           int      `json:"stack"`
	} `json:"settings"`
	Theme     string    `json:"theme"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	URL       string    `json:"url"`
}

// WerckerRun is the format of an individual run wwhen listed
type WerckerRun struct {
	Branch     string `json:"branch"`
	CommitHash string `json:"commitHash"`
	Commits    []struct {
		ID      string `json:"_id"`
		By      string `json:"by"`
		Commit  string `json:"commit"`
		Message string `json:"message"`
	} `json:"commits"`
	CreatedAt time.Time     `json:"createdAt"`
	EnvVars   []interface{} `json:"envVars"`
	ID        string        `json:"id"`
	Message   string        `json:"message"`
	Pipeline  struct {
		CreatedAt            time.Time `json:"createdAt"`
		ID                   string    `json:"id"`
		ManualApproval       bool      `json:"manualApproval"`
		Name                 string    `json:"name"`
		Permissions          string    `json:"permissions"`
		PipelineName         string    `json:"pipelineName"`
		SetScmProviderStatus bool      `json:"setScmProviderStatus"`
		Type                 string    `json:"type"`
		URL                  string    `json:"url"`
	} `json:"pipeline"`
	Progress    int `json:"progress"`
	PullRequest struct {
	} `json:"pullRequest"`
	Result    string `json:"result"`
	SourceRun struct {
		Branch     string    `json:"branch"`
		CommitHash string    `json:"commitHash"`
		CreatedAt  time.Time `json:"createdAt"`
		FinishedAt time.Time `json:"finishedAt"`
		ID         string    `json:"id"`
		Message    string    `json:"message"`
		Pipeline   struct {
			CreatedAt            time.Time `json:"createdAt"`
			ID                   string    `json:"id"`
			ManualApproval       bool      `json:"manualApproval"`
			Name                 string    `json:"name"`
			Permissions          string    `json:"permissions"`
			PipelineName         string    `json:"pipelineName"`
			SetScmProviderStatus bool      `json:"setScmProviderStatus"`
			Type                 string    `json:"type"`
			URL                  string    `json:"url"`
		} `json:"pipeline"`
		Progress  int       `json:"progress"`
		Result    string    `json:"result"`
		StartedAt time.Time `json:"startedAt"`
		Status    string    `json:"status"`
		URL       string    `json:"url"`
		User      struct {
			Avatar struct {
				Gravatar string `json:"gravatar"`
			} `json:"avatar"`
			Meta struct {
				Type     string `json:"type"`
				Username string `json:"username"`
			} `json:"meta"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"user"`
	} `json:"sourceRun"`
	Status string `json:"status"`
	URL    string `json:"url"`
	User   struct {
		Avatar struct {
			Gravatar string `json:"gravatar"`
		} `json:"avatar"`
		Meta struct {
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"meta"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"user"`
	Workflow struct {
		CreatedAt time.Time `json:"createdAt"`
		Data      struct {
			Branch     string `json:"branch"`
			CommitHash string `json:"commitHash"`
			Message    string `json:"message"`
			Scm        struct {
				DevcsProjectID string `json:"devcsProjectId"`
				DevcsRoot      string `json:"devcsRoot"`
				Domain         string `json:"domain"`
				Owner          string `json:"owner"`
				Repository     string `json:"repository"`
				ScmProvider    string `json:"scmProvider"`
				Type           string `json:"type"`
			} `json:"scm"`
		} `json:"data"`
		ID    string `json:"id"`
		Items []struct {
			Data struct {
				CurrentStep int    `json:"currentStep"`
				PipelineID  string `json:"pipelineId"`
				Restricted  bool   `json:"restricted"`
				RunID       string `json:"runId"`
				StepName    string `json:"stepName"`
				TargetName  string `json:"targetName"`
				TotalSteps  int    `json:"totalSteps"`
			} `json:"data"`
			ID          string        `json:"id"`
			ParentItems []interface{} `json:"parentItems"`
			Progress    int           `json:"progress"`
			Result      string        `json:"result"`
			Status      string        `json:"status"`
			Type        string        `json:"type"`
			UpdatedAt   time.Time     `json:"updatedAt"`
			ParentItem  string        `json:"parentItem,omitempty"`
		} `json:"items"`
		StartedAt time.Time `json:"startedAt"`
		Trigger   string    `json:"trigger"`
		UpdatedAt time.Time `json:"updatedAt"`
		URL       string    `json:"url"`
		User      struct {
			Avatar struct {
				Gravatar string `json:"gravatar"`
			} `json:"avatar"`
			Meta struct {
				Type     string `json:"type"`
				Username string `json:"username"`
			} `json:"meta"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"user"`
	} `json:"workflow"`
}

// WerckerTrigger is the response to triggering a specific pipeline
type WerckerTrigger struct {
	Branch     string `json:"branch"`
	CommitHash string `json:"commitHash"`
	Commits    []struct {
		By         string `json:"by"`
		CommitHash string `json:"commitHash"`
		Message    string `json:"message"`
	} `json:"commits"`
	CreatedAt time.Time `json:"createdAt"`
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Pipeline  struct {
		CreatedAt    time.Time `json:"createdAt"`
		ID           string    `json:"id"`
		Name         string    `json:"name"`
		Permissions  string    `json:"permissions"`
		PipelineName string    `json:"pipelineName"`
		Type         string    `json:"type"`
		URL          string    `json:"url"`
	} `json:"pipeline"`
	Result    string `json:"result"`
	SourceRun struct {
		Branch     string `json:"branch"`
		CommitHash string `json:"commitHash"`
		Commits    []struct {
			By         string `json:"by"`
			CommitHash string `json:"commitHash"`
			Message    string `json:"message"`
		} `json:"commits"`
		CreatedAt time.Time `json:"createdAt"`
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Pipeline  struct {
			CreatedAt    time.Time `json:"createdAt"`
			ID           string    `json:"id"`
			Name         string    `json:"name"`
			Permissions  string    `json:"permissions"`
			PipelineName string    `json:"pipelineName"`
			Type         string    `json:"type"`
			URL          string    `json:"url"`
		} `json:"pipeline"`
		Result string `json:"result"`
		Status string `json:"status"`
		URL    string `json:"url"`
		User   struct {
			Avatar struct {
				Gravatar string `json:"gravatar"`
			} `json:"avatar"`
			Meta struct {
				Username string `json:"username"`
			} `json:"meta"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"user"`
	} `json:"sourceRun"`
	Status string `json:"status"`
	URL    string `json:"url"`
	User   struct {
		Avatar struct {
			Gravatar string `json:"gravatar"`
		} `json:"avatar"`
		Meta struct {
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"meta"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"user"`
	Workflow struct {
		CreatedAt time.Time `json:"createdAt"`
		Data      struct {
			Branch  string `json:"branch"`
			Message string `json:"message"`
		} `json:"data"`
		ID    string `json:"id"`
		Items []struct {
			Data struct {
				PipelineID string `json:"pipelineId"`
				RunID      string `json:"runId"`
				TargetName string `json:"targetName"`
			} `json:"data"`
			ID        string    `json:"id"`
			Result    string    `json:"result"`
			Status    string    `json:"status"`
			Type      string    `json:"type"`
			UpdatedAt time.Time `json:"updatedAt"`
		} `json:"items"`
		Trigger   string    `json:"trigger"`
		UpdatedAt time.Time `json:"updatedAt"`
		URL       string    `json:"url"`
		User      struct {
			Avatar struct {
				Gravatar string `json:"gravatar"`
			} `json:"avatar"`
			Meta struct {
				Type     string `json:"type"`
				Username string `json:"username"`
			} `json:"meta"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"user"`
	} `json:"workflow"`
}
