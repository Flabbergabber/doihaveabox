'use strict';

/* Controllers */

angular.module('myApp.controllers', [])
  .controller('MainController', ['$scope', '$http', function($scope, $http) {
      $scope.hideEarnedBoxes = true;
      $scope.doihaveabox = function(summonerName) {

          if (summonerName === undefined)
          {
              $("#sumNameEmptyErrMsg").show();
          }
          else
          {
              $("#sumNameEmptyErrMsg").hide();

              $http.get("/api/doihaveabox?summonerName=" + summonerName)
                  .then(function(response) {
                      $(".champ").removeClass("earned");
                      $(".champ").show();

                      $.each(response.data, function(index, champMastery){
                            if (champMastery.chestGranted == true) {
                                $("#champ" + champMastery.championId).addClass("earned");
                            }
                      });

                      if ($scope.hideEarnedBoxes) {
                          $(".earned").hide();
                      }
                  });
              }
      }

      $scope.toggleHideEarnedBoxes = function() {
          $scope.hideEarnedBoxes = !$scope.hideEarnedBoxes;

          if ($scope.hideEarnedBoxes) {
              $(".earned").hide();
          } else {
              $(".earned").show();
          }
      }
  }])
