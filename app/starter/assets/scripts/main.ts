import 'npm:htmx.org'
import LazyLoad from 'npm:vanilla-lazyload'
import * as Stimulus from 'npm:@hotwired/stimulus'
import RevealController from './controllers/reveal_controller.ts'


document.addEventListener('DOMContentLoaded', () => {
	const app = Stimulus.Application.start()
	app.register('reveal', RevealController)
	const lazyload = new LazyLoad({})
})

