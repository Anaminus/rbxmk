// Tool to increment the version number, read and write rbxmk/version.go, then
// create and tag a commit. Must run in root of repo.
//
// Increment major version:
//
//     go run _tools/version.go +major
//
// Increment minor version:
//
//     go run _tools/version.go +minor
//
// Increment patch version:
//
//     go run _tools/version.go +patch
//
// Set version directly:
//
//     go run _tools/version.go 0.0.0
//
// After confirming commit, push to remote:
//
//     go run _tools/version.go push
//
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const versiongo = `rbxmk/version.go`
const readme = `README.md`

const versiongoContent = `package main

const Version = "%d.%d.%d"
`

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v Version) IncMajor() Version {
	return Version{v.Major + 1, 0, 0}
}

func (v Version) IncMinor() Version {
	return Version{v.Major, v.Minor + 1, 0}
}

func (v Version) IncPatch() Version {
	return Version{v.Major, v.Minor, v.Patch + 1}
}

func Fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func IfFatalf(err error, format string, v ...interface{}) {
	if err == nil {
		return
	}
	v = append(v, err)
	fmt.Fprintf(os.Stderr, format+": %s\n", v...)
	os.Exit(1)
}

func main() {
	flag.Parse()
	var prev, next Version
	switch s := strings.ToLower(flag.Arg(0)); s {
	case "+major":
		prev = readVersion()
		next = prev.IncMajor()
	case "+minor":
		prev = readVersion()
		next = prev.IncMinor()
	case "+patch":
		prev = readVersion()
		next = prev.IncPatch()
	case "push":
		push()
		return
	default:
		_, err := fmt.Sscanf(s, "%d.%d.%d", &next.Major, &next.Minor, &next.Patch)
		IfFatalf(err, "argument must be a version string, +major, +minor, +patch, or push")
		prev = readVersion()
	}
	fmt.Printf("%s => %s\n", prev, next)
	writeVersion(next)
	updateREADME(next)
	commit(next)
}

func readVersion() (v Version) {
	b, err := ioutil.ReadFile(versiongo)
	IfFatalf(err, "read version.go")
	_, err = fmt.Sscanf(string(b), versiongoContent, &v.Major, &v.Minor, &v.Patch)
	IfFatalf(err, "parse version.go")
	return v
}

func writeVersion(v Version) {
	b := []byte(fmt.Sprintf(versiongoContent, v.Major, v.Minor, v.Patch))
	err := ioutil.WriteFile(versiongo, b, 0666)
	IfFatalf(err, "write version.go")
}

var versionTag = regexp.MustCompile(`<version>(.+)</version>`)

func updateREADME(v Version) {
	content, err := os.ReadFile(readme)
	IfFatalf(err, "read README")

	matches := versionTag.FindSubmatch(content)
	if len(matches) == 0 {
		Fatalf("failed to find version in README")
	}
	content = bytes.ReplaceAll(content, matches[1], []byte("v"+v.String()))

	IfFatalf(os.WriteFile(readme, content, 0666), "write README")
}

func commit(v Version) {
	err := exec.Command("git", "add", versiongo, readme).Run()
	IfFatalf(err, "git add")
	err = exec.Command("git", "commit", "-m", fmt.Sprintf("Release version v%s.", v)).Run()
	IfFatalf(err, "git commit")
	err = exec.Command("git", "tag", fmt.Sprintf("v%s", v), "-m", fmt.Sprintf("Release version v%s.", v)).Run()
	IfFatalf(err, "git tag")
}

func push() {
	err := exec.Command("git", "push", "--follow-tags").Run()
	IfFatalf(err, "git push")
}
