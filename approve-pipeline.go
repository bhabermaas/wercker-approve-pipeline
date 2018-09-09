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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	args := os.Args
	if len(args) != 6 {
		fmt.Println("There must be five arguments")
		fmt.Println("username application pipeline-name commitHash wercker-token")
		return
	}

	// Assuming branch=master

	// Setup a convienent place to keep all our goodies.
	demo := &WerckerTestDemo{
		client:       &http.Client{}, // http client for talking to wercker
		branch:       "master",       // what branch
		username:     args[1],        // name of the wercker user
		application:  args[2],        // name of the application
		pipeline:     args[3],        // pipeline name to be approved
		commitHash:   args[4],        // commit ID we are interested in
		werckerToken: args[5],        // auth token
	}

	// Get the applicationID for the application
	err := demo.getApplicationId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s (%s) pipeline=%s commitHash=%s\n", demo.application, demo.appID, demo.pipeline, demo.commitHash)

	// Get the specific run for the pipeline to be triggered
	err = demo.getPipelineToApprove()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Trigger and approve the manual pipeline
	if demo.runID == "" {

		fmt.Println("Successfully triggered and approved the pipeline.")
		return
	}
}

// Fetch and save the application Id
func (p *WerckerTestDemo) getApplicationId() error {
	url := fmt.Sprintf("https://app.wercker.com/api/v3/applications/%s/%s", p.username, p.application)
	body, err := p.getWercker(url)
	if err != nil {
		return err
	}
	application := WerckerApplication{}
	err = json.Unmarshal(body, &application)
	if err != nil {
		return err
	}
	p.appID = application.ID
	return nil
}

// Get the pipeline to be approved and store it's run object
func (p *WerckerTestDemo) getPipelineToApprove() error {
	url := fmt.Sprintf("https://app.wercker.com/api/v3/runs?applicationId=%s&commitHash=%s", p.appID, p.commitHash)
	body, err := p.getWercker(url)
	if err != nil {
		return err
	}
	runs := []WerckerRun{}
	err = json.Unmarshal(body, &runs)
	if err != nil {
		return err
	}

	hasManual := false

	// Looking through all the runs for our commitHash
	for _, thisRun := range runs {
		if thisRun.CommitHash != p.commitHash || thisRun.Pipeline.PipelineName != p.pipeline {
			continue
		}

		// At this point the run object from the list is not fully populated so we need to
		// specifically request it to get all the goodies.
		url = fmt.Sprintf("https://app.wercker.com/api/v3/runs/%s", thisRun.ID)
		body, err := p.getWercker(url)
		if err != nil {
			return err
		}
		run := WerckerRun{}
		err = json.Unmarshal(body, &run)
		if err != nil {
			return err
		}

		// Save sourceID in case a trigger is necessary.
		p.sourceID = run.SourceRun.ID

		msg := fmt.Sprintf("runId=%s pipeline=%s (%s) status=%s", run.ID, run.Pipeline.PipelineName,
			run.Pipeline.ID, run.Status)
		fmt.Println(msg)

		if run.Pipeline.Name == p.pipeline && run.Pipeline.ManualApproval {
			hasManual = true
			p.pipelineID = run.Pipeline.ID
			if run.Status != "pendingapproval" {
				continue
			}
			p.triggerObject = &run
			p.sourceID = run.SourceRun.ID
			if run.SourceRun.Result != "passed" {
				return errors.New("previous pipeline didn't pass so approval is not done.")
			}
			p.runID = run.ID
			err = p.approveRun()
			if err != nil {
				return err
			}
			return nil
		}
	}

	// Fell out of loop. Either there is no manual pipeline or no
	// approval pending.
	if !hasManual {
		return errors.New("approval pipeline does not exist for this run")
	}
	fmt.Println("There is no pending approval: so trigger the pipeline again")

	// Trigger the pipeline and then approve it.
	err = p.triggerAndApprovePipeline()
	if err != nil {
		return err
	}
	return nil
}

// Trigger and approve the triggerObject (WerckerRun) for the pipeline
func (p *WerckerTestDemo) triggerAndApprovePipeline() error {
	// Now get issue the trigger so we can get runId from response.
	// Build the body of the request to do the triger.
	url := fmt.Sprintf("https://app.wercker.com/api/v3/runs")
	pid := fmt.Sprintf("\"pipelineId\":\"%s\"", p.pipelineID)
	msg := fmt.Sprintf("\"message\":\"%s\"", "auto-triggered-002")
	brc := fmt.Sprintf("\"branch\":\"%s\"", p.branch)
	cmt := fmt.Sprintf("\"commitHash\":\"%s\"", p.commitHash)
	src := fmt.Sprintf("\"sourceRunId\":\"%s\"", p.sourceID)
	trg := fmt.Sprintf("{%s,%s,%s,%s,%s}", pid, msg, brc, cmt, src)

	bdy, err := p.postWercker(url, trg)
	if err != nil {
		return err
	}
	trigger := WerckerTrigger{}
	err = json.Unmarshal(bdy, &trigger)
	if err != nil {
		return err
	}
	// Pickup the correct runId from the response that is needed for approve.
	p.runID = trigger.Workflow.Items[0].Data.RunID
	det := fmt.Sprintf("Triggered pipeline=%s (%s), runID=%s", p.pipeline, p.pipelineID, p.runID)
	fmt.Println(det)
	return p.approveRun()
}

// Approve this so it will run the pending pipeline
func (p *WerckerTestDemo) approveRun() error {
	url := "https://app.wercker.com/api/v3/trigger/runs/approve"
	app := fmt.Sprintf("{\"runId\":\"%s\"}", p.runID)
	bdy, err := p.postWercker(url, app)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("Approved pipeline=%s (%s), runID=%s", p.pipeline, p.pipelineID, p.runID)
	fmt.Println(msg)
	showResponse(bdy, url)
	return nil
}

// Issue a http get
func (p *WerckerTestDemo) getWercker(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.werckerToken))
	req.Header.Add("Content-Type", "application/json")
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Issue a http post
func (p *WerckerTestDemo) postWercker(url string, postbody string) ([]byte, error) {
	bbody := []byte(postbody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bbody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.werckerToken))
	req.Header.Add("Content-Type", "application/json")
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// This takes apart a response body to show it as json.
func showResponse(body []byte, url string) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(body, &data)
	if err == nil {
		fmt.Println(fmt.Sprintf("Request: URL is %s", url))
		str, err := json.MarshalIndent(data, "", "  ")
		if err == nil {
			msg := fmt.Sprintf("Response: %s", string(str))
			fmt.Println(msg)
		}
	}
	return err
}
