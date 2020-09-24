function divide(e) {
  let divNum = Math.floor(100 / e.value)
  let progress = document.getElementById('progress')
  if (!Number.isNaN(e.value) && e.value > 0) {
    let inner = ""
    let divnumSum = 0
    let divnums = []
    for (let i = 0;i < e.value; i++) {
      divnumSum += divNum
      divnums.push(divNum)
    }
    if (100 - divnumSum >= 0) {
      for(let i = 0;i < 100 - divnumSum; i++) {
        divnums[i]++
      }
    }
    for (let i = 0;i < e.value; i++) {
      inner += "<div class=\"progress-bar\" role=\"progressbar\" style=\"width: " +divnums[i] + "%\" aria-valuenow=" + divnums[i] + "aria-valuemin=\"0\" aria-valuemax=\"100\">" + divnums[i] + "%</div>"
      inner += "<div class=\"progress-bar bg-dark\" role=\"progressbar\" style=\"width: 1px;\" aria-valuemin=\"0\" aria-valuemax=\"100\"></div>"
    }
    progress.innerHTML = inner
  } else {
    progress.innerHTML = ""
  }
}