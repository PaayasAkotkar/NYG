import { Directive, ElementRef, HostListener, OnInit } from '@angular/core';

@Directive({
  selector: '[appScrollBehvaior]'
})
export class ScrollBehvaiorDirective {

  constructor(private _ref: ElementRef) { }
  @HostListener('wheel', ['$event'])
  onWheelScroll(event: WheelEvent) {
    // Prevent default vertical scrolling
    event.preventDefault();

    // important for touchup on lappy
    const amount = event.deltaX !== 0 ? event.deltaX : event.deltaY
    this._ref.nativeElement.scrollLeft += amount
  }
}
