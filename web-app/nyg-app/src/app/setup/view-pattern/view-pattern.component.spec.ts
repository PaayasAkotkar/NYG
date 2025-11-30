import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewPatternComponent } from './view-pattern.component';

describe('ViewPatternComponent', () => {
  let component: ViewPatternComponent;
  let fixture: ComponentFixture<ViewPatternComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ViewPatternComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewPatternComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
