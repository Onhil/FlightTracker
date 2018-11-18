# Opensky Flight Tracker
A web service to track fligths

# Map
* `GET /flight-tracker` Map with all the planes and airports. 
* `GET /flight-tracker/map/country/<country>` Shows all planes from the given country.
* `GET /flight-tracker/map/plane/<icao24>` Shows only the given plane on a google map.

# Plane
* `GET /flight-tracker/plane` List with all the planes in the database.
* `GET /flight-tracker/country/<country>` List of all planes with `<country>` as their origin country.
* `GET /flight-tracker/plane/<icao24>` Displays info about the plane with `<icao>` as their ICAO.
* `GET /flight-tracker/plane/<icao24>/<field>` Displays just the information about the given `<field>`.

# Airport
* `GET /flight-tracker/airport` Lists all airports.
* `GET /flight-tracker/airport/<icao>` Displays info about the given airport.
* `GET /flight-tracker/airport/<icao>/field>` Displays info about the <field> from the airport with the given <icao>.
* `GET /flight-tracker/airport/country` Lists all countries with an airport.
* `GET /flight-tracker/airport/country/<country>` Lists all airports in a countries.

Demo of app
* `https://plane-tracker-assi3.herokuapp.com/flight-tracker`
