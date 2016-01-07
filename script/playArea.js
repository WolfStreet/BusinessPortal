var gridX;
var gridY;

function playAreaBind() {
	console.log("Caught playAreaBind");
	$('#playGrid').bind('click', function(event) {
		tempfunc(event);
	});
}

function tempfunc(tempevent) {
	console.log($(tempevent.target).parent());
}

function playAreaStyle() {

}

function playAreaClick(buildEvent) {
	userName = $('body').attr('id');
	gridX = $(buildEvent).attr('data-x');
	gridY = $(buildEvent).attr('data-y');
	$.ajax({
		url : "/world01/fetch/player/",
		type : "post",
		dataType : "json",
		data : {
			UserName : userName,
		},
		success : function(data) {
			if (data.Buildings[gridX][gridY].Type == 0) {
				$('#buildPanel').show();
				renderTemplate({
					name : "buildPanel",
					selector : '#buildPanel',
					playerData : "",
				});
			} else {
				$('#updatePanel').show();
				renderTemplate({
					name : "updatePanel",
					selector : "#updatePanel",
					playerData : data.Buildings[gridX][gridY],
				});
			}
		}
	});

}

function loadBuilding(BuildingType) {
	switch (BuildingType) {
	case "0", 0:
		return "<p style='color:rgb(50, 20, 250);'> Empty </p>";
		break;
	case "1", 1:
		return "<p style='color:rgb(234, 234, 108);'> House </p>";
		break;
	case "2", 2:
		return "<p style='color:rgb(234, 234, 108);'> Flat </p>";
		break;
	case "3", 3:
		return "<p style='color:rgb(234, 234, 108);'> Shop </p>";
		break;
	case "4", 4:
		return "<p style='color:rgb(50, 234, 50);'> Farm </p>";
		break;
	case "5", 5:
		return "<p style='color:rgb(176, 88, 0);'> Saw Mill </p>";
		break;
	case "6", 6:
		return "<p style='color:rgb(190, 190, 190);'> Mine </p>";
		break;
	case "7", 7:
		return "<p style='color:rgb(100, 100, 100);'> Cement Kiln </p>";
		break;
	case "8", 8:
		return "<p style='color:rgb(128, 0, 255);'> Textile </p>";
		break;
	case "9", 9:
		return "<p style='color:rgb(0, 0, 0);'> Oil Drill </p>";
		break;
	}
}

function refreshBuildings() {
	var userName = $('body').attr('id');
	$.ajax({
		url : "/world01/fetch/player/",
		type : "post",
		dataType : "json",
		async : false,
		data : {
			UserName : userName,
		},
		success : function(data) {
			renderTemplate({
				name : 'playArea',
				selector : '#playArea',
				playerData : data,
			});
		}
	});
	playAreaBind();
	playAreaStyle();
}