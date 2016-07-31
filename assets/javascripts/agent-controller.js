// Able to delete newly created records

newApp.controller('agentController', function newAppController($scope, $http, $httpParamSerializer) {
  $scope.types = ["condo", "apartment"]

  $scope.propertyCreate = function() {
    var property = $httpParamSerializer($scope.property)
    property.address = $('#address').val();
    debugger;

    $http.post('/api/property', property, config).success(function(data, status, header){
      $scope.property = {}
      console.log("created");
    }).error(function(data, status, header){
      console.log("Failed");
    })
  };

  var config = {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8;'
    }
  }

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
});
