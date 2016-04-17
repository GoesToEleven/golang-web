var myApp = angular.module('myApp', []);

myApp.controller('MyController', function MyController($scope){
    $scope.author = {
        name: 'Todd McLeod',
        title: 'Staff Author',
        company: 'Lynda'
    }
});