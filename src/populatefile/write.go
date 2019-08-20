package populatefile

import (
	"os"
	"fmt"
)

func WriteContentToFile(path, content string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()
	
	_, err = file.WriteString(content)
	if isError(err) { return }

	err = file.Sync()
	if isError(err) { return }
	
	fmt.Println("Done")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
