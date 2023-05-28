
function log(message) {
    document.getElementById("log").innerHTML += "<p>" + message + "</p>";
}

function fetchJson(method, url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open(method, url);
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var resp = JSON.parse(xhr.responseText);
            callback(resp);
        }
    };
    xhr.send();
}
