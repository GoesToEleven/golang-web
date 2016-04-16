var calcPerimeter = function() {
    return this.sideLength * 4;
};

var calcArea = function() {
    return this.sideLength * this.sideLength;
};

var square = {};
square.sideLength = 6;
square.calcPerimeter = calcPerimeter;
// help us define an calcArea method here
square.calcArea = calcArea;

var p = square.calcPerimeter();
var a = square.calcArea();

console.log(p);
console.log(a);

// less preferred start ...

var square = new Object();
square.sideLength = 6;
square.calcPerimeter = function() {
    return this.sideLength * 4;
};
// help us define an area method here
square.calcArea = function() {
    return this.sideLength * this.sideLength;
};

var p = square.calcPerimeter();
var a = square.calcArea();