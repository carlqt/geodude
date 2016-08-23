newApp.controller('registrationsController', function newAppController($scope, $http, $httpParamSerializer) {
  $scope.user = { type: 'user', types: ['user', 'agent']}
  $scope.createUser = function(isValid) {
    if (isValid) {
      alert("Amazing");
    };
  };
});