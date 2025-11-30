import { Directive, ElementRef, OnInit, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appOnFreeze]'
})
export class OnFreezeDirective implements OnInit {

  ngOnInit(): void {
    this.Freeze()
    console.log('freeze: ', true)
  }

  constructor(private el: ElementRef<HTMLDivElement>, private renderer: Renderer2) { }

  Freeze() {
    // this.el.nativeElement.style.width = '100%'
    // this.el.nativeElement.style.height = '100%'
    this.renderer.setStyle(this.el.nativeElement, 'top', '0')
    this.renderer.setStyle(this.el.nativeElement, 'bottom', '0')
    this.renderer.setStyle(this.el.nativeElement, 'left', '0')
    this.renderer.setStyle(this.el.nativeElement, 'right', '0px')
    this.renderer.setStyle(this.el.nativeElement, 'background', 'rgba(255,255,255,0)')
    this.renderer.setStyle(this.el.nativeElement, 'borderRadius', '16px')
    this.renderer.setStyle(this.el.nativeElement, 'boxShadow', '0 4px 30px rgba(0,0,0,0.1)')
    this.renderer.setStyle(this.el.nativeElement, 'backdropFilter', 'blur(0px)')
    this.renderer.setStyle(this.el.nativeElement, 'border', '1px solid rgba(255,255,255,0.42)')
    this.renderer.setStyle(this.el.nativeElement, 'position', 'absolute')
    this.renderer.addClass(this.el.nativeElement, 'freeze-it')
  }

}