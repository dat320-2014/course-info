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

	nxtGrp := loadExistingTeams()
	createTeams(teamDir, nxtGrp)
}

// load existing teams and users and return the next group number to be added
func loadExistingTeams() (nxtGrp int) {
	ts, _, err := client.Organizations.ListTeams(courseOrg, nil)
	if err != nil {
		// ListTeams failed; not sure if this can happen if there are no teams??
		panic(err)
	}
	for _, tname := range ts {
		if *tname.Name == "Owners" {
			// ignore the 'Owners' team
			continue
		}
		fmt.Println("Found existing team:", *tname.Name)
		m, err := strconv.Atoi(strings.TrimPrefix(*tname.Name, groupName))
		if err != nil {
			// ignore teams without a number at the end
			continue
		}
		if m > nxtGrp {
			nxtGrp = m
		}
		teams[*tname.Name] = *tname.ID
		ms, _, err := client.Organizations.ListTeamMembers(*tname.ID, nil)
		if err != nil {
			panic(err)
		}
		for _, u := range ms {
			if prevTID, exists := users[*u.Login]; exists {
				fmt.Printf("Unable to add '%s' to team %d, already member of %d",
					*u.Login, *tname.ID, prevTID)
				continue
			}
			users[*u.Login] = *tname.ID
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

	for i, file := range teamFiles {
		nxtGrp += i
		// group name, such as GroupXX
		var name string
		// the github assigned teamID
		var teamID int
		var err error
		//TODO Optimize this regexp: number of digits in GroupXX is two
		var validTeamFile = regexp.MustCompile(groupName + "[0-9][0-9]-[0-9]+.txt")
		if validTeamFile.MatchString(file) {
			// Check if the file has a valid groupXX name and teamID.
			k := strings.TrimSuffix(file, ".txt")
			l := strings.Split(k, "-")
			name = l[0]
			teamID, err = strconv.Atoi(l[1])
			if err != nil {
				// ignore files without group number or if the teamID is
				// larger than the nxtGrp to be created
				fmt.Printf("Ignoring file: %s\n", file)
				continue
			}
		} else {
			// Number of digits in XX in the GroupXX name is 2 (will use
			// leading zeros if XX < 10)
			name = fmt.Sprintf("%s%.2d", groupName, nxtGrp)
			teamID = createTeam(name)
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
