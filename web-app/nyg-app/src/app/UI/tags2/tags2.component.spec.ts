import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Tags2Component } from './tags2.component';

describe('Tags2Component', () => {
  let component: Tags2Component;
  let fixture: ComponentFixture<Tags2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Tags2Component]
    })
      .compileComponents();

    fixture = TestBed.createComponent(Tags2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should set mat-card bg based on mode', () => {
    const src = component.NYGMode = 'mysterium'
    const bg = component.bg
    const fromColor = component.fromColor
    const toColor = component.toColor

    component.update()
    fixture.detectChanges()
    const dummy = fixture.nativeElement as HTMLElement
    const header = dummy.querySelector('.nyg-header') as HTMLSpanElement
    expect(header).toBeTruthy()
    const cardElem = dummy.querySelector('mat-card .nyg-tag') as HTMLElement
    expect(cardElem).toBeTruthy()
    const cardStyle = window.getComputedStyle(cardElem)

    const _bg =
      `linear-gradient(to left, ${fromColor}, ${toColor})`;
    expect(cardStyle.background).toBe(bg)
    expect(cardStyle.background).toBe(_bg)
  })
});
