$(document).ready(function() {
	var userName = $('body').attr('id');
	$.views.settings.allowCode = true;
	$.views.tags("loadBuilding", function(BuildingType) {
		return loadBuilding(BuildingType);
	});
	$.views.tags("loadPlayers", function(career) {
		return loadPlayers(career);
	});
	$.views.converters("floor", function(val) {
		return Math.floor(val);
	});
	$.ajax({
		url : "/world01/fetch/player/",
		type : "post",
		dataType : "json",
		data : {
			UserName : userName,
		},
		success : function(data) {
			firstPlay(data);
			renderTemplate({
				name : 'buttonPanel',
				selector : '#buttonPanel',
				playerData : data,
			});
			renderTemplate({
				name : 'userPanel',
				selector : '#userPanel',
				playerData : data,
			});
			renderTemplate({
				name : 'queue',
				selector : '#queue',
				playerData : data,
			});
			renderTemplate({
				name : 'playArea',
				selector : '#playArea',
				playerData : data,
			});
			renderTemplate({
				name : 'rightPanel',
				selector : '#rightPanel',
				playerData : "",
			});
			bindings();
			cssStyling();
		}
	});

	setInterval(refreshResources, 3000);
	setInterval(refreshQueue, 3000);
	setInterval(refreshBuildings, 25000);
});

function renderTemplate(fileStuff) {
	var file = '/templates/' + fileStuff.name + '.htm';
	var data = fileStuff.playerData;
	$.when($.get(file)).done(function(tmplData) {
		$.templates({
			tmpl : tmplData
		});
		$(fileStuff.selector).html($.render.tmpl(data).trim());
	});
};

function firstPlay(playerData) {
	if (playerData.Career == "") {
		$('#playArea').toggle();
		$('#firstPlayPopup').toggle();
		$('#careerSelection').on('change', function() {
			var career = $('#careerSelection').val();
			var careerConfirm = window.confirm("Are you sure you want to  be " + career);
			if (careerConfirm == true) {
				$('#playArea').toggle();
				$('#firstPlayPopup').toggle();
				$.ajax({
					url : "/world01/save/player/",
					type : "post",
					dataType : "json",
					data : {
						UserName : playerData.UserName,
						Career : career,
						FirstPlay : "FirstPlay",
					},
				});
				window.location.reload();
			}
		});
	}

};

function bindings() {
	buttonPanelBind();
	userPanelBind();
	rightPanelBind();
	leftPanelBind();
	playAreaBind();
}

function cssStyling() {
	buttonPanelStyle();
	userPanelStyle();
	rightPanelStyle();
	leftPanelStyle();
	playAreaStyle();
}
