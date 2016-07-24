
var newApp = angular.module('newApp', []);

newApp.config(function($interpolateProvider){
  $interpolateProvider.startSymbol('<%');
  $interpolateProvider.endSymbol('%>');
});

newApp.controller('newAppController', function newAppController($scope, $http) {
  $http.get("/api/properties").success(function(data) {
    $scope.locations = data;
  });

  $scope.searchNearby = function($event){
    var keyCode = $event.which || $event.keyCode;
    if (keyCode === 13) {
      alert("Hey guys!");
    }
  };
});
