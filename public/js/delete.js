function deletePage(){
  location.href = "/delete"
}

function viewPage() {
  location.href = "/home"
}

function deleteCheck(patientID) {
  checkbox = document.getElementById("patientCheck" + patientID)
  checkbox.checked = !checkbox.checked
}