import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsManualComponent } from './ocs-manual.component';

describe('OcsManualComponent', () => {
  let component: OcsManualComponent;
  let fixture: ComponentFixture<OcsManualComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsManualComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsManualComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
