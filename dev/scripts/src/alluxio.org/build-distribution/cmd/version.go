package cmd

import (
	"fmt"
	"regexp"
	"strconv"
)

var versionRE = regexp.MustCompile("^(\\d+)\\.(\\d+)\\.(\\d+)(.*)?$")

type version struct {
	major  int
	minor  int
	patch  int
	suffix string
}

func parseVersion(v string) version {
	matches := versionRE.FindStringSubmatch(v)
	major, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(fmt.Sprintf("failed to parse %v as number", matches[1]))
	}
	minor, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(fmt.Sprintf("failed to parse %v as number", matches[2]))
	}
	patch, err := strconv.Atoi(matches[3])
	if err != nil {
		panic(fmt.Sprintf("failed to parse %v as number", matches[3]))
	}
	return version{
		major:  major,
		minor:  minor,
		patch:  patch,
		suffix: matches[4],
	}
}

func (v version) String() string {
	return fmt.Sprintf("%v.%v.%v%v", v.major, v.minor, v.patch, v.suffix)
}

func (v version) hadoopProfile() string {
	switch v.major {
	case 1:
		return "hadoop-1"
	case 2:
		return "hadoop-2"
	default:
		panic(fmt.Sprintf("unexpected hadoop major version %v", v.major))
	}
}
