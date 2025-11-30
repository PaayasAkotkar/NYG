import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TagsComponent } from './tags.component';

describe('TagsComponent', () => {
  let component: TagsComponent;
  let fixture: ComponentFixture<TagsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TagsComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(TagsComponent);
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
    const singleColor = component.singleColor

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
