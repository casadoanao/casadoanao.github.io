const upperLinks = document.querySelectorAll('a.upperLink')
const actualPage = window.location.pathname

function addElementClass(el, classToAdd) {
	return function () {
		el.classList.add(classToAdd)
	}
}

function removeElementClass(el, classToRemove) {
	return function () {
		el.classList.remove(classToRemove)
	}
}

for (const link in upperLinks) {
	const referredKey = upperLinks[link]
	referredKey.onmouseenter = addElementClass(referredKey, 'active')
	referredKey.onmouseout = removeElementClass(referredKey, 'active')

	if (referredKey.href === actualPage) {
		referredKey.innerText = "Você está aqui!"
	}
}
