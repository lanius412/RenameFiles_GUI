package rename

import(
	"path/filepath"
	"io/ioutil"

	"strings"
	"strconv"

	"os"

	"log"
)

func Rename(dir string) {

	basePath := filepath.Dir(dir) + "/" + filepath.Base(dir) + "/"

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	fileNum := len(fileInfos)

	files := make([]string, 0, 1000)
	for _, fileInfo := range fileInfos {
		if strings.HasPrefix(fileInfo.Name(), "."){
			fileNum--
			continue
		} else {
			files = append(files, basePath+fileInfo.Name())
		}
	}
	
	if fileNum >= 1000 {
		log.Println("Sry, Support up to 3digits")
		os.Exit(0)
	}

	for num, file := range files {
		fileName := strings.Replace(file, basePath, "", -1)
		
		extPos := strings.LastIndex(fileName, ".")
		ext := fileName[extPos:] //ex).txt
		
		renamedPath := basePath
		num++
		if fileNum >= 100 && fileNum < 1000{ //Hundreds of Files
			renamedPath = ChangeNameandPath(num, renamedPath, ext)
		} else if fileNum >= 10 { //Dozens of Files
			renamedPath = ChangeNameandPath(num, renamedPath, ext)
		} else if fileNum < 10 { //Several Files
			renamedPath += "0" + strconv.Itoa(num) + ext
		}

		log.Println(fileName + " -> " + strings.Replace(renamedPath, basePath, "", -1))

		os.Rename(file, renamedPath)
		}
	}

	log.Println(strconv.Itoa(fileNum) + " Files Rename Completed")

}

func ChangeNameandPath(fileCount int, renamedPath string, ext string) string{
	switch {
		case fileCount >= 100:
			renamedPath += strconv.Itoa(fileCount) + ext //3digits
		case fileCount >= 10:
			renamedPath += "0" + strconv.Itoa(fileCount) + ext //2digits
		case fileCount < 10:
			renamedPath += "00" + strconv.Itoa(fileCount) + ext //1digit
	}
	return renamedPath
}
