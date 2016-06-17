// Run some code on play.golang.org and display the result
package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

// Errors are ignored for brevity.
func main() {
	var webDriver selenium.WebDriver
	var err error
	// set browser as firefox
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "firefox"})
	// remote to selenium server
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:8081/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("http://www.amazon.com/gp/goldbox?ie=UTF8&ref_=nav_cs_gb")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}
	// sleep for a while for fully loaded javascript
	time.Sleep(4 * time.Second)
	// get title
	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return
	}

	var elem selenium.WebElement
	elem, err = webDriver.FindElement(selenium.ByCSSSelector, "#widgetContent")
	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return
	}

	var firstElem selenium.WebElement
	firstElem, err = elem.FindElement(selenium.ByCSSSelector, ".a-section .dealContainer")
	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return
	}
	// get image
	image, err := firstElem.FindElement(selenium.ByCSSSelector, "img")
	if err == nil {
		img, _ := image.GetAttribute("src")
		fmt.Println(img)
	}
}
