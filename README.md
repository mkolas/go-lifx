# go-lifx
simple go lifx test

This project is a working example of changing a LIFX bulb's color via the REST API through a web interface without OAuth2. A sqlite database is attached to track entries for fun.

Possible Extensions
* Support for more than one bulb- currently it uses the API's "all" bulb selection parameter, but I only have one anyway.
* Allow for changing brightness, setting pulses,  etc.
* We rely on LIFX to rate limit us, which might result in being banned from using the API.
