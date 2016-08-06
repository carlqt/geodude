newApp.controller('loginRegisterController', function newAppController($scope, $http, $httpParamSerializer) {
  $scope.registerForm = false;
  $scope.loginForm = true;

  $scope.showRegister = function() {
    $scope.registerForm = true;
    $scope.loginForm = false;
  };

  $scope.showLogin = function() {
    $scope.registerForm = false;
    $scope.loginForm = true;
  };
});