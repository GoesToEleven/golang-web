var request = new XMLHttpRequest();

request.open('GET', 'js/04_data.json');
//request.open('GET', 'js/04_data_two.json');

request.addEventListener('readystatechange', function () {
    if ((request.status === 200) && (request.readyState === 4)) {
        var items = JSON.parse(request.responseText);
        console.log(items.speakers);
        console.log(items.speakers.length);
        for (var i in items.speakers) {
            console.log(items.speakers[i].name);
            console.log(items.speakers[i].name);
        }
    }
})

request.send();