import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerTagComponent } from './power-tag.component';

describe('PowerTagComponent', () => {
  let component: PowerTagComponent;
  let fixture: ComponentFixture<PowerTagComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerTagComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PowerTagComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
