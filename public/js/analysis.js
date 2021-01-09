function setTab(tab) {
  let params = new URLSearchParams()
  params.append('tab', tab)
  axios.post('/ptv3/api/TabSelect', params).then(() => {
    window.location.reload()
  })
}
