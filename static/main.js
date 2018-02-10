var map = L.map('map').setView([47.1601278, 4.9538808], 5);

L.tileLayer(
    'http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: 'Map data &copy; <a href="http://openstreetmap.org">OpenStreetMap</a>',
        maxZoom: 18
    }
).addTo(map);

var popup = L.popup();

function onMapClick(e) {
    popup.setLatLng(e.latlng)
        .setContent('<a href="/add/'+e.latlng.lat+','+e.latlng.lng+'">追加</a>')
        .openOn(map);
}
map.on('click', onMapClick);
