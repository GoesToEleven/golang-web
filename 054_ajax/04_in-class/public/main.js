var el = document.querySelector("#enter-area");

el.addEventListener("input", function(e){

    // create AJAX request
    var xhr = new XMLHttpRequest();

    // send AJAX request
    xhr.open("post", "/api/check");
    console.log("sending:", e.target.value);
    xhr.send(e.target.value);

    // receive AJAX response
    xhr.addEventListener("readystatechange", function(){
        if (xhr.readyState === 4 && xhr.status === 200) {
            console.log("received from server:", xhr.responseText);
        }
    });
});
