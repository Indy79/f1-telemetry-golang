package car

type CarStatusData struct {
	M_tractionControl   uint8   // 0 (off) - 2 (high)
	M_antiLockBrakes    uint8   // 0 (off) - 1 (on)
	M_fuelMix           uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	M_frontBrakeBias    uint8   // Front brake bias (percentage)
	M_pitLimiterStatus  uint8   // Pit limiter status - 0 = off, 1 = on
	M_fuelInTank        float32 // Current fuel mass
	M_fuelCapacity      float32 // Fuel capacity
	M_fuelRemainingLaps float32 // Fuel remaining in terms of laps (value on MFD)
	M_maxRPM            uint16  // Cars max RPM, point of rev limiter
	M_idleRPM           uint16  // Cars idle RPM
	M_maxGears          uint8   // Maximum number of gears
	M_drsAllowed        uint8   // 0 = not allowed, 1 = allowed, -1 = unknown

	// Added in Beta3:
	M_drsActivationDistance uint16 // 0 = DRS not available, non-zero - DRS will be available
	// in [X] metres

	M_tyresWear          [4]uint8 // Tyre wear percentage
	M_actualTyreCompound uint8    // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1
	// 7 = inter, 8 = wet
	// F1 Classic - 9 = dry, 10 = wet
	// F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard
	// 15 = wet
	M_visualTyreCompound uint8 // F1 visual (can be different from actual compound)
	// 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet
	// F1 Classic – same as above
	// F2 – same as above
	M_tyresAgeLaps         uint8    // Age in laps of the current set of tyres
	M_tyresDamage          [4]uint8 // Tyre damage (percentage)
	M_frontLeftWingDamage  uint8    // Front left wing damage (percentage)
	M_frontRightWingDamage uint8    // Front right wing damage (percentage)
	M_rearWingDamage       uint8    // Rear wing damage (percentage)

	// Added Beta 3:
	M_drsFault uint8 // Indicator for DRS fault, 0 = OK, 1 = fault

	M_engineDamage    uint8 // Engine damage (percentage)
	M_gearBoxDamage   uint8 // Gear box damage (percentage)
	M_vehicleFiaFlags int8  // -1 = invalid/unknown, 0 = none, 1 = green
	// 2 = blue, 3 = yellow, 4 = red
	M_ersStoreEnergy float32 // ERS energy store in Joules
	M_ersDeployMode  uint8   // ERS deployment mode, 0 = none, 1 = medium
	// 2 = overtake, 3 = hotlap
	M_ersHarvestedThisLapMGUK float32 // ERS energy harvested this lap by MGU-K
	M_ersHarvestedThisLapMGUH float32 // ERS energy harvested this lap by MGU-H
	M_ersDeployedThisLap      float32 // ERS energy deployed this lap
}

type PacketCarStatusData struct {
	M_carStatusData [22]CarStatusData
}
