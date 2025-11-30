import { isPlatformBrowser } from '@angular/common';
import { AfterViewInit, Directive, ElementRef, Inject, Input, NgZone, OnChanges, OnInit, PLATFORM_ID, Renderer2, SimpleChanges } from '@angular/core';

@Directive({
  selector: '[appOnProgress]'
})
export class OnProgressDirective implements OnChanges {
  @Input({ required: true }) NYGProgress: number

  constructor(private renderer: Renderer2, private ele: ElementRef,
    @Inject(PLATFORM_ID) private platformId: Object, private run: NgZone) { }
  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGProgress']) {
      this.renderer.setProperty(this.ele.nativeElement, 'style', `--progress:${this.NYGProgress}%`);
      this.renderer.setStyle(this.ele.nativeElement, 'width', `${this.NYGProgress}%`)
      console.log('progress: ', this.NYGProgress)

    }
  }
}
