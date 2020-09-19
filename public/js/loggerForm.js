function setMenu(ID, name) {
  patient = document.getElementById("patient")
  patientID = document.getElementById("patientID")
  patient.value = name
  patientID.value = ID
  console.log(patientID.value)
}

function listMenu(e) {
  let params = new URLSearchParams()
  params.append('searchText', e.value)
  patient = document.getElementById("patient")
  axios.post("/api/readPastPatient", params).then(res => {
    let menu = document.getElementById('menu')
    let inner = ""
    for (let i = 0;i < res.data.length; i++) {
      inner += "<button class=\"dropdown-item\" type=\"button\" onClick=\"setMenu(" + res.data[i].ID + ", '" + res.data[i].Name + "')\"> " + res.data[i].Name + "</button>"
    }
    menu.innerHTML = inner
  })
}