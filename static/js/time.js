
function updateTime() {
    fetchJson("GET", "/api/time", function(data) {
        document.getElementById("time-time").innerText = data.time;
        document.getElementById("time-date").innerText = data.date;
    });
}

setInterval(updateTime, 59000);
