package bno055

import (
	"math"
	"time"
	"tinygo.org/x/drivers"
)

// Accl 加速度センサデータ構造体
type Accl struct {
	XData uint16 // X方向
	YData uint16 // Y方向
	ZData uint16 // Z方向
}

// Gyro ジャイロセンサデータ構造体
type Gyro struct {
	XData uint16 // X方向
	YData uint16 // Y方向
	ZData uint16 // Z方向
}

// Mag 地磁気センサデータ構造体
type Mag struct {
	XData uint16 // X方向
	YData uint16 // Y方向
	ZData uint16 // Z方向
}

// Euler オイラー角データ構造体
type Euler struct {
	HData uint16
	RData uint16
	PData uint16
}

// Quaternion 四元数データ構造体
type Quaternion struct {
	WData uint16
	XData uint16
	YData uint16
	ZData uint16
}

// Sensor センサ情報格納用構造体
type Sensor struct {
	AcclData  Accl
	GyroData  Gyro
	MagData   Mag
	EulerData Euler
	QutaData  Quaternion
}

// Device デバイス情報格納用構造体
type Device struct {
	bus           drivers.I2C
	deviceAddress uint8
	chipId        uint8
	SensorData    Sensor
}

//--- 外部公開レシーバ ---

// New レシーバ作成
func New(bus drivers.I2C, deviceAddress ...uint8) Device {
	d := Device{
		bus:           bus,
		deviceAddress: BNO055_Address_A,
		chipId:        0,
		SensorData:    Sensor{},
	}

	// アドレスが指定されていればそのアドレスに設定
	if len(deviceAddress) > 0 {
		d.deviceAddress = deviceAddress[0]
	}

	return d
}

