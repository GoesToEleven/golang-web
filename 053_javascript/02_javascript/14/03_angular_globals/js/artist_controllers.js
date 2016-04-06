"use strict";

var artistControllers = angular.module('artistControllers', ["firebase"]);

artistControllers.controller('ListController', function ($scope, GetData) {
    $scope.data = GetData;
});

artistControllers.controller('DetailsController', function ($scope, GetData, $routeParams) {
    $scope.data = GetData;
    $scope.whichItem = $routeParams.itemId;
});