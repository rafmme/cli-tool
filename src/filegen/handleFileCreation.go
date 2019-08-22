package filegen

import (
	"../runcmd"
	"../populatefile"
)

func HandleFileCreation(path, generatedFileName, fileTmpl string)  {
	fullPath, err := runcmd.Execute(path, generatedFileName)
	populatefile.IsError(err)
	populatefile.WriteContentToFile(*fullPath, fileTmpl)
}
