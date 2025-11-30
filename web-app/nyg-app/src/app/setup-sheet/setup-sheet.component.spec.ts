import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SetupSheetComponent } from './setup-sheet.component';

describe('SetupSheetComponent', () => {
  let component: SetupSheetComponent;
  let fixture: ComponentFixture<SetupSheetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SetupSheetComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SetupSheetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
