package TestOSPackage

import (
	"log"
	"os"
	"os/user"
)

func Test() {

	//Print environment
	log.Println("================Print environment(================")
	envs := os.Environ()
	for _, s := range envs {
		log.Println(s)
	}

	//Executable
	log.Println("================Print executable(================")
	if path, err := os.Executable(); err != nil {
		log.Println(err)
	} else {
		log.Println(path)
	}

	log.Println("================Print page size================")
	log.Println(os.Getpagesize())

	log.Println("================Print wd(================")
	if wd, err := os.Getwd(); err != nil {
		log.Println(err)
	} else {
		log.Println(wd)
	}

	log.Println("================Print hostname(================")
	if hn, err := os.Hostname(); err != nil {
		log.Println(err)
	} else {
		log.Println(hn)
	}

	log.Println("================Print current user(================")
	if user, err := user.Current(); err != nil {
		log.Println(err)
	} else {
		log.Println(user)
	}

}
