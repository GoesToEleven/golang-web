var artistControllers = angular.module('artistControllers', ["firebase"]);
var ref = new Firebase("https://test-swbc-14-02-vari.firebaseio.com/");

//Now the $firebaseObject, $firebaseArray, and $firebaseAuth services are available to be injected into any controller, service, or factory.

artistControllers.controller('ListController', function ($scope, $firebaseArray) {
    $scope.data = $firebaseArray(ref);
});

artistControllers.controller('DetailsController', function ($scope, $firebaseArray, $routeParams) {
    $scope.data = $firebaseArray(ref);
    $scope.whichItem = $routeParams.itemId;
});