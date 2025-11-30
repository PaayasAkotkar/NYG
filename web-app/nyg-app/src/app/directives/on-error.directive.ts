import { isPlatformBrowser } from '@angular/common';
import { afterNextRender, afterRender, Directive, effect, ElementRef, Inject, Injector, Input, OnChanges, OnInit, PLATFORM_ID, Renderer2, signal, SimpleChanges, WritableSignal } from '@angular/core';

@Directive({
  selector: '[appOnError]'
})
export class OnErrorDirective implements OnChanges {

  @Input({ required: true }) isError: boolean
  ngOnChanges(changes: SimpleChanges): void {
    if (changes['isError']) {
      if (isPlatformBrowser(this.pid)) {
        const element = this.el.nativeElement;
        if (document.activeElement !== element) {
          setTimeout(() => element.focus(), 0);
        }
      }
    }
  }
  constructor(@Inject(PLATFORM_ID) private pid: Object, private el: ElementRef<HTMLInputElement>, private renderer: Renderer2) {

  }


}
