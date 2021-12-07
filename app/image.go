package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pborman/uuid"
)

type Image struct {
	Base64 string `json:"base64"`
}

type Response struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func SaveImage(w http.ResponseWriter, r *http.Request) {
	var img Image

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&img)
	if err != nil {
		res := Response{
			Status:  false,
			Code:    500,
			Title:   "An Exception Throwed",
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
	} else {
		success := Response{
			Status:  true,
			Code:    200,
			Title:   "Success",
			Message: DecodeAndSave(img.Base64),
		}
		json.NewEncoder(w).Encode(success)
	}
}

func DecodeAndSave(s string) string {
	strArrey := strings.Split(s, ";")
	extension := strings.Split(strArrey[0], "/")[1]

	strRandom := strings.Replace(uuid.New(), "-", "", -1)
	imgName := strRandom + "." + extension

	sta := strings.Split(strArrey[len(strArrey)-1], ",")
	base64Body := sta[len(sta)-1]

	file, err := os.Create("img/" + imgName)
	if err != nil {
		return err.Error()
	} else {
		bytes, dcodeErr := base64.StdEncoding.DecodeString(base64Body)
		if dcodeErr != nil {
			return dcodeErr.Error()
		} else {
			os.Mkdir("/img", os.ModeDir)
			n, werr := file.Write(bytes)
			if werr != nil {
				return werr.Error()
			} else {
				fmt.Println(n)
				return imgName
			}
		}
	}
}
