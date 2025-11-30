import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoadLogoComponent } from './load-logo.component';

describe('LoadLogoComponent', () => {
  let component: LoadLogoComponent;
  let fixture: ComponentFixture<LoadLogoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LoadLogoComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoadLogoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
