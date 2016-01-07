function userPanelBind() {
	
}

function userPanelStyle() {
	
}

function refreshResources() {
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
				name : 'userPanel',
				selector : '#userPanel',
				playerData : data,
			});

		}
	});

}