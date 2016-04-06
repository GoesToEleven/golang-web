var artistControllers = angular.module('artistControllers', ["firebase"]);

//Now the $firebaseObject, $firebaseArray, and $firebaseAuth services are available to be injected into any controller, service, or factory.

artistControllers.controller('ListController', function ($scope, $firebaseArray) {
    var ref = new Firebase("https://test-swbc-14-01-rout.firebaseio.com/");
    // download the data into a local object
    console.log($firebaseArray(ref));
    $scope.data = $firebaseArray(ref);
});