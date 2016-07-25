
var newApp = angular.module('newApp', []);

newApp.config(function($interpolateProvider){
  $interpolateProvider.startSymbol('<%');
  $interpolateProvider.endSymbol('%>');
});

newApp.controller('newAppController', function newAppController($scope, $http) {
  $scope.locations = {}

  $http.get("/api/properties").success(function(data) {
    $scope.locations = data;
  });

  $scope.searchNearby = function($event){
    var keyCode = $event.which || $event.keyCode;
    if (keyCode === 13) {

      var config = {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8;'
        }
      }

      $http.post('/api/property', "location=" + $event.currentTarget.value, config)
      .success(function(data, status, header){
        debugger;
        console.log("created");
      })
      .error(function(data, status, header){
        console.log("Failed");
      })
    }
  };
});
