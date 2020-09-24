
function readPatientData(patientid) {
  let params = new URLSearchParams()
  params.append('patientID', patientid)
  axios.post('/api/readPatientData', params).then(res => {
    var pastData = res.data
    var tbody = document.getElementById("pastDataTable")
    var inner = ""
    var explogs = []
    var worklogs = []
    var checklogs = []
    for (let i = 0; i < pastData.length; i++) {
      inner += "<tr>"
      inner += "<td>" + pastData[i].Date + "</td>"
      inner += "<td>" + 
              "<button onClick=\"download('" + pastData[i].Date + "', '" + pastData[i].Name + "'," + replaceAll(pastData[i].Exp) + ")\" class=\"btn btn-dark\">説明時</button></td>"
      inner += "<td><button onClick=\"download('" + pastData[i].Date + "', '" + pastData[i].Mame + "'," + replaceAll(pastData[i].Work) + ")\" class=\"btn btn-dark\">作業時</button></td>"
      inner += "<td><button onClick=\"download('" + pastData[i].Date + "', '" + pastData[i].Name + "'," + replaceAll(pastData[i].Chk) + ")\" class=\"btn btn-dark\">チェック時</button></td>"
      inner += "</tr>"
      explogs.push({'date': pastData[i].Date, 'log': pastData[i].Exp })
      worklogs.push({'date': pastData[i].Date, 'log': pastData[i].Work })
      checklogs.push({'date': pastData[i].Date, 'log': pastData[i].Chk})
    }
    localStorage.setItem('exp', JSON.stringify(explogs))
    localStorage.setItem('work', JSON.stringify(worklogs))
    localStorage.setItem('chk', JSON.stringify(checklogs))
    tbody.innerHTML = inner
    init = document.getElementById('pills-home-tab')
    init.click()
  })
}

function openDialog(patientid) {
  readPatientData(patientid)
}

function download(date, name, log) {
  let csv = ""
  for (let i = 0; i < log.length; i++ ) {
    csv += log[i][0] + "," + log[i][1] + "\n"
  }
  let bom = new Uint8Array([0xEF, 0xBB, 0xBF])
  let blob = new Blob([bom, csv], { type: 'text/csv' })
  let link = document.createElement('a')
  link.href = window.URL.createObjectURL(blob)
  link.download = name + "-" + date + ".csv"
  link.click()
}

function replaceAll(str) {
  while (str.indexOf("\"") > 0) {
    str = str.replace("\"", "'")
  }
  return str
}

function displayExp() {
  let exp = document.getElementById("exp")
  exp.innerHTML = "<canvas id=\"expChart\"></canvas>"
  let ctx = document.getElementById('expChart')
  let explainJson = localStorage.getItem('exp')
  let explainObj = JSON.parse(explainJson)
  let type = 'line'
  let data = []
  for (let i = 0; i < explainObj.length; i++) {
    let tmp = {
      data: conpressTime100Percent(explainObj[i].log),
      label: explainObj[i].date,
      borderColor: "rgba(255,0,0,1)",
      backgroundColor: "rgba(0,0,0,0)"
    }
    data.push(tmp)
  }
  let options = {
    title: {
      display: true,
      text: '作業時ログ'
    },
    scales: {
      xAxes: [
        {
          type: 'linear',
          ticks: {
            suggestedMax: 100,
            suggestedMin: 0,
            stepSize: 10,
          }
        }
      ],
      yAxes: [
        {
          ticks: {
            suggestedMax: 4,
            suggestedMin: 1,
            stepSize: 1,
          }
        }
      ]
    }
  }
  new Chart(ctx, {type: type, data: { datasets: data }, options: options})
}

function displayWork() {
  let work = document.getElementById("wo")
  work.innerHTML = "<canvas id=\"workChart\"></canvas>"
  let ctx = document.getElementById('workChart')
  let worklainJson = localStorage.getItem('work')
  let worklainObj = JSON.parse(worklainJson)
  let type = 'line'
  let data = []
  for (let i = 0; i < worklainObj.length; i++) {
    let tmp = {
      data: conpressTime100Percent(worklainObj[i].log),
      label: worklainObj[i].date,
      borderColor: "rgba(255,0,0,1)",
      backgroundColor: "rgba(0,0,0,0)"
    }
    data.push(tmp)
  }
  let options = {
    title: {
      display: true,
      text: '作業時ログ'
    },
    scales: {
      xAxes: [
        {
          type: 'linear',
          ticks: {
            suggestedMax: 100,
            suggestedMin: 0,
            stepSize: 10,
          }
        }
      ],
      yAxes: [
        {
          ticks: {
            suggestedMax: 4,
            suggestedMin: 1,
            stepSize: 1,
          }
        }
      ]
    }
  }
  new Chart(ctx, {type: type, data: { datasets: data }, options: options})
}

