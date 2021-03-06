// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
	dlppb "google.golang.org/genproto/googleapis/privacy/dlp/v2"
)

func TestTemplateSamples(t *testing.T) {
	testutil.SystemTest(t)
	buf := new(bytes.Buffer)
	fullID := "projects/" + projectID + "/inspectTemplates/golang-samples-test-template"
	// Delete template before trying to create it since the test uses the same name every time.
	listInspectTemplates(buf, client, projectID)
	got := buf.String()
	if strings.Contains(got, fullID) {
		buf.Reset()
		deleteInspectTemplate(buf, client, fullID)
		if got := buf.String(); !strings.Contains(got, "Successfully deleted inspect template") {
			t.Fatalf("failed to delete template")
		}
	}
	buf.Reset()
	createInspectTemplate(buf, client, projectID, dlppb.Likelihood_POSSIBLE, 0, "golang-samples-test-template", "Test Template", "Template for testing", nil)
	got = buf.String()
	if !strings.Contains(got, "Successfully created inspect template") {
		t.Fatalf("failed to createInspectTemplate: %s", got)
	}
	buf.Reset()
	listInspectTemplates(buf, client, projectID)
	got = buf.String()
	if !strings.Contains(got, fullID) {
		t.Fatalf("failed to list newly created template (%s): %q", fullID, got)
	}
	buf.Reset()
	deleteInspectTemplate(buf, client, fullID)
	if got := buf.String(); !strings.Contains(got, "Successfully deleted inspect template") {
		t.Fatalf("failed to delete template")
	}
}
