var request = new XMLHttpRequest();

request.open('GET', 'js/01_data.json');

request.addEventListener('readystatechange', function () {
    if ((request.status === 200) && (request.readyState === 4)) {
        var data = JSON.parse(request.responseText);
        //console.log(data.social[1].name);
        //console.log(data.social[1].twitter);
        //console.log(data.social[1].facebook);
        var template = document.querySelector('#myTemplate').innerHTML;
        var html = Mustache.to_html(template, data);
        document.querySelector('#socialInfo').innerHTML=html;
    }
});

request.send();