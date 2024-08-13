// This example shows how to use 128x64 display over I2C
// Tested on Seeeduino XIAO Expansion Board https://wiki.seeedstudio.com/Seeeduino-XIAO-Expansion-Board/
//
// According to manual, I2C address of the display is 0x78, but that's 8-bit address.
// TinyGo operates on 7-bit addresses and respective 7-bit address would be 0x3C, which we use below.
//
// To learn more about different types of I2C addresses, please see following page
// https://www.totalphase.com/support/articles/200349176-7-bit-8-bit-and-10-bit-I2C-Slave-Addressing

package main

import (
	"bno055/bno055"
	"fmt"
	"machine"
	"time"
)

func main() {
	var chk bool
	var roll, pich, yaw float64
	var proc uint8
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
	})

	led := machine.LED

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d := bno055.New(machine.I2C0)

	if d.Init() {
		proc = 6
		for {
			led.Low()
			time.Sleep(time.Millisecond * 50)
			switch proc {
			case 1:
				chk = d.GetAccl()
				if !chk {
					println("Error")
				} else {
					fmt.Printf("Accl xData=%d, Ydata=%d, Zdata=%d \n",
						d.SensorData.AcclData.XData, d.SensorData.AcclData.YData, d.SensorData.AcclData.ZData)
				}
				break
			case 2:
				chk = d.GetGyro()
				if !chk {
					println("Error")
				} else {
					fmt.Printf("Gyro xData=%d, Ydata=%d, Zdata=%d \n",
						d.SensorData.GyroData.XData, d.SensorData.GyroData.YData, d.SensorData.GyroData.ZData)
				}
				break
			case 3:
				chk = d.GetMag()
				if !chk {
					println("Error")
				} else {
					fmt.Printf("Mag xData=%d, Ydata=%d, Zdata=%d \n",
						d.SensorData.MagData.XData, d.SensorData.MagData.YData, d.SensorData.MagData.ZData)
				}
				break
			case 4:
				chk = d.GetEuler()
				if !chk {
					println("Error")
				} else {
					fmt.Printf("Quaternion RData=%d, Pdata=%d, Hdata=%d \n",
						d.SensorData.EulerData.RData, d.SensorData.EulerData.PData, d.SensorData.EulerData.HData,
					)
				}
				break
			case 5:
				chk = d.GetQuaternion()
				if !chk {
					println("Error")
				} else {
					fmt.Printf("Quaternion wData=%d, xData=%d, Ydata=%d, Zdata=%d \n",
						d.SensorData.QutaData.WData, d.SensorData.QutaData.XData,
						d.SensorData.QutaData.YData, d.SensorData.QutaData.ZData,
					)
				}
				break
			case 6:
				_, temp := d.GetTemp()
				time.Sleep(time.Millisecond * 10)

				for {
					chk, roll, pich, yaw = d.QuaternionToEuler()
					if chk {
						break
					}
					time.Sleep(time.Millisecond * 100)
				}

				fmt.Printf("Euler temp=%d roll=%f, pich=%f, yaw=%f \n", temp, roll, pich, yaw)
				break
			}

			led.High()
			time.Sleep(time.Millisecond * 50)
		}
	}
}
