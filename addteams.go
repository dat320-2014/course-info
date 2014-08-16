// Copyright (C) 2014 Hein Meling.
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"code.google.com/p/goauth2/oauth"

	"github.com/google/go-github/github"
)

// the default group name prefix; will be appended by a sequential group number
const groupName = "Group"

var usageString = `
Program to add student teams and create repositories on github.com for use in
courses that rely on github for handling lab exercise submissions (via pull
requests) etc.

This program relys on setting some environment variables.

The GITHUB_AUTH_TOKEN env variable must be set to one of your personal access
tokens generated on github. You will need to generate such a token on github:
Go to Edit profile, Applications, Personal access tokens, Generate new token.
You should save the token to a file .github-access-token, and set it in your
.zshenv (or similar for your shell) file using:

    export GITHUB_AUTH_TOKEN=` + "`" + `cat $HOME/.github-auth-token` + "`" + `

You may also want to run:

    chmod 600 $HOME/.github-auth-token

Next, you should also set the env variable GITHUB_ORG to the name of the
organization that you have created for the course on github, e.g.:

    export GITHUB_ORG=uis-dat320-fall2014

Typically, this env will be updated once for every semester, but if you have
multiple courses in one semester you can run the script prefixed with the env
variable like so:

    GITHUB_ORG=uis-dat320-fall2014 addteams

Troubleshooting:

** If you get a '404 Not Found []' message, please check that your personal
   access token has the necessary priviliges, e.g. to delete a repository.

Author:

** Copyright (C) 2014 Hein Meling.

`

// globally accessible
var (
	courseOrg string
	client    *github.Client
)

// team name -> team id
var teams = make(map[string]int)

// user name -> team id to which the user is already added
var users = make(map[string]int)

