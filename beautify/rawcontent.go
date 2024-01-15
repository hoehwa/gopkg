package beautify

import (
	"log"
	"os"

	"github.com/alecthomas/chroma/quick"
)

// Read and return contents of the file at the specific absoulte path.
func readFileAt(fileAbsPath string) string {
	readStream, err := os.ReadFile(fileAbsPath)
	if err != nil {
		log.Fatal(err)
	}

	return string(readStream)
}

func RawContent(path string, file string) {
	srcPath := path + "/" + file
	contents := readFileAt(srcPath)
	quick.Highlight(os.Stdout, contents, file, "terminal", "monokai")
}
