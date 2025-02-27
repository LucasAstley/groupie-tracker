const map = L.map('map').setView([20, 0], 1);

L.tileLayer('https://{s}.basemaps.cartocdn.com/dark_all/{z}/{x}/{y}{r}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors &copy; <a href="https://carto.com/attributions">CARTO</a>',
    maxZoom: 19,
    minZoom: 1.5,
    bounds: [[-90, -180], [90, 180]]
}).addTo(map);

map.setMaxBounds([[-85, -180], [85, 180]]);
const whiteIcon = L.icon({
    iconUrl: '/static/images/white-map-pin.png',
    iconSize: [40, 35],
    iconAnchor: [10, 30],
    popupAnchor: [2, -30],
});

const urlParams = new URLSearchParams(window.location.search);
const artistId = urlParams.get('id');

fetch(`/api/concerts${artistId ? `?artistId=${artistId}` : ''}`)
    /*
		This fetch request retrieves the concert data from the server and geo-codes the location of each concert
	 */
    .then(response => response.json())
    .then(data => {
        console.log(data);
        data.forEach(concert => {
            const [city, country] = concert.Place.split('-');
            const query = `${city}, ${country}`;
            const storedCoordinates = localStorage.getItem(query);

            if (storedCoordinates) {
                const [lat, lon] = JSON.parse(storedCoordinates);
                addMarker(lat, lon, query, concert.Dates);
            } else {
                fetch(`https://nominatim.openstreetmap.org/search?format=geojson&q=${encodeURIComponent(query)}`)
                    .then(response => response.json())
                    .then(geoData => {
                        if (geoData.features.length > 0) {
                            const {coordinates} = geoData.features[0].geometry;
                            const [lon, lat] = coordinates;
                            localStorage.setItem(query, JSON.stringify([lat, lon]));
                            addMarker(lat, lon, query, concert.Dates);
                        } else {
                            console.error(`No results found for location: ${query}`);
                        }
                    })
                    .catch(error => console.error('Error fetching geocode data:', error));
            }
        });
    })
    .catch(error => console.error('Error fetching concert data:', error));

function addMarker(lat, lon, query, dates) {
    /*
		This function adds a marker to the map at the given latitude and longitude.
	 */
    const marker = L.marker([lat, lon], {icon: whiteIcon}).addTo(map);
    const formattedDates = dates.join(', ');
    marker.bindPopup(`<b>${query.replace(/_/g, ' ').replace(/(\w)(\w*)/g, (_, first, rest) => first.toUpperCase() + rest.toLowerCase())}</b><br>${formattedDates.replace(/, /g, '<br>')}`);
}
