import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewPattern3Component } from './view-pattern3.component';

describe('ViewPattern3Component', () => {
  let component: ViewPattern3Component;
  let fixture: ComponentFixture<ViewPattern3Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ViewPattern3Component]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewPattern3Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
