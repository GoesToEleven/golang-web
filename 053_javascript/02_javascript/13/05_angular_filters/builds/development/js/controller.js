var myApp = angular.module('myApp', ["firebase"]);

//Now the $firebaseObject, $firebaseArray, and $firebaseAuth services are available to be injected into any controller, service, or factory.

myApp.controller('MyController', function ($scope, $firebaseArray) {
    var ref = new Firebase("https://test-swbc-13-02-04.firebaseio.com");
    // download the data into a local object
    console.log($firebaseArray(ref));
    $scope.data = $firebaseArray(ref);
});