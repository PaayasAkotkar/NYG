import { Directive, ElementRef, HostListener, Input, OnInit, Renderer2, signal, WritableSignal } from '@angular/core';

@Directive({
  selector: '[appOnTick]'
})
export class OnTickDirective implements OnInit {

  // colorToApplyOn = '#FFB02D'
  // colorToApplyOff = '#792B3A'

  // @Input({ required: true }) isOn: boolean = false
  constructor(private renderer: Renderer2, private ele: ElementRef) { }
  // @HostListener('change', ['$event.target.checked'])
  // onTick(e: Event) {
  //   const input = this.ele.nativeElement; // the host element (assume checkbox)
  //   // null-safe check
  //   const isOn = input?.checked ?? false;

  //   this.Update(isOn); console.log("tick")
  // }
  ngOnInit(): void {
    // this.Update(this.isOn)
  }

  // Update(isOn: boolean) {
  //   var apply = this.colorToApplyOn
  //   if (isOn) {
  //     this.renderer.setStyle(this.ele.nativeElement, 'backgroundColor', apply)
  //   }

  // }
}
