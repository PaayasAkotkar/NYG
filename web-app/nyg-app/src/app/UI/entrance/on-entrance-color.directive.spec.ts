import { ElementRef, Renderer2 } from '@angular/core';
import { OnEntranceColorDirective } from './on-entrance-color.directive';

describe('OnEntranceColorDirective', () => {
  let dummy: HTMLDivElement
  let ref: ElementRef<HTMLDivElement>
  let ren: jasmine.SpyObj<Renderer2>
  it('should create an instance', () => {
    ren = jasmine.createSpyObj('Renderer2', ['setStyle'])
    dummy = document.createElement('div')
    ref = new ElementRef<HTMLDivElement>(dummy)
    const directive = new OnEntranceColorDirective(ref, ren);
    expect(directive).toBeTruthy();
  });
});
