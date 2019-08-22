package populatefile

import (
	"os"
	"log"
	"fmt"
	"strings"
)

func WriteContentToFile(path, content string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if IsError(err) { return }
	defer file.Close()
	
	_, err = file.WriteString(content)
	if IsError(err) { return }

	err = file.Sync()
	if IsError(err) { return }
	
	fmt.Println(strings.TrimSpace(fmt.Sprintf("File %s has been successfully created", path)))
}

func IsError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
	}

	return (err != nil)
}
