package file

import (
	"io/ioutil"
)

// Load the file named name
func Load(name string) string {
	content, err := ioutil.ReadFile(name)
	check(err)
	return string(content)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
