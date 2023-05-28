
fetchJson("GET", "/api/day-counter", function(data) {
    Object.keys(data).forEach(function(key) {
        document.getElementById(key).innerText = data[key];
    });
});
