package github

import (
	"testing"
)

func TestSearchTopRepos(t *testing.T) {
	// emptyQuery := NewQuery("")
	awesomeQuery := NewQuery("awesome")
	// dockerQueryWrittenInGo := NewQuery("docker").WithLanguage("go")
	// reactQueryWrittenInTsWithFrontend := NewQuery("react").WithLanguage("typescript").WithTopic("frontend")
	// SearchFamousRepos(emptyQuery)
	SearchTopRepos(awesomeQuery)
	// SearchFamousRepos(dockerQueryWrittenInGo)
	// SearchFamousRepos(reactQueryWrittenInTsWithFrontend)
}
