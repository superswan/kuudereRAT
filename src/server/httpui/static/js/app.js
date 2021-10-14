// Get the modal
var taskSelect = document.getElementById("tasks")
var shellModal = document.getElementById("shellModal");

taskSelect.addEventListener('change', function() {
  var options = taskSelect.querySelectorAll("option");
  var count = options.length;
  if(taskSelect.value == "shell") {
    shellModal.style.display = "block";
  }
}, false);

window.onclick = function(event) {
  if (event.target == shellModal) {
    shellModal.style.display = "none";
  }
}