$(document).ready(function() {
	$('input[type="submit"]').on('click', function(event) {
		event.preventDefault();
		var userName = $("#userName").val();
		var password = $("#password").val();
		if (userName == "" || password == "") {
		} else {
			$.ajax({
				url : "/login/",
				type : "post",
				dataType : "json",
				data : {
					UserName : userName,
					Password : password,
				},
				success : function(data) {
					if (data.Check == "true") {
						$('#user').attr('action', "/world01/");
						$('#user').serialize();
						$('#user').submit();
					} else {
						$('.Message').html(data.Message);
						$('.Message').show();
						$('#userName').attr('value', "");
						$('#password').attr('value', "");
					}
				}
			});
		}
	});
});
