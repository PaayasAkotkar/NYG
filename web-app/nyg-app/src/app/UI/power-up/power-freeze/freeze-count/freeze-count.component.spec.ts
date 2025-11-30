import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FreezeCountComponent } from './freeze-count.component';

describe('FreezeCountComponent', () => {
  let component: FreezeCountComponent;
  let fixture: ComponentFixture<FreezeCountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FreezeCountComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FreezeCountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
