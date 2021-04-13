package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n The last three days\n the rain was un-stoppable\n It was always cold\n no sunshine\n Yeah runnin' down a dream\n That never would come to me\n Workin' on a mystery\n goin' wherever it leads\n Runnin' down a dream"

	scanner := bufio.NewScanner(strings.NewReader(s))

	/*
		the Scan() method returns a bool everytime it scans

		So as long as the bool returned is true, the loop will keep iterating which means we are not at the end of text yet!

		End of Input returns false, so we can exit the loop
	*/
	for scanner.Scan() {
		/*
		 Text() method gets you a new line in your text. eg "I felt so good like anything was possible" is first line for our string. \n signifies end of line which you already know

		 By default, the scanner reads your text line by line
		*/
		line := scanner.Text()
		fmt.Println(line)
	}

}
