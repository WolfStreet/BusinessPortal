function rightPanelBind() {
	
}

function rightPanelStyle() {
	
}

function loadBank(bson_id, resources, users, actioner) {
	var postAt = '/players/career/';
	var response = $.ajax({
		url : "/world01/fetch/bank/",
		type : 'post',
		dataType : 'json',
		async : false,
		data : {
			ID : bson_id,
			Resources : resources,
			Users : users,
			Actioner : actioner,
		}
	}).done();
	return response.responseJSON;
}

function loadMarket(bson_id, resources, users, actioner) {
	var response = $.ajax({
		url : "/world01/fetch/market/",
		dataType : 'json',
		type : 'post',
		async : false,
		data : {
			ID : bson_id,
			Resources : resources,
			Users : users,
			Actioner : actioner,
		}
	}).done();
	return response.responseJSON;
}

function loadJobs(bson_id, career, users, actioner) {
	var response = $.ajax({
		url : "/world01/fetch/job/",
		dataType : 'json',
		type : 'post',
		async : false,
		data : {
			ID : bson_id,
			Career : career,
			Users : users,
			Actioner : actioner,
		}
	}).done();
	return response.responseJSON;
}

/*---right panel post popup---*/

function displaySliderValue(element) {
	$(element).parent().children('p:first').text(element.value + "/" + element.max);
	$('#postPopup .Message').hide();
}

function loadList(element) {
	$(element).siblings('div:eq(0)').show();
	$(element).siblings('div:eq(0)').children('input[type=button]').val(element.value);
	if (element.value == "Candidate") {
		$(element).siblings('div:eq(0)').children('input[type=checkbox]').hide();
		$(element).siblings('div:eq(0)').children('br').hide();
	} else {
		$(element).siblings('div:eq(0)').children('input[type=checkbox]').show();
		$(element).siblings('div:eq(0)').children('br').show();
	}
}

function postMarket(element) {
	$('#postMarketResource').val($('#postMarketResource').siblings('input[type=checkbox]:checked').val());
	var userName = $('body').attr('id');
	var temp = $('#player tr:eq(1) td b').text();
	temp = temp.split(' ');
	var name = temp[1];
	var actioner = element.value;
	var resource = $('#postMarketResource').val();
	var amount = $('#postMarketAmount').val();
	var rate = $('#postmarketRate').val();
	if (!$.isNumeric(amount)) {
		$('#postPopup .Message').show();
		$('#postPopup .Message').text('Amount Entered is Not Number');
		return;
	} else {
		if (parseInt(amount, 10) <= 0) {
			$('#postPopup .Message').show();
			$('#postPopup .Message').text('Amount Entered is Not Positive');
			return;
		}
	}

	if (!$.isNumeric(rate)) {
		$('#postPopup .Message').show();
		$('#postPopup .Message').text('Rate Entered is Not Number');
		return;
	} else {
		if (parseInt(rate, 10) <= 0) {
			$('#postPopup .Message').show();
			$('#postPopup .Message').text('Rate Entered is Not Positive');
			return;
		}
	}
	$.ajax({
		url : '/world01/post/market/',
		type : 'post',
		dataType : 'json',
		data : {
			UserName : userName,
			Name : name,
			Actioner : actioner,
			Resource : resource,
			Amount : amount,
			Rate : rate
		},
		success : function(data) {
			if (data.check == "true") {
				$('#postPopup .Message').hide();
			} else {
				$('#postPopup .Message').show();
				$('#postPopup .Message').text(data.Message);
			}
		}
	});
}

