package wiringPi

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lwiringPi
#include <stdlib.h>
#include <wiringPi.h>
#include <wiringSerial.h>
extern void ISRCallback()
static inline int wiringPiISR1 (int pin, int mode) {
	return wiringPiISR(pin, mode, ISRCallback);
}
 */
import "C"
import (
	_ "unsafe"
	"unsafe"
)

//export ISRCallback
func ISRCallback() {
	callBackFunction()
}

var callBackFunction = func() {

}
func Register(fn func()) {
	callBackFunction = fn
}

func Unregister(name string) {
	callBackFunction = func() {
		log.Println("default isr")
	}
}

// set up
func WiringPiSetup() int {
	return int(C.wiringPiSetup())
}
func WiringPiSetupGpio() int {
	return int(C.wiringPiSetupGpio())
}
func WiringPiSetupPhys() int {
	return int(C.wiringPiSetupPhys())
}
func WiringPiSetupSys() int {
	return int(C.wiringPiSetupSys())
}

//core function
func PinMode(pin, mode int) {
	C.pinMode(C.int(pin), C.int(mode))
}
func PullUpDnControl(pin, pud int) {
	C.pullUpDnControl(C.int(pin), C.int(pud))
}
func DigitalWrite(pin, value int) {
	C.digitalWrite(C.int(pin), C.int(value))
}
func PwmWrite(pin, value int) {
	C.pwmWrite(C.int(pin), C.int(value))
}
func DigitalRead(pin int) int {
	return int(C.digitalRead(C.int(pin)))
}
func AnalogRead(pin int) {
	C.analogRead(C.int(pin))
}
func AnalogWrite(pin, value int) {
	C.analogWrite(C.int(pin), C.int(value))
}

// thread priority
func PiHiPri(priority int) int {
	return int(C.piHiPri(C.int(priority)))
}

// interrupts
func WaitForInterrupt(pin, timeOut int) int {
	return int(C.waitForInterrupt(C.int(pin), C.int(timeOut)))
}

func WiringPiISR(pin, edgeType int) int {
	return int(C.wiringPiISR1(C.int(pin), C.int(edgeType)))
}

// timing
func Millis() uint {
	return uint(C.millis())
}
func Micros() uint {
	return uint(C.micros())
}
func Delay(howLong uint) {
	C.delay(C.uint(howLong))
}
func DelayMicroseconds(howLong uint) {
	C.delayMicroseconds(C.uint(howLong))
}

//serial
func SerialOpen(device string, baud int) int {
	c := C.CString(device)
	defer C.free(unsafe.Pointer(c))
	return int(C.serialOpen(c, C.int(baud)))
}
func SerialClose(fd int) {
	C.serialClose(C.int(fd))
}
func SerialPutchar(fd int, c uint8) {
	C.serialPutchar(C.int(fd), C.uchar(c))
}

func SerialPuts(fd int, s string) {
	c := C.CString(s)
	defer C.free(unsafe.Pointer(c))
	C.serialPuts(C.int(fd), c)

}
func SerialPrintf(fd int, message ...string) {
	len := len(message)
	for i := 0; i < len; i++ {
		d := C.CString(message[i])
		C.serialPuts(C.int(fd), d)
		C.free(unsafe.Pointer(d))
	}
}
func SerialDataAvail(fd int) int {
	return int(C.serialDataAvail(C.int(fd)))
}

func SerialGetchar(fd int) int {
	return int(C.serialGetchar(C.int(fd)))
}
func SerialFlush(fd int) {
	C.serialFlush(C.int(fd))
}
