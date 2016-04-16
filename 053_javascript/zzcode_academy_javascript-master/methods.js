// here is our method to set the height
var setHeight = function (newHeight) {
    this.height = newHeight;
};

var setWidth = function (newWidth) {
    this.width = newWidth;
};

var rectangle = {};
rectangle.height = 3;
rectangle.width = 4;
// here is our method to set the height
rectangle.setHeight = setHeight;

// help by finishing this method
rectangle.setWidth = setWidth;

// here change the width to 8 and height to 6 using our new methods
rectangle.setHeight(6);
rectangle.setWidth(8);

console.log(rectangle.height);
console.log(rectangle.width);

// LESS PREFERRED WAY BELOW

var rectangle = new Object();
rectangle.height = 3;
rectangle.width = 4;
// here is our method to set the height
rectangle.setHeight = function (newHeight) {
    this.height = newHeight;
};
// help by finishing this method
rectangle.setWidth = function (newWidth) {
    this.width = newWidth
};

// here change the width to 8 and height to 6 using our new methods
rectangle.setHeight(6);
rectangle.setWidth(8);

console.log(rectangle.height);
console.log(rectangle.width);