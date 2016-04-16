"use strict";

var myApp = angular.module('myApp', [
    'ngRoute',
    'artistControllers'
]).constant('FIREBASE_URL', 'https://test-swbc-14-03.firebaseio.com/');

myApp.config(['$routeProvider', function ($routeProvider) {
    $routeProvider.
        when('/list', {
            templateUrl: 'partials/list.html',
            controller: 'ListController'
        }).when('/details/:itemId', {
            templateUrl: 'partials/details.html',
            controller: 'DetailsController'
        }).when('/bio', {
            templateUrl: 'partials/bio.html',
            controller: 'ListController'
        }).otherwise({
            redirectTo: '/list'
        });
}]);

myApp.factory('GetData', function ($firebaseArray, FIREBASE_URL) {
    var ref = new Firebase(FIREBASE_URL);
    return $firebaseArray(ref);
});


