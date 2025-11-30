import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsCostSheetComponent } from './ocs-cost-sheet.component';
import { write } from 'fs';

describe('OcsCostSheetComponent', () => {
  let component: OcsCostSheetComponent;
  let fixture: ComponentFixture<OcsCostSheetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsCostSheetComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(OcsCostSheetComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should set the header', () => {
    const writer1 = 'test'
    const writer2 = 'completed'
    component.OCSheader = writer1
    component.OCScostSheet = writer2
    fixture.detectChanges()

    const dummy = fixture.nativeElement as HTMLElement
    const headers = dummy.querySelectorAll('.ocs-header') as NodeListOf<HTMLSpanElement>
    const header = Array.from(headers).map(h => (h.textContent ?? '').trim())

    expect(header.length).toBe(2)
    expect(header[0]).toBe(writer1)
    expect(header[1]).toBe(writer2)


  })
});
