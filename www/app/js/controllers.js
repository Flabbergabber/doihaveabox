'use strict';

/* Controllers */

angular.module('myApp.controllers', [])
  .controller('MainController', ['$scope', '$http', function($scope, $http) {
      $scope.count = 0;
      $scope.doihaveabox = function(summonerName) {
          $scope.count++;
          $http.get("/api/doihaveabox?summonerName=" + summonerName)
              .then(function(response) {
                  //$scope.summonerResult = response.data;
              });
      }
  }])
