var searchBox = document.querySelector('#searchBox');
var template = document.querySelector('#speakerstmpl').innerHTML;
var results = document.querySelector('#speakers');
var allData;

function onKeyUp() {
    var modifiedData = allData;
    var regex = new RegExp(searchBox.value.toLowerCase());
    for (var i = 0; i < modifiedData.speakers.length; i++) {
        var obj = modifiedData.speakers[i];
        console.log('name: ', obj.name);
        if (obj.name.toLowerCase().search(regex) === -1) {
            // no match, don't include
            modifiedData.speakers.splice(i, 1);
        }
    }
    var html = Mustache.to_html(template, modifiedData);
    results.innerHTML = html;
}

function getData() {
    var request = new XMLHttpRequest();
    request.open('GET', 'js/01_data.json');
    request.addEventListener('readystatechange', function () {
        if ((request.status === 200) && (request.readyState === 4)) {
            allData = JSON.parse(request.responseText);
            console.log(allData);
            console.log(allData.speakers.length);
            console.log(allData.speakers[0].name);
            var html = Mustache.to_html(template, allData);
            results.innerHTML = html;
        }
    })
    request.send();
}

function onLoad() {
    getData();
    searchBox.addEventListener('keyup', onKeyUp);
}

window.addEventListener('load', onLoad);