const digaTchau = document.querySelector('button#tchau')
const msgTchau = document.querySelector('div#msgTchau')
var jaRespondido = false

digaTchau.addEventListener('click', () => {
	let adeus = confirm('Isso é um adeus?')
	if (adeus) {
		document.body.innerHTML = '<a href="https://duckduckgo.com">ADEUS!</a>'
	} else {
		if (!jaRespondido) {
			msgTchau.innerText = 'Por que disse adeus se não ia sair?'
			jaRespondido = true
		}
	}
})