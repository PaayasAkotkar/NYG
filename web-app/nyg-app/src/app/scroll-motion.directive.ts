import { Directive, ElementRef, HostListener, inject, Input } from "@angular/core";

@Directive({
  selector: 'appScrollMotion',
})
export class ScrollMotion {
  track = 0
  @HostListener('windows.scroll', ['$event.target'])
  Scroll(events: any) {
    const len = events[0].length
    this.track += 1
    if (this.track == len) {
      this.track = 1
    }
  }
}