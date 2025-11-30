import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewPattern4Component } from './view-pattern-4.component';

describe('ViewPattern4Component', () => {
  let component: ViewPattern4Component;
  let fixture: ComponentFixture<ViewPattern4Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ViewPattern4Component]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewPattern4Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
