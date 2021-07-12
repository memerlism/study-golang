package yamltohtml_test

import (
	yamltohtml "mytesting/yamltohtml"
	"testing"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestYamlToHTML(t *testing.T) {
	testCases := []TestCase{
		TestCase{
			desc:     "Test Case 1",
			path:     "testdata/test_01.yml",
			expected: "<html><head><title>My Awesome Page</title></head><body>This is my awesome content</body></html>",
		},
		TestCase{
			desc:     "Test Case 2",
			path:     "testdata/test_02.yml",
			expected: "<html><head><title>My Awesome Page2</title></head><body>This is my awesome content2</body></html>",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result, err := yamltohtml.YamltoHTML(test.path)
			if err != nil {
				t.Fail()
			}

			t.Log(result)

			if result != test.expected {
				t.Fail()
			}
		})
	}
}
