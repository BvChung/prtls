package fstraversal

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

type MockDirEntry struct {
	name string
}

func (m MockDirEntry) Name() string               { return m.name }
func (m MockDirEntry) IsDir() bool                { return false }
func (m MockDirEntry) Type() fs.FileMode          { return 0 }
func (m MockDirEntry) Info() (fs.FileInfo, error) { return nil, nil }

func TestFsTreeTraversal(t *testing.T) {
	assert := assert.New(t)

	testData := []struct {
		expected []string
		path     string
	}{
		{expected: []string{
			"└── " + dirStyle.Render("b"),
			"    └── " + dirStyle.Render("c"),
			"        ├── cc.py",
			"        └── " + dirStyle.Render("d"),
			"            ├── a.txt",
			"            └── b.txt",
		},
			path: "../testdata/mockDirectory"},
	}

	for _, d := range testData {
		t.Run("Test tree traversal", func(t *testing.T) {
			lines := []string{}

			if err := traverseFsTree(d.path, "", true, false, &lines); err != nil {
				t.Fatal(err)
			}

			assert.ElementsMatch(d.expected, lines)
		})
	}

	t.Run("Test invalid file path", func(t *testing.T) {
		path := "./invalidpath"
		lines := []string{}
		expected := fmt.Sprintf("error reading directory, open %s: The system cannot find the file specified.", path)

		err := traverseFsTree(path, "", true, false, &lines)
		assert.EqualError(err, expected)
	})
}

func TestTreeBuilding(t *testing.T) {
	assert := assert.New(t)

	testData := []struct {
		expected []string
		path     string
	}{
		{
			expected: []string{
				rootStyle.Render("."),
				"└── " + dirStyle.Render("b"),
				"    └── " + dirStyle.Render("c"),
				"        ├── cc.py",
				"        └── " + dirStyle.Render("d"),
				"            ├── a.txt",
				"            └── b.txt",
			},
			path: "../testdata/mockDirectory"},
	}

	for _, d := range testData {
		t.Run("Test tree building", func(t *testing.T) {
			lines := []string{}

			if err := traverseFsTree(d.path, "", true, false, &lines); err != nil {
				t.Fatal(err)
			}

			expectedTree := lipgloss.JoinVertical(lipgloss.Left, d.expected...)
			actualTree := createTree(d.path, false, ".")

			assert.Equal(expectedTree, actualTree)
		})
	}
}

func TestTextFormatting(t *testing.T) {
	assert := assert.New(t)

	testData := []struct {
		expected  string
		delimiter string
		fileName  string
		isDir     bool
	}{
		{expected: "├── file.txt", delimiter: "├── ", fileName: "file.txt", isDir: false},
		{expected: "│   file.txt", delimiter: "│   ", fileName: "file.txt", isDir: false},
		{expected: "└── file.txt", delimiter: "└── ", fileName: "file.txt", isDir: false},
		{expected: "└── " + dirStyle.Render("folder"), delimiter: "└── ", fileName: "folder", isDir: true},
	}

	for _, d := range testData {
		t.Run("Test text formatting", func(t *testing.T) {
			output := formatOutputText(d.delimiter, d.fileName, d.isDir)
			assert.Equal(d.expected, output, "Expected %s, got %s", d.expected, output)
		})
	}
}

func TestFindLastFileIndex(t *testing.T) {
	assert := assert.New(t)

	testData := []struct {
		expected   int
		files      []fs.DirEntry
		showHidden bool
	}{
		{expected: 2, files: []fs.DirEntry{
			MockDirEntry{name: "file1.txt"},
			MockDirEntry{name: ".hiddenfile"},
			MockDirEntry{name: "file2.txt"},
		}, showHidden: false},
		{expected: -1, files: []fs.DirEntry{}, showHidden: false},
	}

	for _, d := range testData {
		t.Run("Test finding last index", func(t *testing.T) {
			output := getLastFileIndex(d.files, d.showHidden)
			assert.Equal(d.expected, output, "Expected %s, got %b", d.expected, output)
		})
	}
}

func TestHiddenFile(t *testing.T) {
	assert := assert.New(t)

	testData := []struct {
		expected bool
		fileName string
	}{
		{expected: false, fileName: "file.txt"},
		{expected: false, fileName: "__file.txt"},
		{expected: true, fileName: ".env"},
	}

	for _, d := range testData {
		t.Run("Test if hidden file", func(t *testing.T) {
			output := isHiddenFile(d.fileName)
			assert.Equal(d.expected, output, "Expected %s, got %b", d.expected, output)
		})
	}
}
