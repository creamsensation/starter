import 'htmx.org'
import LazyLoad from 'vanilla-lazyload'
import * as Stimulus from '@hotwired/stimulus'

document.addEventListener('DOMContentLoaded', () => {
	const app = Stimulus.Application.start()
	const lazyload = new LazyLoad({})
})
