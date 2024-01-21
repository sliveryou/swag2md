package markdown

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	input := []byte(`Title
=

This is a new paragraph. I wonder    if I have too     many spaces.
What about new paragraph.
But the next one...

  Is really new.

1. Item one.
1. Item TWO.

~~Mistaken text.~~

This (**should** be *fine*).

A \> B.
The year was 1986. What a great season.

\*literal asterisks\*.

---

http://example.com

Now a [link](www.github.com) in a paragraph. End with [link_underscore.go](www.github.com).

-	[Link](www.example.com)

### An h3 header

Here's a numbered list:

1.	first item
2.	second item
3.	third item

Here's a table.

| Name  | Age |
|-------|-----|
| Bob   | 27  |
| Alice | 23  |

Colons can be used to align columns.

| Tables        | Are           | Cool      |
|---------------|:-------------:|----------:|
| col 3 is      | right-aligned |     $1600 |
| col 2 is      |   centered!   |       $12 |
| zebra stripes |   are neat    |        $1 |
| support for   | サブタイトル  | priceless |

The outer pipes (|) are optional, and you don't need to make the raw Markdown line up prettily. You can also use inline Markdown.

| Markdown | More      | Pretty     |
|----------|-----------|------------|
| *Still*  | renders | **nicely** |
| 1        | 2         | 3          |

Nested Lists
============

Final paragraph.
`)

	output, err := Process("", input, nil)
	assert.NoError(t, err)

	os.Stdout.Write(output)
}
