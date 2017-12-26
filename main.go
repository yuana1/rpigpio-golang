package rpigpio

/*
#cgo CFLAGS: -I./wiringPi/include
#cgo LDFLAGS: -L./wiringPi/lib -lprint
#include "print.h"
#include <stdlib.h>
int foo ( void (*fn) (void)) {
	fn();
}

 */
import "C"
import (
	"unsafe"
	"fmt"
)
//
//// set up
//func WiringPiSetup() int {
//	return C.wiringPiSetup()
//}
//func WiringPiSetupGpio() int {
//	return int(C.wiringPiSetupGpio())
//}
//func WiringPiSetupPhys() int {
//	return int(C.wiringPiSetupPhys())
//}
//func WiringPiSetupSys() int {
//	return int(C.wiringPiSetupSys())
//}
//
////core function
//func PinMode(pin, mode int) {
//	C.pinMode(C.int(pin), C.int(mode))
//}
//func PullUpDnControl(pin, pud int) {
//	C.pullUpDnControl(C.int(pin), C.int(pud))
//}
//func DigitalWrite(pin, value int) {
//	C.digitalWrite(C.int(pin), C.int(value))
//}
//func PwmWrite(pin, value int) {
//	C.pwmWrite(C.int(pin), C.int(value))
//}
//func DigitalRead(pin int) int {
//	return int(C.digitalRead(C.int(pin)))
//}
//func AnalogRead(pin int) {
//	C.analogRead(C.int(pin))
//}
//func AnalogWrite(pin, value int) {
//	C.analogWrite(C.int(pin), C.int(value))
//}
//
//// thread priority
//func piHiPri(priority int) int {
//	return int(C.piHiPri(C.int(priority)))
//}
//
//// interrupts
//func waitForInterrupt(pin, timeOut int) int {
//	return int(C.waitForInterrupt(C.int(pin), C.int(timeOut)))
//}
//
//func wiringPiISR(pin, edgeType int, function func()) int {
//	return int(C.wiringPiISR(C.int(pin), C.int(edgeType), unsafe.Pointer(&function)))
//}
//
//// timing
//func Millis() uint {
//	return uint(C.millis())
//}
//func Micros() uint {
//	return uint(C.micros())
//}
//func Delay(howLong uint) {
//	C.delay(C.uint(howLong))
//}
//func DelayMicroseconds(howLong uint) {
//	C.delayMicroseconds(C.uint(howLong))
//}
//
////serial
//func SerialOpen(device string, baud int) int {
//	c := C.C.CString(device)
//	defer C.free(unsafe.Pointer(c))
//	return C.serialOpen(c, baud)
//}
//func SerialClose(fd int) {
//	C.serialClose(C.int(fd))
//}
//func SerialPuts(fd int, s string) {
//	c := C.CString(s)
//	defer C.free(unsafe.Pointer(c))
//	C.serialPuts(C.int(fd), c)
//
//}
//func SerialPrintf(fd int, message ...string) {
//	len := len(message)
//	for i := 0; i < len; i++ {
//		d := C.CString(message[i])
//		C.serialPrintf(C.int(fd), d)
//		C.free(unsafe.Pointer(d))
//	}
//}
//func SerialDataAvail(fd int) int {
//	return int(C.serialDataAvail(C.int(fd)))
//}
//
//func SerialGetchar(fd int) int {
//	return int(C.serialGetchar(C.int(fd)))
//}
//func SerialFlush(fd int) {
//	C.serialFlush(C.int(fd))
//}

//export cal
func cal () {
	fmt.Print("aaa")
}
func main() {

	fn := cal
	C.foo(unsafe.Pointer(&fn))
}
