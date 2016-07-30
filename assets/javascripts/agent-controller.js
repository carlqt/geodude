// Able to delete newly created records

newApp.controller('agentController', function newAppController($scope, $http) {
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
});
