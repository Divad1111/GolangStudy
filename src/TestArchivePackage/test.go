package TestArchivePackage

import (
	"archive/tar"
	//"archive/zip"
	"log"
	"os"
)

func Test() {

	log.Println("===================FileInfoHeader Info===================")
	if file, err := os.Stat("E:/WorkSpace/OpenSource/GolangStudy/src/TestArchivePackage/testdata/test.tar"); err == nil {
		if head, err1 := tar.FileInfoHeader(file, ""); err1 == nil {
			log.Printf("%+v", head)
		} else {
			log.Fatal(err1)
		}

	} else {
		log.Fatal(err)
	}

	log.Println("===================Read tar file===================")
	if file, err := os.Open("E:/WorkSpace/OpenSource/GolangStudy/src/TestArchivePackage/testdata/test.tar"); err == nil {
		if reader := tar.NewReader(file); reader != nil {
			if hd, err2 := reader.Next(); err2 == nil {
				log.Printf("HeadInfo:%+v", hd)
				b := make([]byte, 400)
				if ret, err1 := reader.Read(b); err1 == nil {
					log.Printf("ReadSize:%d, ReadContent:%s", ret, b)
				} else {
					log.Fatal(err1)
				}

			} else {
				log.Fatal(err2)
			}
		} else {

		}

		defer file.Close()
	} else {
		log.Fatal(err)
	}
}
