package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days\n the rain was un-stoppable\n It was always cold\n no sunshine\n Yeah runnin' down a dream\n That never would come to me\n Workin' on a mystery\n goin' wherever it leads\n Runnin' down a dream"

	scanner := bufio.NewScanner(strings.NewReader(s))

	scanner.Split(bufio.ScanWords) //We split our text on words. So we get each word in the text. If we did not cal this split method, by default, the scanner will read the text line by line instead of word by word

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text()) //NOTE: now we call .Text() method we are not getting string passed line by line anymore because we have now changed the default to get the text word by word
	}

}
