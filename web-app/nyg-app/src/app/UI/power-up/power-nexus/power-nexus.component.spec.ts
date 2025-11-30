import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerNexusComponent } from './power-nexus.component';

describe('PowerNexusComponent', () => {
  let component: PowerNexusComponent;
  let fixture: ComponentFixture<PowerNexusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerNexusComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PowerNexusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
