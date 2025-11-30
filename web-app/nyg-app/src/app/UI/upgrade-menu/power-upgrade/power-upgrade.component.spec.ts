import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PowerUpgradeComponent } from './power-upgrade.component';
import { ElementRef, QueryList } from '@angular/core';

describe('PowerUpgradeComponent', () => {
  let component: PowerUpgradeComponent;
  let fixture: ComponentFixture<PowerUpgradeComponent>;
  let dummySlots: QueryList<ElementRef<HTMLDivElement>>
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PowerUpgradeComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(PowerUpgradeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should update the style of elem as per donatedSpur', () => {
    fixture.detectChanges()

    component.slot
    const elems = component.slot.toArray()
    component.OCSdonatedSpurs = 40
    expect(component.slot).toBeTruthy()
    expect(component.slot.length).toBe(10)
    component.update()
    fixture.detectChanges()

    var only = component.OCSdonatedSpurs / 5
    var defaultFrom = '#C2BD00'
    var from = '#CED0D0'
    var percentage = `${component.OCSdonatedSpurs}%`
    var percentage2 = `${component.OCSdonatedSpurs / 2}%`
    var fromPer = '--fromPer'
    var toPer = '--toPer'
    var to = '#F0ECF0'

    for (let i = 0; i < only; i++) {
      expect(elems[i].nativeElement.style.getPropertyValue('--from')).toBe(defaultFrom)
      expect(elems[i].nativeElement.style.getPropertyValue('--to')).toBe(from)
      expect(elems[i].nativeElement.style.getPropertyValue(fromPer)).toBe(percentage)
      expect(elems[i].nativeElement.style.getPropertyValue(toPer)).toBe(percentage2)
    }

    for (let i = only; i < component.slot.length; i++) {
      expect(elems[i].nativeElement.style.getPropertyValue('--from')).toBe(from)
      expect(elems[i].nativeElement.style.getPropertyValue('--to')).toBe(to)
      expect(elems[i].nativeElement.style.getPropertyValue(fromPer)).toBe(percentage)
      expect(elems[i].nativeElement.style.getPropertyValue(toPer)).toBe(percentage2)
    }


  })
});
