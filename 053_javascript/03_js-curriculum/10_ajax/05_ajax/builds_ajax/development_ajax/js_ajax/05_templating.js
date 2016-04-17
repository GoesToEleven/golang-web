"use strict";

var request;

request = new XMLHttpRequest();

request.open('GET', 'js/data.json');

request.addEventListener('onreadystatechange', function () {
    if ((request.status === 200) && (request.readyState === 4)) {
        var data = JSON.parse(request.responseText);
        var template = document.querySelector('#speakerstpl').innerHTML;
        var html = Mustache.to_html(template, data);
        document.querySelector('#speakers').innerHTML = html;
    } //ready
})

request.send();