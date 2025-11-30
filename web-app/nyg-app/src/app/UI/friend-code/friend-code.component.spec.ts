import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FriendCodeComponent } from './friend-code.component';

describe('FriendCodeComponent', () => {
  let component: FriendCodeComponent;
  let fixture: ComponentFixture<FriendCodeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FriendCodeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FriendCodeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
