const upperLinks = document.getElementsByClassName('upperLink')
const elements = Object.keys(upperLinks)

function addElementClass(el, classToAdd) {
	return function () {
		el.classList.add(classToAdd)
	}
}

function removeElementClass(el, classToAdd) {
	return function () {
		el.classList.remove(classToAdd)
	}
}


for (const actualElement in elements) {
	actualElement.onmouseenter = addElementClass(actualElement, 'active')
	actualElement.onmouseout = removeElementClass(actualElement, 'active')
}