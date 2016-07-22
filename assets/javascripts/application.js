
var newApp = angular.module('newApp', []);

newApp.config(function($interpolateProvider){
  $interpolateProvider.startSymbol('<%');
  $interpolateProvider.endSymbol('%>');
});

newApp.controller('newAppController', function newAppController($scope) {
  $scope.locations = [{
    address: "Bugis",
    latitude: 1.234,
    longitude: 23.344
  },
  { address: "Ubi Avenue 1",
    latitude: 24.4223,
    longitude: 11.22123 
  }];
});
