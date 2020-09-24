function lossCheck(id) {
  patient = document.getElementById('patient' + id)
  patient.checked = true
}

function changeSearch(e) {
  let patients = JSON.parse(window.localStorage.getItem('allPatients'))
  if (e.value !== '') {
    result = Search(e.value, patients)
    updateTable(result)
  } else {
    updateTableless(patients)
  }
}

function readAllPatients() {
  axios.post('/api/readAllPatients', {}).then(res => {
    jsonStr = JSON.stringify(res.data)
    localStorage.setItem('allPatients', jsonStr)
  })
}

function getEachCol(cols) {
  let name = cols.getElementsByTagName("th")[0].innerHTML
  let gender = cols.getElementsByTagName("td")[0].innerHTML
  let age = cols.getElementsByTagName("td")[1].innerHTML
  return {
    "name": name,
    "gender": gender,
    "age": age
  }
}

function Search(target, patients) {
  let result = []
  for (let i = 0;i < patients.length; i++){
    if (patients[i].Name.indexOf(target) !== -1) {
      result.push(patients[i])
    }
  }
  return result
}


function updateTable(dispdata) {
  tbody = document.getElementById("patientDataTable")
  let inner = ""
  for (let i = 0;i < dispdata.length; i++) {
  inner += "<tr onClick=\"lossCheck(" + dispdata[i].PatientID + ")\" data-toggle=\"modal\" data-target=\"#fsmodal\"" + ">"
  inner += "<td><input id=\"patient" + dispdata[i].PatientID + "\" type=\"radio\" name=\"patient\" value=\"" + dispdata[i].PatientID + "\" /></td>"
  inner += "<th>" + dispdata[i].Name + "</th>"
  inner += "<td>" + dispdata[i].Gender + "</td>"
  inner += "<td>" + dispdata[i].Age + "</td>"
  inner += "</tr>"
  }
  tbody.innerHTML = inner
}

function updateTableless(dispdata) {
  tbody = document.getElementById("patientDataTable")
  let inner = ""
  for (let i = 0;i < dispdata.length; i++) {
  inner += "<tr onClick=\"lossCheck(" + dispdata[i].PatientID + ")\" data-toggle=\"modal\" data-target=\"#fsmodal\"" + ">"
  inner += "<td><input id=\"patient" + dispdata[i].PatientID + "\" type=\"radio\" name=\"patient\" value=\"" + dispdata[i].PatientID + "\" /></td>"
  inner += "<th>" + dispdata[i].Name + "</th>"
  inner += "<td>" + dispdata[i].Gender + "</td>"
  inner += "<td>" + dispdata[i].Age + "</td>"
  inner += "</tr>"
  }
  tbody.innerHTML = inner
}

readAllPatients()