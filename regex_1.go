package main

import(
    "fmt"
    "regexp"
)

func main() {
		str := "ConntrackAgent_install-1.0.0-20210417.tar.gz install Success ! InstallVersion:1.0.0 !"
		var digitsRegexp = regexp.MustCompile(`InstallVersion:(\S+) `)
		fmt.Println(digitsRegexp.FindStringSubmatch(str))

}
