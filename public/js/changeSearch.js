var pagination = document.getElementById('pagination').innerHTML
function changeSearch(e) {
  let patients = JSON.parse(window.localStorage.getItem('allPatients'))
  if (e.value !== '') {
    document.getElementById('pagination').innerHTML = ''
    result = Search(e.value, patients)
    updateTable(result)
  } else {
    patients = JSON.parse(window.localStorage.getItem('patients'))
    updateTableless(patients)
    document.getElementById('pagination').innerHTML = pagination
  }
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

function extractPatientID(str) {
  str = str.replace("function onclick(event) {\nopenDialog( ", "")
  str = str.replace(" )\n}", "")
  return str
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

function readAllPatients() {
  axios.post('/ptv3/api/readAllPatients', {}).then(res => {
    jsonStr = JSON.stringify(res.data)
    localStorage.setItem('allPatients', jsonStr)
  })
}

function storePatients() {
  let i = 0
  let patients = []
  while (document.getElementById("patient" + i) !== null) {
    let row = document.getElementById("patient" + i)
    functionStr = row.onclick.toString()
    functionStr = functionStr.replace("function onclick(event) {\nopenDialog( ", "")
    patientID = extractPatientID(functionStr)
    patient = getEachCol(row)
    patient.patientID = patientID
    patients.push(patient)
    i++
  }
  jsonStr = JSON.stringify(patients)
  window.localStorage.setItem('patients', jsonStr)
}

function updateTable(dispdata) {
  tbody = document.getElementById("patientDataTable")
  let inner = ""
  for (let i = 0;i < dispdata.length; i++) {
  inner += "<tr onClick=\"openDialog(" + dispdata[i].PatientID + ")\" data-toggle=\"modal\" data-target=\"#fsmodal\"" + ">"
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
  inner += "<tr onClick=\"openDialog(" + dispdata[i].patientID + ")\" data-toggle=\"modal\" data-target=\"#fsmodal\"" + ">"
  inner += "<th>" + dispdata[i].name + "</th>"
  inner += "<td>" + dispdata[i].gender + "</td>"
  inner += "<td>" + dispdata[i].age + "</td>"
  inner += "</tr>"
  }
  tbody.innerHTML = inner
}

storePatients()
readAllPatients()