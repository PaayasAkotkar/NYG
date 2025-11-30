import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RefSetupBoardComponent } from './ref-setup-board.component';

describe('RefSetupBoardComponent', () => {
  let component: RefSetupBoardComponent;
  let fixture: ComponentFixture<RefSetupBoardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RefSetupBoardComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RefSetupBoardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
