package main

/*
// some symbols (within the Zabbix agent) won't resolve at link-time
// we can ignore these and resolve at runtime
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
*/
import "C"

func main() {

}
