$(document).ready(function() {
	"use strict";

	// Variables when searching for agents
	var selectPlatform = $('#select-platform'),
		selectActivity = $('#select-activity'),
		selectMicrophone = $('#select-microphone'),
		selectLookingFor = $('#select-lookingfor'),
		selectLevel = $('#select-level'),
		selectDZLevel = $('#select-dzlevel'),
		selectSearch = $('#select-search');

	// Variables when adding a new agent
	var addUsername = $('#aa-username'),
		addPlatform = $('#aa-platform'),
		addActivity = $('#aa-activity'),
		addMicrophone = $('#aa-microphone'),
		addLookingFor = $('#aa-lookingfor'),
		addLevel = $('#aa-level'),
		addDZLevel = $('#aa-dzlevel'),
		addDescription = $('#aa-description');

	// Buttons
	var saveAgentBtn = $('#save-agent'),
		searchAgentsBtn = $('#search-agents');

	// Button contents
	var saveAgentBtnOn = '<i class="fa fa-plus-circle"></i> Agregar agente',
		saveAgentBtnOff = '<i class="fa fa-spinner fa-spin"></i> Agregando, espera...',
		searchAgentsBtnOn = '<i class="fa fa-search"></i> Buscar Agentes',
		searchAgentsBtnOff = '<i class="fa fa-spinner fa-spin"></i> Buscando, espera...';

	// Places for alerts
	var addAlertPlacement = $('#add-alert-placement');

	// Modals
	var modalAddAgent = $('#modal-add-agent');

	// Error alert to use here
	var errAlert = $('<div>').attr({ "class": "alert alert-danger", "role": "alert", "style": "display:none" }).append(
		$('<button>').attr({ "type": "button", "class": "close", "data-dismiss": "alert", "aria-label": "close" }).append(
			$('<span>').attr({ "aria-hidden": "true" }).html('&times;')
		)
	);

	// Messages
	var eUserIncorrect = '<strong>Oh no!</strong> El <u>nombre de usuario</u> es incorrecto o es demasiado corto. Verifícalo y vuelve a intentarlo.',
		ePlatformIncorrect = '<strong>Oh no!</strong> No has seleccionado <u>una plataforma</u>. Elige una y vuelve a intentarlo.',
		eActivityIncorrect = '<strong>Oh no!</strong> No has seleccionado <u>una actividad</u>. Elige una y vuelve a intentarlo.',
		eHasMicrophone = '<strong>Oh no!</strong> No has indicado si <u>tienes o no un micrófono</u>. Elige una opción y vuelve a intentarlo.',
		eWhatAreYouLooking = '<strong>Oh no!</strong> No has indicado <u>qué buscas en otros agentes</u>. Elige una opción y vuelve a intentarlo.',
		eLevelNormal = '<strong>Oh no!</strong> No has indicado <u>qué nivel tienes en el modo historia</u>. Elige una opción y vuelve a intentarlo.',
		eLevelDarkZone = '<strong>Oh no!</strong> No has indicado <u>qué nivel tienes en la Zona Oscura</u>. Elige una opción y vuelve a intentarlo.',
		eNoAdditionalInfo = '<strong>Oh no!</strong> No has entregado <u>información adicional</u> o escribiste poca información, como tu DPS o tu nivel de Salud. Escribe un detalle y vuelve a intentarlo.';

	// Show errors functions
	var fnShowError = function(placement, message) {
		// Scroll to the top of the modal window
		modalAddAgent.scrollTop(0);

		// Append error message
		placement.empty().append(
			errAlert.clone().append(message).fadeIn('fast')
		);
	};

	// Toggle button state
	var fnToggleButton = function(btn, message) {
		// Check if element is disabled
		if (!btn.is(':disabled')) {
			// Disable element
			btn.prop('disabled', true);

			// Store the previous value
			btn.attr('data-original', btn.html());

			// Write the new value
			btn.html(message);
		} else {
			// Disable element
			btn.prop('disabled', false);

			// Recover the content from the attribute
			btn.html(btn.attr('data-original'));

			// Clean the value
			btn.attr('data-original', '');
		}
	};

	// When saving an agent
	saveAgentBtn.click(function(e) {
		// Create a placeholder func so DRY
		var fnSaveAgentError = function(errormsg) {
			fnShowError(addAlertPlacement, errormsg);
			fnToggleButton(saveAgentBtn, saveAgentBtnOn);
		};

		// Disable the button and animate it
		fnToggleButton(saveAgentBtn, saveAgentBtnOff)

		// Check if the username is not empty
		if ($.trim(addUsername.val()) === "" || $.trim(addUsername.val()).length < 6) {
			fnSaveAgentError(eUserIncorrect);
			return;
		}

		// Check if the platform is selected
		if (addPlatform.val() == null) {
			fnSaveAgentError(ePlatformIncorrect);
			return
		}

		// Check if the activity is selected
		if (addActivity.val() == null) {
			fnSaveAgentError(eActivityIncorrect);
			return
		}

		// Check if the user selected microphone or not
		if (addMicrophone.val() == null) {
			fnSaveAgentError(eHasMicrophone);
			return
		}

		// Check if the looking for is selected
		if (addLookingFor.val() == null) {
			fnSaveAgentError(eWhatAreYouLooking);
			return
		}

		// Check if the normal level is selected
		if (addLevel.val() == null) {
			fnSaveAgentError(eLevelNormal);
			return
		}

		// Check if the DZ level is selected
		if (addDZLevel.val() == null) {
			fnSaveAgentError(eLevelDarkZone);
			return
		}

		// Check if the description is not empty
		if ($.trim(addDescription.val()) === "" || $.trim(addDescription.val()).length < 6) {
			fnSaveAgentError(eNoAdditionalInfo);
			return
		}

		setTimeout(function() {
			fnToggleButton(saveAgentBtn, saveAgentBtnOn);
		}, 5000);
	});

	// When searching for agents
	searchAgentsBtn.click(function(e) {
		// Toggle button state
		fnToggleButton(searchAgentsBtn, searchAgentsBtnOff);

		setTimeout(function() {
			// Toggle button state
			fnToggleButton(searchAgentsBtn, searchAgentsBtnOn);
		}, 5000);
	});
});