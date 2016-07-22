
var newApp = angular.module('newApp', []);

newApp.config(function($interpolateProvider){
  $interpolateProvider.startSymbol('<<');
  $interpolateProvider.endSymbol('>>');
});