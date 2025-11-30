import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WatcherViewBoxComponent } from './view-box.component';

describe('WatcherViewBoxComponent', () => {
  let component: WatcherViewBoxComponent;
  let fixture: ComponentFixture<WatcherViewBoxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WatcherViewBoxComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(WatcherViewBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
