# Opensky Flight Tracker
A web service to track fligths

## Specification:
1. Show all planes currently tracked by https://opensky-network.org/ and all airports we got in airports.json
2. Show single airplane with departure and arrival airport if available
3. Show all airplanes from a country
4. A map to display all planes and airports
5. Click on a plane/airport to get more info

## Went well
* No problems setting up the database and connecting to it
* Map not lagging
* Planes rotate
* Openstack updates every 15 min

## Went less well
* Not all airports were added
* Not layers in map (Airports can appear over planes)
* Not every plane has assigned airports 
* 360 pictures to make the planes rotate

## Hard aspects ###### ~~That's what she said~~
1. We had some problems making an elegant solution so that the planes rotate, google maps does not let us rotate images in a easy way 
unless they are svg which makes the map run at 2 frames per minute for a whole minute before the browser crashes. This is only a 
problem when rendering 1000+ object on the map. Also we noticed a tiny bit to late that the heading of planes are given with decimal points, so when trying to get all the images it started to lag as it was trying to fetch images that didn't exist. Easy solution was to 
just floor the heading.

## What we learned
One thing we all learned was the value of communication in a team, there where some moments when we were unsure who was doing what
so that lead to some confusion, we fixed this by having meetings often to update eachother on problems and tried to solve them togheter.

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




# Time
~ 88 hours
