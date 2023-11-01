package main

import (
	"log"
	"time"

	"github.com/googolgl/go-i2c"
	"github.com/googolgl/go-pca9685"
)

const (
	i2cAddress = 0x40
	i2cDevice  = "/dev/i2c-1"

	MaxPulse = pca9685.ServoMaxPulseDef
	MinPulse = pca9685.ServoMinPulseDef
	AcRange  = pca9685.ServoRangeDef

	//Steer Cfg
	steerMaxPulse = 2000
	steerMinPulse = 1000
	steerChannel  = 3

	//ESC Cfg
	escMaxPulse = 2000
	escMinPulse = 1000
	escChannel  = 2
)

func main() {
	i2c, err := i2c.New(i2cAddress, i2cDevice)
	if err != nil {
		log.Printf("error: failed starting i2c with address - %s", err.Error())
		return
	}

	driver, err := pca9685.New(i2c, nil)
	if err != nil {
		log.Printf("error: failed getting servo driver - %s", err.Error())
		return
	}

	steerServo := driver.ServoNew(steerChannel, &pca9685.ServOptions{
		AcRange:  AcRange,
		MinPulse: float32(steerMinPulse),
		MaxPulse: float32(steerMaxPulse),
	})

	// esc := driver.ServoNew(escChannel, &pca9685.ServOptions{
	// 	AcRange:  AcRange,
	// 	MinPulse: float32(escMinPulse),
	// 	MaxPulse: float32(escMaxPulse),
	// })

	log.Printf("turning left: %0.2f\n", 0.25)
	steerServo.Fraction(0.25)
	time.Sleep(2 * time.Second)

	log.Printf("turning right: %0.2f\n", 0.75)
	steerServo.Fraction(0.75)
	time.Sleep(2 * time.Second)

	log.Printf("center: %0.2f\n", 0.5)
	steerServo.Fraction(0.5)
	time.Sleep(2 * time.Second)
}
