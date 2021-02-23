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
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Files to be updated.
const (
	versiongo = `rbxmk/version.go`
	readme    = `README.md`
	changelog = `CHANGELOG.md`
)

// Content of version.go file.
const versiongoContent = `package main

const Version = "%d.%d.%d"
`

// Section in CHANGELOG indicating unreleased version.
const latest = `imperative`

// Template that produces a new CHANGELOG section. First argument is the
// previous version, second argument is the next version.
const changelogTemplate = `
## %[2]s
See a [comparison with the previous version][cmp-%[2]s] for a thorough list of changes.

The [Documentation page][doc-%[2]s] provides a complete reference for this version of rbxmk.

[doc-%[2]s]: https://github.com/Anaminus/rbxmk/blob/%[2]s/doc/README.md#user-content-rbxmk-reference
[cmp-%[2]s]: https://github.com/Anaminus/rbxmk/compare/%[1]s...%[2]s
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
	updateCHANGELOG(prev, next)
	updateREADME(next)
	commit(next)
}

func readVersion() (v Version) {
	b, err := os.ReadFile(versiongo)
	IfFatalf(err, "read version.go")
	_, err = fmt.Sscanf(string(b), versiongoContent, &v.Major, &v.Minor, &v.Patch)
	IfFatalf(err, "parse version.go")
	return v
}

func writeVersion(v Version) {
	b := []byte(fmt.Sprintf(versiongoContent, v.Major, v.Minor, v.Patch))
	err := os.WriteFile(versiongo, b, 0666)
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

func updateCHANGELOG(prev, next Version) {
	content, err := os.ReadFile(changelog)
	IfFatalf(err, "read CHANGELOG")

	match := regexp.MustCompile(fmt.Sprintf(`\n## %s\n(?s:.)*?\n## `, regexp.QuoteMeta(latest)))
	loc := match.FindIndex(content)
	if loc == nil {
		Fatalf("failed to find section %q in CHANGELOG", latest)
	}
	prefix := content[:loc[0]]
	section := content[loc[0]:loc[1]]
	suffix := content[loc[1]:]

	ver := "v" + next.String()
	section = bytes.ReplaceAll(section, []byte(latest), []byte(ver))

	var buf bytes.Buffer
	buf.Write(prefix)
	buf.WriteString(fmt.Sprintf(changelogTemplate, ver, latest))
	buf.Write(section)
	buf.Write(suffix)

	err = os.WriteFile(changelog, buf.Bytes(), 0666)
	IfFatalf(err, "write CHANGELOG")
}

func commit(v Version) {
	err := exec.Command("git", "add", versiongo, readme, changelog).Run()
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
