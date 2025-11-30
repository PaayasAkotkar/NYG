import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewCheatSheetComponent } from './view-cheat-sheet.component';

describe('ViewCheatSheetComponent', () => {
  let component: ViewCheatSheetComponent;
  let fixture: ComponentFixture<ViewCheatSheetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ViewCheatSheetComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewCheatSheetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
