var min = 0
var sec = 0
var step = 1
var countUp
var loggerType = ""
var we = [['time', 'eval'], [0,4]]
var ee = [['time', 'eval'], [0,2]]

function secCountUp() {
  sec++
  time = document.getElementById("time")
  inner = ""
  if (Math.floor(sec / 60) < 10) {
    inner += '0' + Math.floor(sec / 60).toString()
  } else {
    inner += Math.floor(sec / 60).toString()
  }
  inner += ":"
  if (Math.floor(sec % 60) < 10) {
    inner += '0' + Math.floor(sec % 60).toString()
  } else {
    inner += Math.floor(sec % 60).toString()
  }
  renderChart()
  time.innerHTML = inner
}

function renderChart() {
  canvas = document.getElementById("lefttime")
  new Chart(canvas, {
    type: 'pie',
    data: {
      labels: ["残り", "現在"],
      datasets: [{
          backgroundColor: [
            "#BB5179",
            "#58A27C",
          ],
          data: [900, sec]
      }]
    },
    options: {
      animation: false,
      title: {
        display: true,
        text: '残り時間'
      }
    }
  });
}

function start(logType) {
  loggerType = logType
  countUp = setInterval(secCountUp, 1000)
  step++
  btnControl()
}

function end() {
  clearInterval(countUp)
  step++
  btnControl()
}

function btnControl() {
  let startbtn = document.getElementById('start')
  let endbtn = document.getElementById('end')
  let eval4btn = document.getElementById('eval4')
  let eval3btn = document.getElementById('eval3')
  let eval2btn = document.getElementById('eval2')
  let eval1btn = document.getElementById('eval1')
  let popbtn = document.getElementById('pop')
  let nextbtn = document.getElementById('next')
  if(step === 1){
    startbtn.disabled = false
    endbtn.disabled = true
    eval4btn.disabled = true
    eval3btn.disabled = true
    eval2btn.disabled = true
    eval1btn.disabled = true
    popbtn.disabled = true
    nextbtn.disabled = true
  } else if (step === 2) {
    startbtn.disabled = true
    endbtn.disabled = false
    eval4btn.disabled = false
    eval3btn.disabled = false
    eval2btn.disabled = false
    eval1btn.disabled = false
    popbtn.disabled = false
    nextbtn.disabled = true
  } else if (step === 3) {
    startbtn.disabled = true
    endbtn.disabled = true
    eval4btn.disabled = true
    eval3btn.disabled = true
    eval2btn.disabled = true
    eval1btn.disabled = true
    popbtn.disabled = true
    nextbtn.disabled = false
  }
}

function evaluation(eval, msg) {
  if (loggerType === 'exp') {
    ee.push([sec, eval])
  } else if(loggerType === 'work') {
    we.push([sec, eval])
  }
  window.alert(msg + "を押しました．")
}

function pop() {
  window.alert("取り消しを押しました．")
  if (loggerType === 'exp') {
    ee.pop()
  } else if(loggerType === 'work') {
    we.pop()
  }
}

function next() {
  if (loggerType === "exp") {
    window.localStorage.setItem("exp", JSON.stringify(ee))
    console.log('aaa')
    nextStep()
  } else if (loggerType === "work") {
    window.localStorage.setItem("work", JSON.stringify(we))
    storeData()
  }
}

function setTimerLog() {
  let exp = document.getElementById("exp")
  let work = document.getElementById("work")
  exp.value = localStorage.getItem("exp")
  work.value = localStorage.getItem("work")
}

function storeData() {
  let exp = localStorage.getItem("exp")
  let work = localStorage.getItem("work")
  let param = new URLSearchParams()
  param.append('exp', exp)
  param.append('work', work)
  axios.post("/api/storeLog", param).then(() => {
    location.reload()
  })
}

function nextStep() {
  axios.post("/api/nextstep", {}).then(() => {
    location.reload()
  })
}