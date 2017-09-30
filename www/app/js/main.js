$(document).ready(function () {
    $.getJSON("data/champs.json", function(champs) {
        $.each(champs.data, function(index, champ){
            var champFullImageName = champ.image.full;
            var champImgUrl = "//ddragon.leagueoflegends.com/cdn/7.19.1/img/champion/" + champFullImageName;
            $("#divChamps").append("<div id=\"champ" + champ.id + "\"" + "class=\"champ col-xs-4 col-sm-3 col-md-2 col-lg-1\"></div>");
            $("#champ" + champ.id).append("<img class=\"champImg\" src=\"" + champImgUrl + "\"/>");
        });
    });

    $("#sumNameEmptyErrMsg").hide();
});


