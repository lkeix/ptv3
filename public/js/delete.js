function deletePage(){
  location.href = "/ptv3/delete"
}

function viewPage() {
  location.href = "/ptv3/home"
}

function deleteCheck(patientID) {
  checkbox = document.getElementById("patientCheck" + patientID)
  checkbox.checked = !checkbox.checked
}