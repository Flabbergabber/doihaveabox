'use strict';

/* Controllers */

angular.module('myApp.controllers', [])
  .controller('MainController', ['$scope', '$http', function($scope, $http) {

      function doihaveabox(summonerName) {
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
                          if (champMastery.chestGranted === true) {
                              $("#champ" + champMastery.championId).addClass("earned");
                          }
                      });

                      applyFilter();
                  });
          }
      }

      function applyFilter() {
          $(".champ").show();

          if ($scope.role === undefined || $scope.role.code === 0) {
              if ($scope.hideEarnedBoxes) {
                  $(".champ").not(".earned").show();
                  $(".champ").filter(".earned").hide();
              } else {
                  $(".champ").show();
              }
          } else {
              if ($scope.hideEarnedBoxes) {
                  $(".champ").not(".earned").not("." + $scope.role.name).hide();
                  $(".champ").filter(".earned").hide();
              } else {
                  $(".champ").not("." + $scope.role.name).hide();
              }
          }
      }

      $scope.hideEarnedBoxes = true;
      $scope.doihaveabox = function(summonerName) {
          doihaveabox(summonerName);
      };

      $scope.enterPressed = function(keyEvent) {
          if (keyEvent.which === 13) { //Enter
              doihaveabox($scope.summonerName);
          }
      };

      $scope.toggleHideEarnedBoxes = function() {
          $scope.hideEarnedBoxes = !$scope.hideEarnedBoxes;

          applyFilter();
      };

      $scope.roles = [ {code: 0, name: 'All'}, {code: 1, name: 'Assassin'}, {code: 2, name: 'Fighter'}, {code: 3, name: 'Mage'}, {code: 4, name: 'Marksman'}, {code: 5, name: 'Support'}, {code: 6, name: 'Tank'}];
      $scope.changeRoleFilter = function() {
          applyFilter();
      }
  }]);
