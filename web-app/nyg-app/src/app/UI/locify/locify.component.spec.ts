import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LocifyComponent } from './locify.component';

describe('LocifyComponent', () => {
  let component: LocifyComponent;
  let fixture: ComponentFixture<LocifyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LocifyComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LocifyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
