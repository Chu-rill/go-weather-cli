# go-weather-cli

This is a simple CLI application that fetches and displays weather information for a specified location using the WeatherAPI.

## Prerequisites

- Go 1.24 or later
- An internet connection to fetch weather data from the WeatherAPI

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Chu-rill/go-weather-cli.git
cd go-weather-cli
```

2. Install dependencies:

```bash
go mod tidy
```

## Building

To build the project, run the following command:

```bash
go build -o go-weather-cli main.go
```

This will create an executable named go-weather-cli in the current directory.

## Running

To run the application, use the following command:

```bash
./go-weather-cli [location]
```

Replace [location] with the name of the location you want to fetch the weather for. If no location is provided, it defaults to "Abuja".

Example:

```bash
./go-weather-cli London
```

## Usage

The application fetches weather data from the WeatherAPI and displays the current weather and hourly forecast for the specified location. The output includes temperature, chance of rain, and weather conditions. The output is colorized based on the chance of rain.

## Example Output

```
📍 Abuja, Nigeria: 🌡️ 25.0°C, ☁️ Sunny
📅 Hourly Forecast:
🕒 07:00 - 🌡️ 25.0°C, ☔ 0%, ☁️ Sunny
🕒 08:00 - 🌡️ 26.9°C, ☔ 0%, ☁️ Sunny
🕒 09:00 - 🌡️ 29.5°C, ☔ 0%, ☁️ Sunny
🕒 10:00 - 🌡️ 31.8°C, ☔ 0%, ☁️ Sunny
🕒 11:00 - 🌡️ 34.4°C, ☔ 0%, ☁️ Sunny
🕒 12:00 - 🌡️ 36.6°C, ☔ 0%, ☁️ Sunny
🕒 13:00 - 🌡️ 38.3°C, ☔ 0%, ☁️ Sunny
🕒 14:00 - 🌡️ 39.3°C, ☔ 0%, ☁️ Sunny
🕒 15:00 - 🌡️ 39.7°C, ☔ 0%, ☁️ Sunny
🕒 16:00 - 🌡️ 39.6°C, ☔ 0%, ☁️ Sunny
🕒 17:00 - 🌡️ 39.1°C, ☔ 0%, ☁️ Sunny
🕒 18:00 - 🌡️ 37.8°C, ☔ 0%, ☁️ Sunny
🕒 19:00 - 🌡️ 34.1°C, ☔ 0%, ☁️ Clear
🕒 20:00 - 🌡️ 32.7°C, ☔ 0%, ☁️ Clear
🕒 21:00 - 🌡️ 31.9°C, ☔ 0%, ☁️ Clear
🕒 22:00 - 🌡️ 31.2°C, ☔ 0%, ☁️ Clear
🕒 23:00 - 🌡️ 30.7°C, ☔ 0%, ☁️ Clear
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## Acknowledgements

- **WeatherAPI** for providing the weather data.
- **Fatih Color** for colorizing the output.

## Contact

- Twiiter: chu_rill
- Email: churchilldaniel687@gmail.com
