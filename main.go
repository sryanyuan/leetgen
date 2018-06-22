package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

var (
	flagLanguage string
	flagOutput   string
	flagProblem  string
)

func getDescription(problemName string) string {
	doc, err := goquery.NewDocument(fmt.Sprintf("https://leetcode.com/problems/%s/description/", problemName))
	if err != nil {
		log.Fatal(err)
	}

	var desc string
	doc.Find("meta[name=description]").Each(func(i int, selection *goquery.Selection) {
		desc, _ = selection.Attr("content")
	})
	return desc
}

func getCodeDefinition(problemName string) string {
	var err error
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	cdp, err := chromedp.New(ctx)
	if nil != err {
		panic(err)
	}

	var definition string
	err = cdp.Run(ctx, createCodeDefinitionTask(problemName, &definition, lanValueFromString(flagLanguage)))
	if nil != err {
		panic(err)
	}

	err = cdp.Shutdown(ctx)
	if nil != err {
		panic(err)
	}

	err = cdp.Wait()
	if nil != err {
		panic(err)
	}

	return definition
}

func createCodeDefinitionTask(problemName string, definition *string, lanv int) chromedp.Tasks {
	textarea := `//textarea`
	btn := `#question-detail-app > div > div:nth-child(3) > div > div > div.row.control-btn-bar > div > div > div > div > span.Select-arrow-zone`
	lanSel := fmt.Sprintf(`#react-select-2--option-%d`, lanv)
	return chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("https://leetcode.com/problems/%s/description/", problemName)),
		chromedp.Click(btn, chromedp.ByID),
		chromedp.Click(lanSel, chromedp.ByID),
		chromedp.Text(textarea, definition),
	}
}

func initFlag() error {
	flag.StringVar(&flagLanguage, "lan", "", "Problem solve language [C++/Go]")
	flag.StringVar(&flagOutput, "output", "", "Solve problem template source file output directory")
	flag.StringVar(&flagProblem, "problem", "", "Problem name, in slug format")
	flag.Parse()

	if "" == flagLanguage {
		return errors.New("Invalid lan")
	}
	lanv := lanValueFromString(flagLanguage)
	if -1 == lanv {
		return fmt.Errorf("Unsupport language %s", flagLanguage)
	}
	if "" == flagProblem {
		return errors.New("Invalid problem name")
	}
	return nil
}

func lanValueFromString(lan string) int {
	llan := strings.ToLower(lan)
	lans := []string{
		"c++",
		"java",
		"python",
		"python3",
		"c",
		"c#",
		"javascript",
		"ruby",
		"swift",
		"go",
		"scala",
		"kotlin",
	}
	for i, v := range lans {
		if v == llan {
			return i
		}
	}
	return -1
}

func lanExtensionFromString(lan string) string {
	lanv := lanValueFromString(lan)
	lans := []string{
		"h",
		"java",
		"python",
		"python3",
		"c",
		"c#",
		"javascript",
		"ruby",
		"swift",
		"go",
		"scala",
		"kotlin",
	}
	return lans[lanv]
}

func makefile(desc string, code string) error {
	filePath := ""
	if flagOutput != "" {
		filePath = flagOutput + "/"
	}
	filePath += fmt.Sprintf("%s.%s", flagProblem, lanExtensionFromString(flagLanguage))
	of, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if nil != err {
		return fmt.Errorf("Open file error: %v", err)
	}
	defer of.Close()

	llan := strings.ToLower(flagLanguage)
	switch llan {
	case "c++":
		{
			return makefileCC(of, desc, code)
		}
	default:
		{
			return fmt.Errorf("Language %s not support now", llan)
		}
	}
}

func main() {
	err := initFlag()
	if nil != err {
		fmt.Println(err)
		flag.PrintDefaults()
		return
	}

	desc := getDescription(flagProblem)
	if "" == desc {
		fmt.Printf("Get problem description failed\r\n")
		return
	}
	codeDefinition := getCodeDefinition(flagProblem)

	err = makefile(desc, codeDefinition)
}
