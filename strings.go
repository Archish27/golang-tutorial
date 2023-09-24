package main

import(
	"fmt"
	"strings"
)

var pl  = fmt.Println

func main() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl(sV1)
	pl("Length sV1", len(sV1))
	pl("Length", len(sV2))
	pl("Contains Another", strings.Contains(sV2, "word"))
	pl(" o index :", strings.Index(sV2, "o"))
	pl(" Replace :", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome Words \n" // \t \" \\
	sV3 = strings.TrimSpace(sV3)
	pl("Split : ", strings.Split("a-b-c-d", "-"))
	pl("Lower: ", strings.ToLower(sV3))
	pl("Upper: ", strings.ToUpper(sV3))
	pl("Prefix: ", strings.HasPrefix("TacoCat", "Taco"))
	pl("Suffix: ", strings.HasSuffix("TacoCat", "Cat"))

}
