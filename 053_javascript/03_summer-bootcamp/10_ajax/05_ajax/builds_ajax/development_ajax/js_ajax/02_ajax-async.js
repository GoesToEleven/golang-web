var request = new XMLHttpRequest();
request.open('GET', 'data.txt');
request.onreadystatechange = function() {
    if ((request.status === 200) &&
        (request.readyState === 4)) {
        myNode = document.querySelector('body');
        myNode.innerHTML = request.responseText;
    } //ready
} //event
request.send();