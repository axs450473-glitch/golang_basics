package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	/*
		tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you?")
		if err != nil {
			panic(err)
		}

		// Define data for the Welcome message template

		data := map[string]interface{}{
			"name": "John",
		}
		err = tmpl.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
		fmt.Println()
		// without handling errors by ourselves
		templt := template.Must(template.New("Example1").Parse("Welcome, {{.name}}! All good ?"))
		templt.Execute(os.Stdout, data)

	*/
	// CONSOLE APPlication

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	// Define named template for different types of

	tempaltes := map[string]string{
		"welcome":      "Welcome, {{.name}}! We are glad you joined",
		"notification": "{{.name}}, you have new notification: {{.notification}}",
		"error":        "Oops ! An error occured : {{.errorMessage}}",
	}
	//initialize templates using for loop
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range tempaltes {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}
	for {
		// Show Menu
		fmt.Println("\nMenu")
		fmt.Println("\n1. Join")
		fmt.Println("\n2. Get Notification")
		fmt.Println("\n3.Get Error")
		fmt.Println("\n4. Exit")
		fmt.Println("Choose an option")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Println("Enter your Notification Message :)")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"name": name, "notification": notification}
		case "3":
			fmt.Println("Enter your Error Message")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"errorMessage": errorMessage}
		case "4":
			fmt.Println("Exiting .....")
			return
		default:
			fmt.Println("Invalid choice")
			continue
		}

		// render and print the template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing the template")
		}
	}
}
