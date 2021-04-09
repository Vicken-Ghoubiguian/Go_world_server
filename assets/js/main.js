// Function of execution of all the JS code of 'main.js' file...
$( function() {
	
	// For 'accordion' jQuery UI effect...
	$( "#tz_displayer_as_accordion" ).accordion();

	// Define the tooltip's effects...
	$('[data-toggle="tooltip"]').tooltip({

			show: {

        		effect: "slideDown",
        		delay: 250
      		},

      		hide: {

        		effect: "slideUp",
        		delay: 250
      		}
  	});
});