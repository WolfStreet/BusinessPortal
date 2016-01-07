function leftPanelBind() {
	
}

function leftPanelStyle() {
	
}

function makeBuilding(clickedButton) {
	$('.Message').hide();
	var building = $(clickedButton).attr('id');
	var userName = $('body').attr('id');
	switch (building) {
	case "House":
		var buildingIndex = 1;
		break;
	case "Flat":
		var buildingIndex = 2;
		break;
	case "Shop":
		var buildingIndex = 3;
		break;
	case "Farm":
		var buildingIndex = 4;
		break;
	case "Mill":
		var buildingIndex = 5;
		break;
	case "Mine":
		var buildingIndex = 6;
		break;
	case "Kiln":
		var buildingIndex = 7;
		break;
	case "Textile":
		var buildingIndex = 8;
		break;
	case "Drill":
		var buildingIndex = 9;
		break;
	default:
		var buildingIndex = 0;
		break;
	}
	$.ajax({
		url : "/world01/building/build/",
		type : "post",
		dataType : "json",
		data : {
			UserName : userName,
			BuildType : buildingIndex,
			BuildLevel : 1,
			CoordsX : gridX,
			CoordsY : gridY,
		},
		success : function(data) {
			if (data.Check == "true") {
				$('#buildPanel').hide();
				refreshQueue();
				$('.Message').hide();
			} else {
				$('.Message').show();
				$('.Message').html(data.Message);
			}

		}
	});

}

function refreshQueue() {
	var userName = $('body').attr('id');
	$.ajax({
		url : "/world01/fetch/player/",
		type : "post",
		dataType : "json",
		data : {
			UserName : userName,
		},
		success : function(data) {
			renderTemplate({
				name : 'queue',
				selector : '#queue',
				playerData : data,
			});

		}
	});
}

function updateBuilding(buildType, buildLevel) {
	var userName = $('body').attr('id');
	$.ajax({
		url : "/world01/building/build/",
		type : "post",
		dataType : "json",
		data : {
			UserName : userName,
			BuildType : buildType,
			BuildLevel : buildLevel,
			CoordsX : gridX,
			CoordsY : gridY,
		},
		success : function(data) {
			if (data.Check == "true") {
				$('#updatePanel').hide();
				refreshQueue();
				$('.Message').hide();
			} else {
				$('.Message').show();
				$('.Message').html(data.Message);
			}

		}
	});
}

function cancelQueue(buildType, buildLevel, CoordsX, CoordsY) {
	var userName = $('body').attr('id');
	$.ajax({
		url : "/world01/building/cancelQueue/",
		type : "post",
		dataType : "json",
		data : {
			UserName : userName,
			BuildType : buildType,
			BuildLevel : buildLevel,
			CoordsX : CoordsX,
			CoordsY : CoordsY,
		},
		success : function(data) {
			if (data.Check == "true") {
				$('#buildPanel').hide();
				refreshQueue();
				$('.Message').hide();
			} else {
				$('.Message').show();
				$('.Message').html(data.Message);
			}

		}
	});
}
