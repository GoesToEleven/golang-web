var myNode = document.querySelector('body');

// ajax
var request = new XMLHttpRequest();
request.open('GET', 'data.txt', false);
request.send();
if (request.status === 200) {
    var myDiv = document.createElement('div');
    myDiv.innerHTML = request.responseText;
    myNode.appendChild(myDiv);
    mySeparator();
}

// ajax
var secondRequest = new XMLHttpRequest();
secondRequest.open('GET', 'js/people-data.js', false);
secondRequest.send();
if (secondRequest.status === 200) {
    var myDiv = document.createElement('div');
    myDiv.innerHTML = secondRequest.responseText;
    myNode.appendChild(myDiv);
    mySeparator();
}

// formatting output
function mySeparator() {
    var myDiv = document.createElement('div');
    myDiv.innerHTML = '----------------------';
    myNode.appendChild(myDiv);
}