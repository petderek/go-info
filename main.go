package main

import (
	"log"

	"golang.org/x/sys/unix"
)

func main() {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		log.Fatal("syscall failed: ", err)
	}

	log.Println("Sysname:", string(uname.Sysname[:]))
	log.Println("Nodename:", string(uname.Nodename[:]))
	log.Println("Release:", string(uname.Release[:]))
	log.Println("Version:", string(uname.Version[:]))
	log.Println("Machine:", string(uname.Machine[:]))
	log.Println("Domainname:", string(uname.Domainname[:]))
}
