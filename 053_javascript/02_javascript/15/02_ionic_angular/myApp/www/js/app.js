// Ionic Starter App

// angular.module is a global place for creating, registering and retrieving Angular modules
// 'starter' is the name of this angular module example (also set in a <body> attribute in index.html)
// the 2nd parameter is an array of 'requires'
angular.module('starter', ['ionic', 'firebase'])

    .run(function ($ionicPlatform) {
        $ionicPlatform.ready(function () {
            // Hide the accessory bar by default (remove this to show the accessory bar above the keyboard
            // for form inputs)
            if (window.cordova && window.cordova.plugins.Keyboard) {
                cordova.plugins.Keyboard.hideKeyboardAccessoryBar(true);
            }
            if (window.StatusBar) {
                StatusBar.styleDefault();
            }
        });
    })

    .controller('ListController', ['$scope', '$firebaseArray', function ($scope, $firebaseArray) {
        var ref = new Firebase("https://test-swbc-15-02.firebaseio.com/");
        $scope.artists = $firebaseArray(ref);
        $scope.moveItem = function (item, fromIndex, toIndex) {
            $scope.artists.splice(fromIndex, 1);
            $scope.artists.splice(toIndex, 0, item);
        };
    }]);
