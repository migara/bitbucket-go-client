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

type WorkspaceLinks struct {
	Avatar *Link `json:"avatar,omitempty"`
	Html *Link `json:"html,omitempty"`
	Members *Link `json:"members,omitempty"`
	Owners *Link `json:"owners,omitempty"`
	Projects *Link `json:"projects,omitempty"`
	Repositories *Link `json:"repositories,omitempty"`
	Snippets *Link `json:"snippets,omitempty"`
	Self *Link `json:"self,omitempty"`
}
