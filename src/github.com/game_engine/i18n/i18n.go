package i18n

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var (
	path string
	Data map[string]string
)

func init() {
	Data = make(map[string]string)
	path = "server/global/locale.ini"
}

func add2Map(str string) {
	strs := strings.Split(str, "=")
	if len(strs) == 2 {
		key := strings.TrimSpace(strs[0]) //utf8 无 bom
		value := strings.TrimSpace(strs[1])
		Data[key] = value
	}
}

func GetInit(path_ ...string) error {
	if len(path_) == 1 {
		path = path_[0]
	}
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		return inputError
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if (readerError == io.EOF) && (len(inputString) == 0) {
			return nil
		} else {
			add2Map(inputString)
		}
	}
	return nil
}