function displayChk() {
  base = JSON.parse(localStorage.getItem('chk'))
  let setup = document.getElementById('setup')
  let comprehension = document.getElementById('comprehension')
  let concentration = document.getElementById('concentration')
  let politeness = document.getElementById('politeness')
  let speed = document.getElementById('speed')
  let resistance = document.getElementById('resistance')
  let setupData = []
  let comprehensionData = []
  let concentrationData = []
  let politenessData = []
  let speedinner =  ""
  let cancelinner = ""
  let timeinner = ""
  let percentinner = ""
  let resistanceinnerdeep1 = ""
  let resistanceinner = ""
  for (let i = 0; i < base.length; i++) {
    let log = JSON.parse(base[i].log)
    let tmp = {
      label: base[i].date,
      data: [log.q1, log.q2, log.q3],
      pointBackgroundColor: 'RGB(46,106,177)'
    }
    setupData.push(tmp)
    tmp = {
      label: base[i].date,
      data: [log.q4, log.q5, log.q6],
      pointBackgroundColor: 'RGB(46,106,177)'
    }
    comprehensionData.push(tmp)
    tmp = {
      label: base[i].date,
      data: [log.q7, log.q8, log.q9],
      pointBackgroundColor: 'RGB(46,106,177)'
    }
    concentrationData.push(tmp)
    tmp = {
      label: base[i].date,
      data: [log.q10, log.q11, log.q12],
      pointBackgroundColor: 'RGB(46,106,177)'
    }
    politenessData.push(tmp)
    cancelinner += "<p>" + base[i].date + " " + log.q13 +  "</p>"
    timeinner += "<p>" + base[i].date + " " + log.q14 +  "分</p>"
    percentinner += "<p>" + base[i].date + " " + log.q15 + "</p>"
    resistanceinnerdeep1 += "<p>" + base[i].date + " " + log.q16 + "</p>"
  }
  speedinner += "<h4>不必要な中断がない</h4><div style=\"padding: 15px\">" + 
  cancelinner + "</div><h4>所要時間</h4><div style=\"padding: 15px\">" + 
  timeinner + "</div><h4>塗り率</h4><div style=\"padding: 15px\">" + percentinner + "</div>"
  resistanceinner += "<h4>30分程度塗り絵に取り組める</h4><div style=\"padding: 15px\">" + resistanceinnerdeep1 + "</div>"
  speed.innerHTML = speedinner
  resistance.innerHTML = resistanceinner
  let chartConfig = {
    setup: {
      data: {
        labels: ['塗り絵を見て描かれている絵が何かわかる', '必要な色鉛筆が確認できる(12色)', '適切な配置ができる(色鉛筆 色番号用紙 パズル塗り絵)'],
        datasets: setupData
      },
      option: {
        title: {
          display: true,
          text: '段取り'
        },
        scale:{
          ticks:{
            suggestedMin: 1,
            suggestedMax: 5,
            stepSize: 1
          }
        }
      }
    },
    comprehension: {
      data: {
        labels: ['色番号と色鉛筆の色を選別できる', '下絵の番号が区別できる', '塗り方が理解できる'],
        datasets: comprehensionData
      },
      option: {
        title: {
          display: true,
          text: '持続・集中'
        },
        scale:{
          ticks:{
            suggestedMin: 1,
            suggestedMax: 5,
            stepSize: 1
          }
        }
      }
    },
    concentration: {
      data: {
        labels: ['集中して塗り絵に取り組む', '色番号毎に適切に塗れている', '塗り残しがない'],
        datasets: concentrationData
      },
      option: {
        title: {
          display: true,
          text: '持続・集中'
        },
        scale:{
          ticks:{
            suggestedMin: 1,
            suggestedMax: 5,
            stepSize: 1
          }
        }
      }
    },
    politeness: {
      data: {
        labels: ['線からはみ出さずに塗れている', '各々のピースを隙間なく塗っている', '効率良く塗るため　紙の方向を変えている'],
        datasets: politenessData
      },
      option: {
        title: {
          display: true,
          text: '丁寧さ'
        },
        scale:{
          ticks:{
            suggestedMin: 1,
            suggestedMax: 5,
            stepSize: 1
          }
        }
      }
    }
  }
  new Chart(setup, { type: 'radar', data: chartConfig.setup.data, options: chartConfig.setup.option })
  new Chart(comprehension, { type: 'radar', data: chartConfig.comprehension.data, options: chartConfig.comprehension.option})
  new Chart(concentration, { type: 'radar', data: chartConfig.concentration.data, options: chartConfig.concentration.option})
  new Chart(politeness, { type: 'radar', data: chartConfig.politeness.data, options: chartConfig.politeness.option})
}

function conpressTime100Percent(logStr) {
  log = JSON.parse(logStr)
  maxTime = log[log.length - 1][0]
  var res = []
  for (let i = 1; i < log.length; i++) {
    res.push({x: Math.floor((log[i][0] / maxTime) * 100), y: log[i][1]})
  }
  return res
}