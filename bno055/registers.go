package bno055

// BNO055_ADDRESS_A
const BNO055_Address_A = 0x28

// bno055 Address B
const BNO055_Address_B = 0x29

// bno055 ID
const BNO055_ID = 0xa0

/** Offsets registers **/
const NUM_BNO055_Offset_Registers = 22

/** A structure to represent offsets **/
type bno055offsets struct {
	accel_offset_x int16 /**< x acceleration offset */
	accel_offset_y int16 /**< y acceleration offset */
	accel_offset_z int16 /**< z acceleration offset */

	mag_offset_x int16 /**< x magnetometer offset */
	mag_offset_y int16 /**< y magnetometer offset */
	mag_offset_z int16 /**< z magnetometer offset */

	gyro_offset_x int16 /**< x gyroscrope offset */
	gyro_offset_y int16 /**< y gyroscrope offset */
	gyro_offset_z int16 /**< z gyroscrope offset */

	accel_radius int16 /**< acceleration radius */

	mag_radius int16 /**< magnetometer radius */
}

type bno055_opmode uint8

/** Operation mode settings **/
const (
	OPERATION_MODE_CONFIG       bno055_opmode = 0x00
	OPERATION_MODE_ACCONLY      bno055_opmode = 0x01
	OPERATION_MODE_MAGONLY      bno055_opmode = 0x02
	OPERATION_MODE_GYRONLY      bno055_opmode = 0x03
	OPERATION_MODE_ACCMAG       bno055_opmode = 0x04
	OPERATION_MODE_ACCGYRO      bno055_opmode = 0x05
	OPERATION_MODE_MAGGYRO      bno055_opmode = 0x06
	OPERATION_MODE_AMG          bno055_opmode = 0x07
	OPERATION_MODE_IMUPLUS      bno055_opmode = 0x08
	OPERATION_MODE_COMPASS      bno055_opmode = 0x09
	OPERATION_MODE_M4G          bno055_opmode = 0x0a
	OPERATION_MODE_NDOF_FMC_OFF bno055_opmode = 0x0b
	OPERATION_MODE_NDOF         bno055_opmode = 0x0c
)

type bno055reg uint8

