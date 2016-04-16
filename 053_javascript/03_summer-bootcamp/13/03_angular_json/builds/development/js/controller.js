var myApp = angular.module('myApp', []);

myApp.controller('MyController', ['$scope', '$http', function($scope, $http){
    $http.get('js/data.json').success(function(data) {
        console.log(data);
        $scope.artists = data.artists;
        console.log($scope.artists);
    });
}]);