/*----------------------------------------------------------------------------------*/
/*                  BNO055 Library For Tingo Test Program                           */
/*                      Programed By Keisuke Hongyo                                 */
/*                           Date 13.Aug.2024                                       */
/*                 Special Thanks for  Suzupy , Shun and Yukipy                     */
/*----------------------------------------------------------------------------------*/

package main

import (
	"bno055/bno055"
	"fmt"
	"image/color"
	"machine"
	"time"
	"tinygo.org/x/drivers/ssd1306"
)

func getSensor(ch1 chan<- bool) {
	var chk bool
	var roll, pich, yaw float64

	d := bno055.New(machine.I2C0)

	_ = d.Init(bno055.OPERATION_MODE_NDOF)

	for {
		for {
			chk, roll, pich, yaw = d.QuaternionToEuler()
			if chk {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}

		fmt.Printf("Euler roll=%f, pich=%f, yaw=%f \n", roll, pich, yaw)

		ch1 <- true
		time.Sleep(time.Millisecond * 100)
	}
}

func ctrlLed(ch2 chan<- bool) {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 100)

		led.High()
		time.Sleep(time.Millisecond * 100)

		ch2 <- true
	}
}

func procDisp(ch3 chan<- bool) {
	// Display
	display := ssd1306.NewI2C(machine.I2C0)

	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	display.ClearDisplay()

	x := int16(0)
	y := int16(0)
	deltaX := int16(1)
	deltaY := int16(1)

	for {
		pixel := display.GetPixel(x, y)
		c := color.RGBA{255, 255, 255, 255}
		if pixel {
			c = color.RGBA{0, 0, 0, 255}
		}
		display.SetPixel(x, y, c)
		_ = display.Display()

		x += deltaX
		y += deltaY

		if x == 0 || x == 127 {
			deltaX = -deltaX
		}

		if y == 0 || y == 63 {
			deltaY = -deltaY
		}
		time.Sleep(time.Millisecond * 50)
		ch3 <- true
	}
}

func init() {
	err := machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
	})
	if err != nil {
		return
	}
}

func main() {

	// 起動待ち
	time.Sleep(time.Millisecond * 100)

	// チャネル作成
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	ch3 := make(chan bool, 1)

	go getSensor(ch1)
	go ctrlLed(ch2)
	go procDisp(ch3)

	for {
		// Receive channel data from Goroutine
		select {
		case <-ch1:
			break
		case <-ch2:
			break
		case <-ch3:
			break
		}
	}
}