/* PAGE0 REGISTER DEFINITION START*/
const (

	/* Page id register definition */
	BNO055_PAGE_ID_ADDR bno055reg = 0x07

	/* PAGE0 REGISTER DEFINITION START*/
	BNO055_CHIP_ID_ADDR       bno055reg = 0x00
	BNO055_ACCEL_REV_ID_ADDR  bno055reg = 0x01
	BNO055_MAG_REV_ID_ADDR    bno055reg = 0x02
	BNO055_GYRO_REV_ID_ADDR   bno055reg = 0x03
	BNO055_SW_REV_ID_LSB_ADDR bno055reg = 0x04
	BNO055_SW_REV_ID_MSB_ADDR bno055reg = 0x05
	BNO055_BL_REV_ID_ADDR     bno055reg = 0x06

	/* Accel data register */
	BNO055_ACCEL_DATA_X_LSB_ADDR bno055reg = 0x08
	BNO055_ACCEL_DATA_X_MSB_ADDR bno055reg = 0x09
	BNO055_ACCEL_DATA_Y_LSB_ADDR bno055reg = 0x0A
	BNO055_ACCEL_DATA_Y_MSB_ADDR bno055reg = 0x0B
	BNO055_ACCEL_DATA_Z_LSB_ADDR bno055reg = 0x0C
	BNO055_ACCEL_DATA_Z_MSB_ADDR bno055reg = 0x0D

	/* Mag data register */
	BNO055_MAG_DATA_X_LSB_ADDR bno055reg = 0x0E
	BNO055_MAG_DATA_X_MSB_ADDR bno055reg = 0x0F
	BNO055_MAG_DATA_Y_LSB_ADDR bno055reg = 0x10
	BNO055_MAG_DATA_Y_MSB_ADDR bno055reg = 0x11
	BNO055_MAG_DATA_Z_LSB_ADDR bno055reg = 0x12
	BNO055_MAG_DATA_Z_MSB_ADDR bno055reg = 0x13

	/* Gyro data registers */
	BNO055_GYRO_DATA_X_LSB_ADDR bno055reg = 0x14
	BNO055_GYRO_DATA_X_MSB_ADDR bno055reg = 0x15
	BNO055_GYRO_DATA_Y_LSB_ADDR bno055reg = 0x16
	BNO055_GYRO_DATA_Y_MSB_ADDR bno055reg = 0x17
	BNO055_GYRO_DATA_Z_LSB_ADDR bno055reg = 0x18
	BNO055_GYRO_DATA_Z_MSB_ADDR bno055reg = 0x19

	/* Euler data registers */
	BNO055_EULER_H_LSB_ADDR bno055reg = 0x1A
	BNO055_EULER_H_MSB_ADDR bno055reg = 0x1B
	BNO055_EULER_R_LSB_ADDR bno055reg = 0x1C
	BNO055_EULER_R_MSB_ADDR bno055reg = 0x1D
	BNO055_EULER_P_LSB_ADDR bno055reg = 0x1E
	BNO055_EULER_P_MSB_ADDR bno055reg = 0x1F

	/* Quaternion data registers */
	BNO055_QUATERNION_DATA_W_LSB_ADDR bno055reg = 0x20
	BNO055_QUATERNION_DATA_W_MSB_ADDR bno055reg = 0x21
	BNO055_QUATERNION_DATA_X_LSB_ADDR bno055reg = 0x22
	BNO055_QUATERNION_DATA_X_MSB_ADDR bno055reg = 0x23
	BNO055_QUATERNION_DATA_Y_LSB_ADDR bno055reg = 0x24
	BNO055_QUATERNION_DATA_Y_MSB_ADDR bno055reg = 0x25
	BNO055_QUATERNION_DATA_Z_LSB_ADDR bno055reg = 0x26
	BNO055_QUATERNION_DATA_Z_MSB_ADDR bno055reg = 0x27

	/* Linear acceleration data registers */
	BNO055_LINEAR_ACCEL_DATA_X_LSB_ADDR bno055reg = 0x28
	BNO055_LINEAR_ACCEL_DATA_X_MSB_ADDR bno055reg = 0x29
	BNO055_LINEAR_ACCEL_DATA_Y_LSB_ADDR bno055reg = 0x2A
	BNO055_LINEAR_ACCEL_DATA_Y_MSB_ADDR bno055reg = 0x2B
	BNO055_LINEAR_ACCEL_DATA_Z_LSB_ADDR bno055reg = 0x2C
	BNO055_LINEAR_ACCEL_DATA_Z_MSB_ADDR bno055reg = 0x2D

	/* Gravity data registers */
	BNO055_GRAVITY_DATA_X_LSB_ADDR bno055reg = 0x2E
	BNO055_GRAVITY_DATA_X_MSB_ADDR bno055reg = 0x2F
	BNO055_GRAVITY_DATA_Y_LSB_ADDR bno055reg = 0x30
	BNO055_GRAVITY_DATA_Y_MSB_ADDR bno055reg = 0x31
	BNO055_GRAVITY_DATA_Z_LSB_ADDR bno055reg = 0x32
	BNO055_GRAVITY_DATA_Z_MSB_ADDR bno055reg = 0x33

	/* Temperature data register */
	BNO055_TEMP_ADDR bno055reg = 0x34

	/* Status registers */
	BNO055_CALIB_STAT_ADDR      bno055reg = 0x35
	BNO055_SELFTEST_RESULT_ADDR bno055reg = 0x36
	BNO055_INTR_STAT_ADDR       bno055reg = 0x37

	BNO055_SYS_CLK_STAT_ADDR bno055reg = 0x38
	BNO055_SYS_STAT_ADDR     bno055reg = 0x39
	BNO055_SYS_ERR_ADDR      bno055reg = 0x3A

	/* Unit selection register */
	BNO055_UNIT_SEL_ADDR bno055reg = 0x3B

	/* Mode registers */
	BNO055_OPR_MODE_ADDR bno055reg = 0x3D
	BNO055_PWR_MODE_ADDR bno055reg = 0x3E

	BNO055_SYS_TRIGGER_ADDR bno055reg = 0x3F
	BNO055_TEMP_SOURCE_ADDR bno055reg = 0x40

	/* Axis remap registers */
	BNO055_AXIS_MAP_CONFIG_ADDR bno055reg = 0x41
	BNO055_AXIS_MAP_SIGN_ADDR   bno055reg = 0x42

	/* SIC registers */
	BNO055_SIC_MATRIX_0_LSB_ADDR bno055reg = 0x43
	BNO055_SIC_MATRIX_0_MSB_ADDR bno055reg = 0x44
	BNO055_SIC_MATRIX_1_LSB_ADDR bno055reg = 0x45
	BNO055_SIC_MATRIX_1_MSB_ADDR bno055reg = 0x46
	BNO055_SIC_MATRIX_2_LSB_ADDR bno055reg = 0x47
	BNO055_SIC_MATRIX_2_MSB_ADDR bno055reg = 0x48
	BNO055_SIC_MATRIX_3_LSB_ADDR bno055reg = 0x49
	BNO055_SIC_MATRIX_3_MSB_ADDR bno055reg = 0x4A
	BNO055_SIC_MATRIX_4_LSB_ADDR bno055reg = 0x4B
	BNO055_SIC_MATRIX_4_MSB_ADDR bno055reg = 0x4C
	BNO055_SIC_MATRIX_5_LSB_ADDR bno055reg = 0x4D
	BNO055_SIC_MATRIX_5_MSB_ADDR bno055reg = 0x4E
	BNO055_SIC_MATRIX_6_LSB_ADDR bno055reg = 0x4F
	BNO055_SIC_MATRIX_6_MSB_ADDR bno055reg = 0x50
	BNO055_SIC_MATRIX_7_LSB_ADDR bno055reg = 0x51
	BNO055_SIC_MATRIX_7_MSB_ADDR bno055reg = 0x52
	BNO055_SIC_MATRIX_8_LSB_ADDR bno055reg = 0x53
	BNO055_SIC_MATRIX_8_MSB_ADDR bno055reg = 0x54

	/* Accelerometer Offset registers */
	ACCEL_OFFSET_X_LSB_ADDR bno055reg = 0x55
	ACCEL_OFFSET_X_MSB_ADDR bno055reg = 0x56
	ACCEL_OFFSET_Y_LSB_ADDR bno055reg = 0x57
	ACCEL_OFFSET_Y_MSB_ADDR bno055reg = 0x58
	ACCEL_OFFSET_Z_LSB_ADDR bno055reg = 0x59
	ACCEL_OFFSET_Z_MSB_ADDR bno055reg = 0x5A

	/* Magnetometer Offset registers */
	MAG_OFFSET_X_LSB_ADDR bno055reg = 0x5B
	MAG_OFFSET_X_MSB_ADDR bno055reg = 0x5C
	MAG_OFFSET_Y_LSB_ADDR bno055reg = 0x5D
	MAG_OFFSET_Y_MSB_ADDR bno055reg = 0x5E
	MAG_OFFSET_Z_LSB_ADDR bno055reg = 0x5F
	MAG_OFFSET_Z_MSB_ADDR bno055reg = 0x60

	/* Gyroscope Offset register s*/
	GYRO_OFFSET_X_LSB_ADDR bno055reg = 0x61
	GYRO_OFFSET_X_MSB_ADDR bno055reg = 0x62
	GYRO_OFFSET_Y_LSB_ADDR bno055reg = 0x63
	GYRO_OFFSET_Y_MSB_ADDR bno055reg = 0x64
	GYRO_OFFSET_Z_LSB_ADDR bno055reg = 0x65
	GYRO_OFFSET_Z_MSB_ADDR bno055reg = 0x66

	/* Radius registers */
	ACCEL_RADIUS_LSB_ADDR bno055reg = 0x67
	ACCEL_RADIUS_MSB_ADDR bno055reg = 0x68
	MAG_RADIUS_LSB_ADDR   bno055reg = 0x69
	MAG_RADIUS_MSB_ADDR   bno055reg = 0x6A
)

