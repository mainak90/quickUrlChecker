package checker

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

const maxCapacity = 50

func CheckUrlCmdLine(seperator string) {
	if (strings.Contains(seperator, ",")) {
		urls := strings.Split(seperator, ",")
		for _, url := range urls {
			go getWebsite(strings.TrimSpace(url))
			wg.Add(1)
		}
	} else {
		go getWebsite(seperator)
		wg.Add(1)
	}
	wg.Wait()
}

func CheckUrlConFile(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Sprintf("Error reading file %s : %s", f, err)
	}
	defer f.Close()
	r := bufio.NewScanner(f)
	buf := make([]byte, maxCapacity)
	r.Buffer(buf, maxCapacity)
	for r.Scan() {
		go getWebsite(r.Text())
		wg.Add(1)
	}
	wg.Wait()
}

func getWebsite(website string) {
	defer wg.Done()
	if res, err := http.Get(website); err != nil {
		fmt.Printf("Website %s is down!\n", website)
	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
	}
}