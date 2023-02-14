package upload_picture

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download_Image(URL, fileName string) bool {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return false
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	return err == nil
}
