import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsNmeComponent } from './ocs-nme.component';

describe('OcsNmeComponent', () => {
  let component: OcsNmeComponent;
  let fixture: ComponentFixture<OcsNmeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsNmeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsNmeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