function postJobs(element) {
	if (element.value == "Candidate") {
		$('#postJobsCareer').val($('body').attr('data-career'));
		console.log($('body').attr('data-career'));
		return;
	} else {
		$('#postJobsCareer').val($('#postJobsCareer').siblings('input[type=checkbox]:checked').val());
	}
	var userName = $('body').attr('id');
	var actioner = element.value;
	var career = $('#postJobsCareer').val();
	var fee = $('#postJobsFee').val();
	var hours = $('#postJobsHours').val();
	if (!$.isNumeric(hours)) {
		$('#postPopup .Message').show();
		$('#postPopup .Message').text('Hours Entered is Not Number');
		return;
	} else {
		if (parseInt(hours, 10) <= 0) {
			$('#postPopup .Message').show();
			$('#postPopup .Message').text('Hours Entered is Not Positive');
			return;
		}
	}
	if (!$.isNumeric(fee)) {
		$('#postPopup .Message').show();
		$('#postPopup .Message').text('Fee Entered is Not Number');
		return;
	} else {
		if (parseInt(fee, 10) <= 0) {
			$('#postPopup .Message').show();
			$('#postPopup .Message').text('Fee Entered is Not Positive');
			return;
		}
	}
	$.ajax({
		url : '/world01/post/job/',
		type : 'post',
		dataType : 'json',
		data : {
			UserName : userName,
			Actioner : actioner,
			Career : career,
			Fee : fee,
			Hours : hours,
		},
		success : function(data) {
			if (data.Check == "true") {
				$('#postPopup .Message').hide();
			} else {
				$('#postPopup .Message').show();
				$('#postPopup .Message').text(data.Message);
			}
		}
	});
}

function postBank(element) {
	var userName = $('body').attr('id');
	var actioner = element.value;
	var amount = $('#postBankAmount').val();
	var rate = $('#postBankRate').val();
	if (!$.isNumeric(amount)) {
		$('#postPopup .Message').show();
		$('#postPopup .Message').text('Amount Entered is Not Number');
		return;
	} else {
		if (parseInt(amount, 10) <= 0) {
			$('#postPopup .Message').show();
			$('#postPopup .Message').text('Amount Entered is Not Positive');
			return;
		}
	}

	if (!$.isNumeric(rate)) {
		$('#postPopup .Message').show();
		$('#postPopup .Message').text('Rate Entered is Not Number');
		return;
	} else {
		if (parseInt(rate, 10) <= 0) {
			$('#postPopup .Message').show();
			$('#postPopup .Message').text('Rate Entered is Not Positive');
			return;
		}
	}
	$.ajax({
		url : '/world01/post/bank/',
		type : 'post',
		dataType : 'json',
		data : {
			UserName : userName,
			Actioner : actioner,
			Amount : amount,
			Rate : rate,
		},
		success : function(data) {
			if (data.Check == "true") {
				$('#postPopup .Message').hide();
			} else {
				$('#postPopup .Message').show();
				$('#postPopup .Message').text(data.Message);
			}
		}
	});
}

/*---right panel post popup ends---*/
/*---right panel accept popup---*/

function acceptJobs(element) {
	var userName = $('body').attr('id');
	console.log(element.id);
	var bson_id = element.id;
	$.ajax({
		url : '/world01/accept/job/',
		type : 'post',
		dataType : 'json',
		data : {
			ID : bson_id,
			UserName : userName,
		},
		success : function(data) {
			console.log(data);
		}
	});

}

function acceptMarket(element) {
	var userName = $('body').attr('id');
	console.log(element.id);
	var bson_id = element.id;
	$.ajax({
		url : 'worl01/accept/market/',
		type : 'post',
		dataType : 'json',
		data : {
			ID : bson_id,
			UserName : userName,
		},
		success : function(data) {
			console.log(data);
		}
	});
}

function acceptBank(element) {
	var userName = $('body').attr('id');
	console.log($(element).attr('id'));
	var bson_id = $(element).attr('id');
	$.ajax({
		url : 'wprld01/accept/bank/',
		type : 'post',
		dataType : 'json',
		data : {
			ID : bson_id,
			UserName : userName,
		},
		success : function(data) {
			console.log(data);
		}
	});
}
/*---right panel accept popup ends---*/

