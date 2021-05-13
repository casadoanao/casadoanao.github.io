(function() {
	const upperLinks = document.querySelectorAll('a.upperLink')
	const currentPage = window.location.href

	function addElementClass(el, classToAdd) {
		return () => {
			el.classList.add(classToAdd)
		}
	}

	function removeElementClass(el, classToRemove) {
		return () => {
			el.classList.remove(classToRemove)
		}
	}

	for (const link in upperLinks) {
		const element = upperLinks[link]
		element.onmouseenter = addElementClass(element, 'active')
		element.onmouseout = removeElementClass(element, 'active')

		if (element.href === currentPage) {
			element.title = "Você está aqui"
		}
	}
})()
