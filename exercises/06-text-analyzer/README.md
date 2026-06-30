# Exercise 06 - Text Analyzer

## What You Are Learning

Text processing is a pipeline:

```text
raw text -> normalize -> tokenize -> count -> select result
```

Each step has one responsibility. Bugs usually come from skipping a step or
using an undefined rule for punctuation, case, or ties.

## Beginner Bridge: Why Text Needs A Pipeline

Raw text is messy. The same idea can appear in different shapes:

```text
Go
go
GO
go,
```

If you count raw strings directly, those may become four different tokens. A
pipeline makes each decision explicit before counting.

```text
normalize case -> split into words -> count words -> choose the largest count
```

That is the same as cleaning ingredients before cooking. The count step should
not also decide punctuation rules. The "most common" step should not also split
text. One step, one job.

Similar programs:

| Program | Normalize | Tokenize | Count |
|---|---|---|---|
| Tag analyzer | lowercase tags | split by spaces | count tags |
| Log analyzer | normalize levels | read entries | count levels |
| Search helper | lowercase query | split words | count terms |

When stuck, print the output of each pipeline stage. If `Fields` produced the
wrong words, fixing the count loop will not solve the real problem.

## Strings Are Values, Not Editable Character Arrays

Go strings are immutable byte sequences that conventionally contain UTF-8 text.
Functions in the `strings` package return new values:

```go
lower := strings.ToLower(original)
```

`original` is unchanged.

Useful operations:

```go
strings.TrimSpace(text) // remove leading and trailing whitespace
strings.ToLower(text)   // normalize letter case
strings.Fields(text)    // split around runs of whitespace
```

Example:

```go
text := "  Go   is clear  "
words := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
```

Result:

```text
["go", "is", "clear"]
```

`Fields` handles multiple spaces, tabs, and newlines better than splitting only
on `" "`.

## Define What Counts As A Word

For this exercise, case should not matter:

```text
Go == go == GO
```

But `strings.Fields` does not remove punctuation:

```text
"go," and "go" are different tokens
```

Choose and document a rule. A simple first version may count punctuation as part
of a word. An extension can trim punctuation with `strings.Trim`.

## Frequency Counting

Pseudocode:

```text
FUNCTION frequency(text)
    normalized text = lowercase text
    words = split normalized text by whitespace
    counts = empty string-to-integer map

    FOR each word
        counts[word] = counts[word] plus 1

    RETURN counts
```

Trace `"Go is go"`:

| Word | Previous count | New count |
|---|---:|---:|
| `go` | 0 | 1 |
| `is` | 0 | 1 |
| `go` | 1 | 2 |

## Selecting The Most Common Word

Counting and selecting are different operations.

```text
FUNCTION most common word(text)
    counts = frequency(text)
    best word = empty string
    best count = 0

    FOR each word and count in counts
        IF count is greater than best count
            best word = word
            best count = count

    RETURN best word and best count
```

Map order is unspecified. If two words tie and you do not define a tie-breaker,
either may be returned. A deterministic rule could choose the alphabetically
smaller word when counts are equal.

## Worked Example: Tag Analyzer

```go
package main

import (
    "fmt"
    "strings"
)

func tagFrequency(input string) map[string]int {
    counts := make(map[string]int)

    for _, tag := range strings.Fields(strings.ToLower(input)) {
        counts[tag]++
    }

    return counts
}

func mostUsedTag(input string) (string, int) {
    bestTag := ""
    bestCount := 0

    for tag, count := range tagFrequency(input) {
        if count > bestCount {
            bestTag = tag
            bestCount = count
        }
    }

    return bestTag, bestCount
}

func main() {
    input := "Go backend GO testing backend go"
    tag, count := mostUsedTag(input)
    fmt.Println(tag, count)
}
```

The result should be `go 3`. The algorithm is the same as the text analyzer,
but the domain is message tags.

## Your Program

Build:

1. `wordCount(text string) int`
2. `frequency(text string) map[string]int`
3. `mostCommonWord(text string) (string, int)`
4. A `main` function that analyzes one paragraph

Case must be normalized. Empty text must return an empty word and zero count.

## Build In Checkpoints

1. Import `strings` and print the result of `Fields`.
2. Return the number of fields.
3. Count every normalized word in a map.
4. Print the frequency map.
5. Scan the map for the highest count.
6. Decide and implement a tie rule.

```powershell
gofmt -w .\main.go
go run .
go doc strings.Fields
```

## Test Table

| Input | Word count | Important frequency |
|---|---:|---|
| `""` | 0 | none |
| `"   "` | 0 | none |
| `"Go"` | 1 | go: 1 |
| `"Go go GO"` | 3 | go: 3 |
| `"one   two"` | 2 | one: 1, two: 1 |
| `"one\ntwo"` | 2 | one: 1, two: 1 |
| `"go, go"` | 2 | define punctuation behavior |

## Failure Experiments

1. Use `strings.Split(text, " ")` on repeated spaces. Inspect the empty tokens.
2. Remove `ToLower` and compare counts.
3. Initialize `bestCount` to an impossibly high value. Explain why no candidate
   replaces it.
4. Run a tied case repeatedly and observe that map order is not a stable
   tie-breaker.

## Unicode Reality

`range` over a string produces Unicode code points, while indexing a string
produces bytes. `strings.Fields` and `ToLower` are reasonable for this exercise,
but full natural-language tokenization is much more complex. Do not claim this
is a production search engine or language analyzer.

## You Understand This Exercise When

You can explain every pipeline stage, distinguish counting from selecting,
describe punctuation and tie behavior, and build the same algorithm for tags or
log event names.

References:

- [`strings`](https://pkg.go.dev/strings)
- [`range`](https://go.dev/tour/moretypes/16)
- [Maps](https://go.dev/tour/moretypes/19)
