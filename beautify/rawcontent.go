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

func RawContent(basePath string, subPath string) {
	srcPath := basePath + "/" + subPath
	contents := readFileAt(srcPath)
	quick.Highlight(os.Stdout, contents, subPath, "terminal", "monokai")
}
