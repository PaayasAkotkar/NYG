import { ElementRef, Renderer2 } from '@angular/core';
import { OnFreezeDirective } from './on-freeze.directive';

describe('OnFreezeDirective', () => {
  let ref: ElementRef<HTMLDivElement>
  let dummy: HTMLDivElement
  let re: jasmine.SpyObj<Renderer2>
  it('should create an instance', () => {
    dummy = document.createElement('div')
    ref = new ElementRef(dummy)
    re = jasmine.createSpyObj('Renderer2', ['setStyle'])
    const directive = new OnFreezeDirective(ref, re);
    expect(directive).toBeTruthy();
  });
});
