function setTab(tab) {
  let params = new URLSearchParams()
  params.append('tab', tab)
  axios.post('/api/TabSelect', params).then(() => {
    window.location.reload()
  })
}
