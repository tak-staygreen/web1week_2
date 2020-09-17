var request = new XMLHttpRequest();
request.open("GET", "http://localhost:9000/getcat", true);
request.responseType = "json";
request.send();

var data = "";
request.onload = function() {
    data = JSON.stringify(this.response);
    changeBackground();
    updateClock();
};

window.addEventListener("load", function() {
    setInterval( function() {
        updateClock();
    }, 1000);
}, false)

function changeBackground() {
    var num = Math.round(Math.random()*29);
    var urlList = data.split(",");
    var url = urlList[num].substring(2, urlList[num].length-2);
    if (url.split(".").pop() === "jpg") {
        document.body.style.backgroundImage="url(" + url + ")";
    } else {
        changeBackground();
    }
}

function updateClock() {
    var element = document.getElementById("clock");
    var date = new Date();
    var dispDate = date.getFullYear() + "/" + date.getMonth() + "/" + date.getDate() + " " + date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();
    element.innerHTML = dispDate;
    if (date.getSeconds() % 20 == 0) {
        changeBackground()
    }
}