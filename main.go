package main

import (
	"context"
	"github/Tamiquell/mongol-lessons-tg/internal/tg"
	"log"
	"os"

	"github/Tamiquell/mongol-lessons-tg/internal/messages"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tgClient, err := tg.New(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}
	msgModel := messages.New(tgClient)
	tgClient.ListenUpdates(ctx, msgModel)
}

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"reflect"
// 	"strings"
// )

// func main() {
// 	var verbs_rus = []string{"взять", "принести", "работать", "бояться", "охотиться", "убить", "шагать", "отдыхать", "попробовать на вкус", "жить"}
// 	var verbs_mong = []string{"авах", "авчрах", "ажиллах", "айх", "агнах", "алах", "алхах", "амрах", "амсах", "амьдрах"}

// 	first, second := "а", "б"

// 	if string(verbs_mong[0][0]) > "б" {
// 		fmt.Println("AAAAAAAA")
// 	}

// 	for i, _ := range verbs_mong {
// 		// fmt.Println(i, strings.ToLower(string(verbs_mong[i][1])), []rune(string(verbs_mong[i][1]))[0])
// 		fmt.Println(i, verbs_mong[i], reflect.TypeOf(verbs_mong[i][0]))

// 		// if strings.ToLower(string(verbs_mong[i][0])) >= first {
// 		// 	fmt.Println("INSIDE 1")
// 		// }
// 	}

// 	b := new(bytes.Buffer)
// 	for i, _ := range verbs_mong {
// 		if (strings.ToLower(string(verbs_mong[i][0])) >= first) &&
// 			(strings.ToLower(string(verbs_mong[i][0])) <= second) {
// 			fmt.Println("INSIDE")
// 			fmt.Fprintf(b, "%s - %s\n", verbs_mong[i], verbs_rus[i])
// 		}
// 	}
// 	fmt.Println(b.String())
// }
