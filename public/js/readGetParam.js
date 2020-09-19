function readGetParam() {
  let link = location.search
  let value = getParam(link, "page")
  let prev = document.getElementById("prev")
  let next = document.getElementById("next")
  console.log(prev)
  console.log(next)
  return value
}

function getParam(url, name) {
  if (!url) url = window.location.href;
  name = name.replace(/[\[\]]/g, "\\$&");
  var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
      results = regex.exec(url);
  if (!results) return null;
  if (!results[2]) return '';
  return decodeURIComponent(results[2].replace(/\+/g, " "));
}
