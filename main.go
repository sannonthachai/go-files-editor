package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// cwd, _ := os.Getwd()
	items, _ := ioutil.ReadDir(".")

	for _, item := range items {
		if item.IsDir() {
			myFile := item.Name() + "/deployment.yaml"

			lines := forloopFile(myFile)

			for i, line := range lines {
				if strings.Contains(line, "fieldPath: metadata.labels['group']") {
					lines[i] = `                  fieldPath: metadata.labels['group']
            - name: ServiceName
              valueFrom:
                fieldRef:
                  fieldPath: metadata.labels['app']`
				}
			}

			output := strings.Join(lines, "\n")
			err := ioutil.WriteFile(myFile, []byte(output), 0644)

			if err != nil {
				log.Fatalln(err)
			}

			// createFile(cwd, item.Name(), "deployment.yaml")

			// subitems, _ := ioutil.ReadDir(item.Name())
			// for _, subitem := range subitems {
			// 	if !subitem.IsDir() {
			// 		// handle file there
			// 		fmt.Println(item.Name() + "/" + subitem.Name())
			// 	}
			// }
		}
	}
}

func createFile(cwd, dirName, name string) {
	path := filepath.Join(cwd, dirName, name)
	newFilePath := filepath.FromSlash(path)
	_, err := os.Create(newFilePath)
	if err != nil {
		log.Fatalln(err)
	}
}

func forloopFile(path string) []string {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(f), "\n")
	return lines
}
