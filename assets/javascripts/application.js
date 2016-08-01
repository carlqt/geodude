// Able to delete newly created records

var newApp = angular.module('newApp', ['ngRoute']);

// 2 ways to create a service
newApp.service('demoService', function() {
  this.users = ['John', 'James', 'Jake']
});

newApp.factory('demoFactory', function() {
  var fac = {};
  fac.users = ['John', "James", "Jake"]
  return fac;
});

//--------------------------

newApp.config(function($interpolateProvider, $routeProvider){
  $interpolateProvider.startSymbol('<%');
  $interpolateProvider.endSymbol('%>');

  $routeProvider.when('/', {
    templateUrl: '/assets/templates/propertiesTable.html',
    controller: 'newAppController'
  }).when('/agent', {
    templateUrl: '/assets/templates/agent.html',
    controller: 'agentController'
  });
});

newApp.controller('newAppController', function newAppController($scope, $http, demoService) {
  console.log(demoService.users); // Example on how to use a service in your controller

  $scope.locations = []
  var config = {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8;'
    }
  }


  $http.get("/api/properties").success(function(data) {
    $scope.locations = data;
  });

  $scope.addAddress = function($event){
    var keyCode = $event.which || $event.keyCode;
    if (keyCode === 13) {

      $http.post('/api/property', "location=" + $event.currentTarget.value, config).success(function(data, status, header){
        $scope.locations.push(data)
        $scope.createField = ""
        console.log("created");
      }).error(function(data, status, header){
        console.log("Failed");
      })
    }
  };

  $scope.delete = function(id) {
    $http.delete('/api/property/' + id, config).success(function(data, status, header){
      console.log(data);
      removedItemIndex = $scope.locations.findIndex(x=> x.id === id);
      $scope.locations.splice(removedItemIndex, 1);
    });
  };
});
