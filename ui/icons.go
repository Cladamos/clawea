package ui

// Codes are based on WMO weather codes

type DecodedWeather struct {
	Label string
	Icon  string
}

func WeatherCodeDecoder(code int, isNight bool) DecodedWeather {
	switch code {
	// Clear sky
	case 0:
		if isNight {
			return DecodedWeather{"Open Sky", ClearNight}
		}
		return DecodedWeather{"Sunny", Sunny}

	// Mainly clear, partly cloudy
	case 1, 2:
		if isNight {
			return DecodedWeather{"Partly Cloudy", PartlyCloudyNight}
		}
		return DecodedWeather{"Partly Cloudy", PartlyCloudy}

	// Overcast
	case 3:
		return DecodedWeather{"Cloudy", Cloudy}

	// Fog and depositing rime fog
	case 45, 48:
		return DecodedWeather{"Foggy", Foggy}

	// Drizzle (Light) and Rain (Slight/Moderate)
	case 51, 53, 55, 56, 57, 61, 63, 66, 67, 80, 81:
		return DecodedWeather{"Rainy", Rainy}

	// Rain (Heavy) and Rain Showers (Violent)
	case 65, 82:
		return DecodedWeather{"Heavy Rain", HeavyRain}

	// Snow (Slight/Moderate/Heavy) and Snow Showers
	case 71, 73, 75, 77, 85, 86:
		return DecodedWeather{"Snowing", Snowing}

	// Thunderstorm (Slight/Moderate)
	case 95, 96, 99:
		return DecodedWeather{"Thunderstorm", Thunderstorm}

	default:
		return DecodedWeather{"Unknown Weather", Cloudy}
	}
}

var (
	Cloudy = `
      .--.
   .-(    ).
  (___.__)__)
`

	Rainy = `
      .--.
   .-(    ).
  (___.__)__)` + blueStyle(`
    ' ' ' '
`)
	Snowing = `
      .--.
   .-(    ).
  (___.__)__)
    * * * * 
    * * * * 
`
	// PartlyCloudy = `
	//    \  /
	//  _ /"".-.
	//    \_(   ).
	//    /(___(__)
	//`

	PartlyCloudy = yellowStyle(`
	\  /`) + yellowStyle(`
  _ /""`) + `.-.` + yellowStyle(`
	\_`) + `(   ).` + yellowStyle(`
	/`) + `(___(__)`

	HeavyRain = `
      .--.
   .-(    ).
  (___.__)__)` + blueStyle(`
    / / / /
   / / / /
`)
	Thunderstorm = `
      .--.
   .-(    ).
  (___.__)__)` + yellowStyle(`
    /_   /_
     /    /
`)
	Sunny = yellowStyle(`
    \   /
     .-.
  - (   ) -
     '-'
    /   \
`)
	Foggy = `
  ~   ~   ~   
    ~   ~   ~   
  ~   ~   ~    
    ~   ~   ~    
`
	// PartlyCloudyNight = `
	//    .  .
	//  . /"".-.
	//  . \_(   ).
	//    .(___(__)
	//`

	PartlyCloudyNight = lightBlueStyle(`
    .  .`) + lightBlueStyle(`
  . /""`) + `.-.` + lightBlueStyle(`
  . \_`) + `(   ).` + lightBlueStyle(`
	.`) + `(___(__)`

	ClearNight = lightBlueStyle(`
    .   .
     .-.
  . (   ) .
     '-'
    .   .
   `)
)
