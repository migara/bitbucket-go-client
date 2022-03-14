/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Issue struct {
	Type_ string `json:"type"`
	Links *IssueLinks `json:"links,omitempty"`
	Id int32 `json:"id,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
	Title string `json:"title,omitempty"`
	Reporter *User `json:"reporter,omitempty"`
	Assignee *User `json:"assignee,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
	UpdatedOn time.Time `json:"updated_on,omitempty"`
	EditedOn time.Time `json:"edited_on,omitempty"`
	State string `json:"state,omitempty"`
	Kind string `json:"kind,omitempty"`
	Priority string `json:"priority,omitempty"`
	Milestone *Milestone `json:"milestone,omitempty"`
	Version *Version `json:"version,omitempty"`
	Component *Component `json:"component,omitempty"`
	Votes int32 `json:"votes,omitempty"`
	Content *RenderedPullRequestMarkupTitle `json:"content,omitempty"`
}
