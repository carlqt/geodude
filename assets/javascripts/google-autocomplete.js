function initialize() {
  var input = document.getElementById('address');
  var autocomplete = new google.maps.places.Autocomplete(input);
}

google.maps.event.addDomListener(window, 'load', initialize);