// Init センサ初期化関数
func (d *Device) Init(mode ...bno055_opmode) bool {
	var _mode bno055_opmode

	// モードチェック
	if len(mode) > 0 {
		_mode = mode[0]
	} else {
		_mode = OPERATION_MODE_IMUPLUS
	}

	// 起動チェック
	// 起動まで850ms
	time.Sleep(1000 * time.Millisecond) // 1秒待ち

	if !d.isConnected() {
		// 再度接続
		time.Sleep(1000 * time.Millisecond)
		if !d.isConnected() {
			return false
		}
	}

	d.setmode(OPERATION_MODE_CONFIG)

	// リセット
	_ = d.writeRegister(d.deviceAddress, uint8(BNO055_SYS_TRIGGER_ADDR), []byte{0x20})
	time.Sleep(100 * time.Millisecond)

	// 接続確認
	for !d.isConnected() {
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(50 * time.Millisecond)

	_ = d.writeRegister(d.deviceAddress, uint8(BNO055_PWR_MODE_ADDR), []byte{byte(POWER_MODE_NORMAL)})
	time.Sleep(50 * time.Millisecond)

	_ = d.writeRegister(d.deviceAddress, uint8(BNO055_PAGE_ID_ADDR), []byte{0x00})

	_ = d.writeRegister(d.deviceAddress, uint8(BNO055_SYS_TRIGGER_ADDR), []byte{0x00})
	time.Sleep(50 * time.Millisecond)

	d.setmode(_mode)
	time.Sleep(50 * time.Millisecond)

	return true
}

// GetGyro ジャイロセンサデータ取得
func (d *Device) GetGyro() bool {
	data := make([]byte, 6)
	err := d.readRegister(d.deviceAddress, uint8(BNO055_GYRO_DATA_X_LSB_ADDR), data)
	if err != nil {
		return false
	}

	d.SensorData.GyroData.XData = uint16(data[1])<<8 | uint16(data[0])
	d.SensorData.GyroData.YData = uint16(data[3])<<8 | uint16(data[2])
	d.SensorData.GyroData.ZData = uint16(data[5])<<8 | uint16(data[4])

	return true
}

// GetAccl 加速度センサデータ取得
func (d *Device) GetAccl() bool {
	data := make([]byte, 6)

	err := d.readRegister(d.deviceAddress, uint8(BNO055_ACCEL_DATA_X_LSB_ADDR), data)
	if err != nil {
		return false
	}

	d.SensorData.AcclData.XData = uint16(data[1])<<8 | uint16(data[0])
	d.SensorData.AcclData.YData = uint16(data[3])<<8 | uint16(data[2])
	d.SensorData.AcclData.ZData = uint16(data[5])<<8 | uint16(data[4])

	return true
}

// GetAccl 地磁気センサデータ取得
func (d *Device) GetMag() bool {
	data := make([]byte, 6)

	err := d.readRegister(d.deviceAddress, uint8(BNO055_MAG_DATA_X_LSB_ADDR), data)
	if err != nil {
		return false
	}

	d.SensorData.MagData.XData = uint16(data[1])<<8 | uint16(data[0])
	d.SensorData.MagData.YData = uint16(data[3])<<8 | uint16(data[2])
	d.SensorData.MagData.ZData = uint16(data[5])<<8 | uint16(data[4])

	return true
}

// GetAccl オイラー角データ取得
func (d *Device) GetEuler() bool {
	data := make([]byte, 6)

	err := d.readRegister(d.deviceAddress, uint8(BNO055_EULER_H_LSB_ADDR), data)
	if err != nil {
		return false
	}

	d.SensorData.EulerData.HData = uint16(data[1])<<8 | uint16(data[0])
	d.SensorData.EulerData.RData = uint16(data[3])<<8 | uint16(data[2])
	d.SensorData.EulerData.PData = uint16(data[5])<<8 | uint16(data[4])

	return true
}

// GetQuaternion 四元数データ取得
func (d *Device) GetQuaternion() bool {
	data := make([]byte, 8)

	err := d.readRegister(d.deviceAddress, uint8(BNO055_QUATERNION_DATA_W_LSB_ADDR), data)
	if err != nil {
		return false
	}

	d.SensorData.QutaData.WData = uint16(data[1])<<8 | uint16(data[0])
	d.SensorData.QutaData.XData = uint16(data[3])<<8 | uint16(data[2])
	d.SensorData.QutaData.YData = uint16(data[5])<<8 | uint16(data[4])
	d.SensorData.QutaData.ZData = uint16(data[7])<<8 | uint16(data[6])

	return true
}

// QuaternionToEuler 四元数データからオイラー角を計算
func (d *Device) QuaternionToEuler() (chk bool, roll float64, pich float64, yaw float64) {
	var w, x, y, z float64
	chk = d.GetQuaternion()
	if !chk {
		return chk, 0.0, 0.0, 0.0
	} else {
		w = float64(int16(d.SensorData.QutaData.WData)) / float64(16384.0)
		x = float64(int16(d.SensorData.QutaData.XData)) / float64(16384.0)
		y = float64(int16(d.SensorData.QutaData.YData)) / float64(16384.0)
		z = float64(int16(d.SensorData.QutaData.ZData)) / float64(16384.0)

		// roll (x-axis rotation)
		ysqr := y * y
		t0 := +2.0 * (w*x + y*z)
		t1 := +1.0 - 2.0*(x*x+ysqr)
		roll = math.Atan2(t0, t1) * (180 / math.Pi)

		// pitch (y-axis rotation)
		t2 := +2.0 * (w*y - z*x)
		if t2 > 1.0 {
			t2 = 1.0
		} else if t2 < -1.0 {
			t2 = -1.0
		}
		pich = math.Asin(t2) * (180 / math.Pi)

		// yaw (z-axis rotation)
		t3 := +2.0 * (w*z + x*y)
		t4 := +1.0 - 2.0*(ysqr+z*z)
		yaw = math.Atan2(t3, t4) * (180 / math.Pi)
	}
	return chk, roll, pich, yaw
}

// PrintchipId チップID表示
func (d *Device) PrintchipId() {
	println(d.chipId)
}

// PrintchipId チップID表示
func (d *Device) GetTemp() (chk bool, temp int8) {
	data := make([]byte, 1)
	err := d.readRegister(d.deviceAddress, uint8(BNO055_TEMP_ADDR), data)
	if err != nil {
		return false, 0
	}
	temp = int8(data[0])
	return true, temp
}

// --------------------------------- 内部処理用
// isConnected センサとの接続を確認
func (d *Device) isConnected() bool {
	data := make([]byte, 1)
	time.Sleep(100 * time.Millisecond)
	err := d.readRegister(d.deviceAddress, uint8(BNO055_CHIP_ID_ADDR), data)
	if err != nil {
		return false
	}
	d.chipId = data[0]
	return data[0] == BNO055_ID
}

// setmode
func (d *Device) setmode(mode bno055_opmode) {
	var wData []byte
	wData = append(wData, byte(mode))
	_ = d.writeRegister(d.deviceAddress, uint8(BNO055_OPR_MODE_ADDR), wData)
	time.Sleep(100 * time.Millisecond)
}

// readRegister I2C用レジスタ読み込み関数
func (d *Device) readRegister(addr uint8, reg uint8, data []byte) error {
	return d.bus.Tx(uint16(addr), []byte{reg}, data)
}

// writeRegister I2C用レジスタ書き込み関数
func (d *Device) writeRegister(addr uint8, reg uint8, data []byte) error {
	buf := make([]uint8, len(data)+1)
	buf[0] = reg
	copy(buf[1:], data)
	return d.bus.Tx(uint16(addr), buf, nil)
}

// --------------------------------- 内部処理用
