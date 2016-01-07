$(document).ready(function() {
	$("#registerUser").prop('disabled', true);
	$("#checkUser").on('click', function() {
		$("#userLabel").html("");
		var userName = $("#userName").val();
		var name = $("#name").val();
		var password = $("#password").val();
		if (userName == "" || name == "" || password == "") {
			$("#userLabel").show();
			$("#userLabel").html("Please fill all the details before proceding further");
			$("#userLabel").fadeTo(2000, 0.4);
		} else {
			$.ajax({
				url : '/register/check/',
				type : 'post',
				dataType : 'json',
				data : {
					UserName : userName
				},
				success : function(data) {
					console.log(data);
					if (data.Check == "false") {
						$("#userLabel").show();
						$("#userLabel").html(data.Message);
						$("#registerUser").prop('disabled', true);
					} else {
						$("#userLabel").show();
						$("#userLabel").html(data.Message);
						$("#registerUser").prop('disabled', false);
					}
				}
			});
		}
	});

	$("#registerUser").on('click', function() {
		var userName = $("#userName").val();
		var name = $("#name").val();
		var password = $("#password").val();
		$.ajax({
			url : '/register/save/',
			type : 'post',
			datatype : 'html',
			data : {
				UserName : userName,
				Name : name,
				Password : password
			},
			success : function(data) {
				window.location.href = "/";
			}
		});

	});
});
