$(document).ready(function () {
    $.getJSON("data/champs.json", function(champs) {
        $.each(champs.data, function(index, champ){
            var champFullImageName = champ.image.full;
            var champImgUrl = "//ddragon.leagueoflegends.com/cdn/9.24.2/img/champion/" + champFullImageName;
            $("#divChamps").append("<div id=\"champ" + champ.key + "\"" + "class=\"champ col-xs-4 col-sm-3 col-md-2 col-lg-1\"></div>");
            $.each(champ.tags, function(index, tag) {
                $("#champ" + champ.key).addClass(tag);
            });
            $("#champ" + champ.key).append("<img class=\"champImg\" src=\"" + champImgUrl + "\"/>");
        });

        var imageCounter = 0;
        $(".champImg").one("load", function() {
            imageCounter++;
            if (imageCounter === Object.keys(champs.data).length) {
                $(".loader").hide();
                $("#divChamps").show();
            }
        }).each(function() {
            if(this.complete) {
                $(this).trigger('load');
            }
        });
    })

    $("#sumNameEmptyErrMsg").hide();
});


