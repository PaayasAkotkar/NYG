import { afterRender, AfterViewInit, Directive, effect, ElementRef, Input, OnChanges, OnInit, Renderer2, SimpleChanges, WritableSignal } from '@angular/core';

@Directive({
  selector: '[appScrollView]'
})
export class ScrollViewDirective {
  @Input({ required: true }) Position: WritableSignal<number>


  constructor(private ref: ElementRef<HTMLDivElement>, private render: Renderer2) {
    effect(() => {
      if (this.ref.nativeElement) {
        this.render.setProperty(this.ref.nativeElement, "scrollTop", this.Position())
      }

    })

  }

}
