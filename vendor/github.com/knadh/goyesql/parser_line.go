package goyesql

import (
	"regexp"
	"strings"
)

// A line may be blank, a tag, a comment or a query
const (
	lineBlank = iota
	lineQuery
	lineComment
	lineTag
)

// ParsedLine stores line type and value
//
// For example: parsedLine{Type=lineTag, Value="foo"}
type parsedLine struct {
	Type  int
	Tag   string
	Value string
}

var (
	// -- tag: $value
	reTag = regexp.MustCompile("^\\s*--\\s*(.+)\\s*:\\s*(.+)")

	// -- $comment
	reComment = regexp.MustCompile("^\\s*--\\s*(.+)")
)

func parseLine(line string) parsedLine {
	line = strings.Trim(line, " ")

	if line == "" {
		return parsedLine{lineBlank, "", ""}
	} else if matches := reTag.FindStringSubmatch(line); len(matches) > 1 {
		return parsedLine{lineTag, matches[1], matches[2]}
	} else if matches := reComment.FindStringSubmatch(line); len(matches) > 0 {
		return parsedLine{lineComment, "", matches[1]}
	}

	return parsedLine{lineQuery, "", line}
}
