package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	/* Book's example */
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
	/* Book's example */

	/* Exercise 7.1 */
	var w WordsCounter
	w.Write([]byte("My ass is on fire"))
	fmt.Println(w)

	w = 0
	var words = "So many words I can't even bother to count"
	fmt.Fprintf(&w, words)
	fmt.Println(w)
	/* Exercise 7.1 */

	/* Exercise 7.2 */
	newWriter, count := CountingWriter(os.Stdout)
	newWriter.Write([]byte(words))
	fmt.Println(*count)
	/* Exercise 7.2 */
}

type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	count := c.countWords(p)
	*c += count
	return len(p), nil
}

func (c *WordsCounter) countWords(p []byte) WordsCounter {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}
	return WordsCounter(count)
}

type byteCounter struct {
	w io.Writer
	count int64
}

func (b *byteCounter) Write(p []byte) (int, error) {
	n, err := b.w.Write(p)
	b.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &byteCounter{w, 0}
	return c, &c.count
}