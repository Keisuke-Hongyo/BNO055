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
	"machine"
	"time"
)

func main() {
	var chk bool
	var roll, pich, yaw float64
	var proc uint8
	//var cab bno055.SensorCalibration
	err := machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
	})
	if err != nil {
		return
	}

	led := machine.LED

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d := bno055.New(machine.I2C0)

	if d.Init(bno055.OPERATION_MODE_NDOF) {
		/*fmt.Println("Cab Start")
		for {
			time.Sleep(10 * time.Millisecond)
			chk, cab = d.GetCalibration()
			if !chk {
				continue
			} else if cab.Gyro != 0x03 || cab.Mag != 0x03 {
				fmt.Println(cab)
				continue
			} else {
				break
			}
		}
		fmt.Println("Cab End")*/
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
