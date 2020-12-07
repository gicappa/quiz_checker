package file;
	
import (
	"io/ioutil"
)

func LoadQuiz(name string) []byte {
	content, err := ioutil.ReadFile(name)
	check(err)
	return content
}

func check(e error) {
	if e != nil {
			panic(e)
	}
}