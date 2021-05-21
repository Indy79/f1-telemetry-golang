package packets

type MarshalZone struct
{
	M_zoneStart float32   // Fraction (0..1) of way through the lap the marshal zone starts
	M_zoneFlag  int8   // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

type WeatherForecastSample struct
{
 	M_sessionType      uint8                     // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1
												// 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2
												// 12 = Time Trial
	M_timeOffset  	   uint8                    // Time in minutes the forecast is for
	M_weather     	   uint8                    // Weather - 0 = clear, 1 = light cloud, 2 = overcast
												// 3 = light rain, 4 = heavy rain, 5 = storm
	M_trackTemperature int8               // Track temp. in degrees celsius
	M_airTemperature   int8               // Air temp. in degrees celsius
};

type PacketSessionData struct
{


	M_weather uint8                   // Weather - 0 = clear, 1 = light cloud, 2 = overcast
										// 3 = light rain, 4 = heavy rain, 5 = storm
	M_trackTemperature int8         // Track temp. in degrees celsius
	M_airTemperature  int8         // Air temp. in degrees celsius
	M_totalLaps     uint8            // Total number of laps in this race
	M_trackLength   uint16           // Track length in metres
	M_sessionType   uint8            // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P
									// 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ
									// 10 = R, 11 = R2, 12 = Time Trial
	M_trackId             int8             // -1 for unknown, 0-21 for tracks, see appendix
	M_formula             uint8            // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2,
					 					 // 3 = F1 Generic
	M_sessionTimeLeft     uint16           // Time left in session in seconds
	M_sessionDuration     uint16           // Session duration in seconds
	M_pitSpeedLimit       uint8            // Pit speed limit in kilometres per hour
	M_gamePaused          uint8            // Whether the game is paused
	M_isSpectating        uint8            // Whether the player is spectating
	M_spectatorCarIndex   uint8            // Index of the car being spectated
	M_sliProNativeSupport uint8      	     // SLI Pro support, 0 = inactive, 1 = active
	M_numMarshalZones     uint8            // Number of marshal zones to follow
	M_marshalZones[21]    MarshalZone      // List of marshal zones – max 21
	M_safetyCarStatus     uint8            // 0 = no safety car, 1 = full safety car
											// 2 = virtual safety car
	M_networkGame                uint8 // 0 = offline, 1 = online
	M_numWeatherForecastSamples  uint8 // Number of weather samples to follow
	M_weatherForecastSamples[20] WeatherForecastSample   // Array of weather forecast samples
};
