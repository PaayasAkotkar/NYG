import { ChangeDetectorRef, ElementRef, NgZone, PLATFORM_ID, Renderer2 } from '@angular/core';
import { OnStyleDirective } from './on-style.directive';
import { TestBed } from '@angular/core/testing';

describe('OnStyleDirective', () => {
  let dummyRenderer: Renderer2
  let zone: NgZone
  let cdr: ChangeDetectorRef
  beforeEach(() => {
    dummyRenderer = {
      addClass: jasmine.createSpy('addClass')
    } as unknown as Renderer2
    TestBed.configureTestingModule({
      imports: [OnStyleDirective],
      providers: [{ proivde: PLATFORM_ID, use_value: 'browser' }, NgZone, ChangeDetectorRef]
    })
    zone = TestBed.inject(NgZone)
    cdr = TestBed.inject(ChangeDetectorRef)
  })


  it('should create an instance', () => {
    const divE = document.createElement('div')
    let dummy: ElementRef = new ElementRef(divE)
    const directive = new OnStyleDirective(zone, cdr, dummy, dummyRenderer);
    expect(directive).toBeTruthy();
  });
});