type bno055_powermode int

/** bno055 power settings */
const (
	POWER_MODE_NORMAL   bno055_powermode = 0x00
	POWER_MODE_LOWPOWER bno055_powermode = 0x01
	POWER_MODE_SUSPEND  bno055_powermode = 0x02
)

type bno055_axis_remap_config int

/** Remap settings **/
const (
	REMAP_CONFIG_P0 bno055_axis_remap_config = 0x21
	REMAP_CONFIG_P1 bno055_axis_remap_config = 0x24 // default
	REMAP_CONFIG_P2 bno055_axis_remap_config = 0x24
	REMAP_CONFIG_P3 bno055_axis_remap_config = 0x21
	REMAP_CONFIG_P4 bno055_axis_remap_config = 0x24
	REMAP_CONFIG_P5 bno055_axis_remap_config = 0x21
	REMAP_CONFIG_P6 bno055_axis_remap_config = 0x21
	REMAP_CONFIG_P7 bno055_axis_remap_config = 0x24
)

type bno055_axis_remap_sign int

/** Remap Signs **/
const (
	REMAP_SIGN_P0 bno055_axis_remap_sign = 0x04
	REMAP_SIGN_P1 bno055_axis_remap_sign = 0x00 // default
	REMAP_SIGN_P2 bno055_axis_remap_sign = 0x06
	REMAP_SIGN_P3 bno055_axis_remap_sign = 0x02
	REMAP_SIGN_P4 bno055_axis_remap_sign = 0x03
	REMAP_SIGN_P5 bno055_axis_remap_sign = 0x01
	REMAP_SIGN_P6 bno055_axis_remap_sign = 0x07
	REMAP_SIGN_P7 bno055_axis_remap_sign = 0x05
)

/** A structure to represent revisions **/
type BNO055_REV_INFO struct {
	accel_rev uint8  /**< acceleration rev */
	mag_rev   uint8  /**< magnetometer rev */
	gyro_rev  uint8  /**< gyroscrope rev */
	ww_rev    uint16 /**< SW rev */
	bl_rev    uint8  /**< bootloader rev */
}

/** Vector Mappings **/
const (
	VECTOR_ACCELEROMETER = BNO055_ACCEL_DATA_X_LSB_ADDR
	VECTOR_MAGNETOMETER  = BNO055_MAG_DATA_X_LSB_ADDR
	VECTOR_GYROSCOPE     = BNO055_GYRO_DATA_X_LSB_ADDR
	VECTOR_EULER         = BNO055_EULER_H_LSB_ADDR
	VECTOR_LINEARACCEL   = BNO055_LINEAR_ACCEL_DATA_X_LSB_ADDR
	VECTOR_GRAVITY       = BNO055_GRAVITY_DATA_X_LSB_ADDR
)
