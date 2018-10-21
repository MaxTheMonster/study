package main

import (
	"bytes"
	"fmt"
)

type subject struct {
	name   string
	topics []string
}

// getSubjects returns a formatted string containing the user's subjects
func getSubjects(subjects []subject) string {
	var buffer bytes.Buffer
	var formattedSubjects string
	for _, s := range subjects {
		buffer.WriteString("\n" + s.name + "\n")
		for _, topic := range s.topics {
			buffer.WriteString("\t" + topic + "\n")
		}
	}
	formattedSubjects = buffer.String()
	return formattedSubjects
}

func main() {
	subjects := []subject{{"biology", []string{"infection and disease", "cell biology"}}, {"religious studies", []string{"the trinity", "creation", "atonement", "eschatology", "the incarnation"}}}
	fmt.Println(getSubjects(subjects))
}
