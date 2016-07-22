
var newApp = angular.module('newApp', []);

newApp.config(function($interpolateProvider){
  $interpolateProvider.startSymbol('<%');
  $interpolateProvider.endSymbol('%>');
});

newApp.controller('newAppController', function newAppController($scope, $http) {
  $http.get("/properties").success(function(data) {
    $scope.locations = data;
  });
});
