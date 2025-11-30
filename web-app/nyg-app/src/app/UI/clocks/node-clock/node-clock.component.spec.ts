import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NodeClockComponent } from './node-clock.component';

describe('NodeClockComponent', () => {
  let component: NodeClockComponent;
  let fixture: ComponentFixture<NodeClockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NodeClockComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NodeClockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
