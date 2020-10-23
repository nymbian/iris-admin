package libs

import (
	"io"
	"github.com/nymbian/iris-admin/common"
	"log"

	"os"
	"path"
	"strconv"

	//"strings"

	"github.com/kataras/iris"
	//config "github.com/spf13/viper"
)

func UploadFile(key string, Ctx iris.Context) (bool, string) {
	file, info, err := Ctx.FormFile(key)
	filePath := ""
	if err != nil {
		return false, "Error while uploading: <b>" + err.Error() + "</b>"
	}

	var minSize int64 = 0
	//maxSize := config.GetInt64("UploadSize") * 1024 * 1024
	var maxSize int64 = 5 * 1024 * 1024
	log.Printf("%d Max", maxSize)
	if info.Size > minSize {
		if info.Size > maxSize {
			log.Printf("%d UploadSize ToMax", info.Size)
			return false, "Error while uploading: UploadSize ToMax"
		}
		fname := strconv.Itoa(common.GenerateRangeNum(100, 9999)) + "_" + info.Filename

		fileSuffix := path.Ext(fname)

		fileSuffixExists := false
		CanFileSuffix := [...]string{".jpg", ".png", ".jpeg", ".gif"}
		//CanFileSuffix := strings.Split(config.GetString("UploadSuffixExists"), ",")
		for _, v := range CanFileSuffix {
			//log.Fatalln(fileSuffix)
			if v == fileSuffix {
				fileSuffixExists = true
			}
		}

		if fileSuffixExists == false {
			return false, "fileSuffix error: <b>" + fileSuffix + "</b>"
		}

		filePath = "./uploads/avatar/" + fname
		out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return false, "Error while uploading: <b>" + err.Error() + "</b>"
		}
		defer out.Close()
		io.Copy(out, file)
		filePath = filePath[1:]
	}
	defer file.Close()
	return true, filePath
}
