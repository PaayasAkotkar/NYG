import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsCardComponent } from './ocs-card.component';

describe('OcsCardComponent', () => {
  let component: OcsCardComponent;
  let fixture: ComponentFixture<OcsCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsCardComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
