import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SetChallengeComponent } from './set-challenge.component';

describe('SetChallengeComponent', () => {
  let component: SetChallengeComponent;
  let fixture: ComponentFixture<SetChallengeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SetChallengeComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(SetChallengeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
