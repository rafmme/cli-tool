package filegen

import (
	"fmt"
	"flag"
	"os/exec"
	"log"
)

func Create(kind, fileName, path string)  {
	if kind == "" {
		log.Print("Error! Expects the `kind` flag")
		return
	}
	
	switch kind {
	case "namespace", "ns":
		fmt.Println(kind)
		break
	case "deployment":
	default:
		log.Print(`doee`)
		return
	}

	_, err := exec.Command(
		"bash", "-c",  "cd "+path+" && touch "+fileName+".yaml",
	).Output()

	if err != nil {
		log.Fatal(err)
		return
	}
}

func Run() {
	kind := flag.String("kind", "", "Kubernetes API Object")
	path := flag.String("path", "./", "")
	fileName := flag.String("filename", "default", "")
	
	flag.Parse()

	Create(*kind, *fileName, *path)
}
