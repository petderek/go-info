package main

import (
	"log"
	"syscall"
)

func main() {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		log.Fatal("syscall failed: ", err)
	}

	log.Println("Sysname:", transcribe(uname.Sysname[:]))
	log.Println("Nodename:", transcribe(uname.Nodename[:]))
	log.Println("Release:", transcribe(uname.Release[:]))
	log.Println("Version:", transcribe(uname.Version[:]))
	log.Println("Machine:", transcribe(uname.Machine[:]))
	log.Println("Domainname:", transcribe(uname.Domainname[:]))
}

func transcribe(val []int8) string {
	var buf [65]byte
	for i, v := range val {
		if v == 0 {
			break
		}
		buf[i] = uint8(v)
	}
	return string(buf[:])
}
