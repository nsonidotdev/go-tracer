package tracer

import (
	"context"
	"fmt"
	"strings"
)

type printOptions struct {
	isRoot      bool
	isLastChild bool
	prefix      string
}

const (
	emptyPrefix = "    "
	midPrefix   = "├── "
	endPrefix   = "└── "
	downPrefix  = "│   "
)

func PrintTrace(ctx context.Context) {
	s, err := getCtxSpan(ctx)
	if err != nil {
		fmt.Println("error extracting span from context", err)
		return
	}

	fmt.Printf("TRACE: %s\n", s.id)
	recursivePrint(s, printOptions{isRoot: true, prefix: ""})
}

func recursivePrint(s *Span, opts printOptions) {
	lines := formatSpan(s)

	if !opts.isRoot {
		printSpanLine(opts.prefix, downPrefix, "")

		for i, line := range lines {
			var firstPrefix string
			var defaultPrefix string

			if opts.isLastChild {
				firstPrefix = endPrefix
				defaultPrefix = emptyPrefix
			} else {
				firstPrefix = midPrefix
				defaultPrefix = downPrefix
			}

			if i == 0 {
				printSpanLine(opts.prefix, firstPrefix, line)
			} else {
				printSpanLine(opts.prefix, defaultPrefix, line)
			}
		}
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}

	if len(s.children) == 0 {
		return
	}

	for i, childSpan := range s.children {
		isLastChild := i == len(s.children)-1

		var newPrefix string
		if opts.isRoot {
			newPrefix = ""
		} else if !opts.isLastChild {
			newPrefix = opts.prefix + downPrefix
		} else {
			newPrefix = opts.prefix + emptyPrefix
		}

		recursivePrint(childSpan, printOptions{
			isRoot:      false,
			isLastChild: isLastChild,
			prefix:      newPrefix,
		})
	}
}

func formatSpan(s *Span) []string {
	lines := make([]string, 0, 3+len(s.meta))

	lines = append(
		lines,
		s.name,
		fmt.Sprintf("status: %s | duration: %dms", strings.ToUpper(string(s.status)), s.duration.Milliseconds()),
	)

	if s.reason != "" {
		lines = append(lines, fmt.Sprintf("reason: %s", s.reason))
	}

	if len(s.meta) != 0 {
		metaLines := formatMeta(s)
		lines = append(lines, metaLines...)
	}

	return lines
}

func formatMeta(s *Span) []string {
	lines := make([]string, 0, len(s.meta))
	for key, value := range s.meta {
		lines = append(lines, fmt.Sprintf("%s: %s", key, value))
	}

	return lines
}

func printSpanLine(prevPrefix string, newPrefix string, line string) {
	formattedLine := fmt.Sprintf("%s%s%s", prevPrefix, newPrefix, line)
	fmt.Println(formattedLine)
}
