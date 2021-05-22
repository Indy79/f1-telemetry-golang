package car

type CarTelemetryData struct {
	M_speed                   uint16     // Speed of car in kilometres per hour
	M_throttle                float32    // Amount of throttle applied (0.0 to 1.0)
	M_steer                   float32    // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	M_brake                   float32    // Amount of brake applied (0.0 to 1.0)
	M_clutch                  uint8      // Amount of clutch applied (0 to 100)
	M_gear                    int8       // Gear selected (1-8, N=0, R=-1)
	M_engineRPM               uint16     // Engine RPM
	M_drs                     uint8      // 0 = off, 1 = on
	M_revLightsPercent        uint8      // Rev lights indicator (percentage)
	M_brakesTemperature       [4]uint16  // Brakes temperature (celsius)
	M_tyresSurfaceTemperature [4]uint8   // Tyres surface temperature (celsius)
	M_tyresInnerTemperature   [4]uint8   // Tyres inner temperature (celsius)
	M_engineTemperature       uint16     // Engine temperature (celsius)
	M_tyresPressure           [4]float32 // Tyres pressure (PSI)
	M_surfaceType             [4]uint8   // Driving surface, see appendices
}

type PacketCarTelemetryData struct {
	M_carTelemetryData [22]CarTelemetryData

	M_buttonStatus uint32 // Bit flags specifying which buttons are being pressed
	// currently - see appendices

	// Added in Beta 3:
	M_mfdPanelIndex uint8 // Index of MFD panel open - 255 = MFD closed
	// Single player, race – 0 = Car setup, 1 = Pits
	// 2 = Damage, 3 =  Engine, 4 = Temperatures
	// May vary depending on game mode
	M_mfdPanelIndexSecondaryPlayer uint8 // See above
	M_suggestedGear                int8  // Suggested gear for the player (1-8)
	// 0 if no gear suggested
}
