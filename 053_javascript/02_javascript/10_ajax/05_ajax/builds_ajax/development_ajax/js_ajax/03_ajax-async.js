var request = new XMLHttpRequest();
request.open('GET', 'data.txt');
request.addEventListener('readystatechange', function() {
    if ((request.status === 200) && (request.readyState === 4)) {
        myNode = document.querySelector('body');
        myNode.innerHTML = request.responseText;
        console.log('new way');
    } //ready
});

request.send();