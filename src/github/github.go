package github

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
)

const (
  api_base = "https://api.github.com"
)

type Client struct {
  username string
  password string
  htclient *http.Client
}

type Team struct {
  Id            int
  Name          string
  Url           string
}
type Teams []Team

type User struct {
  Id            int
  Login         string
  Avatar_url    string
  Gravatar_id   string
  Url           string
}

type Repo struct {
  Id            int
  Url           string
  Html_url      string
  Clone_url     string
  Git_url       string
  Ssh_url       string
  Svn_url       string
  Mirror_url    string
  Owner         User
  Name          string
  Full_name     string
  Description   string
  Homepage      string
  Language      string
  Private       bool
  Fork          bool
  Forks         int
  Watchers      int
  Size          int
  Master_branch string
  Open_issues   int
  Pushed_at     string
  Created_at    string
  Updated_at    string
}
type Repos []Repo

type Label struct {
  Url           string
  Name          string
  Color         string
}
type Labels []Label

type Milestone struct {
  Url           string
  Number        int
  State         string
  Title         string
  Description   string
  Creator       User
  Open_issues   int
  Closed_issues int
  Created_at    string
  Due_on        string
}

type PullRequest struct {
  Html_url      string
  Diff_url      string
  Patch_url     string
}

type Issue struct {
  Url           string
  Html_url      string
  Number        int
  State         string
  Title         string
  Body          string
  User          User
  Labels        Labels
  Assignee      User
  Milestone     Milestone
  Comments      int
  Pull_request  PullRequest
  Closed_at     string
  Created_at    string
  Updated_at    string
}
type Issues []Issue

func NewClient(username, password string) *Client {
  return &Client{username, password, new(http.Client)}
}

func (c *Client) get(url string) ([]byte, error) {
  req, err := http.NewRequest("GET", url, nil)
  req.SetBasicAuth(c.username, c.password)
  resp, err := c.htclient.Do(req)
  if err != nil { return make([]byte,0), err }

  if resp.StatusCode != 200 {
    return make([]byte,0), errors.New(
      fmt.Sprintf("Got a %v status code on fetch of %v.", resp.StatusCode, url))
  }
  b, err2 := ioutil.ReadAll(resp.Body)
  return b, err2
}

func (c *Client) OrgTeams(org string) (Teams, error) {
  var teams = Teams{}

  url := fmt.Sprintf("%v/orgs/%v/teams", api_base, org)
  body, err := c.get(url)
  if err != nil { return teams, err }

  err2 := json.Unmarshal(body, &teams)
  return teams, err2
}

func (c *Client) TeamRepos(id int) (Repos, error) {
  var repos = Repos{}

  url := fmt.Sprintf("%v/teams/%d/repos", api_base, id)
  body, err := c.get(url)
  if err != nil { return repos, err }

  err2 := json.Unmarshal(body, &repos)
  return repos, err2
}

func (c *Client) RepoIssues(org, repo string) (Issues, error) {
  var issues = Issues{}

  url := fmt.Sprintf("%v/repos/%v/%v/issues", api_base, org, repo)
  body, err := c.get(url)
  if err != nil { return issues, err }

  err2 := json.Unmarshal(body, &issues)
  return issues, err2
} 
