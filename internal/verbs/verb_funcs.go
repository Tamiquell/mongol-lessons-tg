package verbs

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// var verbs_rus = []string{"взять", "принести", "работать", "бояться", "охотиться", "убить", "шагать", "отдыхать", "попробовать на вкус", "жить"}
// var verbs_mong = []string{"авах", "авчрах", "ажиллах", "айх", "агнах", "алах", "алхах", "амрах", "амсах", "амьдрах"}

var verbs_mong []string
var verbs_rus []string
var match = map[string]int{
	"а": 0,
	"б": 28,
	"г": 29,
	"з": 60,
	"и": 61,
	"ө": 92,
	"с": 93,
	"т": 122,
	"у": 123,
	"х": 167,
	"ц": 168,
	"я": 185,
}

func VerbsList(first, second string) (string, error) {
	b := new(bytes.Buffer)
	left, _ := match[first]
	right, _ := match[second]
	for i, _ := range verbs_mong {
		if (i >= left) && (i <= right) {
			fmt.Fprintf(b, "%s - %s\n", verbs_mong[i], verbs_rus[i])
		}
	}
	return b.String(), nil
}

func RandomVerbs() (string, string, error) {
	var resultRus, resultMong []string
	ids := rand.Perm(len(verbs_mong))
	for _, i := range ids[:5] {
		resultRus = append(resultRus, verbs_rus[i])
		resultMong = append(resultMong, verbs_mong[i])
	}

	resRus := strings.Join(resultRus[:], "\n")
	resMong := strings.Join(resultMong[:], "\n")

	return resRus, resMong, nil
}

func FillVerbs() error {
	f, err := os.Open("data/verbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " - ")
		verbs_mong = append(verbs_mong, row[0])
		verbs_rus = append(verbs_rus, row[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func FullVerbsList(firstLetter string) (string, error) {
	if firstLetter == "" {
		b := new(bytes.Buffer)
		for i, _ := range verbs_mong {
			fmt.Fprintf(b, "%s - %s\n", verbs_mong[i], verbs_rus[i])
		}
		return b.String(), nil
	} else {
		b := new(bytes.Buffer)
		for i, _ := range verbs_mong {
			if strings.ToLower(string(verbs_mong[i][0])) != firstLetter {
				break
			}
			fmt.Fprintf(b, "%s - %s\n", verbs_mong[i], verbs_rus[i])
		}
		return b.String(), nil
	}
}
