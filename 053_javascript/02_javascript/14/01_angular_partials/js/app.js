var myApp = angular.module('myApp', [
    'ngRoute',
    'artistControllers'
]);

myApp.config(['$routeProvider', function ($routeProvider) {
    $routeProvider.
        when('/list', {
            templateUrl: 'partials/list.html',
            controller: 'ListController'
        }).when('/bio', {
            templateUrl: 'partials/bio.html',
            controller: 'ListController'
        }).otherwise({
            redirectTo: '/list'
        });
}]);