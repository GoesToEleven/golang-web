/**
 * Created by tm on 10/4/2014.
 */
var userLetter = "a";
var userNum = 7;
var userNum2 = 12;


var conditionals = function() {
    if (userNum < 10 && userNum2 < 10)
        console.log("Both less than 10");
    else
        console.log("Both not less than 10");

    if (userNum < 10 || userNum2 < 10)
        console.log("One or both less than 10");
    else
        console.log("Neither less than 10");
};

switch (userLetter) {
    case 'a':
        console.log("aaaaaaaa");
        conditionals();
        break;
    case 'b':
        console.log("bbbbbbbbbb");
        conditionals();
        break;
    case 'c':
        console.log("ccccccccccc");
        conditionals();
        break;
    default:
        console.log("zzzzzzzzzzzzzz");
        conditionals();
}
