import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsDetComponent } from './ocs-det.component';

describe('OcsDetComponent', () => {
  let component: OcsDetComponent;
  let fixture: ComponentFixture<OcsDetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsDetComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsDetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
