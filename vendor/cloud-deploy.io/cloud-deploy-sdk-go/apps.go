package ghost

import "encoding/json"

// GetApps returns all apps
//
// Cloud Deploy API docs
// https://docs.cloud-deploy.io/docs/_static/api.html#tag/app%2Fpaths%2F~1apps%2Fget
func (c *Client) GetApps() (apps Apps, err error) {
	res, err := c.get("/apps")
	if err == nil {
		err = json.NewDecoder(res.Body).Decode(&apps)
	}
	return
}

// CreateApp creates a new app
//
// Cloud Deploy API docs:
// https://docs.cloud-deploy.io/docs/_static/api.html#tag/app%2Fpaths%2F~1apps%2Fpost
func (c *Client) CreateApp(app App) (metadata EveItemMetadata, err error) {
	res, err := c.post("/apps", app)
	if err == nil {
		err = json.NewDecoder(res.Body).Decode(&metadata)
	}
	return
}

// GetApp returns the requested app
//
// Cloud Deploy API docs:
// https://docs.cloud-deploy.io/docs/_static/api.html#tag/app%2Fpaths%2F~1apps~1%7BappId%7D%2Fget
func (c *Client) GetApp(id string) (app App, err error) {
	res, err := c.get("/apps/" + id)
	if err == nil {
		err = json.NewDecoder(res.Body).Decode(&app)
	}
	return
}

// UpdateApp updates an existing app
//
// Cloud Deploy API docs:
// https://docs.cloud-deploy.io/docs/_static/api.html#tag/app%2Fpaths%2F~1apps~1%7BappId%7D%2Fpatch
func (c *Client) UpdateApp(app *App, id string, etag string) (metadata EveItemMetadata, err error) {
	res, err := c.patch("/apps/"+id, app, map[string]string{"If-Match": etag})
	if err == nil {
		err = json.NewDecoder(res.Body).Decode(&metadata)
	}
	return
}

// DeleteApp deletes an existing app
//
// Cloud Deploy API docs:
// https://docs.cloud-deploy.io/docs/_static/api.html#tag/app%2Fpaths%2F~1apps%2Fdelete
func (c *Client) DeleteApp(id string, etag string) (err error) {
	_, err = c.delete("/apps/"+id, map[string]string{"If-Match": etag})
	return
}
