var searchBox = document.querySelector('#searchBox');

function onKeyUp() {
    var regex = new RegExp(searchBox.value.toLowerCase());
    var names = document.querySelectorAll('.names');
    for (var i = 0; i < names.length; i++) {
        var obj = names[i];
        if (obj.innerHTML.toLowerCase().search(regex) === -1) {
            obj.parentNode.style.display = 'none';
        } else {
            obj.parentNode.style.display = 'block';
        }

    }
}

function getData() {
    var request = new XMLHttpRequest();
    request.open('GET', 'js/01_data.json');
    request.addEventListener('readystatechange', function () {
        if ((request.status === 200) && (request.readyState === 4)) {
            var data = JSON.parse(request.responseText);
            var template = document.querySelector('#speakerstmpl').innerHTML;
            var html = Mustache.to_html(template, data);
            document.querySelector('#speakers').innerHTML = html;
        }
    })
    request.send();
}

function onLoad() {
    getData();
    searchBox.addEventListener('keyup', onKeyUp);
}

window.addEventListener('load', onLoad);