package main

import (
  "github"
  "fmt"
  "sync"
)

func main() {
  var username = ""
  var password = ""
  var orgname = ""
  var ghids = []int{}

  c := github.NewClient(username, password)

  teams, err := c.OrgTeams(orgname)
  if (err != nil) {
    fmt.Println(err)
  }
  fmt.Println(teams)
  for _,team := range teams {
    ghids = append(ghids, team.Id)
    fmt.Println(team)
  }

  var ghrepos = []string{}

  var wg sync.WaitGroup
  for _, id := range ghids {
    wg.Add(1)
    fmt.Printf("Fetching %d\n", id)
    go func(id int) {
      repos, err2 := c.TeamRepos(id)
      wg.Done()
      if (err2 != nil) {
        fmt.Println(err2)
      }

      for _,repo := range repos{
        var found = false
        for _, r := range ghrepos {
          if r == repo.Name { found = true }
        }
        if (!found) { ghrepos = append(ghrepos, repo.Name) }
      }
    }(id)
  }
  wg.Wait()

  var wg2 sync.WaitGroup
  fmt.Println("Repo list: ")
  fmt.Println(ghrepos)
  for _, name := range ghrepos {
    wg2.Add(1)
    fmt.Printf("... Fetching %s\n", name)
    go func(name string) {
      issues, err3 := c.RepoIssues(orgname, name)
      wg2.Done()
      if (err3 != nil) {
        fmt.Println(err3)
      }

      for _,issue := range issues{
        fmt.Println("Issue: "+issue.Title)
      }
    }(name)
  }
  wg2.Wait()
}
