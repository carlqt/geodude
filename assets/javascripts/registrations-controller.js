newApp.controller('registrationsController', function newAppController($scope, $http, $httpParamSerializer) {
  $scope.user = { type: 'user', types: ['user', 'agent']}
});