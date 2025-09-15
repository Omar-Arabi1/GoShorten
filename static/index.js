function copyText() {
  let shortUrl = document.getElementById("shortened-url");
  console.log(shortUrl.innerText);
  navigator.clipboard.writeText(shortUrl.innerText);
}
