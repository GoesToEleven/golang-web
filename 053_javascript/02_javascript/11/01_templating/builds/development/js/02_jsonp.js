// this does not work
// the whole trick of jsonp is to not to have to use ajax

var request = new XMLHttpRequest();

request.open('GET', 'http://api.flickr.com/services/feeds/photos_public.gne?id=73845487@N00&format=json&tags=portfolio');

request.addEventListener('readystatechange', function() {
    if((request.state === 200) && (request.readyState === 4)) {
        var data = JSON.parse(request.responseText);
        console.log(data);
    }
});

request.send();