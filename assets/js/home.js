var config;

const LEVEL = {
	STUDENT: 0,
	TUTOR: 1,
	SUPERVISOR: 2,
	HEAD: 3,
	ADMIN: 4,
	ROOT: 5
}

String.prototype.capitalize = function() {
	return this.charAt(0).toUpperCase() + this.substring(1);
};

function showWait() {
	$("#cnt").html(waitingBlock);
}

function loadSuccess(data) {
	myself = data;
	$("#fullname").html(myself.Person.Lastname + ", " + myself.Person.Firstname);

	let highest = levelHighest(myself.Roles);
	//my options
	for (i = 0; i <= highest; i++) {
		$(".role-" + i).removeClass("hidden");
	}

	//homepage
	if (levelContains(myself.Roles, LEVEL["STUDENT"])) {
		showStudent();
	} else if (levelHighest(myself.Roles) >= LEVEL["SUPERVISOR"]) {
		showWatchlist();
	} else {
		showTutored();
	}
}

function getLevel(role){
	//because "supervisor" can have a suffix after "-"
	return LEVEL[role.match(/([A-Za-z])*/)[0].toUpperCase()];
}

function allLevels(roles){
	var levels = [];
	for (i = 0; i < roles.length; i++) {
		levels.push(getLevel(roles[i]));
	}
	return levels;
}

function levelContains(roles, role) {
	for (i = 0; i < roles.length; i++) {
		if (getLevel(roles[i]) == role) {
			return true;
		}
	}
	return false;
}

function levelHighest(roles) {
	max = Number.NEGATIVE_INFINITY;
	for (i = 0; i < roles.length; i++) {
		max = Math.max(max, getLevel(roles[i]));
	}
	return max;
}

function showModal(next) {
	$("#modal").modal("show");
	$('#modal').find('[data-toggle="popover"]').popover();
	$('#modal').unbind('shown.bs.modal');
	if (next) {
		$('#modal').on('shown.bs.modal', function() {
			next();
		});
	}
	$("#modal").find(".date").datetimepicker();
	$('#modal').find('[data-toggle="confirmation"]').confirmation();
}

function tableCount() {
	count = $("tbody").find("tr").length;
	$(".count").html(count);
}

function ui() {

	$("#cnt").find(".tablesorter").tablesorter();
	$('#cnt').find('[data-toggle="popover"]').popover();
	$('#cnt').find('[data-toggle="confirmation"]').confirmation();

	$("#cnt").find(".shiftSelectable").shiftSelectable();
	$("#cnt").find("table").bind("sortEnd", function() {
		$("#cnt").find('.shiftSelectable').shiftSelectable();
	});

	$(".date").datetimepicker({
		format: "DD MMM YYYY"
	});

	tableCount();

	$(".globalSelect").change(function() {
		var ctx = $(this).data("context");
		$("#" + ctx).find("input:checkbox").prop("checked", this.checked);
	});
}

function showAlumni(student) {
	internship(student).done(function(i) {
		$("#modal").render("alumni-modal", i.Convention.Student, function() {
			var val = i.Convention.Student.Alumni.Position;

			if (val == "sabbatical") {
				$("#country").addClass("hidden");
				$("#contract").addClass("hidden");
				$("#company").addClass("hidden");
			} else if (val == "entrepreneurship") {
				$("#country").removeClass("hidden");
				$("#contract").addClass("hidden");
				$("#company").addClass("hidden");
			} else if (val == "study") {
				$("#country").removeClass("hidden");
				$("#contract").addClass("hidden");
				$("#company").addClass("hidden");
			} else if (val == "company") {
				$("#country").removeClass("hidden");
				$("#contract").removeClass("hidden");
				$("#company").removeClass("hidden");
		} else if (val == "looking") {
				$("#country").addClass("hidden");
				$("#contract").addClass("hidden");
				$("#company").addClass("hidden");
		}
		showModal();
		});
	});
}

function hideModal() {
	$('#modal').find('[data-toggle="popover"]').popover('destroy');
	$("#modal").modal("hide");
	$("#modal").html("");
}

function updateInternshipRow(em) {
	var partial = $("table").data("partial");
	var row = $("table").find(`tr[data-email="${em}"]`);
	internship(em).done(function(u) {
		var cnt = Handlebars.partials[partial](u);
		row.replaceWith(cnt);
		$('table').trigger("update").trigger("updateCache");
	});
}

function showInternship(em, edit) {
	if (!edit || levelHighest(myself.Roles) < LEVEL["ADMIN"]) {
		internship(em).done(function(i) {
			internshipModal(i, [], edit && levelHighest(myself.Roles) >= LEVEL["ADMIN"]);
		}).fail(logFail);
	} else {
		$.when(internship(em), users()).done(function(i, uss) {
			i = i[0];
			uss = uss[0].filter(function(u) {
				return !levelContains(u.Roles, LEVEL["STUDENT"]) && u.Person.Email != i.Convention.Tutor.Person.Email;
			});
			internshipModal(i, uss, edit);
		}).fail(logFail);
	}
}

function resetSurvey(btn, student, kind) {
	postResetSurvey(student, kind).done(function(data, status, xhr) {
		$(btn).attr("disabled","disabled");
		defaultSuccess(data, status, xhr);
	}).fail(notifyError);
}

function requestSurvey(btn, student, kind) {
	postRequestSurvey(student, kind).done(function(data, status, xhr) {
		defaultSuccess(data, status, xhr);
		$(btn).html(kind);
	}).fail(notifyError);
}

function internshipModal(i, uss, edit) {
	var dta = {
		I: i,
		Editable: edit,
		Teachers: uss
	}
	$("#modal").render("convention-detail", dta, showModal);
}