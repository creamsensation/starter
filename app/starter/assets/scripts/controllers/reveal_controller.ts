import { Controller } from "npm:@hotwired/stimulus"

import { Class } from '../class.ts'

export default class RevealController extends Controller {
	static targets = ['open', 'close', 'item']
	declare readonly openTarget: HTMLButtonElement
	declare readonly closeTarget: HTMLButtonElement
	declare readonly itemTargets: HTMLElement[]
	
	
	handleOpen() {
		this.openTarget.classList.remove(Class.isActive)
		this.closeTarget.classList.add(Class.isActive)
		for (const item of this.itemTargets) {
			item.classList.add(Class.isActive)
		}
	}
	
	handleClose() {
		this.closeTarget.classList.remove(Class.isActive)
		this.openTarget.classList.add(Class.isActive)
		for (const item of this.itemTargets) {
			item.classList.remove(Class.isActive)
		}
	}
	
	handleClickOutside(e: Event) {
		if (this.element == e.target || this.element.contains(e.target)) {
			return
		}
		this.handleClose()
	}
}