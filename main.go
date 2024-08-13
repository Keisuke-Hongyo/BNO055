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
	font "github.com/Nondzu/ssd1306_font"
	"machine"
	"time"
	"tinygo.org/x/drivers/ssd1306"
)

type sensor struct {
	roll float64
	pich float64
	yaw  float64
}

func getSensor(snrCh chan<- sensor) {
	var chk bool
	var snr sensor

	d := bno055.New(machine.I2C0)

	_ = d.Init(bno055.OPERATION_MODE_NDOF)

	for {
		for {
			chk, snr.roll, snr.pich, snr.yaw = d.QuaternionToEuler()
			if chk {
				break
			}
			time.Sleep(time.Millisecond * 100)
		}

		fmt.Printf("Euler roll=%f, pich=%f, yaw=%f \n", snr.roll, snr.pich, snr.yaw)

		snrCh <- snr
		time.Sleep(time.Millisecond * 10)
	}
}

func ctrlLed(ledCh chan<- bool) {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 100)

		led.High()
		time.Sleep(time.Millisecond * 100)

		ledCh <- true
	}
}

func procDisp(snrCh <-chan sensor, ch3 chan<- bool) {

	dev := ssd1306.NewI2C(machine.I2C0)

	dev.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	dev.ClearBuffer()
	dev.ClearDisplay()

	//font library init
	display := font.NewDisplay(dev)
	display.Configure(font.Config{FontType: font.FONT_11x18})

	lcdprint := func(x int16, y int16, str string) {
		display.XPos = x // set position X
		display.YPos = y // set position Y
		display.PrintText(str)
	}

	for {
		select {
		case s := <-snrCh:
			lcdprint(0, 2, fmt.Sprintf("Roll=%f", s.roll))
			lcdprint(0, 22, fmt.Sprintf("Pich=%f", s.pich))
			lcdprint(0, 42, fmt.Sprintf("Yaw=%f", s.yaw))
			break
		}
		time.Sleep(time.Millisecond * 10)
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
	snrCh := make(chan sensor, 1)
	ledCh := make(chan bool, 1)
	devCh := make(chan bool, 1)

	go getSensor(snrCh)
	go ctrlLed(ledCh)
	go procDisp(snrCh, devCh)

	for {
		// Receive channel data from Goroutine
		select {
		case <-ledCh:
			break
		case <-devCh:
			break
		}
	}
}
