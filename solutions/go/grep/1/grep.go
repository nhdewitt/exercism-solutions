package grep

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Options struct {
    LineNumbers		bool
    OnlyFilenames	bool
    CaseInsensitive	bool
    MatchAll		bool
    Invert			bool
}

var flagMap = map[string]func(*Options){
    "-n": func(o *Options) { o.LineNumbers = true },
    "-l": func(o *Options) { o.OnlyFilenames = true },
    "-i": func(o *Options) { o.CaseInsensitive = true },
    "-x": func(o *Options) { o.MatchAll = true },
    "-v": func(o *Options) { o.Invert = true },
}

func ParseFlags(flags []string) Options {
    var o Options
    for _, f := range flags {
        if set, ok := flagMap[f]; ok {
            set(&o)
        } else {
            continue
        }
    }
    return o
}

type Grep struct {
    pattern		string
    opts		Options
    multiFile	bool
    lowerPat	string
    matchFn		func(line string) bool
}

func NewGrep(pattern string, opts Options, fileCount int) *Grep {
    g := &Grep{
        pattern: 	pattern,
        opts:		opts,
        multiFile:	fileCount > 1,
    }
    if opts.CaseInsensitive {
        g.lowerPat = strings.ToLower(pattern)
    }

    if opts.MatchAll {
        if opts.CaseInsensitive {
            g.matchFn = func(line string) bool { return strings.ToLower(line) == g.lowerPat }
        } else {
            g.matchFn = func(line string) bool { return line == g.pattern }
        }
    } else {
        if opts.CaseInsensitive {
            g.matchFn = func(line string) bool { return strings.Contains(strings.ToLower(line), g.lowerPat) }
        } else {
            g.matchFn = func(line string) bool { return strings.Contains(line, g.pattern) }
        }
    }
    return g
}

func (g *Grep) buildOutput(file string, lineNo int, line string) string {
    var b strings.Builder
    if g.multiFile {
        b.WriteString(file)
        b.WriteByte(':')
    }
    if g.opts.LineNumbers {
        b.WriteString(fmt.Sprintf("%d:", lineNo))
    }
    b.WriteString(line)
    return b.String()
}

func (g *Grep) processFile(path string) (matches, nonMatches []string, matchedFile bool) {
    f, err := os.Open(path)
    if err != nil {
        return nil, nil, false
    }
    defer f.Close()

    s := bufio.NewScanner(f)
    lineNo := 1

    for s.Scan() {
        line := s.Text()

        if g.opts.OnlyFilenames {
            if g.matchFn(line) {
                return nil, nil, true
            }
            lineNo++
            continue
        }

        out := g.buildOutput(path, lineNo, line)
        if g.matchFn(line) {
            matches = append(matches, out)
        } else {
            nonMatches = append(nonMatches, out)
        }
        lineNo++
    }
    return matches, nonMatches, false
}

func Search(pattern string, flags, files []string) []string {
	opts := ParseFlags(flags)
    g := NewGrep(pattern, opts, len(files))

    var allMatches, allNonMatches, matchedFiles []string
    for _, f := range files {
        m, n, mf := g.processFile(f)
        if opts.OnlyFilenames && mf {
            matchedFiles = append(matchedFiles, f)
            continue
        }
        allMatches = append(allMatches, m...)
        allNonMatches = append(allNonMatches, n...)
    }

    if opts.OnlyFilenames {
        return matchedFiles
    }
    if opts.Invert {
        return allNonMatches
    }
    return allMatches
}
