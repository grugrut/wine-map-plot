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
        .setContent('<a href="/add?lat='+e.latlng.lat+'&lng='+e.latlng.lng+'">追加</a>')
        .openOn(map);
}
map.on('click', onMapClick);

var req = new XMLHttpRequest();
req.onreadystatechange = function() {
    if (req.readyState == 4 && req.status == 200) {
        var wineries = JSON.parse(req.responseText);
        for (var winery of wineries) {
            var marker = L.marker([winery['Latitude'], winery['Longitude']]).addTo(map);
            marker.bindPopup(winery['Name']+'<br>'+winery['NameJa']);
        }
    }
};
req.open('GET', 'http://localhost:8080/api/fetch');
req.send(null);
