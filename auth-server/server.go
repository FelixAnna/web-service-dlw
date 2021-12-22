package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/felixanna/web-service-dlw/auth-server/auth"
)

var (
	portvar int
)

func init() {
	flag.IntVar(&portvar, "p", 9096, "the base port for the server")
}

func main() {
	chanTest()
	flag.Parse()

	log.Println("Dumping requests")

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/auth", auth.AuthHandler)

	http.HandleFunc("/oauth/authorize", auth.OAuthAuthorize)

	http.HandleFunc("/oauth/token", auth.OAuthToken)

	http.HandleFunc("/test", auth.Test)

	log.Printf("Server is running at %d port.\n", portvar)
	log.Printf("Point your OAuth client Auth endpoint to %s:%d%s", "http://localhost", portvar, "/oauth/authorize")
	log.Printf("Point your OAuth client Token endpoint to %s:%d%s", "http://localhost", portvar, "/oauth/token")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portvar), nil))
}

func chanTest() {
	fmt.Println(solution(""))
	och := make(chan int)

	list := []int{1, 2, 1, 2, 3, 4, 3, 4, 5, 6, 7}

	start := time.Now().UnixNano()
	var fl int = len(list) / 2
	if len(list)%2 != 0 {
		fl += 1
	}
	ch1 := make([]int, fl)
	ch2 := make([]int, len(list)/2)

	go func() {
		for _, val := range list {
			och <- val
		}

		close(och)
	}()

	i, j := 0, 0
	for {
		el1, ok := <-och
		if ok {
			ch1[i] = el1
			i++
			//ch1 = append(ch1, el1)
		} else {
			break
		}

		el2, ok := <-och
		if ok {
			ch2[j] = el2
			j++
			//ch2 = append(ch2, el2)
		} else {
			break
		}

	}

	end := time.Now().UnixNano()
	fmt.Println(end - start)
	fmt.Println(ch1)
	fmt.Println(ch2)
}

func solution(source string) bool {

	kuohao := make([]int, 0)

	// 在这⾥写代码
	for _, val := range source {
		if val == '(' {
			kuohao = append(kuohao, 1)
			continue
		}

		if val == ')' {
			kuohao = append(kuohao, -1)
		}
	}

	sum := 0
	for _, val := range kuohao {
		sum += val

		if sum < 0 {
			return false
		}
	}

	return sum == 0
}
