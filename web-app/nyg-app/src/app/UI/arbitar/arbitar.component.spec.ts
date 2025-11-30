import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ArbitarComponent } from './arbitar.component';

describe('ArbitarComponent', () => {
  let component: ArbitarComponent;
  let fixture: ComponentFixture<ArbitarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ArbitarComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ArbitarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
