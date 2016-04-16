var url = 'https://test-swbc-13-01-form.firebaseio.com/';
var ref = new Firebase(url);

// on is like addEventListener in firebase
ref.on('value', function (snapshot) {
    var data = snapshot.val();
    var output = '';
    for (var i in data) {
        console.log(i);
        console.log(data[i]);
        output += '<section class="wrapBox">';
        output += '<h1>' + data[i].title + '</h1>';
        output += '<p>' + data[i].message + '</p>';
        output += '<button class="deleter" id="' + i + '">delete me</button>';
        output += '</section>';
    }
    document.querySelector('main').innerHTML = output;
});

document.querySelector('#mySubmit').addEventListener('click', function () {
    ref.push({
        "title": (document.forms[0].myFrmTitle.value),
        "message": (document.forms[0].myTextArea.value)
    });
});

document.querySelector('main').addEventListener('click', function(e){
    console.log(e);
    console.log(e.target);
    console.log(e.target.className == 'deleter');
    if(e.target.className == 'deleter') {
        console.log(e.target.id);
        var deleteURL = url + e.target.id;
        console.log(deleteURL);
        var deleteMe = new Firebase(deleteURL);
        deleteMe.remove();
    }
});

// YOU COULD ALSO DO ERROR CHECKING
// SEE HERE:
// http://planetoftheweb.com/presentations/fresno/W03D02.html#/15
//
//document.querySelector('main').addEventListener('click', function(e){
//    console.log(e);
//    console.log(e.target);
//    console.log(e.target.className == 'deleter');
//    if(e.target.className == 'deleter') {
//        console.log(e.target.id);
//        var deleteURL = url + e.target.id;
//        console.log(deleteURL);
//        var deleteMe = new Firebase(deleteURL);
//        var onDeleteError = function(error) {
//            if (error) {
//                console.log('Synchronization failed');
//            } else {
//                console.log('Synchronization succeeded');
//            }
//        };
//        deleteMe.remove(onDeleteError);
//    }
//});


document.querySelector('#removeAllDate').addEventListener('click', function () {
    ref.remove();
});