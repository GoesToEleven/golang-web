for (var i = 1; i <= 20; i++) {
    var remainderDiv3 = i % 3;
    var remainderDiv5 = i % 5;

    if (!remainderDiv3 && !remainderDiv5)
        console.log("FizzBuzz");
    else if (!remainderDiv3)
        console.log("Fizz");
    else if (!remainderDiv5)
        console.log("Buzz");
    else
        console.log(i);
}