func main() {
	teamDir := flag.String("dir", "teams", "directory with team files")
	delTeams := flag.Bool("delete", false, "delete all teams/repos in course (for debugging)")
	listRepos := flag.Bool("list", false, "list repos in course (for debugging)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, usageString)
	}
	flag.Parse()

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	courseOrg = os.Getenv("GITHUB_ORG")
	if token == "" || courseOrg == "" {
		flag.Usage()
		os.Exit(0)
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	client = github.NewClient(t.Client())

	if *listRepos {
		org, _, err := client.Organizations.Get(courseOrg)
		if err != nil {
			panic(err)
		}
		fmt.Println("Found organization:", *org.Name)

		repos, _, err := client.Repositories.ListByOrg(courseOrg, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("With the following repositories:")
		for _, r := range repos {
			fmt.Printf("\t%s\n", *r.FullName)
		}
		os.Exit(0)
	}

	nxtGrp := loadExistingTeams()
	if *delTeams {
		deleteTeamsAndRepos()
		os.Exit(0)
	}

	createTeams(teamDir, nxtGrp)
}

// delete teams and repos for those teams; make sure the the auth token has
// the appropriate privileges before using this function.
func deleteTeamsAndRepos() {
	if len(teams) == 0 {
		fmt.Println("Nothing to delete")
		return
	}
	for name, id := range teams {
		fmt.Println("Deleting repo:", name)
		_, err := client.Repositories.Delete(courseOrg, name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Deleting team:", id)
		_, err = client.Organizations.DeleteTeam(id)
		if err != nil {
			panic(err)
		}
	}
}

// load existing teams and users and return the next group number to be added
func loadExistingTeams() (nxtGrp int) {
	ts, _, err := client.Organizations.ListTeams(courseOrg, nil)
	if err != nil {
		// ListTeams failed; not sure if this can happen if there are no teams??
		panic(err)
	}
	for _, team := range ts {
		if *team.Name == "Owners" {
			// ignore the 'Owners' team
			continue
		}
		fmt.Println("Found existing team:", *team.Name)
		m, err := strconv.Atoi(strings.TrimPrefix(*team.Name, groupName))
		if err != nil {
			// ignore teams without a number at the end
			continue
		}
		if m > nxtGrp {
			nxtGrp = m
		}
		teams[*team.Name] = *team.ID
		ms, _, err := client.Organizations.ListTeamMembers(*team.ID, nil)
		if err != nil {
			panic(err)
		}
		for _, u := range ms {
			if prevTID, exists := users[*u.Login]; exists {
				fmt.Printf("Unable to add '%s' to team %d, already member of %d",
					*u.Login, *team.ID, prevTID)
				continue
			}
			users[*u.Login] = *team.ID
		}
	}
	nxtGrp++
	return
}

// createTeams adds teams to github based on team files provided by students.
// If a team file matches the specific format 'GroupXX-TeamID.txt' (a simple
// rename from a .done file), this allows students to add more members to
// their group.
func createTeams(teamDir *string, nxtGrp int) {
	teamFiles := filter(filesIn(*teamDir), func(file string) bool {
		return strings.HasSuffix(file, "txt")
	})
	//TODO Optimize this regexp: number of digits in GroupXX is two
	var validTeamFile = regexp.MustCompile(groupName + "[0-9][0-9]-[0-9]+.txt")

	for _, file := range teamFiles {
		// group name, such as GroupXX
		var name string
		// the github assigned teamID
		var teamID int
		var err error
		if validTeamFile.MatchString(file) {
			// Check if the file has a valid groupXX name and teamID.
			l := strings.Split(strings.TrimSuffix(file, ".txt"), "-")
			name = l[0]
			teamID, _ = strconv.Atoi(l[1])
			// here it is safe to ignore the err (_) since the regexp above
			// will detect a non-number teamID in the file name.
		} else {
			// Number of digits in XX in the GroupXX name is 2 (will use
			// leading zeros if XX < 10)
			name = fmt.Sprintf("%s%.2d", groupName, nxtGrp)
			teamID = createTeam(name)
			createRepo(name, teamID)
			nxtGrp++
		}

		oldpath := filepath.Join(*teamDir, file)
		b, err := ioutil.ReadFile(oldpath)
		if err != nil {
			panic(err)
		}
		members := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")

		addMembers(teamID, members)
		// Mark the processed file by renaming it: 'groupXX-teamID.done'
		newfile := name + "-" + strconv.Itoa(teamID) + ".done"
		newpath := filepath.Join(*teamDir, newfile)
		err = os.Rename(oldpath, newpath)
		if err != nil {
			fmt.Println("Could not rename to:", newpath)
			fmt.Println(err)
		}
	}
}

func createTeam(name string) int {
	if id, exists := teams[name]; exists {
		fmt.Printf("Team '%s' already exists\n", name)
		return id
	}
	fmt.Println("Creating team:", name)
	input := &github.Team{Name: github.String(name)}
	team, _, err := client.Organizations.CreateTeam(courseOrg, input)
	if err != nil {
		// something other than already exist happened
		panic(err)
	}
	teams[name] = *team.ID
	return *team.ID
}

// create a new private repository named name for the given team
func createRepo(name string, team int) {
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(false), // TODO Update once github has given me private repos
	}
	r, _, err := client.Repositories.Create(courseOrg, repo)
	if err != nil {
		fmt.Printf("Failed to create repo: %s: %v\n", name, err)
		return
	}
	fmt.Printf("Created repository: %s for team %d;\n  URL: %s\n", name, team, *r.URL)
	_, err = client.Organizations.AddTeamRepo(team, courseOrg, name)
	if err != nil {
		fmt.Printf("Failed to add team %d to repo: %s: %v\n", team, name, err)
	}
	fmt.Println("Added team to repository:", name)
}

func addMembers(teamID int, members []string) {
	for _, member := range members {
		if prevTID, exists := users[member]; !exists {
			_, err := client.Organizations.AddTeamMember(teamID, member)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Added '%s' to team: %d\n", member, teamID)
			users[member] = teamID
		} else {
			fmt.Printf("'%s' is already member of: %d\n", member, prevTID)
		}
	}
}

// filter returns a new slice with elements of s that satisfy fn()
func filter(s []string, fn func(string) bool) []string {
	var p []string
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

// filesIn returns files in the given directory dir
func filesIn(dir string) []string {
	fh, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	files, err := fh.Readdirnames(0)
	if err != nil {
		panic(err)
	}
	return files
}
