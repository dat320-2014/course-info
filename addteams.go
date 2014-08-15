package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"code.google.com/p/goauth2/oauth"

	"github.com/google/go-github/github"
)

// the default group name prefix; will be appended by a sequential group number
const groupName = "group"

//TODO: change to GITHUB_AUTH_TOKEN

var usageString = `
This program relys on setting some environment variables.

The ACCESS_TOKEN env variable must be set to one of your personal access
tokens generated on github. You will need to generate such a token on github:
Go to Edit profile, Applications, Personal access tokens, Generate new token.
You should save the token to a file .github-access-token, and set it in your
.zshenv (or similar for your shell) file using:

    export ACCESS_TOKEN=` + "`" + `cat $HOME/.github-access-token` + "`" + `

You may also want to run:

    chmod 600 $HOME/.github-access-token

Next, you should also set the env variable COURSE_ORG to the name of the
organization that you have created for the course on github, e.g.:

    export COURSE_ORG=uis-dat320-fall2014

Typically, this env will be updated once for every semester, but if you have
multiple courses in one semester you can run the script prefixed with the env
variable like so:

    COURSE_ORG=uis-dat320-fall2014 addteams

`

// globally accessible
var (
	courseOrg string
	client    *github.Client
)

// team name -> team id
var teams = make(map[string]int)

func main() {
	teamDir := flag.String("dir", "teams", "directory with team files")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, usageString)
	}
	flag.Parse()

	token := os.Getenv("ACCESS_TOKEN")
	courseOrg = os.Getenv("COURSE_ORG")
	if token == "" || courseOrg == "" {
		flag.Usage()
		os.Exit(0)
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	client = github.NewClient(t.Client())

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
		fmt.Printf("\t%v\n", *r.FullName)
	}

	ts, _, err := client.Organizations.ListTeams(courseOrg, nil)
	if err != nil {
		// ListTeams failed; not sure if this can happen if there are no teams??
		panic(err)
	}
	for _, tname := range ts {
		teams[*tname.Name] = *tname.ID
	}

	createTeams(teamDir)
}

func createTeams(teamDir *string) {
	teamFiles := filter(filesIn(*teamDir), func(file string) bool {
		return strings.HasSuffix(file, "txt")
	})
	for grp, file := range teamFiles {
		b, err := ioutil.ReadFile(filepath.Join(*teamDir, file))
		if err != nil {
			panic(err)
		}
		members := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")
		name := fmt.Sprintf("%s%d", groupName, grp+1)
		teamID := createTeam(name)
		addMembers(teamID, members)
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
	return *team.ID

}

func addMembers(teamID int, members []string) {
	ms, _, err := client.Organizations.ListTeamMembers(teamID, nil)
	if err != nil {
		fmt.Println("Could be just no members yet??")
		// panic(err)
	}
	users := make(map[string]bool)
	for _, u := range ms {
		users[*u.Login] = true
	}

	for _, member := range members {
		fmt.Printf("Checking %s for team id: %d\n", member, teamID)
		if !users[member] {
			_, err = client.Organizations.AddTeamMember(teamID, member)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Added %s to team id: %d\n", member, teamID)
		} else {
			fmt.Printf("Didn't add %s to team id: %d\n", member, teamID)
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
