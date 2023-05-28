
function updateSysstat() {
    var elements = document.querySelectorAll(".sysstat");
    for (var i = 0; i < elements.length; i++) {
        document.getElementById(elements[i].id + "-progress").style.width = elements[i].dataset.value + "%";
        elements[i].innerText = elements[i].dataset.value;
    }
}

updateSysstat();

setInterval(function() {
    fetchJson("GET", "/api/sysstat", function(data) {
        Object.keys(data).forEach(function(key) {
            document.getElementById("sysstat-" + key).dataset.value = data[key];
        });
        setTimeout(updateSysstat, 200);
    });
}, 15000);
